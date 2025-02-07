// Code generated by MockGen. DO NOT EDIT.
// Source: internal/repository/repository.go
//
// Generated by this command:
//
//	mockgen -source=internal/repository/repository.go -destination=internal/repository/mocks/mock.go
//

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	domain "poymanov/todo/internal/domain"
	reflect "reflect"

	uuid "github.com/google/uuid"
	gomock "go.uber.org/mock/gomock"
)

// MockTask is a mock of Task interface.
type MockTask struct {
	ctrl     *gomock.Controller
	recorder *MockTaskMockRecorder
	isgomock struct{}
}

// MockTaskMockRecorder is the mock recorder for MockTask.
type MockTaskMockRecorder struct {
	mock *MockTask
}

// NewMockTask creates a new mock instance.
func NewMockTask(ctrl *gomock.Controller) *MockTask {
	mock := &MockTask{ctrl: ctrl}
	mock.recorder = &MockTaskMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTask) EXPECT() *MockTaskMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockTask) Create(task *domain.Task) (*domain.Task, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", task)
	ret0, _ := ret[0].(*domain.Task)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockTaskMockRecorder) Create(task any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockTask)(nil).Create), task)
}

// Delete mocks base method.
func (m *MockTask) Delete(id uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockTaskMockRecorder) Delete(id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockTask)(nil).Delete), id)
}

// GetAllByUserId mocks base method.
func (m *MockTask) GetAllByUserId(id uuid.UUID) *[]domain.Task {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllByUserId", id)
	ret0, _ := ret[0].(*[]domain.Task)
	return ret0
}

// GetAllByUserId indicates an expected call of GetAllByUserId.
func (mr *MockTaskMockRecorder) GetAllByUserId(id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllByUserId", reflect.TypeOf((*MockTask)(nil).GetAllByUserId), id)
}

// IsExistsById mocks base method.
func (m *MockTask) IsExistsById(id uuid.UUID) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsExistsById", id)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsExistsById indicates an expected call of IsExistsById.
func (mr *MockTaskMockRecorder) IsExistsById(id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsExistsById", reflect.TypeOf((*MockTask)(nil).IsExistsById), id)
}

// Update mocks base method.
func (m *MockTask) Update(task *domain.Task) (*domain.Task, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", task)
	ret0, _ := ret[0].(*domain.Task)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockTaskMockRecorder) Update(task any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockTask)(nil).Update), task)
}

// MockUser is a mock of User interface.
type MockUser struct {
	ctrl     *gomock.Controller
	recorder *MockUserMockRecorder
	isgomock struct{}
}

// MockUserMockRecorder is the mock recorder for MockUser.
type MockUserMockRecorder struct {
	mock *MockUser
}

// NewMockUser creates a new mock instance.
func NewMockUser(ctrl *gomock.Controller) *MockUser {
	mock := &MockUser{ctrl: ctrl}
	mock.recorder = &MockUserMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUser) EXPECT() *MockUserMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockUser) Create(user *domain.User) (*domain.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", user)
	ret0, _ := ret[0].(*domain.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockUserMockRecorder) Create(user any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockUser)(nil).Create), user)
}

// FindByEmail mocks base method.
func (m *MockUser) FindByEmail(email string) (*domain.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByEmail", email)
	ret0, _ := ret[0].(*domain.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByEmail indicates an expected call of FindByEmail.
func (mr *MockUserMockRecorder) FindByEmail(email any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByEmail", reflect.TypeOf((*MockUser)(nil).FindByEmail), email)
}
