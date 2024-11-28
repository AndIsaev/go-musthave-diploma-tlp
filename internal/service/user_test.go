package service

import (
	"context"
	"database/sql"
	"github.com/AndIsaev/go-musthave-diploma-tlp/internal/model"
	"github.com/AndIsaev/go-musthave-diploma-tlp/internal/storage/mocks"
	"github.com/golang/mock/gomock"
	"reflect"
	"testing"
)

func TestMethodsRegister(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStorage := mocks.NewMockStorage(ctrl)
	mockUserRepo := mocks.NewMockUserRepository(ctrl)

	mockStorage.EXPECT().User().Return(mockUserRepo).AnyTimes()

	ctx := context.Background()

	type args struct {
		ctx    context.Context
		params *model.AuthParams
	}
	tests := []struct {
		name    string
		setup   func()
		storage *mocks.MockStorage
		args    args
		want    *model.UserWithToken
		wantErr bool
	}{
		{
			name: "success registration",
			setup: func() {
				mockUserRepo.EXPECT().
					GetUserByLogin(ctx, &model.UserLogin{Username: "test_user"}).
					Return(nil, sql.ErrNoRows)

				mockUserRepo.EXPECT().
					CreateUser(ctx, &model.AuthParams{Login: "test_user", Password: "password"}).
					Return(&model.UserWithToken{Token: "token", Login: "test_user", ID: 1}, nil)
			},
			storage: mockStorage,
			args:    args{ctx: ctx, params: &model.AuthParams{Login: "test_user", Password: "password"}},
			want:    &model.UserWithToken{Token: "token", Login: "test_user", ID: 1},
			wantErr: false,
		},
		{
			name: "unsuccessful registration due to existing user",
			setup: func() {
				mockUserRepo.EXPECT().
					GetUserByLogin(ctx, &model.UserLogin{Username: "existing_user"}).
					Return(&model.User{Login: "existing_user"}, nil)
			},
			storage: mockStorage,
			args:    args{ctx: ctx, params: &model.AuthParams{Login: "existing_user", Password: "password"}},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setup()
			s := &Methods{
				Storage: tt.storage,
			}
			got, err := s.Register(tt.args.ctx, tt.args.params)
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
