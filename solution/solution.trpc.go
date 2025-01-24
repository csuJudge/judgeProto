// Code generated by trpc-go/trpc-cmdline v1.0.9. DO NOT EDIT.
// source: solution.proto

package solution

import (
	"context"
	"errors"
	"fmt"

	_ "trpc.group/trpc-go/trpc-go"
	"trpc.group/trpc-go/trpc-go/client"
	"trpc.group/trpc-go/trpc-go/codec"
	_ "trpc.group/trpc-go/trpc-go/http"
	"trpc.group/trpc-go/trpc-go/server"
)

// START ======================================= Server Service Definition ======================================= START

// SolutionServerService defines service.
type SolutionServerService interface {
	// CountUserProblemSolution CountUserProblemSolution 计算用户题目的提交次数
	CountUserProblemSolution(ctx context.Context, req *CountUserProblemSolutionReq) (*CountUserProblemSolutionRsp, error)
	// RejudgeSolution RejudgeSolution 重判
	RejudgeSolution(ctx context.Context, req *RejudgeSolutionReq) (*CommonRsp, error)
	// QueryRuntimeInfo QueryRuntimeInfo 查询信息
	QueryRuntimeInfo(ctx context.Context, req *QueryRuntimeInfoReq) (*QueryRuntimeInfoRsp, error)
	// QuerySolutionResult QuerySolutionResult 查询提交结果
	QuerySolutionResult(ctx context.Context, req *QuerySolutionResultReq) (*QuerySolutionResultRsp, error)
	// QuerySourceCode QuerySourceCode 查询源代码
	QuerySourceCode(ctx context.Context, req *QuerySourceCodeReq) (*QuerySourceCodeRsp, error)
	// QueryLatestCode QueryLatestCode 查询最新的代码
	QueryLatestCode(ctx context.Context, req *QueryLatestCodeReq) (*QueryLatestCodeRsp, error)
	// QuerySimList QuerySimList 查询相似度列表
	QuerySimList(ctx context.Context, req *QuerySimListReq) (*QuerySimListRsp, error)
	// QuerySimSolutionData QuerySimSolutionData 查询相似的代码数据
	QuerySimSolutionData(ctx context.Context, req *QuerySimSolutionDataReq) (*QuerySimSolutionDataRsp, error)
	// AddSolution AddSolution 添加提交数据
	AddSolution(ctx context.Context, req *AddSolutionReq) (*CommonRsp, error)
	// QueryUserProblemSolution QueryUserProblemSolution 查询用户题目的提交数据
	QueryUserProblemSolution(ctx context.Context, req *QueryUserProblemSolutionReq) (*QueryUserProblemSolutionRsp, error)
}

func SolutionServerService_CountUserProblemSolution_Handler(svr interface{}, ctx context.Context, f server.FilterFunc) (interface{}, error) {
	req := &CountUserProblemSolutionReq{}
	filters, err := f(req)
	if err != nil {
		return nil, err
	}
	handleFunc := func(ctx context.Context, reqbody interface{}) (interface{}, error) {
		return svr.(SolutionServerService).CountUserProblemSolution(ctx, reqbody.(*CountUserProblemSolutionReq))
	}

	var rsp interface{}
	rsp, err = filters.Filter(ctx, req, handleFunc)
	if err != nil {
		return nil, err
	}
	return rsp, nil
}

func SolutionServerService_RejudgeSolution_Handler(svr interface{}, ctx context.Context, f server.FilterFunc) (interface{}, error) {
	req := &RejudgeSolutionReq{}
	filters, err := f(req)
	if err != nil {
		return nil, err
	}
	handleFunc := func(ctx context.Context, reqbody interface{}) (interface{}, error) {
		return svr.(SolutionServerService).RejudgeSolution(ctx, reqbody.(*RejudgeSolutionReq))
	}

	var rsp interface{}
	rsp, err = filters.Filter(ctx, req, handleFunc)
	if err != nil {
		return nil, err
	}
	return rsp, nil
}

