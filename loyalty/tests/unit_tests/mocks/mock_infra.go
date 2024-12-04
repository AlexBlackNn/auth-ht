// Code generated by MockGen. DO NOT EDIT.
// Source: ../loyalty/internal/services/loyaltyservice/loyalty.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	domain "github.com/AlexBlackNn/authloyalty/loyalty/internal/domain"
	broker "github.com/AlexBlackNn/authloyalty/loyalty/pkg/broker"
	gomock "github.com/golang/mock/gomock"
)

// MockloyaltyBroker is a mock of loyaltyBroker interface.
type MockloyaltyBroker struct {
	ctrl     *gomock.Controller
	recorder *MockloyaltyBrokerMockRecorder
}

// MockloyaltyBrokerMockRecorder is the mock recorder for MockloyaltyBroker.
type MockloyaltyBrokerMockRecorder struct {
	mock *MockloyaltyBroker
}

// NewMockloyaltyBroker creates a new mock instance.
func NewMockloyaltyBroker(ctrl *gomock.Controller) *MockloyaltyBroker {
	mock := &MockloyaltyBroker{ctrl: ctrl}
	mock.recorder = &MockloyaltyBrokerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockloyaltyBroker) EXPECT() *MockloyaltyBrokerMockRecorder {
	return m.recorder
}

// GetMessageChan mocks base method.
func (m *MockloyaltyBroker) GetMessageChan() chan *broker.MessageReceived {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMessageChan")
	ret0, _ := ret[0].(chan *broker.MessageReceived)
	return ret0
}

// GetMessageChan indicates an expected call of GetMessageChan.
func (mr *MockloyaltyBrokerMockRecorder) GetMessageChan() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMessageChan", reflect.TypeOf((*MockloyaltyBroker)(nil).GetMessageChan))
}

// MockloyaltyStorage is a mock of loyaltyStorage interface.
type MockloyaltyStorage struct {
	ctrl     *gomock.Controller
	recorder *MockloyaltyStorageMockRecorder
}

// MockloyaltyStorageMockRecorder is the mock recorder for MockloyaltyStorage.
type MockloyaltyStorageMockRecorder struct {
	mock *MockloyaltyStorage
}

// NewMockloyaltyStorage creates a new mock instance.
func NewMockloyaltyStorage(ctrl *gomock.Controller) *MockloyaltyStorage {
	mock := &MockloyaltyStorage{ctrl: ctrl}
	mock.recorder = &MockloyaltyStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockloyaltyStorage) EXPECT() *MockloyaltyStorageMockRecorder {
	return m.recorder
}

// AddLoyalty mocks base method.
func (m *MockloyaltyStorage) AddLoyalty(ctx context.Context, loyalty *domain.UserLoyalty) (*domain.UserLoyalty, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddLoyalty", ctx, loyalty)
	ret0, _ := ret[0].(*domain.UserLoyalty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddLoyalty indicates an expected call of AddLoyalty.
func (mr *MockloyaltyStorageMockRecorder) AddLoyalty(ctx, loyalty interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddLoyalty", reflect.TypeOf((*MockloyaltyStorage)(nil).AddLoyalty), ctx, loyalty)
}

// GetLoyalty mocks base method.
func (m *MockloyaltyStorage) GetLoyalty(ctx context.Context, loyalty *domain.UserLoyalty) (*domain.UserLoyalty, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLoyalty", ctx, loyalty)
	ret0, _ := ret[0].(*domain.UserLoyalty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLoyalty indicates an expected call of GetLoyalty.
func (mr *MockloyaltyStorageMockRecorder) GetLoyalty(ctx, loyalty interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLoyalty", reflect.TypeOf((*MockloyaltyStorage)(nil).GetLoyalty), ctx, loyalty)
}

// HealthCheck mocks base method.
func (m *MockloyaltyStorage) HealthCheck(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HealthCheck", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// HealthCheck indicates an expected call of HealthCheck.
func (mr *MockloyaltyStorageMockRecorder) HealthCheck(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HealthCheck", reflect.TypeOf((*MockloyaltyStorage)(nil).HealthCheck), arg0)
}

// Stop mocks base method.
func (m *MockloyaltyStorage) Stop() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Stop")
	ret0, _ := ret[0].(error)
	return ret0
}

// Stop indicates an expected call of Stop.
func (mr *MockloyaltyStorageMockRecorder) Stop() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockloyaltyStorage)(nil).Stop))
}