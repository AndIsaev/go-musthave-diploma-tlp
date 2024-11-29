package tests

import (
	"database/sql"
	"reflect"
	"testing"

	"github.com/AndIsaev/go-musthave-diploma-tlp/internal/model"
	"github.com/AndIsaev/go-musthave-diploma-tlp/internal/service"
)

func TestSetOrderMethod(t *testing.T) {
	login := model.UserLogin{Username: "exists_user"}
	existsUser := &model.User{ID: 1, Login: login.Username}
	userOrder := &model.UserOrder{ID: 1, UserID: existsUser.ID, UserLogin: login, Number: "9278923470"}
	newOrder := &model.Order{ID: &userOrder.ID, Number: userOrder.Number}

	tests := []struct {
		name    string
		setup   func(ts *testSuite)
		args    *model.UserOrder
		want    *model.Order
		wantErr bool
	}{
		{
			name: "success set order of user",
			setup: func(ts *testSuite) {
				ts.mockUserRepo.EXPECT().
					GetUserByLogin(ts.ctx, &model.UserLogin{Username: login.Username}).
					Return(existsUser, nil)
				ts.mockOrderRepo.EXPECT().
					GetOrderByNumber(ts.ctx, userOrder).
					Return(nil, sql.ErrNoRows)

				ts.mockOrderRepo.EXPECT().
					SetUserOrder(ts.ctx, userOrder).
					Return(newOrder, nil)
			},

			args:    userOrder,
			want:    newOrder,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ts := setupTest(t)

			tt.setup(ts)

			s := &service.Methods{
				Storage: ts.mockStorage,
			}
			got, err := s.SetOrder(ts.ctx, tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("SetOrder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetOrder() got = %v, want %v", got, tt.want)
			}
		})
	}
}
