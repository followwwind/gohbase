// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/tsuna/gohbase (interfaces: RPCClient)

// Package mock is a generated GoMock package.
package mock

import (
	hrpc "github.com/followwwind/gohbase/hrpc"
	gomock "github.com/golang/mock/gomock"
	proto "github.com/golang/protobuf/proto"
	reflect "reflect"
)

// MockRPCClient is a mock of RPCClient interface
type MockRPCClient struct {
	ctrl     *gomock.Controller
	recorder *MockRPCClientMockRecorder
}

// MockRPCClientMockRecorder is the mock recorder for MockRPCClient
type MockRPCClientMockRecorder struct {
	mock *MockRPCClient
}

// NewMockRPCClient creates a new mock instance
func NewMockRPCClient(ctrl *gomock.Controller) *MockRPCClient {
	mock := &MockRPCClient{ctrl: ctrl}
	mock.recorder = &MockRPCClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRPCClient) EXPECT() *MockRPCClientMockRecorder {
	return m.recorder
}

// SendRPC mocks base method
func (m *MockRPCClient) SendRPC(arg0 hrpc.Call) (proto.Message, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendRPC", arg0)
	ret0, _ := ret[0].(proto.Message)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SendRPC indicates an expected call of SendRPC
func (mr *MockRPCClientMockRecorder) SendRPC(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendRPC", reflect.TypeOf((*MockRPCClient)(nil).SendRPC), arg0)
}