func SolutionServerService_QueryRuntimeInfo_Handler(svr interface{}, ctx context.Context, f server.FilterFunc) (interface{}, error) {
	req := &QueryRuntimeInfoReq{}
	filters, err := f(req)
	if err != nil {
		return nil, err
	}
	handleFunc := func(ctx context.Context, reqbody interface{}) (interface{}, error) {
		return svr.(SolutionServerService).QueryRuntimeInfo(ctx, reqbody.(*QueryRuntimeInfoReq))
	}

	var rsp interface{}
	rsp, err = filters.Filter(ctx, req, handleFunc)
	if err != nil {
		return nil, err
	}
	return rsp, nil
}

func SolutionServerService_QuerySolutionResult_Handler(svr interface{}, ctx context.Context, f server.FilterFunc) (interface{}, error) {
	req := &QuerySolutionResultReq{}
	filters, err := f(req)
	if err != nil {
		return nil, err
	}
	handleFunc := func(ctx context.Context, reqbody interface{}) (interface{}, error) {
		return svr.(SolutionServerService).QuerySolutionResult(ctx, reqbody.(*QuerySolutionResultReq))
	}

	var rsp interface{}
	rsp, err = filters.Filter(ctx, req, handleFunc)
	if err != nil {
		return nil, err
	}
	return rsp, nil
}

func SolutionServerService_QuerySourceCode_Handler(svr interface{}, ctx context.Context, f server.FilterFunc) (interface{}, error) {
	req := &QuerySourceCodeReq{}
	filters, err := f(req)
	if err != nil {
		return nil, err
	}
	handleFunc := func(ctx context.Context, reqbody interface{}) (interface{}, error) {
		return svr.(SolutionServerService).QuerySourceCode(ctx, reqbody.(*QuerySourceCodeReq))
	}

	var rsp interface{}
	rsp, err = filters.Filter(ctx, req, handleFunc)
	if err != nil {
		return nil, err
	}
	return rsp, nil
}

func SolutionServerService_QueryLatestCode_Handler(svr interface{}, ctx context.Context, f server.FilterFunc) (interface{}, error) {
	req := &QueryLatestCodeReq{}
	filters, err := f(req)
	if err != nil {
		return nil, err
	}
	handleFunc := func(ctx context.Context, reqbody interface{}) (interface{}, error) {
		return svr.(SolutionServerService).QueryLatestCode(ctx, reqbody.(*QueryLatestCodeReq))
	}

	var rsp interface{}
	rsp, err = filters.Filter(ctx, req, handleFunc)
	if err != nil {
		return nil, err
	}
	return rsp, nil
}

func SolutionServerService_QuerySimList_Handler(svr interface{}, ctx context.Context, f server.FilterFunc) (interface{}, error) {
	req := &QuerySimListReq{}
	filters, err := f(req)
	if err != nil {
		return nil, err
	}
	handleFunc := func(ctx context.Context, reqbody interface{}) (interface{}, error) {
		return svr.(SolutionServerService).QuerySimList(ctx, reqbody.(*QuerySimListReq))
	}

	var rsp interface{}
	rsp, err = filters.Filter(ctx, req, handleFunc)
	if err != nil {
		return nil, err
	}
	return rsp, nil
}

func SolutionServerService_QuerySimSolutionData_Handler(svr interface{}, ctx context.Context, f server.FilterFunc) (interface{}, error) {
	req := &QuerySimSolutionDataReq{}
	filters, err := f(req)
	if err != nil {
		return nil, err
	}
	handleFunc := func(ctx context.Context, reqbody interface{}) (interface{}, error) {
		return svr.(SolutionServerService).QuerySimSolutionData(ctx, reqbody.(*QuerySimSolutionDataReq))
	}

	var rsp interface{}
	rsp, err = filters.Filter(ctx, req, handleFunc)
	if err != nil {
		return nil, err
	}
	return rsp, nil
}

func SolutionServerService_AddSolution_Handler(svr interface{}, ctx context.Context, f server.FilterFunc) (interface{}, error) {
	req := &AddSolutionReq{}
	filters, err := f(req)
	if err != nil {
		return nil, err
	}
	handleFunc := func(ctx context.Context, reqbody interface{}) (interface{}, error) {
		return svr.(SolutionServerService).AddSolution(ctx, reqbody.(*AddSolutionReq))
	}

	var rsp interface{}
	rsp, err = filters.Filter(ctx, req, handleFunc)
	if err != nil {
		return nil, err
	}
	return rsp, nil
}

