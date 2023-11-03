// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/erupshis/revtracker/internal/storage (interfaces: BaseStorage)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	data "github.com/erupshis/revtracker/internal/data"
	gomock "github.com/golang/mock/gomock"
)

// MockBaseStorage is a mock of BaseStorage interface.
type MockBaseStorage struct {
	ctrl     *gomock.Controller
	recorder *MockBaseStorageMockRecorder
}

// MockBaseStorageMockRecorder is the mock recorder for MockBaseStorage.
type MockBaseStorageMockRecorder struct {
	mock *MockBaseStorage
}

// NewMockBaseStorage creates a new mock instance.
func NewMockBaseStorage(ctrl *gomock.Controller) *MockBaseStorage {
	mock := &MockBaseStorage{ctrl: ctrl}
	mock.recorder = &MockBaseStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBaseStorage) EXPECT() *MockBaseStorageMockRecorder {
	return m.recorder
}

// DeleteContentByID mocks base method.
func (m *MockBaseStorage) DeleteContentByID(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteContentByID", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteContentByID indicates an expected call of DeleteContentByID.
func (mr *MockBaseStorageMockRecorder) DeleteContentByID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteContentByID", reflect.TypeOf((*MockBaseStorage)(nil).DeleteContentByID), arg0, arg1)
}

// DeleteHomeworkByID mocks base method.
func (m *MockBaseStorage) DeleteHomeworkByID(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteHomeworkByID", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteHomeworkByID indicates an expected call of DeleteHomeworkByID.
func (mr *MockBaseStorageMockRecorder) DeleteHomeworkByID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteHomeworkByID", reflect.TypeOf((*MockBaseStorage)(nil).DeleteHomeworkByID), arg0, arg1)
}

// DeleteHomeworkQuestionByID mocks base method.
func (m *MockBaseStorage) DeleteHomeworkQuestionByID(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteHomeworkQuestionByID", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteHomeworkQuestionByID indicates an expected call of DeleteHomeworkQuestionByID.
func (mr *MockBaseStorageMockRecorder) DeleteHomeworkQuestionByID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteHomeworkQuestionByID", reflect.TypeOf((*MockBaseStorage)(nil).DeleteHomeworkQuestionByID), arg0, arg1)
}

// DeleteQuestionByID mocks base method.
func (m *MockBaseStorage) DeleteQuestionByID(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteQuestionByID", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteQuestionByID indicates an expected call of DeleteQuestionByID.
func (mr *MockBaseStorageMockRecorder) DeleteQuestionByID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteQuestionByID", reflect.TypeOf((*MockBaseStorage)(nil).DeleteQuestionByID), arg0, arg1)
}

// InsertContent mocks base method.
func (m *MockBaseStorage) InsertContent(arg0 context.Context, arg1 *data.Content) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertContent", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// InsertContent indicates an expected call of InsertContent.
func (mr *MockBaseStorageMockRecorder) InsertContent(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertContent", reflect.TypeOf((*MockBaseStorage)(nil).InsertContent), arg0, arg1)
}

// InsertHomework mocks base method.
func (m *MockBaseStorage) InsertHomework(arg0 context.Context, arg1 *data.Homework) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertHomework", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// InsertHomework indicates an expected call of InsertHomework.
func (mr *MockBaseStorageMockRecorder) InsertHomework(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertHomework", reflect.TypeOf((*MockBaseStorage)(nil).InsertHomework), arg0, arg1)
}

// InsertHomeworkQuestion mocks base method.
func (m *MockBaseStorage) InsertHomeworkQuestion(arg0 context.Context, arg1 *data.HomeworkQuestion) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertHomeworkQuestion", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// InsertHomeworkQuestion indicates an expected call of InsertHomeworkQuestion.
func (mr *MockBaseStorageMockRecorder) InsertHomeworkQuestion(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertHomeworkQuestion", reflect.TypeOf((*MockBaseStorage)(nil).InsertHomeworkQuestion), arg0, arg1)
}

// InsertQuestion mocks base method.
func (m *MockBaseStorage) InsertQuestion(arg0 context.Context, arg1 *data.Question) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertQuestion", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// InsertQuestion indicates an expected call of InsertQuestion.
func (mr *MockBaseStorageMockRecorder) InsertQuestion(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertQuestion", reflect.TypeOf((*MockBaseStorage)(nil).InsertQuestion), arg0, arg1)
}

// InsertUser mocks base method.
func (m *MockBaseStorage) InsertUser(arg0 context.Context, arg1 *data.User) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertUser", arg0, arg1)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertUser indicates an expected call of InsertUser.
func (mr *MockBaseStorageMockRecorder) InsertUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertUser", reflect.TypeOf((*MockBaseStorage)(nil).InsertUser), arg0, arg1)
}

