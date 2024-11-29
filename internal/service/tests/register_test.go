package tests

import (
	"database/sql"
	"reflect"
	"testing"

	"github.com/AndIsaev/go-musthave-diploma-tlp/internal/service"

	"github.com/AndIsaev/go-musthave-diploma-tlp/internal/model"
)

func TestMethodsRegister(t *testing.T) {
	tests := []struct {
		name    string
		setup   func(ts *testSuite)
		args    *model.AuthParams
		want    *model.UserWithToken
		wantErr bool
	}{
		{
			name: "success registration",
			setup: func(ts *testSuite) {
				ts.mockUserRepo.EXPECT().
					GetUserByLogin(ts.ctx, &model.UserLogin{Username: "test_user"}).
					Return(nil, sql.ErrNoRows)

				ts.mockUserRepo.EXPECT().
					CreateUser(ts.ctx, &model.AuthParams{Login: "test_user", Password: "password"}).
					Return(&model.UserWithToken{Token: "token", Login: "test_user", ID: 1}, nil)
			},
			args:    &model.AuthParams{Login: "test_user", Password: "password"},
			want:    &model.UserWithToken{Token: "token", Login: "test_user", ID: 1},
			wantErr: false,
		},
		{
			name: "unsuccessful registration due to existing user",
			setup: func(ts *testSuite) {
				ts.mockUserRepo.EXPECT().
					GetUserByLogin(ts.ctx, &model.UserLogin{Username: "existing_user"}).
					Return(&model.User{Login: "existing_user"}, nil)
			},
			args:    &model.AuthParams{Login: "existing_user", Password: "password"},
			want:    nil,
			wantErr: true,
		},
		{
			name: "unsuccessful registration, if connection closed  not defined",
			setup: func(ts *testSuite) {
				ts.mockUserRepo.EXPECT().
					GetUserByLogin(ts.ctx, &model.UserLogin{Username: "existing_user"}).
					Return(nil, sql.ErrConnDone)
			},
			args:    &model.AuthParams{Login: "existing_user", Password: "password"},
			want:    nil,
			wantErr: true,
		},
		{
			name: "unsuccessful registration, if connection closed not defined",
			setup: func(ts *testSuite) {
				ts.mockUserRepo.EXPECT().
					GetUserByLogin(ts.ctx, &model.UserLogin{Username: "existing_user"}).
					Return(nil, sql.ErrConnDone)
			},
			args:    &model.AuthParams{Login: "existing_user", Password: "password"},
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
			got, err := s.Register(ts.ctx, tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("Register() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Register() got = %v, want %v", got, tt.want)
			}
		})
	}
}
