// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/kubeshop/testkube-operator/client/testsuites/v3 (interfaces: Interface)

// Package v3 is a generated GoMock package.
package v3

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v3 "github.com/kubeshop/testkube-operator/apis/testsuite/v3"
	v1 "k8s.io/api/core/v1"
)

// MockInterface is a mock of Interface interface.
type MockInterface struct {
	ctrl     *gomock.Controller
	recorder *MockInterfaceMockRecorder
}

// MockInterfaceMockRecorder is the mock recorder for MockInterface.
type MockInterfaceMockRecorder struct {
	mock *MockInterface
}

// NewMockInterface creates a new mock instance.
func NewMockInterface(ctrl *gomock.Controller) *MockInterface {
	mock := &MockInterface{ctrl: ctrl}
	mock.recorder = &MockInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInterface) EXPECT() *MockInterfaceMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockInterface) Create(arg0 *v3.TestSuite) (*v3.TestSuite, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0)
	ret0, _ := ret[0].(*v3.TestSuite)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockInterfaceMockRecorder) Create(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockInterface)(nil).Create), arg0)
}

// CreateTestsuiteSecrets mocks base method.
func (m *MockInterface) CreateTestsuiteSecrets(arg0 *v3.TestSuite) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTestsuiteSecrets", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateTestsuiteSecrets indicates an expected call of CreateTestsuiteSecrets.
func (mr *MockInterfaceMockRecorder) CreateTestsuiteSecrets(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTestsuiteSecrets", reflect.TypeOf((*MockInterface)(nil).CreateTestsuiteSecrets), arg0)
}

// Delete mocks base method.
func (m *MockInterface) Delete(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockInterfaceMockRecorder) Delete(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockInterface)(nil).Delete), arg0)
}

// DeleteAll mocks base method.
func (m *MockInterface) DeleteAll() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAll")
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAll indicates an expected call of DeleteAll.
func (mr *MockInterfaceMockRecorder) DeleteAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAll", reflect.TypeOf((*MockInterface)(nil).DeleteAll))
}

// DeleteByLabels mocks base method.
func (m *MockInterface) DeleteByLabels(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteByLabels", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteByLabels indicates an expected call of DeleteByLabels.
func (mr *MockInterfaceMockRecorder) DeleteByLabels(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteByLabels", reflect.TypeOf((*MockInterface)(nil).DeleteByLabels), arg0)
}

// Get mocks base method.
func (m *MockInterface) Get(arg0 string) (*v3.TestSuite, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0)
	ret0, _ := ret[0].(*v3.TestSuite)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockInterfaceMockRecorder) Get(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockInterface)(nil).Get), arg0)
}

// GetCurrentSecretUUID mocks base method.
func (m *MockInterface) GetCurrentSecretUUID(arg0 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCurrentSecretUUID", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCurrentSecretUUID indicates an expected call of GetCurrentSecretUUID.
func (mr *MockInterfaceMockRecorder) GetCurrentSecretUUID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCurrentSecretUUID", reflect.TypeOf((*MockInterface)(nil).GetCurrentSecretUUID), arg0)
}

// GetSecretTestSuiteVars mocks base method.
func (m *MockInterface) GetSecretTestSuiteVars(arg0, arg1 string) (map[string]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSecretTestSuiteVars", arg0, arg1)
	ret0, _ := ret[0].(map[string]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSecretTestSuiteVars indicates an expected call of GetSecretTestSuiteVars.
func (mr *MockInterfaceMockRecorder) GetSecretTestSuiteVars(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSecretTestSuiteVars", reflect.TypeOf((*MockInterface)(nil).GetSecretTestSuiteVars), arg0, arg1)
}

// List mocks base method.
func (m *MockInterface) List(arg0 string) (*v3.TestSuiteList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0)
	ret0, _ := ret[0].(*v3.TestSuiteList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockInterfaceMockRecorder) List(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockInterface)(nil).List), arg0)
}

// ListLabels mocks base method.
func (m *MockInterface) ListLabels() (map[string][]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListLabels")
	ret0, _ := ret[0].(map[string][]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListLabels indicates an expected call of ListLabels.
func (mr *MockInterfaceMockRecorder) ListLabels() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListLabels", reflect.TypeOf((*MockInterface)(nil).ListLabels))
}

// LoadTestVariablesSecret mocks base method.
func (m *MockInterface) LoadTestVariablesSecret(arg0 *v3.TestSuite) (*v1.Secret, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoadTestVariablesSecret", arg0)
	ret0, _ := ret[0].(*v1.Secret)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LoadTestVariablesSecret indicates an expected call of LoadTestVariablesSecret.
func (mr *MockInterfaceMockRecorder) LoadTestVariablesSecret(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoadTestVariablesSecret", reflect.TypeOf((*MockInterface)(nil).LoadTestVariablesSecret), arg0)
}

// Update mocks base method.
func (m *MockInterface) Update(arg0 *v3.TestSuite) (*v3.TestSuite, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0)
	ret0, _ := ret[0].(*v3.TestSuite)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockInterfaceMockRecorder) Update(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockInterface)(nil).Update), arg0)
}

// UpdateStatus mocks base method.
func (m *MockInterface) UpdateStatus(testSuite *v3.TestSuite) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateStatus", testSuite)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateStatus indicates an expected call of UpdateStatus.
func (mr *MockInterfaceMockRecorder) UpdateStatus(testSuite interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateStatus", reflect.TypeOf((*MockInterface)(nil).UpdateStatus), testSuite)
}

// UpdateTestsuiteSecrets mocks base method.
func (m *MockInterface) UpdateTestsuiteSecrets(arg0 *v3.TestSuite) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTestsuiteSecrets", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateTestsuiteSecrets indicates an expected call of UpdateTestsuiteSecrets.
func (mr *MockInterfaceMockRecorder) UpdateTestsuiteSecrets(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTestsuiteSecrets", reflect.TypeOf((*MockInterface)(nil).UpdateTestsuiteSecrets), arg0)
}