func SolutionServerService_QueryUserProblemSolution_Handler(svr interface{}, ctx context.Context, f server.FilterFunc) (interface{}, error) {
	req := &QueryUserProblemSolutionReq{}
	filters, err := f(req)
	if err != nil {
		return nil, err
	}
	handleFunc := func(ctx context.Context, reqbody interface{}) (interface{}, error) {
		return svr.(SolutionServerService).QueryUserProblemSolution(ctx, reqbody.(*QueryUserProblemSolutionReq))
	}

	var rsp interface{}
	rsp, err = filters.Filter(ctx, req, handleFunc)
	if err != nil {
		return nil, err
	}
	return rsp, nil
}

// SolutionServerServer_ServiceDesc descriptor for server.RegisterService.
var SolutionServerServer_ServiceDesc = server.ServiceDesc{
	ServiceName: "oj.solution.SolutionServer",
	HandlerType: ((*SolutionServerService)(nil)),
	Methods: []server.Method{
		{
			Name: "/oj.solution.SolutionServer/CountUserProblemSolution",
			Func: SolutionServerService_CountUserProblemSolution_Handler,
		},
		{
			Name: "/oj.solution.SolutionServer/RejudgeSolution",
			Func: SolutionServerService_RejudgeSolution_Handler,
		},
		{
			Name: "/oj.solution.SolutionServer/QueryRuntimeInfo",
			Func: SolutionServerService_QueryRuntimeInfo_Handler,
		},
		{
			Name: "/oj.solution.SolutionServer/QuerySolutionResult",
			Func: SolutionServerService_QuerySolutionResult_Handler,
		},
		{
			Name: "/oj.solution.SolutionServer/QuerySourceCode",
			Func: SolutionServerService_QuerySourceCode_Handler,
		},
		{
			Name: "/oj.solution.SolutionServer/QueryLatestCode",
			Func: SolutionServerService_QueryLatestCode_Handler,
		},
		{
			Name: "/oj.solution.SolutionServer/QuerySimList",
			Func: SolutionServerService_QuerySimList_Handler,
		},
		{
			Name: "/oj.solution.SolutionServer/QuerySimSolutionData",
			Func: SolutionServerService_QuerySimSolutionData_Handler,
		},
		{
			Name: "/oj.solution.SolutionServer/AddSolution",
			Func: SolutionServerService_AddSolution_Handler,
		},
		{
			Name: "/oj.solution.SolutionServer/QueryUserProblemSolution",
			Func: SolutionServerService_QueryUserProblemSolution_Handler,
		},
	},
}

// RegisterSolutionServerService registers service.
func RegisterSolutionServerService(s server.Service, svr SolutionServerService) {
	if err := s.Register(&SolutionServerServer_ServiceDesc, svr); err != nil {
		panic(fmt.Sprintf("SolutionServer register error:%v", err))
	}
}

// START --------------------------------- Default Unimplemented Server Service --------------------------------- START

type UnimplementedSolutionServer struct{}

// CountUserProblemSolution CountUserProblemSolution 计算用户题目的提交次数
func (s *UnimplementedSolutionServer) CountUserProblemSolution(ctx context.Context, req *CountUserProblemSolutionReq) (*CountUserProblemSolutionRsp, error) {
	return nil, errors.New("rpc CountUserProblemSolution of service SolutionServer is not implemented")
}

// RejudgeSolution RejudgeSolution 重判
func (s *UnimplementedSolutionServer) RejudgeSolution(ctx context.Context, req *RejudgeSolutionReq) (*CommonRsp, error) {
	return nil, errors.New("rpc RejudgeSolution of service SolutionServer is not implemented")
}

