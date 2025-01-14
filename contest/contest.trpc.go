// Code generated by trpc-go/trpc-cmdline v1.0.9. DO NOT EDIT.
// source: contest.proto

package contest

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

// NewsServerService defines service.
type NewsServerService interface {
	// QueryContest QueryContest 查询作业
	QueryContest(ctx context.Context, req *QueryContestReq) (*QueryContestRsp, error)
	// QueryContestList QueryContestList 查询作业列表
	QueryContestList(ctx context.Context, req *QueryContestListReq) (*QueryContestListRsp, error)
	// UpdateContestStatus UpdateContestStatus 更新作业状态
	UpdateContestStatus(ctx context.Context, req *UpdateContestStatusReq) (*CommonRsp, error)
}

func NewsServerService_QueryContest_Handler(svr interface{}, ctx context.Context, f server.FilterFunc) (interface{}, error) {
	req := &QueryContestReq{}
	filters, err := f(req)
	if err != nil {
		return nil, err
	}
	handleFunc := func(ctx context.Context, reqbody interface{}) (interface{}, error) {
		return svr.(NewsServerService).QueryContest(ctx, reqbody.(*QueryContestReq))
	}

	var rsp interface{}
	rsp, err = filters.Filter(ctx, req, handleFunc)
	if err != nil {
		return nil, err
	}
	return rsp, nil
}

func NewsServerService_QueryContestList_Handler(svr interface{}, ctx context.Context, f server.FilterFunc) (interface{}, error) {
	req := &QueryContestListReq{}
	filters, err := f(req)
	if err != nil {
		return nil, err
	}
	handleFunc := func(ctx context.Context, reqbody interface{}) (interface{}, error) {
		return svr.(NewsServerService).QueryContestList(ctx, reqbody.(*QueryContestListReq))
	}

	var rsp interface{}
	rsp, err = filters.Filter(ctx, req, handleFunc)
	if err != nil {
		return nil, err
	}
	return rsp, nil
}

func NewsServerService_UpdateContestStatus_Handler(svr interface{}, ctx context.Context, f server.FilterFunc) (interface{}, error) {
	req := &UpdateContestStatusReq{}
	filters, err := f(req)
	if err != nil {
		return nil, err
	}
	handleFunc := func(ctx context.Context, reqbody interface{}) (interface{}, error) {
		return svr.(NewsServerService).UpdateContestStatus(ctx, reqbody.(*UpdateContestStatusReq))
	}

	var rsp interface{}
	rsp, err = filters.Filter(ctx, req, handleFunc)
	if err != nil {
		return nil, err
	}
	return rsp, nil
}

// NewsServerServer_ServiceDesc descriptor for server.RegisterService.
var NewsServerServer_ServiceDesc = server.ServiceDesc{
	ServiceName: "oj.contest.NewsServer",
	HandlerType: ((*NewsServerService)(nil)),
	Methods: []server.Method{
		{
			Name: "/oj.contest.NewsServer/QueryContest",
			Func: NewsServerService_QueryContest_Handler,
		},
		{
			Name: "/oj.contest.NewsServer/QueryContestList",
			Func: NewsServerService_QueryContestList_Handler,
		},
		{
			Name: "/oj.contest.NewsServer/UpdateContestStatus",
			Func: NewsServerService_UpdateContestStatus_Handler,
		},
	},
}

// RegisterNewsServerService registers service.
func RegisterNewsServerService(s server.Service, svr NewsServerService) {
	if err := s.Register(&NewsServerServer_ServiceDesc, svr); err != nil {
		panic(fmt.Sprintf("NewsServer register error:%v", err))
	}
}

// START --------------------------------- Default Unimplemented Server Service --------------------------------- START

type UnimplementedNewsServer struct{}

// QueryContest QueryContest 查询作业
func (s *UnimplementedNewsServer) QueryContest(ctx context.Context, req *QueryContestReq) (*QueryContestRsp, error) {
	return nil, errors.New("rpc QueryContest of service NewsServer is not implemented")
}

// QueryContestList QueryContestList 查询作业列表
func (s *UnimplementedNewsServer) QueryContestList(ctx context.Context, req *QueryContestListReq) (*QueryContestListRsp, error) {
	return nil, errors.New("rpc QueryContestList of service NewsServer is not implemented")
}

