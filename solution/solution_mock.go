// Code generated by MockGen. DO NOT EDIT.
// Source: stub/github.com/csuJudge/judgeProto/solution/solution.trpc.go
//
// Generated by this command:
//
//	mockgen -destination=stub/github.com/csuJudge/judgeProto/solution/solution_mock.go -package=solution -self_package=github.com/csuJudge/judgeProto/solution --source=stub/github.com/csuJudge/judgeProto/solution/solution.trpc.go
//

// Package solution is a generated GoMock package.
package solution

import (
	context "context"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
	client "trpc.group/trpc-go/trpc-go/client"
)

// MockCodeServerService is a mock of CodeServerService interface.
type MockCodeServerService struct {
	ctrl     *gomock.Controller
	recorder *MockCodeServerServiceMockRecorder
}

// MockCodeServerServiceMockRecorder is the mock recorder for MockCodeServerService.
type MockCodeServerServiceMockRecorder struct {
	mock *MockCodeServerService
}

// NewMockCodeServerService creates a new mock instance.
func NewMockCodeServerService(ctrl *gomock.Controller) *MockCodeServerService {
	mock := &MockCodeServerService{ctrl: ctrl}
	mock.recorder = &MockCodeServerServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCodeServerService) EXPECT() *MockCodeServerServiceMockRecorder {
	return m.recorder
}

// ISGOMOCK indicates that this struct is a gomock mock.
func (m *MockCodeServerService) ISGOMOCK() struct{} {
	return struct{}{}
}

