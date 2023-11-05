// Code generated by MockGen. DO NOT EDIT.
// Source: webook/internal/repository/dao/user.go
//
// Generated by this command:
//
//	mockgen -source=webook/internal/repository/dao/user.go -destination=webook/internal/repository/dao/mock/user_mock.go -package=daomock
//
// Package daomock is a generated GoMock package.
package daomock

import (
	context "context"
	reflect "reflect"

	dao "github.com/TangTangHC/basic-go-study/webook/internal/repository/dao"
	gomock "go.uber.org/mock/gomock"
)

// MockUserDao is a mock of UserDao interface.
type MockUserDao struct {
	ctrl     *gomock.Controller
	recorder *MockUserDaoMockRecorder
}

// MockUserDaoMockRecorder is the mock recorder for MockUserDao.
type MockUserDaoMockRecorder struct {
	mock *MockUserDao
}

// NewMockUserDao creates a new mock instance.
func NewMockUserDao(ctrl *gomock.Controller) *MockUserDao {
	mock := &MockUserDao{ctrl: ctrl}
	mock.recorder = &MockUserDaoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserDao) EXPECT() *MockUserDaoMockRecorder {
	return m.recorder
}

// FindByEmail mocks base method.
func (m *MockUserDao) FindByEmail(arg0 context.Context, arg1 string) (dao.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByEmail", arg0, arg1)
	ret0, _ := ret[0].(dao.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByEmail indicates an expected call of FindByEmail.
func (mr *MockUserDaoMockRecorder) FindByEmail(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByEmail", reflect.TypeOf((*MockUserDao)(nil).FindByEmail), arg0, arg1)
}

// FindById mocks base method.
func (m *MockUserDao) FindById(arg0 context.Context, arg1 int64) (dao.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindById", arg0, arg1)
	ret0, _ := ret[0].(dao.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindById indicates an expected call of FindById.
func (mr *MockUserDaoMockRecorder) FindById(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindById", reflect.TypeOf((*MockUserDao)(nil).FindById), arg0, arg1)
}

// FindByPhone mocks base method.
func (m *MockUserDao) FindByPhone(arg0 context.Context, arg1 string) (dao.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByPhone", arg0, arg1)
	ret0, _ := ret[0].(dao.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByPhone indicates an expected call of FindByPhone.
func (mr *MockUserDaoMockRecorder) FindByPhone(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByPhone", reflect.TypeOf((*MockUserDao)(nil).FindByPhone), arg0, arg1)
}

// Insert mocks base method.
func (m *MockUserDao) Insert(ctx context.Context, user dao.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", ctx, user)
	ret0, _ := ret[0].(error)
	return ret0
}

// Insert indicates an expected call of Insert.
func (mr *MockUserDaoMockRecorder) Insert(ctx, user any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockUserDao)(nil).Insert), ctx, user)
}

// UpdateById mocks base method.
func (m *MockUserDao) UpdateById(arg0 context.Context, arg1 dao.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateById", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateById indicates an expected call of UpdateById.
func (mr *MockUserDaoMockRecorder) UpdateById(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateById", reflect.TypeOf((*MockUserDao)(nil).UpdateById), arg0, arg1)
}
