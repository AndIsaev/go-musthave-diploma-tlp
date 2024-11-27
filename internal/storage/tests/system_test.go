package tests

import (
	"context"
	"github.com/AndIsaev/go-musthave-diploma-tlp/internal/storage/mocks"
	"github.com/golang/mock/gomock"
	"testing"
)

func TestSystemRepository_Ping(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockSystemRepository(ctrl)

	ctx := context.Background()

	mockRepo.EXPECT().Ping(ctx).Return(nil)

	err := mockRepo.Ping(ctx)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}

func TestSystemRepository_Close(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockSystemRepository(ctrl)

	ctx := context.Background()

	mockRepo.EXPECT().Close(ctx).Return(nil)

	err := mockRepo.Close(ctx)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}
