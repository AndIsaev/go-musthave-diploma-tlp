package postgres

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"

	"github.com/AndIsaev/go-musthave-diploma-tlp/internal/model"
)

func (p *PgStorage) GetUserByLogin(ctx context.Context, userLogin *model.UserLogin) (*model.User, error) {
	user := &model.User{}
	query := "SELECT id, login, password FROM users WHERE login = $1;"

	err := p.db.GetContext(ctx, user, query, userLogin.Username)
	if err != nil {
		log.Printf("error when try select info of user %v\n", userLogin.Username)
		return nil, err
	}
	return user, nil
}

func (p *PgStorage) CreateUser(ctx context.Context, params *model.AuthParams) (*model.UserWithToken, error) {
	user := &model.User{}
	query := `INSERT INTO users (login, password) VALUES ($1, $2) RETURNING id, login;`

	hashedPassword, _ := HashPassword(params.Password)

	err := p.db.QueryRowContext(ctx, query, params.Login, hashedPassword).Scan(&user.ID, &user.Login)
	if err != nil {
		log.Println("can't insert data of user")
		return nil, err
	}

	token, err := GenerateJWT(user)
	if err != nil {
		log.Println("error when try to generate token")
		return nil, err
	}

	return &model.UserWithToken{Login: user.Login, ID: user.ID, Token: token.Token}, nil
}

func (p *PgStorage) Auth(ctx context.Context, params *model.AuthParams) (*model.UserWithToken, error) {
	user, err := p.GetUserByLogin(ctx, &model.UserLogin{Username: params.Login})
	if err != nil {
		log.Printf("error get user by login: %v\n", err.Error())
		return nil, err
	}

	err = CheckPasswordHash(params.Password, user.Password)
	if err != nil {
		log.Println("incorrect password")
		return nil, err
	}

	token, err := GenerateJWT(user)
	if err != nil {
		return nil, err
	}

	return &model.UserWithToken{Login: user.Login, ID: user.ID, Token: token.Token}, nil
}

func CheckPasswordHash(password, hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func GenerateJWT(user *model.User) (*model.UserToken, error) {
	expirationTime := time.Now().Add(time.Hour * 24 * 7) // 7 days
	claims := &model.Claims{
		Login: user.Login,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	var jwtKey = []byte(os.Getenv("SECRET_KEY"))

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		log.Printf("problem generate token for user: %v\n", err.Error())
		return nil, err
	}
	return &model.UserToken{Token: tokenString}, nil
}
