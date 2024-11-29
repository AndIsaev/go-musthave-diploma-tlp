package tests

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"

	"github.com/AndIsaev/go-musthave-diploma-tlp/internal/storage/mocks"
)

type testSuite struct {
	ctrl             *gomock.Controller
	mockStorage      *mocks.MockStorage
	mockUserRepo     *mocks.MockUserRepository
	mockOrderRepo    *mocks.MockOrderRepository
	mockBalanceRepo  *mocks.MockBalanceRepository
	mockWithdrawRepo *mocks.MockWithdrawRepository
	mockSystemRepo   *mocks.MockSystemRepository
	ctx              context.Context
}

// setupTest - use it to prepare the context for testing service.Service
func setupTest(t *testing.T) *testSuite {
	ctrl := gomock.NewController(t)

	mockStorage := mocks.NewMockStorage(ctrl)
	mockUserRepo := mocks.NewMockUserRepository(ctrl)
	mockOrderRepo := mocks.NewMockOrderRepository(ctrl)
	mockBalanceRepo := mocks.NewMockBalanceRepository(ctrl)
	mockWithdrawRepo := mocks.NewMockWithdrawRepository(ctrl)
	mockSystemRepo := mocks.NewMockSystemRepository(ctrl)

	ctx := context.Background()

	mockStorage.EXPECT().User().Return(mockUserRepo).AnyTimes()
	mockStorage.EXPECT().Order().Return(mockOrderRepo).AnyTimes()
	mockStorage.EXPECT().Balance().Return(mockBalanceRepo).AnyTimes()
	mockStorage.EXPECT().Withdraw().Return(mockWithdrawRepo).AnyTimes()
	mockStorage.EXPECT().System().Return(mockSystemRepo).AnyTimes()

	return &testSuite{
		ctrl:             ctrl,
		mockStorage:      mockStorage,
		mockUserRepo:     mockUserRepo,
		mockOrderRepo:    mockOrderRepo,
		mockBalanceRepo:  mockBalanceRepo,
		mockWithdrawRepo: mockWithdrawRepo,
		mockSystemRepo:   mockSystemRepo,
		ctx:              ctx,
	}
}
