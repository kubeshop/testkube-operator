// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/kubeshop/testkube-operator/pkg/client/testworkflows/v1 (interfaces: TestWorkflowTemplatesInterface)

// Package v1 is a generated GoMock package.
package v1

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "github.com/kubeshop/testkube-operator/api/testworkflows/v1"
)

// MockTestWorkflowTemplatesInterface is a mock of TestWorkflowTemplatesInterface interface.
type MockTestWorkflowTemplatesInterface struct {
	ctrl     *gomock.Controller
	recorder *MockTestWorkflowTemplatesInterfaceMockRecorder
}

// MockTestWorkflowTemplatesInterfaceMockRecorder is the mock recorder for MockTestWorkflowTemplatesInterface.
type MockTestWorkflowTemplatesInterfaceMockRecorder struct {
	mock *MockTestWorkflowTemplatesInterface
}

// NewMockTestWorkflowTemplatesInterface creates a new mock instance.
func NewMockTestWorkflowTemplatesInterface(ctrl *gomock.Controller) *MockTestWorkflowTemplatesInterface {
	mock := &MockTestWorkflowTemplatesInterface{ctrl: ctrl}
	mock.recorder = &MockTestWorkflowTemplatesInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTestWorkflowTemplatesInterface) EXPECT() *MockTestWorkflowTemplatesInterfaceMockRecorder {
	return m.recorder
}

// Apply mocks base method.
func (m *MockTestWorkflowTemplatesInterface) Apply(arg0 *v1.TestWorkflowTemplate) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Apply", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Apply indicates an expected call of Apply.
func (mr *MockTestWorkflowTemplatesInterfaceMockRecorder) Apply(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Apply", reflect.TypeOf((*MockTestWorkflowTemplatesInterface)(nil).Apply), arg0)
}

// Create mocks base method.
func (m *MockTestWorkflowTemplatesInterface) Create(arg0 *v1.TestWorkflowTemplate) (*v1.TestWorkflowTemplate, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0)
	ret0, _ := ret[0].(*v1.TestWorkflowTemplate)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockTestWorkflowTemplatesInterfaceMockRecorder) Create(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockTestWorkflowTemplatesInterface)(nil).Create), arg0)
}

// Delete mocks base method.
func (m *MockTestWorkflowTemplatesInterface) Delete(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockTestWorkflowTemplatesInterfaceMockRecorder) Delete(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockTestWorkflowTemplatesInterface)(nil).Delete), arg0)
}

// DeleteAll mocks base method.
func (m *MockTestWorkflowTemplatesInterface) DeleteAll() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAll")
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAll indicates an expected call of DeleteAll.
func (mr *MockTestWorkflowTemplatesInterfaceMockRecorder) DeleteAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAll", reflect.TypeOf((*MockTestWorkflowTemplatesInterface)(nil).DeleteAll))
}

// DeleteByLabels mocks base method.
func (m *MockTestWorkflowTemplatesInterface) DeleteByLabels(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteByLabels", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteByLabels indicates an expected call of DeleteByLabels.
func (mr *MockTestWorkflowTemplatesInterfaceMockRecorder) DeleteByLabels(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteByLabels", reflect.TypeOf((*MockTestWorkflowTemplatesInterface)(nil).DeleteByLabels), arg0)
}

// Get mocks base method.
func (m *MockTestWorkflowTemplatesInterface) Get(arg0 string) (*v1.TestWorkflowTemplate, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0)
	ret0, _ := ret[0].(*v1.TestWorkflowTemplate)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockTestWorkflowTemplatesInterfaceMockRecorder) Get(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockTestWorkflowTemplatesInterface)(nil).Get), arg0)
}

// List mocks base method.
func (m *MockTestWorkflowTemplatesInterface) List(arg0 string) (*v1.TestWorkflowTemplateList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0)
	ret0, _ := ret[0].(*v1.TestWorkflowTemplateList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockTestWorkflowTemplatesInterfaceMockRecorder) List(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockTestWorkflowTemplatesInterface)(nil).List), arg0)
}

// ListLabels mocks base method.
func (m *MockTestWorkflowTemplatesInterface) ListLabels() (map[string][]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListLabels")
	ret0, _ := ret[0].(map[string][]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListLabels indicates an expected call of ListLabels.
func (mr *MockTestWorkflowTemplatesInterfaceMockRecorder) ListLabels() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListLabels", reflect.TypeOf((*MockTestWorkflowTemplatesInterface)(nil).ListLabels))
}

// Update mocks base method.
func (m *MockTestWorkflowTemplatesInterface) Update(arg0 *v1.TestWorkflowTemplate) (*v1.TestWorkflowTemplate, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0)
	ret0, _ := ret[0].(*v1.TestWorkflowTemplate)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockTestWorkflowTemplatesInterfaceMockRecorder) Update(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockTestWorkflowTemplatesInterface)(nil).Update), arg0)
}

// UpdateStatus mocks base method.
func (m *MockTestWorkflowTemplatesInterface) UpdateStatus(arg0 *v1.TestWorkflowTemplate) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateStatus", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateStatus indicates an expected call of UpdateStatus.
func (mr *MockTestWorkflowTemplatesInterfaceMockRecorder) UpdateStatus(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateStatus", reflect.TypeOf((*MockTestWorkflowTemplatesInterface)(nil).UpdateStatus), arg0)
}