// AddCode mocks base method.
func (m *MockCodeServerService) AddCode(ctx context.Context, req *AddCodeReq) (*CommonRsp, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddCode", ctx, req)
	ret0, _ := ret[0].(*CommonRsp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddCode indicates an expected call of AddCode.
func (mr *MockCodeServerServiceMockRecorder) AddCode(ctx, req any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddCode", reflect.TypeOf((*MockCodeServerService)(nil).AddCode), ctx, req)
}

// MockKeyActionServerService is a mock of KeyActionServerService interface.
type MockKeyActionServerService struct {
	ctrl     *gomock.Controller
	recorder *MockKeyActionServerServiceMockRecorder
}

// MockKeyActionServerServiceMockRecorder is the mock recorder for MockKeyActionServerService.
type MockKeyActionServerServiceMockRecorder struct {
	mock *MockKeyActionServerService
}

// NewMockKeyActionServerService creates a new mock instance.
func NewMockKeyActionServerService(ctrl *gomock.Controller) *MockKeyActionServerService {
	mock := &MockKeyActionServerService{ctrl: ctrl}
	mock.recorder = &MockKeyActionServerServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockKeyActionServerService) EXPECT() *MockKeyActionServerServiceMockRecorder {
	return m.recorder
}

// ISGOMOCK indicates that this struct is a gomock mock.
func (m *MockKeyActionServerService) ISGOMOCK() struct{} {
	return struct{}{}
}

// AddKeyAction mocks base method.
func (m *MockKeyActionServerService) AddKeyAction(ctx context.Context, req *AddKeyActionReq) (*CommonRsp, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddKeyAction", ctx, req)
	ret0, _ := ret[0].(*CommonRsp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddKeyAction indicates an expected call of AddKeyAction.
func (mr *MockKeyActionServerServiceMockRecorder) AddKeyAction(ctx, req any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddKeyAction", reflect.TypeOf((*MockKeyActionServerService)(nil).AddKeyAction), ctx, req)
}

// MockSolutionServerService is a mock of SolutionServerService interface.
type MockSolutionServerService struct {
	ctrl     *gomock.Controller
	recorder *MockSolutionServerServiceMockRecorder
}

// MockSolutionServerServiceMockRecorder is the mock recorder for MockSolutionServerService.
type MockSolutionServerServiceMockRecorder struct {
	mock *MockSolutionServerService
}

// NewMockSolutionServerService creates a new mock instance.
func NewMockSolutionServerService(ctrl *gomock.Controller) *MockSolutionServerService {
	mock := &MockSolutionServerService{ctrl: ctrl}
	mock.recorder = &MockSolutionServerServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSolutionServerService) EXPECT() *MockSolutionServerServiceMockRecorder {
	return m.recorder
}

// ISGOMOCK indicates that this struct is a gomock mock.
func (m *MockSolutionServerService) ISGOMOCK() struct{} {
	return struct{}{}
}

// AddSolution mocks base method.
func (m *MockSolutionServerService) AddSolution(ctx context.Context, req *AddSolutionReq) (*AddSolutionRsp, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddSolution", ctx, req)
	ret0, _ := ret[0].(*AddSolutionRsp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddSolution indicates an expected call of AddSolution.
func (mr *MockSolutionServerServiceMockRecorder) AddSolution(ctx, req any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddSolution", reflect.TypeOf((*MockSolutionServerService)(nil).AddSolution), ctx, req)
}

// AddTestSolution mocks base method.
func (m *MockSolutionServerService) AddTestSolution(ctx context.Context, req *AddSolutionReq) (*AddSolutionRsp, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddTestSolution", ctx, req)
	ret0, _ := ret[0].(*AddSolutionRsp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddTestSolution indicates an expected call of AddTestSolution.
func (mr *MockSolutionServerServiceMockRecorder) AddTestSolution(ctx, req any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddTestSolution", reflect.TypeOf((*MockSolutionServerService)(nil).AddTestSolution), ctx, req)
}

// CountContestProblemSubmission mocks base method.
func (m *MockSolutionServerService) CountContestProblemSubmission(ctx context.Context, req *CountContestSubmissionReq) (*CountContestProblemSubmissionRsp, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CountContestProblemSubmission", ctx, req)
	ret0, _ := ret[0].(*CountContestProblemSubmissionRsp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CountContestProblemSubmission indicates an expected call of CountContestProblemSubmission.
func (mr *MockSolutionServerServiceMockRecorder) CountContestProblemSubmission(ctx, req any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CountContestProblemSubmission", reflect.TypeOf((*MockSolutionServerService)(nil).CountContestProblemSubmission), ctx, req)
}

// CountContestSubmission mocks base method.
func (m *MockSolutionServerService) CountContestSubmission(ctx context.Context, req *CountContestSubmissionReq) (*CountContestSubmissionRsp, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CountContestSubmission", ctx, req)
	ret0, _ := ret[0].(*CountContestSubmissionRsp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CountContestSubmission indicates an expected call of CountContestSubmission.
func (mr *MockSolutionServerServiceMockRecorder) CountContestSubmission(ctx, req any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CountContestSubmission", reflect.TypeOf((*MockSolutionServerService)(nil).CountContestSubmission), ctx, req)
}

// CountUserProblemSolution mocks base method.
func (m *MockSolutionServerService) CountUserProblemSolution(ctx context.Context, req *CountUserProblemSolutionReq) (*CountUserProblemSolutionRsp, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CountUserProblemSolution", ctx, req)
	ret0, _ := ret[0].(*CountUserProblemSolutionRsp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CountUserProblemSolution indicates an expected call of CountUserProblemSolution.
func (mr *MockSolutionServerServiceMockRecorder) CountUserProblemSolution(ctx, req any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CountUserProblemSolution", reflect.TypeOf((*MockSolutionServerService)(nil).CountUserProblemSolution), ctx, req)
}

// QueryContestRankData mocks base method.
func (m *MockSolutionServerService) QueryContestRankData(ctx context.Context, req *QueryContestRankDataReq) (*QueryContestRankDataRsp, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryContestRankData", ctx, req)
	ret0, _ := ret[0].(*QueryContestRankDataRsp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryContestRankData indicates an expected call of QueryContestRankData.
func (mr *MockSolutionServerServiceMockRecorder) QueryContestRankData(ctx, req any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryContestRankData", reflect.TypeOf((*MockSolutionServerService)(nil).QueryContestRankData), ctx, req)
}

// QueryLatestCode mocks base method.
func (m *MockSolutionServerService) QueryLatestCode(ctx context.Context, req *QueryLatestCodeReq) (*QueryLatestCodeRsp, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryLatestCode", ctx, req)
	ret0, _ := ret[0].(*QueryLatestCodeRsp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryLatestCode indicates an expected call of QueryLatestCode.
func (mr *MockSolutionServerServiceMockRecorder) QueryLatestCode(ctx, req any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryLatestCode", reflect.TypeOf((*MockSolutionServerService)(nil).QueryLatestCode), ctx, req)
}

// QueryRuntimeInfo mocks base method.
func (m *MockSolutionServerService) QueryRuntimeInfo(ctx context.Context, req *QueryRuntimeInfoReq) (*QueryRuntimeInfoRsp, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryRuntimeInfo", ctx, req)
	ret0, _ := ret[0].(*QueryRuntimeInfoRsp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryRuntimeInfo indicates an expected call of QueryRuntimeInfo.
func (mr *MockSolutionServerServiceMockRecorder) QueryRuntimeInfo(ctx, req any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryRuntimeInfo", reflect.TypeOf((*MockSolutionServerService)(nil).QueryRuntimeInfo), ctx, req)
}

// QuerySimList mocks base method.
func (m *MockSolutionServerService) QuerySimList(ctx context.Context, req *QuerySimListReq) (*QuerySimListRsp, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QuerySimList", ctx, req)
	ret0, _ := ret[0].(*QuerySimListRsp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QuerySimList indicates an expected call of QuerySimList.
func (mr *MockSolutionServerServiceMockRecorder) QuerySimList(ctx, req any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QuerySimList", reflect.TypeOf((*MockSolutionServerService)(nil).QuerySimList), ctx, req)
}

// QuerySimSolutionData mocks base method.
func (m *MockSolutionServerService) QuerySimSolutionData(ctx context.Context, req *QuerySimSolutionDataReq) (*QuerySimSolutionDataRsp, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QuerySimSolutionData", ctx, req)
	ret0, _ := ret[0].(*QuerySimSolutionDataRsp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QuerySimSolutionData indicates an expected call of QuerySimSolutionData.
func (mr *MockSolutionServerServiceMockRecorder) QuerySimSolutionData(ctx, req any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QuerySimSolutionData", reflect.TypeOf((*MockSolutionServerService)(nil).QuerySimSolutionData), ctx, req)
}

// QuerySolution mocks base method.
func (m *MockSolutionServerService) QuerySolution(ctx context.Context, req *QuerySolutionReq) (*QuerySolutionRsp, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QuerySolution", ctx, req)
	ret0, _ := ret[0].(*QuerySolutionRsp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QuerySolution indicates an expected call of QuerySolution.
func (mr *MockSolutionServerServiceMockRecorder) QuerySolution(ctx, req any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QuerySolution", reflect.TypeOf((*MockSolutionServerService)(nil).QuerySolution), ctx, req)
}

// QuerySolutionResult mocks base method.
func (m *MockSolutionServerService) QuerySolutionResult(ctx context.Context, req *QuerySolutionResultReq) (*QuerySolutionResultRsp, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QuerySolutionResult", ctx, req)
	ret0, _ := ret[0].(*QuerySolutionResultRsp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QuerySolutionResult indicates an expected call of QuerySolutionResult.
func (mr *MockSolutionServerServiceMockRecorder) QuerySolutionResult(ctx, req any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QuerySolutionResult", reflect.TypeOf((*MockSolutionServerService)(nil).QuerySolutionResult), ctx, req)
}

// QuerySourceCode mocks base method.
func (m *MockSolutionServerService) QuerySourceCode(ctx context.Context, req *QuerySourceCodeReq) (*QuerySourceCodeRsp, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QuerySourceCode", ctx, req)
	ret0, _ := ret[0].(*QuerySourceCodeRsp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QuerySourceCode indicates an expected call of QuerySourceCode.
func (mr *MockSolutionServerServiceMockRecorder) QuerySourceCode(ctx, req any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QuerySourceCode", reflect.TypeOf((*MockSolutionServerService)(nil).QuerySourceCode), ctx, req)
}

// QueryUserProblemSolution mocks base method.
func (m *MockSolutionServerService) QueryUserProblemSolution(ctx context.Context, req *QueryUserProblemSolutionReq) (*QueryUserProblemSolutionRsp, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryUserProblemSolution", ctx, req)
	ret0, _ := ret[0].(*QueryUserProblemSolutionRsp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryUserProblemSolution indicates an expected call of QueryUserProblemSolution.
func (mr *MockSolutionServerServiceMockRecorder) QueryUserProblemSolution(ctx, req any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryUserProblemSolution", reflect.TypeOf((*MockSolutionServerService)(nil).QueryUserProblemSolution), ctx, req)
}

// RejudgeSolution mocks base method.
func (m *MockSolutionServerService) RejudgeSolution(ctx context.Context, req *RejudgeSolutionReq) (*CommonRsp, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RejudgeSolution", ctx, req)
	ret0, _ := ret[0].(*CommonRsp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RejudgeSolution indicates an expected call of RejudgeSolution.
func (mr *MockSolutionServerServiceMockRecorder) RejudgeSolution(ctx, req any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RejudgeSolution", reflect.TypeOf((*MockSolutionServerService)(nil).RejudgeSolution), ctx, req)
}

// MockCodeServerClientProxy is a mock of CodeServerClientProxy interface.
type MockCodeServerClientProxy struct {
	ctrl     *gomock.Controller
	recorder *MockCodeServerClientProxyMockRecorder
}

// MockCodeServerClientProxyMockRecorder is the mock recorder for MockCodeServerClientProxy.
type MockCodeServerClientProxyMockRecorder struct {
	mock *MockCodeServerClientProxy
}

// NewMockCodeServerClientProxy creates a new mock instance.
func NewMockCodeServerClientProxy(ctrl *gomock.Controller) *MockCodeServerClientProxy {
	mock := &MockCodeServerClientProxy{ctrl: ctrl}
	mock.recorder = &MockCodeServerClientProxyMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCodeServerClientProxy) EXPECT() *MockCodeServerClientProxyMockRecorder {
	return m.recorder
}

// ISGOMOCK indicates that this struct is a gomock mock.
func (m *MockCodeServerClientProxy) ISGOMOCK() struct{} {
	return struct{}{}
}

// AddCode mocks base method.
func (m *MockCodeServerClientProxy) AddCode(ctx context.Context, req *AddCodeReq, opts ...client.Option) (*CommonRsp, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, req}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AddCode", varargs...)
	ret0, _ := ret[0].(*CommonRsp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddCode indicates an expected call of AddCode.
func (mr *MockCodeServerClientProxyMockRecorder) AddCode(ctx, req any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, req}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddCode", reflect.TypeOf((*MockCodeServerClientProxy)(nil).AddCode), varargs...)
}

// MockKeyActionServerClientProxy is a mock of KeyActionServerClientProxy interface.
type MockKeyActionServerClientProxy struct {
	ctrl     *gomock.Controller
	recorder *MockKeyActionServerClientProxyMockRecorder
}

// MockKeyActionServerClientProxyMockRecorder is the mock recorder for MockKeyActionServerClientProxy.
type MockKeyActionServerClientProxyMockRecorder struct {
	mock *MockKeyActionServerClientProxy
}

// NewMockKeyActionServerClientProxy creates a new mock instance.
func NewMockKeyActionServerClientProxy(ctrl *gomock.Controller) *MockKeyActionServerClientProxy {
	mock := &MockKeyActionServerClientProxy{ctrl: ctrl}
	mock.recorder = &MockKeyActionServerClientProxyMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockKeyActionServerClientProxy) EXPECT() *MockKeyActionServerClientProxyMockRecorder {
	return m.recorder
}

// ISGOMOCK indicates that this struct is a gomock mock.
func (m *MockKeyActionServerClientProxy) ISGOMOCK() struct{} {
	return struct{}{}
}

// AddKeyAction mocks base method.
func (m *MockKeyActionServerClientProxy) AddKeyAction(ctx context.Context, req *AddKeyActionReq, opts ...client.Option) (*CommonRsp, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, req}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AddKeyAction", varargs...)
	ret0, _ := ret[0].(*CommonRsp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddKeyAction indicates an expected call of AddKeyAction.
func (mr *MockKeyActionServerClientProxyMockRecorder) AddKeyAction(ctx, req any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, req}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddKeyAction", reflect.TypeOf((*MockKeyActionServerClientProxy)(nil).AddKeyAction), varargs...)
}

// MockSolutionServerClientProxy is a mock of SolutionServerClientProxy interface.
type MockSolutionServerClientProxy struct {
	ctrl     *gomock.Controller
	recorder *MockSolutionServerClientProxyMockRecorder
}

// MockSolutionServerClientProxyMockRecorder is the mock recorder for MockSolutionServerClientProxy.
type MockSolutionServerClientProxyMockRecorder struct {
	mock *MockSolutionServerClientProxy
}

// NewMockSolutionServerClientProxy creates a new mock instance.
func NewMockSolutionServerClientProxy(ctrl *gomock.Controller) *MockSolutionServerClientProxy {
	mock := &MockSolutionServerClientProxy{ctrl: ctrl}
	mock.recorder = &MockSolutionServerClientProxyMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSolutionServerClientProxy) EXPECT() *MockSolutionServerClientProxyMockRecorder {
	return m.recorder
}

// ISGOMOCK indicates that this struct is a gomock mock.
func (m *MockSolutionServerClientProxy) ISGOMOCK() struct{} {
	return struct{}{}
}

// AddSolution mocks base method.
func (m *MockSolutionServerClientProxy) AddSolution(ctx context.Context, req *AddSolutionReq, opts ...client.Option) (*AddSolutionRsp, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, req}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AddSolution", varargs...)
	ret0, _ := ret[0].(*AddSolutionRsp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddSolution indicates an expected call of AddSolution.
func (mr *MockSolutionServerClientProxyMockRecorder) AddSolution(ctx, req any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, req}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddSolution", reflect.TypeOf((*MockSolutionServerClientProxy)(nil).AddSolution), varargs...)
}

// AddTestSolution mocks base method.
func (m *MockSolutionServerClientProxy) AddTestSolution(ctx context.Context, req *AddSolutionReq, opts ...client.Option) (*AddSolutionRsp, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, req}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AddTestSolution", varargs...)
	ret0, _ := ret[0].(*AddSolutionRsp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddTestSolution indicates an expected call of AddTestSolution.
func (mr *MockSolutionServerClientProxyMockRecorder) AddTestSolution(ctx, req any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, req}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddTestSolution", reflect.TypeOf((*MockSolutionServerClientProxy)(nil).AddTestSolution), varargs...)
}

// CountContestProblemSubmission mocks base method.
func (m *MockSolutionServerClientProxy) CountContestProblemSubmission(ctx context.Context, req *CountContestSubmissionReq, opts ...client.Option) (*CountContestProblemSubmissionRsp, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, req}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CountContestProblemSubmission", varargs...)
	ret0, _ := ret[0].(*CountContestProblemSubmissionRsp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CountContestProblemSubmission indicates an expected call of CountContestProblemSubmission.
func (mr *MockSolutionServerClientProxyMockRecorder) CountContestProblemSubmission(ctx, req any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, req}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CountContestProblemSubmission", reflect.TypeOf((*MockSolutionServerClientProxy)(nil).CountContestProblemSubmission), varargs...)
}

// CountContestSubmission mocks base method.
func (m *MockSolutionServerClientProxy) CountContestSubmission(ctx context.Context, req *CountContestSubmissionReq, opts ...client.Option) (*CountContestSubmissionRsp, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, req}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CountContestSubmission", varargs...)
	ret0, _ := ret[0].(*CountContestSubmissionRsp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CountContestSubmission indicates an expected call of CountContestSubmission.
func (mr *MockSolutionServerClientProxyMockRecorder) CountContestSubmission(ctx, req any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, req}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CountContestSubmission", reflect.TypeOf((*MockSolutionServerClientProxy)(nil).CountContestSubmission), varargs...)
}

// CountUserProblemSolution mocks base method.
func (m *MockSolutionServerClientProxy) CountUserProblemSolution(ctx context.Context, req *CountUserProblemSolutionReq, opts ...client.Option) (*CountUserProblemSolutionRsp, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, req}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CountUserProblemSolution", varargs...)
	ret0, _ := ret[0].(*CountUserProblemSolutionRsp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CountUserProblemSolution indicates an expected call of CountUserProblemSolution.
func (mr *MockSolutionServerClientProxyMockRecorder) CountUserProblemSolution(ctx, req any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, req}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CountUserProblemSolution", reflect.TypeOf((*MockSolutionServerClientProxy)(nil).CountUserProblemSolution), varargs...)
}

// QueryContestRankData mocks base method.
func (m *MockSolutionServerClientProxy) QueryContestRankData(ctx context.Context, req *QueryContestRankDataReq, opts ...client.Option) (*QueryContestRankDataRsp, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, req}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "QueryContestRankData", varargs...)
	ret0, _ := ret[0].(*QueryContestRankDataRsp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryContestRankData indicates an expected call of QueryContestRankData.
func (mr *MockSolutionServerClientProxyMockRecorder) QueryContestRankData(ctx, req any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, req}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryContestRankData", reflect.TypeOf((*MockSolutionServerClientProxy)(nil).QueryContestRankData), varargs...)
}

// QueryLatestCode mocks base method.
func (m *MockSolutionServerClientProxy) QueryLatestCode(ctx context.Context, req *QueryLatestCodeReq, opts ...client.Option) (*QueryLatestCodeRsp, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, req}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "QueryLatestCode", varargs...)
	ret0, _ := ret[0].(*QueryLatestCodeRsp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryLatestCode indicates an expected call of QueryLatestCode.
func (mr *MockSolutionServerClientProxyMockRecorder) QueryLatestCode(ctx, req any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, req}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryLatestCode", reflect.TypeOf((*MockSolutionServerClientProxy)(nil).QueryLatestCode), varargs...)
}

// QueryRuntimeInfo mocks base method.
func (m *MockSolutionServerClientProxy) QueryRuntimeInfo(ctx context.Context, req *QueryRuntimeInfoReq, opts ...client.Option) (*QueryRuntimeInfoRsp, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, req}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "QueryRuntimeInfo", varargs...)
	ret0, _ := ret[0].(*QueryRuntimeInfoRsp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryRuntimeInfo indicates an expected call of QueryRuntimeInfo.
func (mr *MockSolutionServerClientProxyMockRecorder) QueryRuntimeInfo(ctx, req any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, req}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryRuntimeInfo", reflect.TypeOf((*MockSolutionServerClientProxy)(nil).QueryRuntimeInfo), varargs...)
}

// QuerySimList mocks base method.
func (m *MockSolutionServerClientProxy) QuerySimList(ctx context.Context, req *QuerySimListReq, opts ...client.Option) (*QuerySimListRsp, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, req}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "QuerySimList", varargs...)
	ret0, _ := ret[0].(*QuerySimListRsp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QuerySimList indicates an expected call of QuerySimList.
func (mr *MockSolutionServerClientProxyMockRecorder) QuerySimList(ctx, req any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, req}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QuerySimList", reflect.TypeOf((*MockSolutionServerClientProxy)(nil).QuerySimList), varargs...)
}

// QuerySimSolutionData mocks base method.
func (m *MockSolutionServerClientProxy) QuerySimSolutionData(ctx context.Context, req *QuerySimSolutionDataReq, opts ...client.Option) (*QuerySimSolutionDataRsp, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, req}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "QuerySimSolutionData", varargs...)
	ret0, _ := ret[0].(*QuerySimSolutionDataRsp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QuerySimSolutionData indicates an expected call of QuerySimSolutionData.
func (mr *MockSolutionServerClientProxyMockRecorder) QuerySimSolutionData(ctx, req any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, req}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QuerySimSolutionData", reflect.TypeOf((*MockSolutionServerClientProxy)(nil).QuerySimSolutionData), varargs...)
}

// QuerySolution mocks base method.
func (m *MockSolutionServerClientProxy) QuerySolution(ctx context.Context, req *QuerySolutionReq, opts ...client.Option) (*QuerySolutionRsp, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, req}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "QuerySolution", varargs...)
	ret0, _ := ret[0].(*QuerySolutionRsp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QuerySolution indicates an expected call of QuerySolution.
func (mr *MockSolutionServerClientProxyMockRecorder) QuerySolution(ctx, req any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, req}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QuerySolution", reflect.TypeOf((*MockSolutionServerClientProxy)(nil).QuerySolution), varargs...)
}

// QuerySolutionResult mocks base method.
func (m *MockSolutionServerClientProxy) QuerySolutionResult(ctx context.Context, req *QuerySolutionResultReq, opts ...client.Option) (*QuerySolutionResultRsp, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, req}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "QuerySolutionResult", varargs...)
	ret0, _ := ret[0].(*QuerySolutionResultRsp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QuerySolutionResult indicates an expected call of QuerySolutionResult.
func (mr *MockSolutionServerClientProxyMockRecorder) QuerySolutionResult(ctx, req any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, req}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QuerySolutionResult", reflect.TypeOf((*MockSolutionServerClientProxy)(nil).QuerySolutionResult), varargs...)
}

// QuerySourceCode mocks base method.
func (m *MockSolutionServerClientProxy) QuerySourceCode(ctx context.Context, req *QuerySourceCodeReq, opts ...client.Option) (*QuerySourceCodeRsp, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, req}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "QuerySourceCode", varargs...)
	ret0, _ := ret[0].(*QuerySourceCodeRsp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QuerySourceCode indicates an expected call of QuerySourceCode.
func (mr *MockSolutionServerClientProxyMockRecorder) QuerySourceCode(ctx, req any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, req}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QuerySourceCode", reflect.TypeOf((*MockSolutionServerClientProxy)(nil).QuerySourceCode), varargs...)
}

// QueryUserProblemSolution mocks base method.
func (m *MockSolutionServerClientProxy) QueryUserProblemSolution(ctx context.Context, req *QueryUserProblemSolutionReq, opts ...client.Option) (*QueryUserProblemSolutionRsp, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, req}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "QueryUserProblemSolution", varargs...)
	ret0, _ := ret[0].(*QueryUserProblemSolutionRsp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryUserProblemSolution indicates an expected call of QueryUserProblemSolution.
func (mr *MockSolutionServerClientProxyMockRecorder) QueryUserProblemSolution(ctx, req any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, req}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryUserProblemSolution", reflect.TypeOf((*MockSolutionServerClientProxy)(nil).QueryUserProblemSolution), varargs...)
}

// RejudgeSolution mocks base method.
func (m *MockSolutionServerClientProxy) RejudgeSolution(ctx context.Context, req *RejudgeSolutionReq, opts ...client.Option) (*CommonRsp, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, req}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RejudgeSolution", varargs...)
	ret0, _ := ret[0].(*CommonRsp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RejudgeSolution indicates an expected call of RejudgeSolution.
func (mr *MockSolutionServerClientProxyMockRecorder) RejudgeSolution(ctx, req any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, req}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RejudgeSolution", reflect.TypeOf((*MockSolutionServerClientProxy)(nil).RejudgeSolution), varargs...)
}
