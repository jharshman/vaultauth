// Code generated by MockGen. DO NOT EDIT.
// Source: client.go

// Package mock_main is a generated GoMock package.
package mock

import (
	gomock "github.com/golang/mock/gomock"
	api "github.com/hashicorp/vault/api"
	reflect "reflect"
)

// MockLogical is a mock of Logical interface
type MockLogical struct {
	ctrl     *gomock.Controller
	recorder *MockLogicalMockRecorder
}

// MockLogicalMockRecorder is the mock recorder for MockLogical
type MockLogicalMockRecorder struct {
	mock *MockLogical
}

// NewMockLogical creates a new mock instance
func NewMockLogical(ctrl *gomock.Controller) *MockLogical {
	mock := &MockLogical{ctrl: ctrl}
	mock.recorder = &MockLogicalMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockLogical) EXPECT() *MockLogicalMockRecorder {
	return m.recorder
}

// Read mocks base method
func (m *MockLogical) Read(path string) (*api.Secret, error) {
	ret := m.ctrl.Call(m, "Read", path)
	ret0, _ := ret[0].(*api.Secret)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Read indicates an expected call of Read
func (mr *MockLogicalMockRecorder) Read(path interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockLogical)(nil).Read), path)
}

// Write mocks base method
func (m *MockLogical) Write(path string, data map[string]interface{}) (*api.Secret, error) {
	ret := m.ctrl.Call(m, "Write", path, data)
	ret0, _ := ret[0].(*api.Secret)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Write indicates an expected call of Write
func (mr *MockLogicalMockRecorder) Write(path, data interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Write", reflect.TypeOf((*MockLogical)(nil).Write), path, data)
}