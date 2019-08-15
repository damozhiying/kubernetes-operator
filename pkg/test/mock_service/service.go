// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/gosoon/kubernetes-operator/pkg/server/service (interfaces: Interface)

// Package mock_service is a generated GoMock package.
package mock_service

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	types "github.com/gosoon/kubernetes-operator/pkg/types"
)

// MockInterface is a mock of Interface interface
type MockInterface struct {
	ctrl     *gomock.Controller
	recorder *MockInterfaceMockRecorder
}

// MockInterfaceMockRecorder is the mock recorder for MockInterface
type MockInterfaceMockRecorder struct {
	mock *MockInterface
}

// NewMockInterface creates a new mock instance
func NewMockInterface(ctrl *gomock.Controller) *MockInterface {
	mock := &MockInterface{ctrl: ctrl}
	mock.recorder = &MockInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockInterface) EXPECT() *MockInterfaceMockRecorder {
	return m.recorder
}

// CreateCluster mocks base method
func (m *MockInterface) CreateCluster(arg0, arg1, arg2 string, arg3 *types.EcsClient) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCluster", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateCluster indicates an expected call of CreateCluster
func (mr *MockInterfaceMockRecorder) CreateCluster(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCluster", reflect.TypeOf((*MockInterface)(nil).CreateCluster), arg0, arg1, arg2, arg3)
}

// CreateClusterCallback mocks base method
func (m *MockInterface) CreateClusterCallback(arg0, arg1, arg2 string, arg3 *types.Callback) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateClusterCallback", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateClusterCallback indicates an expected call of CreateClusterCallback
func (mr *MockInterfaceMockRecorder) CreateClusterCallback(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateClusterCallback", reflect.TypeOf((*MockInterface)(nil).CreateClusterCallback), arg0, arg1, arg2, arg3)
}

// DeleteCluster mocks base method
func (m *MockInterface) DeleteCluster(arg0, arg1, arg2 string, arg3 *types.EcsClient) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCluster", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCluster indicates an expected call of DeleteCluster
func (mr *MockInterfaceMockRecorder) DeleteCluster(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCluster", reflect.TypeOf((*MockInterface)(nil).DeleteCluster), arg0, arg1, arg2, arg3)
}

// DeleteClusterCallback mocks base method
func (m *MockInterface) DeleteClusterCallback(arg0, arg1, arg2 string, arg3 *types.Callback) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteClusterCallback", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteClusterCallback indicates an expected call of DeleteClusterCallback
func (mr *MockInterfaceMockRecorder) DeleteClusterCallback(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteClusterCallback", reflect.TypeOf((*MockInterface)(nil).DeleteClusterCallback), arg0, arg1, arg2, arg3)
}

// GetClusterOperationLogs mocks base method
func (m *MockInterface) GetClusterOperationLogs(arg0, arg1, arg2 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetClusterOperationLogs", arg0, arg1, arg2)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetClusterOperationLogs indicates an expected call of GetClusterOperationLogs
func (mr *MockInterfaceMockRecorder) GetClusterOperationLogs(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetClusterOperationLogs", reflect.TypeOf((*MockInterface)(nil).GetClusterOperationLogs), arg0, arg1, arg2)
}

// ScaleDown mocks base method
func (m *MockInterface) ScaleDown(arg0, arg1, arg2 string, arg3 *types.EcsClient) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ScaleDown", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// ScaleDown indicates an expected call of ScaleDown
func (mr *MockInterfaceMockRecorder) ScaleDown(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ScaleDown", reflect.TypeOf((*MockInterface)(nil).ScaleDown), arg0, arg1, arg2, arg3)
}

// ScaleDownCallback mocks base method
func (m *MockInterface) ScaleDownCallback(arg0, arg1, arg2 string, arg3 *types.Callback) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ScaleDownCallback", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// ScaleDownCallback indicates an expected call of ScaleDownCallback
func (mr *MockInterfaceMockRecorder) ScaleDownCallback(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ScaleDownCallback", reflect.TypeOf((*MockInterface)(nil).ScaleDownCallback), arg0, arg1, arg2, arg3)
}

// ScaleUp mocks base method
func (m *MockInterface) ScaleUp(arg0, arg1, arg2 string, arg3 *types.EcsClient) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ScaleUp", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// ScaleUp indicates an expected call of ScaleUp
func (mr *MockInterfaceMockRecorder) ScaleUp(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ScaleUp", reflect.TypeOf((*MockInterface)(nil).ScaleUp), arg0, arg1, arg2, arg3)
}

// ScaleUpCallback mocks base method
func (m *MockInterface) ScaleUpCallback(arg0, arg1, arg2 string, arg3 *types.Callback) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ScaleUpCallback", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// ScaleUpCallback indicates an expected call of ScaleUpCallback
func (mr *MockInterfaceMockRecorder) ScaleUpCallback(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ScaleUpCallback", reflect.TypeOf((*MockInterface)(nil).ScaleUpCallback), arg0, arg1, arg2, arg3)
}
