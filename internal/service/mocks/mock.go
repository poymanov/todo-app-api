// Code generated by MockGen. DO NOT EDIT.
// Source: internal/service/service.go
//
// Generated by this command:
//
//	mockgen -source=internal/service/service.go -destination=internal/service/mocks/mock.go
//

// Package mock_service is a generated GoMock package.
package mock_service

import (
	domain "poymanov/todo/internal/domain"
	service "poymanov/todo/internal/service"
	reflect "reflect"

	uuid "github.com/google/uuid"
	gomock "go.uber.org/mock/gomock"
)

// MockAuth is a mock of Auth interface.
type MockAuth struct {
	ctrl     *gomock.Controller
	recorder *MockAuthMockRecorder
	isgomock struct{}
}

// MockAuthMockRecorder is the mock recorder for MockAuth.
type MockAuthMockRecorder struct {
	mock *MockAuth
}

// NewMockAuth creates a new mock instance.
func NewMockAuth(ctrl *gomock.Controller) *MockAuth {
	mock := &MockAuth{ctrl: ctrl}
	mock.recorder = &MockAuthMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuth) EXPECT() *MockAuthMockRecorder {
	return m.recorder
}

// Login mocks base method.
func (m *MockAuth) Login(data service.LoginData) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Login", data)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Login indicates an expected call of Login.
func (mr *MockAuthMockRecorder) Login(data any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockAuth)(nil).Login), data)
}

// Register mocks base method.
func (m *MockAuth) Register(data service.RegisterData) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Register", data)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Register indicates an expected call of Register.
func (mr *MockAuthMockRecorder) Register(data any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockAuth)(nil).Register), data)
}

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
func (m *MockTask) Create(description string, userId uuid.UUID) (*domain.Task, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", description, userId)
	ret0, _ := ret[0].(*domain.Task)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockTaskMockRecorder) Create(description, userId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockTask)(nil).Create), description, userId)
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

// UpdateDescription mocks base method.
func (m *MockTask) UpdateDescription(id uuid.UUID, description string) (*domain.Task, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateDescription", id, description)
	ret0, _ := ret[0].(*domain.Task)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateDescription indicates an expected call of UpdateDescription.
func (mr *MockTaskMockRecorder) UpdateDescription(id, description any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateDescription", reflect.TypeOf((*MockTask)(nil).UpdateDescription), id, description)
}

// UpdateIsCompleted mocks base method.
func (m *MockTask) UpdateIsCompleted(id uuid.UUID, isCompleted bool) (*domain.Task, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateIsCompleted", id, isCompleted)
	ret0, _ := ret[0].(*domain.Task)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateIsCompleted indicates an expected call of UpdateIsCompleted.
func (mr *MockTaskMockRecorder) UpdateIsCompleted(id, isCompleted any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateIsCompleted", reflect.TypeOf((*MockTask)(nil).UpdateIsCompleted), id, isCompleted)
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
func (m *MockUser) Create(name, email, password string) (*domain.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", name, email, password)
	ret0, _ := ret[0].(*domain.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockUserMockRecorder) Create(name, email, password any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockUser)(nil).Create), name, email, password)
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
