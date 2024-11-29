package tests

import (
	"database/sql"
	"github.com/AndIsaev/go-musthave-diploma-tlp/internal/exception"
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
	//anotherUser := &model.User{ID: 2, Login: "another"}

	tests := []struct {
		name        string
		setup       func(ts *testSuite)
		args        *model.UserOrder
		want        *model.Order
		wantErr     bool
		expectedErr error
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
		{
			name: "unsuccessful set order if connection closed",
			setup: func(ts *testSuite) {
				ts.mockUserRepo.EXPECT().
					GetUserByLogin(ts.ctx, &model.UserLogin{Username: login.Username}).
					Return(existsUser, nil)
				ts.mockOrderRepo.EXPECT().
					GetOrderByNumber(ts.ctx, userOrder).
					Return(nil, sql.ErrConnDone)
			},

			args:        userOrder,
			want:        nil,
			wantErr:     true,
			expectedErr: sql.ErrConnDone,
		},
		{
			name: "unsuccessful set order if order already set for another user",
			setup: func(ts *testSuite) {
				newOrder.UserID = 10
				ts.mockUserRepo.EXPECT().
					GetUserByLogin(ts.ctx, &model.UserLogin{Username: login.Username}).
					Return(existsUser, nil)
				ts.mockOrderRepo.EXPECT().
					GetOrderByNumber(ts.ctx, userOrder).
					Return(newOrder, nil)
			},

			args:        userOrder,
			want:        nil,
			wantErr:     true,
			expectedErr: exception.ErrOrderAlreadyExistsAnotherUser,
		},
		{
			name: "unsuccessful set order if order already set for this user",
			setup: func(ts *testSuite) {
				newOrder.UserID = 1
				ts.mockUserRepo.EXPECT().
					GetUserByLogin(ts.ctx, &model.UserLogin{Username: login.Username}).
					Return(existsUser, nil)
				ts.mockOrderRepo.EXPECT().
					GetOrderByNumber(ts.ctx, userOrder).
					Return(newOrder, nil)
			},

			args:        userOrder,
			want:        nil,
			wantErr:     true,
			expectedErr: exception.ErrOrderAlreadyExists,
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
			if !reflect.DeepEqual(err, tt.expectedErr) {
				t.Errorf("SetOrder() err = %v, expectedErr %v", err, tt.expectedErr)
			}
		})
	}
}