// QueryRuntimeInfo QueryRuntimeInfo 查询信息
func (s *UnimplementedSolutionServer) QueryRuntimeInfo(ctx context.Context, req *QueryRuntimeInfoReq) (*QueryRuntimeInfoRsp, error) {
	return nil, errors.New("rpc QueryRuntimeInfo of service SolutionServer is not implemented")
}

// QuerySolutionResult QuerySolutionResult 查询提交结果
func (s *UnimplementedSolutionServer) QuerySolutionResult(ctx context.Context, req *QuerySolutionResultReq) (*QuerySolutionResultRsp, error) {
	return nil, errors.New("rpc QuerySolutionResult of service SolutionServer is not implemented")
}

// QuerySourceCode QuerySourceCode 查询源代码
func (s *UnimplementedSolutionServer) QuerySourceCode(ctx context.Context, req *QuerySourceCodeReq) (*QuerySourceCodeRsp, error) {
	return nil, errors.New("rpc QuerySourceCode of service SolutionServer is not implemented")
}

// QueryLatestCode QueryLatestCode 查询最新的代码
func (s *UnimplementedSolutionServer) QueryLatestCode(ctx context.Context, req *QueryLatestCodeReq) (*QueryLatestCodeRsp, error) {
	return nil, errors.New("rpc QueryLatestCode of service SolutionServer is not implemented")
}

// QuerySimList QuerySimList 查询相似度列表
func (s *UnimplementedSolutionServer) QuerySimList(ctx context.Context, req *QuerySimListReq) (*QuerySimListRsp, error) {
	return nil, errors.New("rpc QuerySimList of service SolutionServer is not implemented")
}

// QuerySimSolutionData QuerySimSolutionData 查询相似的代码数据
func (s *UnimplementedSolutionServer) QuerySimSolutionData(ctx context.Context, req *QuerySimSolutionDataReq) (*QuerySimSolutionDataRsp, error) {
	return nil, errors.New("rpc QuerySimSolutionData of service SolutionServer is not implemented")
}

// AddSolution AddSolution 添加提交数据
func (s *UnimplementedSolutionServer) AddSolution(ctx context.Context, req *AddSolutionReq) (*CommonRsp, error) {
	return nil, errors.New("rpc AddSolution of service SolutionServer is not implemented")
}

// QueryUserProblemSolution QueryUserProblemSolution 查询用户题目的提交数据
func (s *UnimplementedSolutionServer) QueryUserProblemSolution(ctx context.Context, req *QueryUserProblemSolutionReq) (*QueryUserProblemSolutionRsp, error) {
	return nil, errors.New("rpc QueryUserProblemSolution of service SolutionServer is not implemented")
}

// END --------------------------------- Default Unimplemented Server Service --------------------------------- END

// END ======================================= Server Service Definition ======================================= END

// START ======================================= Client Service Definition ======================================= START

// SolutionServerClientProxy defines service client proxy
type SolutionServerClientProxy interface {
	// CountUserProblemSolution CountUserProblemSolution 计算用户题目的提交次数
	CountUserProblemSolution(ctx context.Context, req *CountUserProblemSolutionReq, opts ...client.Option) (rsp *CountUserProblemSolutionRsp, err error)
	// RejudgeSolution RejudgeSolution 重判
	RejudgeSolution(ctx context.Context, req *RejudgeSolutionReq, opts ...client.Option) (rsp *CommonRsp, err error)
	// QueryRuntimeInfo QueryRuntimeInfo 查询信息
	QueryRuntimeInfo(ctx context.Context, req *QueryRuntimeInfoReq, opts ...client.Option) (rsp *QueryRuntimeInfoRsp, err error)
	// QuerySolutionResult QuerySolutionResult 查询提交结果
	QuerySolutionResult(ctx context.Context, req *QuerySolutionResultReq, opts ...client.Option) (rsp *QuerySolutionResultRsp, err error)
	// QuerySourceCode QuerySourceCode 查询源代码
	QuerySourceCode(ctx context.Context, req *QuerySourceCodeReq, opts ...client.Option) (rsp *QuerySourceCodeRsp, err error)
	// QueryLatestCode QueryLatestCode 查询最新的代码
	QueryLatestCode(ctx context.Context, req *QueryLatestCodeReq, opts ...client.Option) (rsp *QueryLatestCodeRsp, err error)
	// QuerySimList QuerySimList 查询相似度列表
	QuerySimList(ctx context.Context, req *QuerySimListReq, opts ...client.Option) (rsp *QuerySimListRsp, err error)
	// QuerySimSolutionData QuerySimSolutionData 查询相似的代码数据
	QuerySimSolutionData(ctx context.Context, req *QuerySimSolutionDataReq, opts ...client.Option) (rsp *QuerySimSolutionDataRsp, err error)
	// AddSolution AddSolution 添加提交数据
	AddSolution(ctx context.Context, req *AddSolutionReq, opts ...client.Option) (rsp *CommonRsp, err error)
	// QueryUserProblemSolution QueryUserProblemSolution 查询用户题目的提交数据
	QueryUserProblemSolution(ctx context.Context, req *QueryUserProblemSolutionReq, opts ...client.Option) (rsp *QueryUserProblemSolutionRsp, err error)
}