// UpdateContestStatus UpdateContestStatus 更新作业状态
func (s *UnimplementedNewsServer) UpdateContestStatus(ctx context.Context, req *UpdateContestStatusReq) (*CommonRsp, error) {
	return nil, errors.New("rpc UpdateContestStatus of service NewsServer is not implemented")
}

// END --------------------------------- Default Unimplemented Server Service --------------------------------- END

// END ======================================= Server Service Definition ======================================= END

// START ======================================= Client Service Definition ======================================= START

// NewsServerClientProxy defines service client proxy
type NewsServerClientProxy interface {
	// QueryContest QueryContest 查询作业
	QueryContest(ctx context.Context, req *QueryContestReq, opts ...client.Option) (rsp *QueryContestRsp, err error)
	// QueryContestList QueryContestList 查询作业列表
	QueryContestList(ctx context.Context, req *QueryContestListReq, opts ...client.Option) (rsp *QueryContestListRsp, err error)
	// UpdateContestStatus UpdateContestStatus 更新作业状态
	UpdateContestStatus(ctx context.Context, req *UpdateContestStatusReq, opts ...client.Option) (rsp *CommonRsp, err error)
}

type NewsServerClientProxyImpl struct {
	client client.Client
	opts   []client.Option
}

var NewNewsServerClientProxy = func(opts ...client.Option) NewsServerClientProxy {
	return &NewsServerClientProxyImpl{client: client.DefaultClient, opts: opts}
}

func (c *NewsServerClientProxyImpl) QueryContest(ctx context.Context, req *QueryContestReq, opts ...client.Option) (*QueryContestRsp, error) {
	ctx, msg := codec.WithCloneMessage(ctx)
	defer codec.PutBackMessage(msg)
	msg.WithClientRPCName("/oj.contest.NewsServer/QueryContest")
	msg.WithCalleeServiceName(NewsServerServer_ServiceDesc.ServiceName)
	msg.WithCalleeApp("")
	msg.WithCalleeServer("")
	msg.WithCalleeService("NewsServer")
	msg.WithCalleeMethod("QueryContest")
	msg.WithSerializationType(codec.SerializationTypePB)
	callopts := make([]client.Option, 0, len(c.opts)+len(opts))
	callopts = append(callopts, c.opts...)
	callopts = append(callopts, opts...)
	rsp := &QueryContestRsp{}
	if err := c.client.Invoke(ctx, req, rsp, callopts...); err != nil {
		return nil, err
	}
	return rsp, nil
}

func (c *NewsServerClientProxyImpl) QueryContestList(ctx context.Context, req *QueryContestListReq, opts ...client.Option) (*QueryContestListRsp, error) {
	ctx, msg := codec.WithCloneMessage(ctx)
	defer codec.PutBackMessage(msg)
	msg.WithClientRPCName("/oj.contest.NewsServer/QueryContestList")
	msg.WithCalleeServiceName(NewsServerServer_ServiceDesc.ServiceName)
	msg.WithCalleeApp("")
	msg.WithCalleeServer("")
	msg.WithCalleeService("NewsServer")
	msg.WithCalleeMethod("QueryContestList")
	msg.WithSerializationType(codec.SerializationTypePB)
	callopts := make([]client.Option, 0, len(c.opts)+len(opts))
	callopts = append(callopts, c.opts...)
	callopts = append(callopts, opts...)
	rsp := &QueryContestListRsp{}
	if err := c.client.Invoke(ctx, req, rsp, callopts...); err != nil {
		return nil, err
	}
	return rsp, nil
}

func (c *NewsServerClientProxyImpl) UpdateContestStatus(ctx context.Context, req *UpdateContestStatusReq, opts ...client.Option) (*CommonRsp, error) {
	ctx, msg := codec.WithCloneMessage(ctx)
	defer codec.PutBackMessage(msg)
	msg.WithClientRPCName("/oj.contest.NewsServer/UpdateContestStatus")
	msg.WithCalleeServiceName(NewsServerServer_ServiceDesc.ServiceName)
	msg.WithCalleeApp("")
	msg.WithCalleeServer("")
	msg.WithCalleeService("NewsServer")
	msg.WithCalleeMethod("UpdateContestStatus")
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

// END ======================================= Client Service Definition ======================================= END
