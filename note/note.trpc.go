// Code generated by trpc-go/trpc-cmdline v1.0.9. DO NOT EDIT.
// source: note.proto

package note

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

// NoteServerService defines service.
type NoteServerService interface {
	// QueryNote QueryNote 查询题解
	QueryNote(ctx context.Context, req *QueryNoteReq) (*QueryNoteRsp, error)
	// OperateNote OperateNote 操作题解
	OperateNote(ctx context.Context, req *OperateNoteReq) (*OperateNoteRsp, error)
}

func NoteServerService_QueryNote_Handler(svr interface{}, ctx context.Context, f server.FilterFunc) (interface{}, error) {
	req := &QueryNoteReq{}
	filters, err := f(req)
	if err != nil {
		return nil, err
	}
	handleFunc := func(ctx context.Context, reqbody interface{}) (interface{}, error) {
		return svr.(NoteServerService).QueryNote(ctx, reqbody.(*QueryNoteReq))
	}

	var rsp interface{}
	rsp, err = filters.Filter(ctx, req, handleFunc)
	if err != nil {
		return nil, err
	}
	return rsp, nil
}

func NoteServerService_OperateNote_Handler(svr interface{}, ctx context.Context, f server.FilterFunc) (interface{}, error) {
	req := &OperateNoteReq{}
	filters, err := f(req)
	if err != nil {
		return nil, err
	}
	handleFunc := func(ctx context.Context, reqbody interface{}) (interface{}, error) {
		return svr.(NoteServerService).OperateNote(ctx, reqbody.(*OperateNoteReq))
	}

	var rsp interface{}
	rsp, err = filters.Filter(ctx, req, handleFunc)
	if err != nil {
		return nil, err
	}
	return rsp, nil
}

// NoteServerServer_ServiceDesc descriptor for server.RegisterService.
var NoteServerServer_ServiceDesc = server.ServiceDesc{
	ServiceName: "oj.note.NoteServer",
	HandlerType: ((*NoteServerService)(nil)),
	Methods: []server.Method{
		{
			Name: "/oj.note.NoteServer/QueryNote",
			Func: NoteServerService_QueryNote_Handler,
		},
		{
			Name: "/oj.note.NoteServer/OperateNote",
			Func: NoteServerService_OperateNote_Handler,
		},
	},
}

// RegisterNoteServerService registers service.
func RegisterNoteServerService(s server.Service, svr NoteServerService) {
	if err := s.Register(&NoteServerServer_ServiceDesc, svr); err != nil {
		panic(fmt.Sprintf("NoteServer register error:%v", err))
	}
}

// START --------------------------------- Default Unimplemented Server Service --------------------------------- START

type UnimplementedNoteServer struct{}

// QueryNote QueryNote 查询题解
func (s *UnimplementedNoteServer) QueryNote(ctx context.Context, req *QueryNoteReq) (*QueryNoteRsp, error) {
	return nil, errors.New("rpc QueryNote of service NoteServer is not implemented")
}

// OperateNote OperateNote 操作题解
func (s *UnimplementedNoteServer) OperateNote(ctx context.Context, req *OperateNoteReq) (*OperateNoteRsp, error) {
	return nil, errors.New("rpc OperateNote of service NoteServer is not implemented")
}

// END --------------------------------- Default Unimplemented Server Service --------------------------------- END

// END ======================================= Server Service Definition ======================================= END

// START ======================================= Client Service Definition ======================================= START

// NoteServerClientProxy defines service client proxy
type NoteServerClientProxy interface {
	// QueryNote QueryNote 查询题解
	QueryNote(ctx context.Context, req *QueryNoteReq, opts ...client.Option) (rsp *QueryNoteRsp, err error)
	// OperateNote OperateNote 操作题解
	OperateNote(ctx context.Context, req *OperateNoteReq, opts ...client.Option) (rsp *OperateNoteRsp, err error)
}

type NoteServerClientProxyImpl struct {
	client client.Client
	opts   []client.Option
}

var NewNoteServerClientProxy = func(opts ...client.Option) NoteServerClientProxy {
	return &NoteServerClientProxyImpl{client: client.DefaultClient, opts: opts}
}

func (c *NoteServerClientProxyImpl) QueryNote(ctx context.Context, req *QueryNoteReq, opts ...client.Option) (*QueryNoteRsp, error) {
	ctx, msg := codec.WithCloneMessage(ctx)
	defer codec.PutBackMessage(msg)
	msg.WithClientRPCName("/oj.note.NoteServer/QueryNote")
	msg.WithCalleeServiceName(NoteServerServer_ServiceDesc.ServiceName)
	msg.WithCalleeApp("")
	msg.WithCalleeServer("")
	msg.WithCalleeService("NoteServer")
	msg.WithCalleeMethod("QueryNote")
	msg.WithSerializationType(codec.SerializationTypePB)
	callopts := make([]client.Option, 0, len(c.opts)+len(opts))
	callopts = append(callopts, c.opts...)
	callopts = append(callopts, opts...)
	rsp := &QueryNoteRsp{}
	if err := c.client.Invoke(ctx, req, rsp, callopts...); err != nil {
		return nil, err
	}
	return rsp, nil
}

func (c *NoteServerClientProxyImpl) OperateNote(ctx context.Context, req *OperateNoteReq, opts ...client.Option) (*OperateNoteRsp, error) {
	ctx, msg := codec.WithCloneMessage(ctx)
	defer codec.PutBackMessage(msg)
	msg.WithClientRPCName("/oj.note.NoteServer/OperateNote")
	msg.WithCalleeServiceName(NoteServerServer_ServiceDesc.ServiceName)
	msg.WithCalleeApp("")
	msg.WithCalleeServer("")
	msg.WithCalleeService("NoteServer")
	msg.WithCalleeMethod("OperateNote")
	msg.WithSerializationType(codec.SerializationTypePB)
	callopts := make([]client.Option, 0, len(c.opts)+len(opts))
	callopts = append(callopts, c.opts...)
	callopts = append(callopts, opts...)
	rsp := &OperateNoteRsp{}
	if err := c.client.Invoke(ctx, req, rsp, callopts...); err != nil {
		return nil, err
	}
	return rsp, nil
}

// END ======================================= Client Service Definition ======================================= END