type SolutionServerClientProxyImpl struct {
	client client.Client
	opts   []client.Option
}

var NewSolutionServerClientProxy = func(opts ...client.Option) SolutionServerClientProxy {
	return &SolutionServerClientProxyImpl{client: client.DefaultClient, opts: opts}
}

func (c *SolutionServerClientProxyImpl) CountUserProblemSolution(ctx context.Context, req *CountUserProblemSolutionReq, opts ...client.Option) (*CountUserProblemSolutionRsp, error) {
	ctx, msg := codec.WithCloneMessage(ctx)
	defer codec.PutBackMessage(msg)
	msg.WithClientRPCName("/oj.solution.SolutionServer/CountUserProblemSolution")
	msg.WithCalleeServiceName(SolutionServerServer_ServiceDesc.ServiceName)
	msg.WithCalleeApp("")
	msg.WithCalleeServer("")
	msg.WithCalleeService("SolutionServer")
	msg.WithCalleeMethod("CountUserProblemSolution")
	msg.WithSerializationType(codec.SerializationTypePB)
	callopts := make([]client.Option, 0, len(c.opts)+len(opts))
	callopts = append(callopts, c.opts...)
	callopts = append(callopts, opts...)
	rsp := &CountUserProblemSolutionRsp{}
	if err := c.client.Invoke(ctx, req, rsp, callopts...); err != nil {
		return nil, err
	}
	return rsp, nil
}

func (c *SolutionServerClientProxyImpl) RejudgeSolution(ctx context.Context, req *RejudgeSolutionReq, opts ...client.Option) (*CommonRsp, error) {
	ctx, msg := codec.WithCloneMessage(ctx)
	defer codec.PutBackMessage(msg)
	msg.WithClientRPCName("/oj.solution.SolutionServer/RejudgeSolution")
	msg.WithCalleeServiceName(SolutionServerServer_ServiceDesc.ServiceName)
	msg.WithCalleeApp("")
	msg.WithCalleeServer("")
	msg.WithCalleeService("SolutionServer")
	msg.WithCalleeMethod("RejudgeSolution")
	msg.WithSerializationType(codec.SerializationTypePB)
	callopts := make([]client.Option, 0, len(c.opts)+len(opts))
	callopts = append(callopts, c.opts...)
	callopts = append(callopts, opts...)
	rsp := &CommonRsp{}
	if err := c.client.Invoke(ctx, req, rsp, callopts...); err != nil {
		return nil, err
	}
	return rsp, nil
}

func (c *SolutionServerClientProxyImpl) QueryRuntimeInfo(ctx context.Context, req *QueryRuntimeInfoReq, opts ...client.Option) (*QueryRuntimeInfoRsp, error) {
	ctx, msg := codec.WithCloneMessage(ctx)
	defer codec.PutBackMessage(msg)
	msg.WithClientRPCName("/oj.solution.SolutionServer/QueryRuntimeInfo")
	msg.WithCalleeServiceName(SolutionServerServer_ServiceDesc.ServiceName)
	msg.WithCalleeApp("")
	msg.WithCalleeServer("")
	msg.WithCalleeService("SolutionServer")
	msg.WithCalleeMethod("QueryRuntimeInfo")
	msg.WithSerializationType(codec.SerializationTypePB)
	callopts := make([]client.Option, 0, len(c.opts)+len(opts))
	callopts = append(callopts, c.opts...)
	callopts = append(callopts, opts...)
	rsp := &QueryRuntimeInfoRsp{}
	if err := c.client.Invoke(ctx, req, rsp, callopts...); err != nil {
		return nil, err
	}
	return rsp, nil
}

