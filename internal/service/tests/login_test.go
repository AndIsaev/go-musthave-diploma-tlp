package tests

import (
	"database/sql"
	"errors"
	"github.com/AndIsaev/go-musthave-diploma-tlp/internal/model"
	"github.com/AndIsaev/go-musthave-diploma-tlp/internal/service"
	"reflect"
	"testing"
)

func TestLoginMethod(t *testing.T) {
	tests := []struct {
		name    string
		setup   func(ts *testSuite)
		args    *model.AuthParams
		want    *model.UserWithToken
		wantErr bool
	}{
		{
			name: "success login",
			setup: func(ts *testSuite) {
				ts.mockUserRepo.EXPECT().
					Auth(ts.ctx, &model.AuthParams{Login: "existing_user", Password: "password"}).
					Return(&model.UserWithToken{ID: 1, Login: "existing_user", Token: "token"}, nil)
			},
			args:    &model.AuthParams{Login: "existing_user", Password: "password"},
			want:    &model.UserWithToken{ID: 1, Login: "existing_user", Token: "token"},
			wantErr: false,
		},
		{
			name: "unsuccessful login",
			setup: func(ts *testSuite) {
				ts.mockUserRepo.EXPECT().
					Auth(ts.ctx, &model.AuthParams{Login: "existing_user", Password: "password"}).
					Return(nil, errors.New("error login"))
			},
			args:    &model.AuthParams{Login: "existing_user", Password: "password"},
			want:    nil,
			wantErr: true,
		},
		{
			name: "unsuccessful connection",
			setup: func(ts *testSuite) {
				ts.mockUserRepo.EXPECT().
					Auth(ts.ctx, &model.AuthParams{Login: "existing_user", Password: "password"}).
					Return(nil, sql.ErrConnDone)
			},
			args:    &model.AuthParams{Login: "existing_user", Password: "password"},
			want:    nil,
			wantErr: true,
		},
		{
			name: "unsuccessful login if user doesn't exists",
			setup: func(ts *testSuite) {
				ts.mockUserRepo.EXPECT().
					Auth(ts.ctx, &model.AuthParams{Login: "existing_user", Password: "password"}).
					Return(nil, sql.ErrNoRows)
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
			got, err := s.Login(ts.ctx, tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("Login() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Login() got = %v, want %v", got, tt.want)
			}
		})
	}
}
