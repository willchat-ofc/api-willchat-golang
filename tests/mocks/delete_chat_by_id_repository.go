// Code generated by MockGen. DO NOT EDIT.
// Source: internal/data/protocols/delete_chat_by_id_repository.go

// Package mock_protocols is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockDeleteChatByIdRepository is a mock of DeleteChatByIdRepository interface.
type MockDeleteChatByIdRepository struct {
	ctrl     *gomock.Controller
	recorder *MockDeleteChatByIdRepositoryMockRecorder
}

// MockDeleteChatByIdRepositoryMockRecorder is the mock recorder for MockDeleteChatByIdRepository.
type MockDeleteChatByIdRepositoryMockRecorder struct {
	mock *MockDeleteChatByIdRepository
}

// NewMockDeleteChatByIdRepository creates a new mock instance.
func NewMockDeleteChatByIdRepository(ctrl *gomock.Controller) *MockDeleteChatByIdRepository {
	mock := &MockDeleteChatByIdRepository{ctrl: ctrl}
	mock.recorder = &MockDeleteChatByIdRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDeleteChatByIdRepository) EXPECT() *MockDeleteChatByIdRepositoryMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockDeleteChatByIdRepository) Delete(chatId string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", chatId)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockDeleteChatByIdRepositoryMockRecorder) Delete(chatId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockDeleteChatByIdRepository)(nil).Delete), chatId)
}
