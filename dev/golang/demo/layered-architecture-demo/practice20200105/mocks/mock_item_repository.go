// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/tabakazu/practice20200105/domain/repository (interfaces: ItemRepository)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	model "github.com/tabakazu/practice20200105/domain/model"
	reflect "reflect"
)

// MockItemRepository is a mock of ItemRepository interface
type MockItemRepository struct {
	ctrl     *gomock.Controller
	recorder *MockItemRepositoryMockRecorder
}

// MockItemRepositoryMockRecorder is the mock recorder for MockItemRepository
type MockItemRepositoryMockRecorder struct {
	mock *MockItemRepository
}

// NewMockItemRepository creates a new mock instance
func NewMockItemRepository(ctrl *gomock.Controller) *MockItemRepository {
	mock := &MockItemRepository{ctrl: ctrl}
	mock.recorder = &MockItemRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockItemRepository) EXPECT() *MockItemRepositoryMockRecorder {
	return m.recorder
}

// FindByID mocks base method
func (m *MockItemRepository) FindByID(arg0 uint) (*model.Item, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByID", arg0)
	ret0, _ := ret[0].(*model.Item)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByID indicates an expected call of FindByID
func (mr *MockItemRepositoryMockRecorder) FindByID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByID", reflect.TypeOf((*MockItemRepository)(nil).FindByID), arg0)
}
