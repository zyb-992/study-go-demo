// Code generated by MockGen. DO NOT EDIT.
// Source: gomock-learn/person (interfaces: Person)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockPerson is a mock of Person interface.
type MockPerson struct {
	ctrl     *gomock.Controller
	recorder *MockPersonMockRecorder
}

// MockPersonMockRecorder is the mock recorder for MockPerson.
type MockPersonMockRecorder struct {
	mock *MockPerson
}

// NewMockPerson creates a new mock instance.
func NewMockPerson(ctrl *gomock.Controller) *MockPerson {
	mock := &MockPerson{ctrl: ctrl}
	mock.recorder = &MockPersonMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPerson) EXPECT() *MockPersonMockRecorder {
	return m.recorder
}

// Eat mocks base method.
func (m *MockPerson) Eat(arg0 string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Eat", arg0)
	ret0, _ := ret[0].(string)
	return ret0
}

// Eat indicates an expected call of Eat.
func (mr *MockPersonMockRecorder) Eat(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Eat", reflect.TypeOf((*MockPerson)(nil).Eat), arg0)
}

// Sleep mocks base method.
func (m *MockPerson) Sleep(arg0 string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Sleep", arg0)
	ret0, _ := ret[0].(string)
	return ret0
}

// Sleep indicates an expected call of Sleep.
func (mr *MockPersonMockRecorder) Sleep(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Sleep", reflect.TypeOf((*MockPerson)(nil).Sleep), arg0)
}
