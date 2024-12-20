// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/AndIsaev/go-musthave-diploma-tlp/internal/storage (interfaces: WithdrawRepository)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	model "github.com/AndIsaev/go-musthave-diploma-tlp/internal/model"
	gomock "github.com/golang/mock/gomock"
)

// MockWithdrawRepository is a mock of WithdrawRepository interface.
type MockWithdrawRepository struct {
	ctrl     *gomock.Controller
	recorder *MockWithdrawRepositoryMockRecorder
}

// MockWithdrawRepositoryMockRecorder is the mock recorder for MockWithdrawRepository.
type MockWithdrawRepositoryMockRecorder struct {
	mock *MockWithdrawRepository
}

// NewMockWithdrawRepository creates a new mock instance.
func NewMockWithdrawRepository(ctrl *gomock.Controller) *MockWithdrawRepository {
	mock := &MockWithdrawRepository{ctrl: ctrl}
	mock.recorder = &MockWithdrawRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWithdrawRepository) EXPECT() *MockWithdrawRepositoryMockRecorder {
	return m.recorder
}

// CreateWithdraw mocks base method.
func (m *MockWithdrawRepository) CreateWithdraw(arg0 context.Context, arg1 *model.Withdraw, arg2 int) (*model.Withdraw, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateWithdraw", arg0, arg1, arg2)
	ret0, _ := ret[0].(*model.Withdraw)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateWithdraw indicates an expected call of CreateWithdraw.
func (mr *MockWithdrawRepositoryMockRecorder) CreateWithdraw(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateWithdraw", reflect.TypeOf((*MockWithdrawRepository)(nil).CreateWithdraw), arg0, arg1, arg2)
}

// GetListWithdrawnBalance mocks base method.
func (m *MockWithdrawRepository) GetListWithdrawnBalance(arg0 context.Context, arg1 int) ([]model.Withdrawal, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetListWithdrawnBalance", arg0, arg1)
	ret0, _ := ret[0].([]model.Withdrawal)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetListWithdrawnBalance indicates an expected call of GetListWithdrawnBalance.
func (mr *MockWithdrawRepositoryMockRecorder) GetListWithdrawnBalance(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetListWithdrawnBalance", reflect.TypeOf((*MockWithdrawRepository)(nil).GetListWithdrawnBalance), arg0, arg1)
}
