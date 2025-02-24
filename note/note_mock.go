// Code generated by MockGen. DO NOT EDIT.
// Source: stub/github.com/csuJudge/judgeProto/note/note.trpc.go
//
// Generated by this command:
//
//	mockgen -destination=stub/github.com/csuJudge/judgeProto/note/note_mock.go -package=note -self_package=github.com/csuJudge/judgeProto/note --source=stub/github.com/csuJudge/judgeProto/note/note.trpc.go
//

// Package note is a generated GoMock package.
package note

import (
	context "context"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
	client "trpc.group/trpc-go/trpc-go/client"
)

// MockNoteServerService is a mock of NoteServerService interface.
type MockNoteServerService struct {
	ctrl     *gomock.Controller
	recorder *MockNoteServerServiceMockRecorder
}

// MockNoteServerServiceMockRecorder is the mock recorder for MockNoteServerService.
type MockNoteServerServiceMockRecorder struct {
	mock *MockNoteServerService
}

// NewMockNoteServerService creates a new mock instance.
func NewMockNoteServerService(ctrl *gomock.Controller) *MockNoteServerService {
	mock := &MockNoteServerService{ctrl: ctrl}
	mock.recorder = &MockNoteServerServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNoteServerService) EXPECT() *MockNoteServerServiceMockRecorder {
	return m.recorder
}

// ISGOMOCK indicates that this struct is a gomock mock.
func (m *MockNoteServerService) ISGOMOCK() struct{} {
	return struct{}{}
}

// OperateNote mocks base method.
func (m *MockNoteServerService) OperateNote(ctx context.Context, req *OperateNoteReq) (*OperateNoteRsp, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OperateNote", ctx, req)
	ret0, _ := ret[0].(*OperateNoteRsp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// OperateNote indicates an expected call of OperateNote.
func (mr *MockNoteServerServiceMockRecorder) OperateNote(ctx, req any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OperateNote", reflect.TypeOf((*MockNoteServerService)(nil).OperateNote), ctx, req)
}

// QueryNote mocks base method.
func (m *MockNoteServerService) QueryNote(ctx context.Context, req *QueryNoteReq) (*QueryNoteRsp, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryNote", ctx, req)
	ret0, _ := ret[0].(*QueryNoteRsp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryNote indicates an expected call of QueryNote.
func (mr *MockNoteServerServiceMockRecorder) QueryNote(ctx, req any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryNote", reflect.TypeOf((*MockNoteServerService)(nil).QueryNote), ctx, req)
}

// MockNoteServerClientProxy is a mock of NoteServerClientProxy interface.
type MockNoteServerClientProxy struct {
	ctrl     *gomock.Controller
	recorder *MockNoteServerClientProxyMockRecorder
}

// MockNoteServerClientProxyMockRecorder is the mock recorder for MockNoteServerClientProxy.
type MockNoteServerClientProxyMockRecorder struct {
	mock *MockNoteServerClientProxy
}

// NewMockNoteServerClientProxy creates a new mock instance.
func NewMockNoteServerClientProxy(ctrl *gomock.Controller) *MockNoteServerClientProxy {
	mock := &MockNoteServerClientProxy{ctrl: ctrl}
	mock.recorder = &MockNoteServerClientProxyMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNoteServerClientProxy) EXPECT() *MockNoteServerClientProxyMockRecorder {
	return m.recorder
}

// ISGOMOCK indicates that this struct is a gomock mock.
func (m *MockNoteServerClientProxy) ISGOMOCK() struct{} {
	return struct{}{}
}

// OperateNote mocks base method.
func (m *MockNoteServerClientProxy) OperateNote(ctx context.Context, req *OperateNoteReq, opts ...client.Option) (*OperateNoteRsp, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, req}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "OperateNote", varargs...)
	ret0, _ := ret[0].(*OperateNoteRsp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// OperateNote indicates an expected call of OperateNote.
func (mr *MockNoteServerClientProxyMockRecorder) OperateNote(ctx, req any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, req}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OperateNote", reflect.TypeOf((*MockNoteServerClientProxy)(nil).OperateNote), varargs...)
}

// QueryNote mocks base method.
func (m *MockNoteServerClientProxy) QueryNote(ctx context.Context, req *QueryNoteReq, opts ...client.Option) (*QueryNoteRsp, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, req}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "QueryNote", varargs...)
	ret0, _ := ret[0].(*QueryNoteRsp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryNote indicates an expected call of QueryNote.
func (mr *MockNoteServerClientProxyMockRecorder) QueryNote(ctx, req any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, req}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryNote", reflect.TypeOf((*MockNoteServerClientProxy)(nil).QueryNote), varargs...)
}