func (c *SolutionServerClientProxyImpl) QuerySolutionResult(ctx context.Context, req *QuerySolutionResultReq, opts ...client.Option) (*QuerySolutionResultRsp, error) {
	ctx, msg := codec.WithCloneMessage(ctx)
	defer codec.PutBackMessage(msg)
	msg.WithClientRPCName("/oj.solution.SolutionServer/QuerySolutionResult")
	msg.WithCalleeServiceName(SolutionServerServer_ServiceDesc.ServiceName)
	msg.WithCalleeApp("")
	msg.WithCalleeServer("")
	msg.WithCalleeService("SolutionServer")
	msg.WithCalleeMethod("QuerySolutionResult")
	msg.WithSerializationType(codec.SerializationTypePB)
	callopts := make([]client.Option, 0, len(c.opts)+len(opts))
	callopts = append(callopts, c.opts...)
	callopts = append(callopts, opts...)
	rsp := &QuerySolutionResultRsp{}
	if err := c.client.Invoke(ctx, req, rsp, callopts...); err != nil {
		return nil, err
	}
	return rsp, nil
}

func (c *SolutionServerClientProxyImpl) QuerySourceCode(ctx context.Context, req *QuerySourceCodeReq, opts ...client.Option) (*QuerySourceCodeRsp, error) {
	ctx, msg := codec.WithCloneMessage(ctx)
	defer codec.PutBackMessage(msg)
	msg.WithClientRPCName("/oj.solution.SolutionServer/QuerySourceCode")
	msg.WithCalleeServiceName(SolutionServerServer_ServiceDesc.ServiceName)
	msg.WithCalleeApp("")
	msg.WithCalleeServer("")
	msg.WithCalleeService("SolutionServer")
	msg.WithCalleeMethod("QuerySourceCode")
	msg.WithSerializationType(codec.SerializationTypePB)
	callopts := make([]client.Option, 0, len(c.opts)+len(opts))
	callopts = append(callopts, c.opts...)
	callopts = append(callopts, opts...)
	rsp := &QuerySourceCodeRsp{}
	if err := c.client.Invoke(ctx, req, rsp, callopts...); err != nil {
		return nil, err
	}
	return rsp, nil
}

func (c *SolutionServerClientProxyImpl) QueryLatestCode(ctx context.Context, req *QueryLatestCodeReq, opts ...client.Option) (*QueryLatestCodeRsp, error) {
	ctx, msg := codec.WithCloneMessage(ctx)
	defer codec.PutBackMessage(msg)
	msg.WithClientRPCName("/oj.solution.SolutionServer/QueryLatestCode")
	msg.WithCalleeServiceName(SolutionServerServer_ServiceDesc.ServiceName)
	msg.WithCalleeApp("")
	msg.WithCalleeServer("")
	msg.WithCalleeService("SolutionServer")
	msg.WithCalleeMethod("QueryLatestCode")
	msg.WithSerializationType(codec.SerializationTypePB)
	callopts := make([]client.Option, 0, len(c.opts)+len(opts))
	callopts = append(callopts, c.opts...)
	callopts = append(callopts, opts...)
	rsp := &QueryLatestCodeRsp{}
	if err := c.client.Invoke(ctx, req, rsp, callopts...); err != nil {
		return nil, err
	}
	return rsp, nil
}

