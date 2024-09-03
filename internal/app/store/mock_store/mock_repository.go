// Code generated by MockGen. DO NOT EDIT.
// Source: repository.go

// Package mock_store is a generated GoMock package.
package mock_store

import (
	models "filmoteka/internal/app/models"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockIFilmRepository is a mock of IFilmRepository interface.
type MockIFilmRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIFilmRepositoryMockRecorder
}

// MockIFilmRepositoryMockRecorder is the mock recorder for MockIFilmRepository.
type MockIFilmRepositoryMockRecorder struct {
	mock *MockIFilmRepository
}

// NewMockIFilmRepository creates a new mock instance.
func NewMockIFilmRepository(ctrl *gomock.Controller) *MockIFilmRepository {
	mock := &MockIFilmRepository{ctrl: ctrl}
	mock.recorder = &MockIFilmRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIFilmRepository) EXPECT() *MockIFilmRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockIFilmRepository) Create(arg0 models.Film) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockIFilmRepositoryMockRecorder) Create(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockIFilmRepository)(nil).Create), arg0)
}

// Delete mocks base method.
func (m *MockIFilmRepository) Delete(id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", id)
	ret1, _ := ret[0].(error)
	return ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockIFilmRepositoryMockRecorder) Delete(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockIFilmRepository)(nil).Delete), id)
}

// Find mocks base method.
func (m *MockIFilmRepository) Find(arg0 int) (models.Film, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", arg0)
	ret0, _ := ret[0].(models.Film)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find.
func (mr *MockIFilmRepositoryMockRecorder) Find(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockIFilmRepository)(nil).Find), arg0)
}

// FindAll mocks base method.
func (m *MockIFilmRepository) FindAll() ([]models.Film, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAll")
	ret0, _ := ret[0].([]models.Film)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAll indicates an expected call of FindAll.
func (mr *MockIFilmRepositoryMockRecorder) FindAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAll", reflect.TypeOf((*MockIFilmRepository)(nil).FindAll))
}

// Update mocks base method.
func (m *MockIFilmRepository) Update(arg0 models.Film) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0)
	ret1, _ := ret[0].(error)
	return ret1
}

// Update indicates an expected call of Update.
func (mr *MockIFilmRepositoryMockRecorder) Update(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockIFilmRepository)(nil).Update), arg0)
}

// MockIActorRepository is a mock of IActorRepository interface.
type MockIActorRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIActorRepositoryMockRecorder
}

// MockIActorRepositoryMockRecorder is the mock recorder for MockIActorRepository.
type MockIActorRepositoryMockRecorder struct {
	mock *MockIActorRepository
}

// NewMockIActorRepository creates a new mock instance.
func NewMockIActorRepository(ctrl *gomock.Controller) *MockIActorRepository {
	mock := &MockIActorRepository{ctrl: ctrl}
	mock.recorder = &MockIActorRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIActorRepository) EXPECT() *MockIActorRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockIActorRepository) Create(arg0 models.Actor) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockIActorRepositoryMockRecorder) Create(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockIActorRepository)(nil).Create), arg0)
}

// Delete mocks base method.
func (m *MockIActorRepository) Delete(id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", id)
	ret1, _ := ret[0].(error)
	return ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockIActorRepositoryMockRecorder) Delete(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockIActorRepository)(nil).Delete), id)
}

// Find mocks base method.
func (m *MockIActorRepository) Find(arg0 int) (models.Actor, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", arg0)
	ret0, _ := ret[0].(models.Actor)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find.
func (mr *MockIActorRepositoryMockRecorder) Find(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockIActorRepository)(nil).Find), arg0)
}

// FindAll mocks base method.
func (m *MockIActorRepository) FindAll() ([]models.Actor, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAll")
	ret0, _ := ret[0].([]models.Actor)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAll indicates an expected call of FindAll.
func (mr *MockIActorRepositoryMockRecorder) FindAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAll", reflect.TypeOf((*MockIActorRepository)(nil).FindAll))
}

// Update mocks base method.
func (m *MockIActorRepository) Update(arg0 models.Actor) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0)
	ret1, _ := ret[0].(error)
	return ret1
}

// Update indicates an expected call of Update.
func (mr *MockIActorRepositoryMockRecorder) Update(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockIActorRepository)(nil).Update), arg0)
}
