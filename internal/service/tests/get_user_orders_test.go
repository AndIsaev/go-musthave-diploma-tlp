package tests

import (
	"database/sql"
	"reflect"
	"testing"

	"github.com/AndIsaev/go-musthave-diploma-tlp/internal/model"
	"github.com/AndIsaev/go-musthave-diploma-tlp/internal/service"
)

func TestGetUserOrdersMethod(t *testing.T) {
	login := &model.UserLogin{Username: "exists_user"}
	existsUser := &model.User{ID: 1, Login: login.Username}
	orders := []model.Order{model.Order{UserID: existsUser.ID}, model.Order{UserID: existsUser.ID}}

	tests := []struct {
		name    string
		setup   func(ts *testSuite)
		args    model.UserLogin
		want    []model.Order
		wantErr bool
	}{
		{
			name: "success get orders of user",
			setup: func(ts *testSuite) {
				ts.mockUserRepo.EXPECT().
					GetUserByLogin(ts.ctx, login).
					Return(existsUser, nil)
				ts.mockOrderRepo.EXPECT().
					ListOrdersByUserID(ts.ctx, existsUser.ID).
					Return(orders, nil)
			},

			args:    *login,
			want:    orders,
			wantErr: false,
		},
		{
			name: "unsuccessful get orders if user doesn't exists",
			setup: func(ts *testSuite) {
				ts.mockUserRepo.EXPECT().
					GetUserByLogin(ts.ctx, login).
					Return(nil, sql.ErrNoRows)
			},

			args:    *login,
			want:    nil,
			wantErr: true,
		},
		{
			name: "success get orders of user if not rows",
			setup: func(ts *testSuite) {
				ts.mockUserRepo.EXPECT().
					GetUserByLogin(ts.ctx, login).
					Return(existsUser, nil)
				ts.mockOrderRepo.EXPECT().
					ListOrdersByUserID(ts.ctx, existsUser.ID).
					Return(nil, sql.ErrNoRows)
			},

			args:    *login,
			want:    nil,
			wantErr: true,
		},
		{
			name: "success get orders of user if not connection",
			setup: func(ts *testSuite) {
				ts.mockUserRepo.EXPECT().
					GetUserByLogin(ts.ctx, login).
					Return(existsUser, nil)
				ts.mockOrderRepo.EXPECT().
					ListOrdersByUserID(ts.ctx, existsUser.ID).
					Return(nil, sql.ErrConnDone)
			},

			args:    *login,
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ts := setupTest(t)

			tt.setup(ts)

			s := &service.Methods{
				Storage: ts.mockStorage,
			}
			got, err := s.GetUserOrders(ts.ctx, &tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUserOrders() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUserOrders() got = %v, want %v", got, tt.want)
			}
		})
	}
}
