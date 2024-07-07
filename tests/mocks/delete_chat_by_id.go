// Code generated by MockGen. DO NOT EDIT.
// Source: internal/domain/usecase/delete_chat_by_id.go

// Package mock_usecase is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockDeleteChatById is a mock of DeleteChatById interface.
type MockDeleteChatById struct {
	ctrl     *gomock.Controller
	recorder *MockDeleteChatByIdMockRecorder
}

// MockDeleteChatByIdMockRecorder is the mock recorder for MockDeleteChatById.
type MockDeleteChatByIdMockRecorder struct {
	mock *MockDeleteChatById
}

// NewMockDeleteChatById creates a new mock instance.
func NewMockDeleteChatById(ctrl *gomock.Controller) *MockDeleteChatById {
	mock := &MockDeleteChatById{ctrl: ctrl}
	mock.recorder = &MockDeleteChatByIdMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDeleteChatById) EXPECT() *MockDeleteChatByIdMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockDeleteChatById) Delete(chatId string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", chatId)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockDeleteChatByIdMockRecorder) Delete(chatId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockDeleteChatById)(nil).Delete), chatId)
}
