// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/Dyleme/image-coverter/pkg/service (interfaces: Requester)

// Package mock_service is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"
	time "time"

	model "github.com/Dyleme/image-coverter/internal/model"
	gomock "github.com/golang/mock/gomock"
)

// MockRequester is a mock of Requester interface.
type MockRequester struct {
	ctrl     *gomock.Controller
	recorder *MockRequesterMockRecorder
}

// MockRequesterMockRecorder is the mock recorder for MockRequester.
type MockRequesterMockRecorder struct {
	mock *MockRequester
}

// NewMockRequester creates a new mock instance.
func NewMockRequester(ctrl *gomock.Controller) *MockRequester {
	mock := &MockRequester{ctrl: ctrl}
	mock.recorder = &MockRequesterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRequester) EXPECT() *MockRequesterMockRecorder {
	return m.recorder
}

// AddImage mocks base method.
func (m *MockRequester) AddImage(arg0 context.Context, arg1 int, arg2 model.Info) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddImage", arg0, arg1, arg2)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddImage indicates an expected call of AddImage.
func (mr *MockRequesterMockRecorder) AddImage(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddImage", reflect.TypeOf((*MockRequester)(nil).AddImage), arg0, arg1, arg2)
}

// AddProcessedImageIDToRequest mocks base method.
func (m *MockRequester) AddProcessedImageIDToRequest(arg0 context.Context, arg1, arg2 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddProcessedImageIDToRequest", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddProcessedImageIDToRequest indicates an expected call of AddProcessedImageIDToRequest.
func (mr *MockRequesterMockRecorder) AddProcessedImageIDToRequest(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddProcessedImageIDToRequest", reflect.TypeOf((*MockRequester)(nil).AddProcessedImageIDToRequest), arg0, arg1, arg2)
}

// AddProcessedTimeToRequest mocks base method.
func (m *MockRequester) AddProcessedTimeToRequest(arg0 context.Context, arg1 int, arg2 time.Time) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddProcessedTimeToRequest", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddProcessedTimeToRequest indicates an expected call of AddProcessedTimeToRequest.
func (mr *MockRequesterMockRecorder) AddProcessedTimeToRequest(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddProcessedTimeToRequest", reflect.TypeOf((*MockRequester)(nil).AddProcessedTimeToRequest), arg0, arg1, arg2)
}

// AddRequest mocks base method.
func (m *MockRequester) AddRequest(arg0 context.Context, arg1 *model.Request, arg2 int) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddRequest", arg0, arg1, arg2)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddRequest indicates an expected call of AddRequest.
func (mr *MockRequesterMockRecorder) AddRequest(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddRequest", reflect.TypeOf((*MockRequester)(nil).AddRequest), arg0, arg1, arg2)
}

// DeleteImage mocks base method.
func (m *MockRequester) DeleteImage(arg0 context.Context, arg1, arg2 int) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteImage", arg0, arg1, arg2)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteImage indicates an expected call of DeleteImage.
func (mr *MockRequesterMockRecorder) DeleteImage(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteImage", reflect.TypeOf((*MockRequester)(nil).DeleteImage), arg0, arg1, arg2)
}

// DeleteRequest mocks base method.
func (m *MockRequester) DeleteRequest(arg0 context.Context, arg1, arg2 int) (int, int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteRequest", arg0, arg1, arg2)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// DeleteRequest indicates an expected call of DeleteRequest.
func (mr *MockRequesterMockRecorder) DeleteRequest(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteRequest", reflect.TypeOf((*MockRequester)(nil).DeleteRequest), arg0, arg1, arg2)
}

// GetRequest mocks base method.
func (m *MockRequester) GetRequest(arg0 context.Context, arg1, arg2 int) (*model.Request, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRequest", arg0, arg1, arg2)
	ret0, _ := ret[0].(*model.Request)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRequest indicates an expected call of GetRequest.
func (mr *MockRequesterMockRecorder) GetRequest(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRequest", reflect.TypeOf((*MockRequester)(nil).GetRequest), arg0, arg1, arg2)
}

// GetRequests mocks base method.
func (m *MockRequester) GetRequests(arg0 context.Context, arg1 int) ([]model.Request, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRequests", arg0, arg1)
	ret0, _ := ret[0].([]model.Request)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRequests indicates an expected call of GetRequests.
func (mr *MockRequesterMockRecorder) GetRequests(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRequests", reflect.TypeOf((*MockRequester)(nil).GetRequests), arg0, arg1)
}

// UpdateRequestStatus mocks base method.
func (m *MockRequester) UpdateRequestStatus(arg0 context.Context, arg1 int, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateRequestStatus", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateRequestStatus indicates an expected call of UpdateRequestStatus.
func (mr *MockRequesterMockRecorder) UpdateRequestStatus(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateRequestStatus", reflect.TypeOf((*MockRequester)(nil).UpdateRequestStatus), arg0, arg1, arg2)
}
