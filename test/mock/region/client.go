// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/tsuna/gohbase/hrpc (interfaces: RegionClient)

// Package mock is a generated GoMock package.
package mock

import (
	hrpc "github.com/followwwind/gohbase/hrpc"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockRegionClient is a mock of RegionClient interface
type MockRegionClient struct {
	ctrl     *gomock.Controller
	recorder *MockRegionClientMockRecorder
}

// MockRegionClientMockRecorder is the mock recorder for MockRegionClient
type MockRegionClientMockRecorder struct {
	mock *MockRegionClient
}

// NewMockRegionClient creates a new mock instance
func NewMockRegionClient(ctrl *gomock.Controller) *MockRegionClient {
	mock := &MockRegionClient{ctrl: ctrl}
	mock.recorder = &MockRegionClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRegionClient) EXPECT() *MockRegionClientMockRecorder {
	return m.recorder
}

// Addr mocks base method
func (m *MockRegionClient) Addr() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Addr")
	ret0, _ := ret[0].(string)
	return ret0
}

// Addr indicates an expected call of Addr
func (mr *MockRegionClientMockRecorder) Addr() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Addr", reflect.TypeOf((*MockRegionClient)(nil).Addr))
}

// Close mocks base method
func (m *MockRegionClient) Close() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close
func (mr *MockRegionClientMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockRegionClient)(nil).Close))
}

// QueueRPC mocks base method
func (m *MockRegionClient) QueueRPC(arg0 hrpc.Call) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "QueueRPC", arg0)
}

// QueueRPC indicates an expected call of QueueRPC
func (mr *MockRegionClientMockRecorder) QueueRPC(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueueRPC", reflect.TypeOf((*MockRegionClient)(nil).QueueRPC), arg0)
}

// String mocks base method
func (m *MockRegionClient) String() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "String")
	ret0, _ := ret[0].(string)
	return ret0
}

// String indicates an expected call of String
func (mr *MockRegionClientMockRecorder) String() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "String", reflect.TypeOf((*MockRegionClient)(nil).String))
}
