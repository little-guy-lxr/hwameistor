// Code generated by MockGen. DO NOT EDIT.
// Source: grpc_server.go

// Package csi is a generated GoMock package.
package csi

import (
	reflect "reflect"

	csi "github.com/container-storage-interface/spec/lib/go/csi"
	gomock "github.com/golang/mock/gomock"
)

// MockServer is a mock of Server interface.
type MockServer struct {
	ctrl     *gomock.Controller
	recorder *MockServerMockRecorder
}

// MockServerMockRecorder is the mock recorder for MockServer.
type MockServerMockRecorder struct {
	mock *MockServer
}

// NewMockServer creates a new mock instance.
func NewMockServer(ctrl *gomock.Controller) *MockServer {
	mock := &MockServer{ctrl: ctrl}
	mock.recorder = &MockServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServer) EXPECT() *MockServerMockRecorder {
	return m.recorder
}

// GracefulStop mocks base method.
func (m *MockServer) GracefulStop() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "GracefulStop")
}

// GracefulStop indicates an expected call of GracefulStop.
func (mr *MockServerMockRecorder) GracefulStop() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GracefulStop", reflect.TypeOf((*MockServer)(nil).GracefulStop))
}

// Init mocks base method.
func (m *MockServer) Init(endpoint string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Init", endpoint)
}

// Init indicates an expected call of Init.
func (mr *MockServerMockRecorder) Init(endpoint interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Init", reflect.TypeOf((*MockServer)(nil).Init), endpoint)
}

// Run mocks base method.
func (m *MockServer) Run(ids csi.IdentityServer, cs csi.ControllerServer, ns csi.NodeServer) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Run", ids, cs, ns)
}

// Run indicates an expected call of Run.
func (mr *MockServerMockRecorder) Run(ids, cs, ns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Run", reflect.TypeOf((*MockServer)(nil).Run), ids, cs, ns)
}

// Stop mocks base method.
func (m *MockServer) Stop() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Stop")
}

// Stop indicates an expected call of Stop.
func (mr *MockServerMockRecorder) Stop() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockServer)(nil).Stop))
}
