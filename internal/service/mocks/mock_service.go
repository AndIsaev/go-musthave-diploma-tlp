// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/AndIsaev/go-musthave-diploma-tlp/internal/service (interfaces: Service)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	model "github.com/AndIsaev/go-musthave-diploma-tlp/internal/model"
	gomock "github.com/golang/mock/gomock"
)

// MockService is a mock of Service interface.
type MockService struct {
	ctrl     *gomock.Controller
	recorder *MockServiceMockRecorder
}

// MockServiceMockRecorder is the mock recorder for MockService.
type MockServiceMockRecorder struct {
	mock *MockService
}

// NewMockService creates a new mock instance.
func NewMockService(ctrl *gomock.Controller) *MockService {
	mock := &MockService{ctrl: ctrl}
	mock.recorder = &MockServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockService) EXPECT() *MockServiceMockRecorder {
	return m.recorder
}

// DeductPoints mocks base method.
func (m *MockService) DeductPoints(arg0 context.Context, arg1 *model.Withdraw, arg2 *model.UserLogin) (*model.Withdraw, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeductPoints", arg0, arg1, arg2)
	ret0, _ := ret[0].(*model.Withdraw)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeductPoints indicates an expected call of DeductPoints.
func (mr *MockServiceMockRecorder) DeductPoints(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeductPoints", reflect.TypeOf((*MockService)(nil).DeductPoints), arg0, arg1, arg2)
}

// GetUserBalance mocks base method.
func (m *MockService) GetUserBalance(arg0 context.Context, arg1 *model.UserLogin) (*model.Balance, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserBalance", arg0, arg1)
	ret0, _ := ret[0].(*model.Balance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserBalance indicates an expected call of GetUserBalance.
func (mr *MockServiceMockRecorder) GetUserBalance(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserBalance", reflect.TypeOf((*MockService)(nil).GetUserBalance), arg0, arg1)
}

// GetUserOrders mocks base method.
func (m *MockService) GetUserOrders(arg0 context.Context, arg1 *model.UserLogin) ([]model.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserOrders", arg0, arg1)
	ret0, _ := ret[0].([]model.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserOrders indicates an expected call of GetUserOrders.
func (mr *MockServiceMockRecorder) GetUserOrders(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserOrders", reflect.TypeOf((*MockService)(nil).GetUserOrders), arg0, arg1)
}

// GetUserWithdrawals mocks base method.
func (m *MockService) GetUserWithdrawals(arg0 context.Context, arg1 *model.UserLogin) ([]model.Withdrawal, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserWithdrawals", arg0, arg1)
	ret0, _ := ret[0].([]model.Withdrawal)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserWithdrawals indicates an expected call of GetUserWithdrawals.
func (mr *MockServiceMockRecorder) GetUserWithdrawals(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserWithdrawals", reflect.TypeOf((*MockService)(nil).GetUserWithdrawals), arg0, arg1)
}

// Login mocks base method.
func (m *MockService) Login(arg0 context.Context, arg1 *model.AuthParams) (*model.UserWithToken, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Login", arg0, arg1)
	ret0, _ := ret[0].(*model.UserWithToken)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Login indicates an expected call of Login.
func (mr *MockServiceMockRecorder) Login(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockService)(nil).Login), arg0, arg1)
}

// Register mocks base method.
func (m *MockService) Register(arg0 context.Context, arg1 *model.AuthParams) (*model.UserWithToken, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Register", arg0, arg1)
	ret0, _ := ret[0].(*model.UserWithToken)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Register indicates an expected call of Register.
func (mr *MockServiceMockRecorder) Register(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockService)(nil).Register), arg0, arg1)
}

// SetOrder mocks base method.
func (m *MockService) SetOrder(arg0 context.Context, arg1 *model.UserOrder) (*model.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetOrder", arg0, arg1)
	ret0, _ := ret[0].(*model.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SetOrder indicates an expected call of SetOrder.
func (mr *MockServiceMockRecorder) SetOrder(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetOrder", reflect.TypeOf((*MockService)(nil).SetOrder), arg0, arg1)
}
