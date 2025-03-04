// Code generated by MockGen. DO NOT EDIT.
// Source: ./rest.go

// Package tests is a generated GoMock package.
package tests

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockRESTInterface is a mock of RESTInterface interface.
type MockRESTInterface struct {
	ctrl     *gomock.Controller
	recorder *MockRESTInterfaceMockRecorder
}

// MockRESTInterfaceMockRecorder is the mock recorder for MockRESTInterface.
type MockRESTInterfaceMockRecorder struct {
	mock *MockRESTInterface
}

// NewMockRESTInterface creates a new mock instance.
func NewMockRESTInterface(ctrl *gomock.Controller) *MockRESTInterface {
	mock := &MockRESTInterface{ctrl: ctrl}
	mock.recorder = &MockRESTInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRESTInterface) EXPECT() *MockRESTInterfaceMockRecorder {
	return m.recorder
}

// WatchUpdates mocks base method.
func (m *MockRESTInterface) WatchUpdates(ctx context.Context, environmentId string, includeInitialData bool) WatcherUpdate {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WatchUpdates", ctx, environmentId, includeInitialData)
	ret0, _ := ret[0].(WatcherUpdate)
	return ret0
}

// WatchUpdates indicates an expected call of WatchUpdates.
func (mr *MockRESTInterfaceMockRecorder) WatchUpdates(ctx, environmentId, includeInitialData interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WatchUpdates", reflect.TypeOf((*MockRESTInterface)(nil).WatchUpdates), ctx, environmentId, includeInitialData)
}