func (c *SolutionServerClientProxyImpl) QuerySimList(ctx context.Context, req *QuerySimListReq, opts ...client.Option) (*QuerySimListRsp, error) {
	ctx, msg := codec.WithCloneMessage(ctx)
	defer codec.PutBackMessage(msg)
	msg.WithClientRPCName("/oj.solution.SolutionServer/QuerySimList")
	msg.WithCalleeServiceName(SolutionServerServer_ServiceDesc.ServiceName)
	msg.WithCalleeApp("")
	msg.WithCalleeServer("")
	msg.WithCalleeService("SolutionServer")
	msg.WithCalleeMethod("QuerySimList")
	msg.WithSerializationType(codec.SerializationTypePB)
	callopts := make([]client.Option, 0, len(c.opts)+len(opts))
	callopts = append(callopts, c.opts...)
	callopts = append(callopts, opts...)
	rsp := &QuerySimListRsp{}
	if err := c.client.Invoke(ctx, req, rsp, callopts...); err != nil {
		return nil, err
	}
	return rsp, nil
}

func (c *SolutionServerClientProxyImpl) QuerySimSolutionData(ctx context.Context, req *QuerySimSolutionDataReq, opts ...client.Option) (*QuerySimSolutionDataRsp, error) {
	ctx, msg := codec.WithCloneMessage(ctx)
	defer codec.PutBackMessage(msg)
	msg.WithClientRPCName("/oj.solution.SolutionServer/QuerySimSolutionData")
	msg.WithCalleeServiceName(SolutionServerServer_ServiceDesc.ServiceName)
	msg.WithCalleeApp("")
	msg.WithCalleeServer("")
	msg.WithCalleeService("SolutionServer")
	msg.WithCalleeMethod("QuerySimSolutionData")
	msg.WithSerializationType(codec.SerializationTypePB)
	callopts := make([]client.Option, 0, len(c.opts)+len(opts))
	callopts = append(callopts, c.opts...)
	callopts = append(callopts, opts...)
	rsp := &QuerySimSolutionDataRsp{}
	if err := c.client.Invoke(ctx, req, rsp, callopts...); err != nil {
		return nil, err
	}
	return rsp, nil
}

func (c *SolutionServerClientProxyImpl) AddSolution(ctx context.Context, req *AddSolutionReq, opts ...client.Option) (*CommonRsp, error) {
	ctx, msg := codec.WithCloneMessage(ctx)
	defer codec.PutBackMessage(msg)
	msg.WithClientRPCName("/oj.solution.SolutionServer/AddSolution")
	msg.WithCalleeServiceName(SolutionServerServer_ServiceDesc.ServiceName)
	msg.WithCalleeApp("")
	msg.WithCalleeServer("")
	msg.WithCalleeService("SolutionServer")
	msg.WithCalleeMethod("AddSolution")
	msg.WithSerializationType(codec.SerializationTypePB)
	callopts := make([]client.Option, 0, len(c.opts)+len(opts))
	callopts = append(callopts, c.opts...)
	callopts = append(callopts, opts...)
	rsp := &CommonRsp{}
	if err := c.client.Invoke(ctx, req, rsp, callopts...); err != nil {
		return nil, err
	}
	return rsp, nil
}

func (c *SolutionServerClientProxyImpl) QueryUserProblemSolution(ctx context.Context, req *QueryUserProblemSolutionReq, opts ...client.Option) (*QueryUserProblemSolutionRsp, error) {
	ctx, msg := codec.WithCloneMessage(ctx)
	defer codec.PutBackMessage(msg)
	msg.WithClientRPCName("/oj.solution.SolutionServer/QueryUserProblemSolution")
	msg.WithCalleeServiceName(SolutionServerServer_ServiceDesc.ServiceName)
	msg.WithCalleeApp("")
	msg.WithCalleeServer("")
	msg.WithCalleeService("SolutionServer")
	msg.WithCalleeMethod("QueryUserProblemSolution")
	msg.WithSerializationType(codec.SerializationTypePB)
	callopts := make([]client.Option, 0, len(c.opts)+len(opts))
	callopts = append(callopts, c.opts...)
	callopts = append(callopts, opts...)
	rsp := &QueryUserProblemSolutionRsp{}
	if err := c.client.Invoke(ctx, req, rsp, callopts...); err != nil {
		return nil, err
	}
	return rsp, nil
}

// END ======================================= Client Service Definition ======================================= END
