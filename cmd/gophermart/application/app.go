package application

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	c "github.com/AndIsaev/go-musthave-diploma-tlp/cmd/gophermart/configuration"
	"github.com/AndIsaev/go-musthave-diploma-tlp/internal/handler"
	"github.com/AndIsaev/go-musthave-diploma-tlp/internal/model"
	"github.com/AndIsaev/go-musthave-diploma-tlp/internal/service"
	"github.com/AndIsaev/go-musthave-diploma-tlp/internal/storage"
	"github.com/AndIsaev/go-musthave-diploma-tlp/internal/storage/postgres"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/go-resty/resty/v2"
)

type App struct {
	Name          string
	Server        *http.Server
	Config        *c.Config
	DBConn        storage.Storage
	Router        chi.Router
	Handler       *handler.Handler
	ErrChan       chan error
	AmountWorkers int
	Client        *resty.Client
}

func NewApp() *App {
	app := &App{Name: "Gophermart", ErrChan: make(chan error), AmountWorkers: 2}
	app.Config = c.NewConfig()
	app.Router = chi.NewRouter()
	return app
}

// StartApp - start app
func (a *App) StartApp() (err error) {
	log.Printf("start app - %v", a.Name)

	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	chOrder := make(chan model.Order)

	conn, err := postgres.NewPgStorage(a.Config.DB)
	if err != nil {
		return err
	}
	a.DBConn = conn
	err = a.upMigrations()
	if err != nil {
		return err
	}
	a.Client = a.initHTTPClient()
	a.Handler = &handler.Handler{Validator: validator.New()}
	a.Handler.UserService = &service.UserMethods{Storage: a.DBConn}

	a.initRouter()

	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(chOrder)
		a.runUpdateOrders(ctx, chOrder)
	}()

	wg.Add(a.AmountWorkers)
	a.runWorkers(ctx, &wg, chOrder)

	wg.Add(1)
	go func() {
		defer wg.Done()
		a.startHTTPServer()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case err := <-a.ErrChan:
				log.Printf("Application encountered an error: %v", err)
				a.Shutdown(ctx)
				close(chOrder)
				cancel()
				return // Выход из горутины
			case <-ctx.Done():
				log.Println("Context was cancelled, shutting down gracefully.")
				a.Shutdown(ctx)
				close(chOrder)
				cancel()
				return
			}
		}
	}()

	wg.Wait()

	return
}

// initHTTPServer - init http server
func (a *App) initHTTPServer() {
	server := &http.Server{}
	server.Addr = a.Config.Address
	server.Handler = a.Router
	a.Server = server
}

// startHTTPServer - start http server
func (a *App) startHTTPServer() {
	a.initHTTPServer()
	log.Printf("start server on: %s\n", a.Config.Address)
	a.ErrChan <- a.Server.ListenAndServe()
}

// Shutdown - close active connections
func (a *App) Shutdown(ctx context.Context) {
	if err := a.DBConn.System().Close(ctx); err != nil {
		log.Println(errors.Unwrap(err))
	}
}

// upMigrations - run migrations of db
func (a *App) upMigrations() error {
	if err := a.DBConn.System().RunMigrations(context.Background()); err != nil {
		return err
	}
	return nil

}

func (a *App) runUpdateOrders(ctx context.Context, ch chan model.Order) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			a.getOrders(ctx, ch)
			time.Sleep(time.Second * 5)
		}
	}
}

func (a *App) getOrders(ctx context.Context, ch chan model.Order) {
	orders, err := a.DBConn.User().ListOrders(ctx)
	if err != nil {
		log.Println("error receiving orders: ", err)
		a.ErrChan <- err
		return // Добавляем возврат, чтобы завершить работу функции при ошибке
	}

	for _, order := range orders {
		select {
		case <-ctx.Done():
			return // Завершите отправку, если контекст отменен
		default:
			ch <- order
		}
	}
}

func (a *App) runWorkers(ctx context.Context, wg *sync.WaitGroup, ch chan model.Order) {
	for w := 1; w <= a.AmountWorkers; w++ {
		go func(w int) {
			defer wg.Done()
			a.worker(ctx, ch, w)
		}(w)
	}

}

func (a *App) worker(ctx context.Context, ch chan model.Order, w int) {
	for {
		select {
		case <-ctx.Done():
			return
		default:

			order := <-ch
			log.Printf("worker #%d sent value - %v to accrual service\n", w, order.Number)
			err := a.getAccrualOrders(&order)
			if err != nil {
				a.ErrChan <- err
			}

			err = a.DBConn.User().UpdateOrder(ctx, &order)
			if err != nil {
				a.ErrChan <- err
			}

			if order.Status == model.PROCESSED {
				// set current balance for user

				err := a.updateCurrentBalance(ctx, order)
				if err != nil {
					a.ErrChan <- err
				}
			}
		}
	}
}

func (a *App) getAccrualOrders(order *model.Order) error {
	_, err := a.Client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(&order).
		SetResult(&order).
		Get(fmt.Sprintf("/api/orders/%v", order.Number))

	if err != nil {
		return errors.Unwrap(err)
	}

	return nil
}

func (a *App) initHTTPClient() *resty.Client {
	cli := resty.New()
	cli.SetBaseURL(a.Config.Accrual)
	return cli
}

func (a *App) updateCurrentBalance(ctx context.Context, order model.Order) error {
	_, err := a.DBConn.User().GetBalance(ctx, order.UserID)
	if errors.Is(err, sql.ErrNoRows) {
		log.Println("try creating new instance of balance")
		_, err := a.DBConn.User().CreateBalance(ctx, *order.Accrual, order.UserID)
		if err != nil {
			log.Println("error when creating a balance")
			return err
		}
		return nil
	}
	err = a.DBConn.User().UpdateBalance(ctx, *order.Accrual, order.UserID)
	if err != nil {
		log.Println("can't update balance")
		return err
	}
	return nil
}