// SelectContentByID mocks base method.
func (m *MockBaseStorage) SelectContentByID(arg0 context.Context, arg1 int64) (*data.Content, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectContentByID", arg0, arg1)
	ret0, _ := ret[0].(*data.Content)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectContentByID indicates an expected call of SelectContentByID.
func (mr *MockBaseStorageMockRecorder) SelectContentByID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectContentByID", reflect.TypeOf((*MockBaseStorage)(nil).SelectContentByID), arg0, arg1)
}

// SelectHomeworkByID mocks base method.
func (m *MockBaseStorage) SelectHomeworkByID(arg0 context.Context, arg1 int64) (*data.Homework, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectHomeworkByID", arg0, arg1)
	ret0, _ := ret[0].(*data.Homework)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectHomeworkByID indicates an expected call of SelectHomeworkByID.
func (mr *MockBaseStorageMockRecorder) SelectHomeworkByID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectHomeworkByID", reflect.TypeOf((*MockBaseStorage)(nil).SelectHomeworkByID), arg0, arg1)
}

// SelectHomeworkQuestionByID mocks base method.
func (m *MockBaseStorage) SelectHomeworkQuestionByID(arg0 context.Context, arg1 int64) (*data.HomeworkQuestion, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectHomeworkQuestionByID", arg0, arg1)
	ret0, _ := ret[0].(*data.HomeworkQuestion)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectHomeworkQuestionByID indicates an expected call of SelectHomeworkQuestionByID.
func (mr *MockBaseStorageMockRecorder) SelectHomeworkQuestionByID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectHomeworkQuestionByID", reflect.TypeOf((*MockBaseStorage)(nil).SelectHomeworkQuestionByID), arg0, arg1)
}

// SelectQuestionByID mocks base method.
func (m *MockBaseStorage) SelectQuestionByID(arg0 context.Context, arg1 int64) (*data.Question, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectQuestionByID", arg0, arg1)
	ret0, _ := ret[0].(*data.Question)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectQuestionByID indicates an expected call of SelectQuestionByID.
func (mr *MockBaseStorageMockRecorder) SelectQuestionByID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectQuestionByID", reflect.TypeOf((*MockBaseStorage)(nil).SelectQuestionByID), arg0, arg1)
}

// SelectUser mocks base method.
func (m *MockBaseStorage) SelectUser(arg0 context.Context, arg1 map[string]interface{}) (*data.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectUser", arg0, arg1)
	ret0, _ := ret[0].(*data.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectUser indicates an expected call of SelectUser.
func (mr *MockBaseStorageMockRecorder) SelectUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectUser", reflect.TypeOf((*MockBaseStorage)(nil).SelectUser), arg0, arg1)
}

// UpdateContentByID mocks base method.
func (m *MockBaseStorage) UpdateContentByID(arg0 context.Context, arg1 int64, arg2 map[string]interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateContentByID", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateContentByID indicates an expected call of UpdateContentByID.
func (mr *MockBaseStorageMockRecorder) UpdateContentByID(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateContentByID", reflect.TypeOf((*MockBaseStorage)(nil).UpdateContentByID), arg0, arg1, arg2)
}

// UpdateHomeworkByID mocks base method.
func (m *MockBaseStorage) UpdateHomeworkByID(arg0 context.Context, arg1 int64, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateHomeworkByID", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateHomeworkByID indicates an expected call of UpdateHomeworkByID.
func (mr *MockBaseStorageMockRecorder) UpdateHomeworkByID(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateHomeworkByID", reflect.TypeOf((*MockBaseStorage)(nil).UpdateHomeworkByID), arg0, arg1, arg2)
}

// UpdateHomeworkQuestionByID mocks base method.
func (m *MockBaseStorage) UpdateHomeworkQuestionByID(arg0 context.Context, arg1 int64, arg2 map[string]interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateHomeworkQuestionByID", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateHomeworkQuestionByID indicates an expected call of UpdateHomeworkQuestionByID.
func (mr *MockBaseStorageMockRecorder) UpdateHomeworkQuestionByID(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateHomeworkQuestionByID", reflect.TypeOf((*MockBaseStorage)(nil).UpdateHomeworkQuestionByID), arg0, arg1, arg2)
}

// UpdateQuestionByID mocks base method.
func (m *MockBaseStorage) UpdateQuestionByID(arg0 context.Context, arg1 int64, arg2 map[string]interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateQuestionByID", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateQuestionByID indicates an expected call of UpdateQuestionByID.
func (mr *MockBaseStorageMockRecorder) UpdateQuestionByID(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateQuestionByID", reflect.TypeOf((*MockBaseStorage)(nil).UpdateQuestionByID), arg0, arg1, arg2)
}
