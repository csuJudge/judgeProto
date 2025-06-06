// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v5.29.3
// source: blog.proto

package blog

import (
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type QueryBlogPageSizeRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string  `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"` // 返回信息
	Code    int32   `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"`      // 返回码
	Blogs   []*Blog `protobuf:"bytes,3,rep,name=blogs,proto3" json:"blogs,omitempty"`     // 博客的集合
	Total   int32   `protobuf:"varint,4,opt,name=total,proto3" json:"total,omitempty"`    // 总数
}

func (x *QueryBlogPageSizeRsp) Reset() {
	*x = QueryBlogPageSizeRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blog_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryBlogPageSizeRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryBlogPageSizeRsp) ProtoMessage() {}

func (x *QueryBlogPageSizeRsp) ProtoReflect() protoreflect.Message {
	mi := &file_blog_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryBlogPageSizeRsp.ProtoReflect.Descriptor instead.
func (*QueryBlogPageSizeRsp) Descriptor() ([]byte, []int) {
	return file_blog_proto_rawDescGZIP(), []int{0}
}

func (x *QueryBlogPageSizeRsp) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *QueryBlogPageSizeRsp) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *QueryBlogPageSizeRsp) GetBlogs() []*Blog {
	if x != nil {
		return x.Blogs
	}
	return nil
}

func (x *QueryBlogPageSizeRsp) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

type QueryBlogReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BlogID int32 `protobuf:"varint,1,opt,name=blogID,proto3" json:"blogID,omitempty"` // 博客的主键
	UserID int32 `protobuf:"varint,2,opt,name=userID,proto3" json:"userID,omitempty"` // 用户编号
}

func (x *QueryBlogReq) Reset() {
	*x = QueryBlogReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blog_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryBlogReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryBlogReq) ProtoMessage() {}

func (x *QueryBlogReq) ProtoReflect() protoreflect.Message {
	mi := &file_blog_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryBlogReq.ProtoReflect.Descriptor instead.
func (*QueryBlogReq) Descriptor() ([]byte, []int) {
	return file_blog_proto_rawDescGZIP(), []int{1}
}

func (x *QueryBlogReq) GetBlogID() int32 {
	if x != nil {
		return x.BlogID
	}
	return 0
}

func (x *QueryBlogReq) GetUserID() int32 {
	if x != nil {
		return x.UserID
	}
	return 0
}

type QueryBlogRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"` // 返回信息
	Code    int32  `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"`      // 返回码
	Blog    *Blog  `protobuf:"bytes,3,opt,name=blog,proto3" json:"blog,omitempty"`       // 博客
}

func (x *QueryBlogRsp) Reset() {
	*x = QueryBlogRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blog_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryBlogRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryBlogRsp) ProtoMessage() {}

func (x *QueryBlogRsp) ProtoReflect() protoreflect.Message {
	mi := &file_blog_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryBlogRsp.ProtoReflect.Descriptor instead.
func (*QueryBlogRsp) Descriptor() ([]byte, []int) {
	return file_blog_proto_rawDescGZIP(), []int{2}
}

func (x *QueryBlogRsp) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *QueryBlogRsp) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *QueryBlogRsp) GetBlog() *Blog {
	if x != nil {
		return x.Blog
	}
	return nil
}

type Blog struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BlogID      int32  `protobuf:"varint,1,opt,name=blogID,proto3" json:"blogID,omitempty"`          // 博客主键
	BlogName    string `protobuf:"bytes,2,opt,name=blogName,proto3" json:"blogName,omitempty"`       // 博客名称
	BlogContent string `protobuf:"bytes,3,opt,name=blogContent,proto3" json:"blogContent,omitempty"` // 博客内容
	Creator     string `protobuf:"bytes,4,opt,name=creator,proto3" json:"creator,omitempty"`         // 创建者
	CreatedTime string `protobuf:"bytes,5,opt,name=createdTime,proto3" json:"createdTime,omitempty"` // 创建时间
	Enable      int32  `protobuf:"varint,6,opt,name=enable,proto3" json:"enable,omitempty"`          // 是否可用
}

func (x *Blog) Reset() {
	*x = Blog{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blog_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Blog) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Blog) ProtoMessage() {}

func (x *Blog) ProtoReflect() protoreflect.Message {
	mi := &file_blog_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Blog.ProtoReflect.Descriptor instead.
func (*Blog) Descriptor() ([]byte, []int) {
	return file_blog_proto_rawDescGZIP(), []int{3}
}

func (x *Blog) GetBlogID() int32 {
	if x != nil {
		return x.BlogID
	}
	return 0
}

func (x *Blog) GetBlogName() string {
	if x != nil {
		return x.BlogName
	}
	return ""
}

func (x *Blog) GetBlogContent() string {
	if x != nil {
		return x.BlogContent
	}
	return ""
}

func (x *Blog) GetCreator() string {
	if x != nil {
		return x.Creator
	}
	return ""
}

func (x *Blog) GetCreatedTime() string {
	if x != nil {
		return x.CreatedTime
	}
	return ""
}

func (x *Blog) GetEnable() int32 {
	if x != nil {
		return x.Enable
	}
	return 0
}

type QueryBlogPageSizeReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page   int32 `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`     // 页
	Limit  int32 `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`   // 页大小
	UserID int32 `protobuf:"varint,3,opt,name=userID,proto3" json:"userID,omitempty"` // 用户主键
}

func (x *QueryBlogPageSizeReq) Reset() {
	*x = QueryBlogPageSizeReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blog_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryBlogPageSizeReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryBlogPageSizeReq) ProtoMessage() {}

func (x *QueryBlogPageSizeReq) ProtoReflect() protoreflect.Message {
	mi := &file_blog_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryBlogPageSizeReq.ProtoReflect.Descriptor instead.
func (*QueryBlogPageSizeReq) Descriptor() ([]byte, []int) {
	return file_blog_proto_rawDescGZIP(), []int{4}
}

func (x *QueryBlogPageSizeReq) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *QueryBlogPageSizeReq) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *QueryBlogPageSizeReq) GetUserID() int32 {
	if x != nil {
		return x.UserID
	}
	return 0
}

type UpdateBlogStatusReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID int32 `protobuf:"varint,1,opt,name=userID,proto3" json:"userID,omitempty"` // 用户主键
	BlogID int32 `protobuf:"varint,2,opt,name=blogID,proto3" json:"blogID,omitempty"` // 博客主键
	Status int32 `protobuf:"varint,3,opt,name=status,proto3" json:"status,omitempty"` // 状态
}

func (x *UpdateBlogStatusReq) Reset() {
	*x = UpdateBlogStatusReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blog_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateBlogStatusReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateBlogStatusReq) ProtoMessage() {}

func (x *UpdateBlogStatusReq) ProtoReflect() protoreflect.Message {
	mi := &file_blog_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateBlogStatusReq.ProtoReflect.Descriptor instead.
func (*UpdateBlogStatusReq) Descriptor() ([]byte, []int) {
	return file_blog_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateBlogStatusReq) GetUserID() int32 {
	if x != nil {
		return x.UserID
	}
	return 0
}

func (x *UpdateBlogStatusReq) GetBlogID() int32 {
	if x != nil {
		return x.BlogID
	}
	return 0
}

func (x *UpdateBlogStatusReq) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

type AddBlogReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`       // 博客名
	Content string `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"` // 内容
	UserID  int32  `protobuf:"varint,3,opt,name=userID,proto3" json:"userID,omitempty"`  // 用户编号
}

func (x *AddBlogReq) Reset() {
	*x = AddBlogReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blog_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddBlogReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddBlogReq) ProtoMessage() {}

func (x *AddBlogReq) ProtoReflect() protoreflect.Message {
	mi := &file_blog_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddBlogReq.ProtoReflect.Descriptor instead.
func (*AddBlogReq) Descriptor() ([]byte, []int) {
	return file_blog_proto_rawDescGZIP(), []int{6}
}

func (x *AddBlogReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AddBlogReq) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *AddBlogReq) GetUserID() int32 {
	if x != nil {
		return x.UserID
	}
	return 0
}

type UpdateBlogReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID  int32  `protobuf:"varint,1,opt,name=userID,proto3" json:"userID,omitempty"`  // 用户主键
	BlogID  int32  `protobuf:"varint,2,opt,name=blogID,proto3" json:"blogID,omitempty"`  // 博客主键
	Name    string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`       // 名称
	Content string `protobuf:"bytes,4,opt,name=content,proto3" json:"content,omitempty"` // 内容
}

func (x *UpdateBlogReq) Reset() {
	*x = UpdateBlogReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blog_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateBlogReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateBlogReq) ProtoMessage() {}

func (x *UpdateBlogReq) ProtoReflect() protoreflect.Message {
	mi := &file_blog_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateBlogReq.ProtoReflect.Descriptor instead.
func (*UpdateBlogReq) Descriptor() ([]byte, []int) {
	return file_blog_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateBlogReq) GetUserID() int32 {
	if x != nil {
		return x.UserID
	}
	return 0
}

func (x *UpdateBlogReq) GetBlogID() int32 {
	if x != nil {
		return x.BlogID
	}
	return 0
}

func (x *UpdateBlogReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateBlogReq) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type CommonRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"` // 返回信息
	Code    int32  `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"`      // 返回码
}

func (x *CommonRsp) Reset() {
	*x = CommonRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blog_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommonRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommonRsp) ProtoMessage() {}

func (x *CommonRsp) ProtoReflect() protoreflect.Message {
	mi := &file_blog_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommonRsp.ProtoReflect.Descriptor instead.
func (*CommonRsp) Descriptor() ([]byte, []int) {
	return file_blog_proto_rawDescGZIP(), []int{8}
}

func (x *CommonRsp) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *CommonRsp) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

var File_blog_proto protoreflect.FileDescriptor

var file_blog_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x62, 0x6c, 0x6f, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x6f, 0x6a,
	0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x22, 0x7f, 0x0a, 0x14, 0x51, 0x75, 0x65, 0x72, 0x79, 0x42, 0x6c,
	0x6f, 0x67, 0x50, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x52, 0x73, 0x70, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x23, 0x0a, 0x05, 0x62,
	0x6c, 0x6f, 0x67, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x6f, 0x6a, 0x2e,
	0x62, 0x6c, 0x6f, 0x67, 0x2e, 0x42, 0x6c, 0x6f, 0x67, 0x52, 0x05, 0x62, 0x6c, 0x6f, 0x67, 0x73,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0x3e, 0x0a, 0x0c, 0x51, 0x75, 0x65, 0x72, 0x79, 0x42,
	0x6c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x6c, 0x6f, 0x67, 0x49, 0x44,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x62, 0x6c, 0x6f, 0x67, 0x49, 0x44, 0x12, 0x16,
	0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x22, 0x5f, 0x0a, 0x0c, 0x51, 0x75, 0x65, 0x72, 0x79, 0x42,
	0x6c, 0x6f, 0x67, 0x52, 0x73, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x12, 0x21, 0x0a, 0x04, 0x62, 0x6c, 0x6f, 0x67, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x6f, 0x6a, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x2e, 0x42, 0x6c, 0x6f,
	0x67, 0x52, 0x04, 0x62, 0x6c, 0x6f, 0x67, 0x22, 0xb0, 0x01, 0x0a, 0x04, 0x42, 0x6c, 0x6f, 0x67,
	0x12, 0x16, 0x0a, 0x06, 0x62, 0x6c, 0x6f, 0x67, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x06, 0x62, 0x6c, 0x6f, 0x67, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x62, 0x6c, 0x6f, 0x67,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x62, 0x6c, 0x6f, 0x67,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x62, 0x6c, 0x6f, 0x67, 0x43, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x62, 0x6c, 0x6f, 0x67, 0x43,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f,
	0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72,
	0x12, 0x20, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x54, 0x69,
	0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x22, 0x58, 0x0a, 0x14, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x42, 0x6c, 0x6f, 0x67, 0x50, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x52,
	0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x44, 0x22, 0x5d, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x6c,
	0x6f, 0x67, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x6c, 0x6f, 0x67, 0x49, 0x44, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x06, 0x62, 0x6c, 0x6f, 0x67, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x22, 0x52, 0x0a, 0x0a, 0x41, 0x64, 0x64, 0x42, 0x6c, 0x6f, 0x67, 0x52, 0x65,
	0x71, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x22, 0x6d, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x42, 0x6c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44,
	0x12, 0x16, 0x0a, 0x06, 0x62, 0x6c, 0x6f, 0x67, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x06, 0x62, 0x6c, 0x6f, 0x67, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x39, 0x0a, 0x09, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x52, 0x73, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x32, 0xf7, 0x03, 0x0a, 0x0a, 0x42, 0x6c, 0x6f, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x12, 0x34, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x42, 0x6c, 0x6f, 0x67, 0x12, 0x13, 0x2e, 0x6f, 0x6a,
	0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x2e, 0x41, 0x64, 0x64, 0x42, 0x6c, 0x6f, 0x67, 0x52, 0x65, 0x71,
	0x1a, 0x12, 0x2e, 0x6f, 0x6a, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x52, 0x73, 0x70, 0x22, 0x00, 0x12, 0x3a, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x42, 0x6c, 0x6f, 0x67, 0x12, 0x16, 0x2e, 0x6f, 0x6a, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x6c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x1a, 0x12, 0x2e, 0x6f,
	0x6a, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x73, 0x70,
	0x22, 0x00, 0x12, 0x46, 0x0a, 0x10, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x6c, 0x6f, 0x67,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1c, 0x2e, 0x6f, 0x6a, 0x2e, 0x62, 0x6c, 0x6f, 0x67,
	0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x6c, 0x6f, 0x67, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x65, 0x71, 0x1a, 0x12, 0x2e, 0x6f, 0x6a, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x2e, 0x43,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x73, 0x70, 0x22, 0x00, 0x12, 0x4d, 0x0a, 0x0b, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x4d, 0x79, 0x42, 0x6c, 0x6f, 0x67, 0x12, 0x1d, 0x2e, 0x6f, 0x6a, 0x2e, 0x62,
	0x6c, 0x6f, 0x67, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x42, 0x6c, 0x6f, 0x67, 0x50, 0x61, 0x67,
	0x65, 0x53, 0x69, 0x7a, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x1d, 0x2e, 0x6f, 0x6a, 0x2e, 0x62, 0x6c,
	0x6f, 0x67, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x42, 0x6c, 0x6f, 0x67, 0x50, 0x61, 0x67, 0x65,
	0x53, 0x69, 0x7a, 0x65, 0x52, 0x73, 0x70, 0x22, 0x00, 0x12, 0x53, 0x0a, 0x11, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x42, 0x6c, 0x6f, 0x67, 0x50, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1d,
	0x2e, 0x6f, 0x6a, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x42, 0x6c,
	0x6f, 0x67, 0x50, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x1d, 0x2e,
	0x6f, 0x6a, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x42, 0x6c, 0x6f,
	0x67, 0x50, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x52, 0x73, 0x70, 0x22, 0x00, 0x12, 0x3b,
	0x0a, 0x09, 0x51, 0x75, 0x65, 0x72, 0x79, 0x42, 0x6c, 0x6f, 0x67, 0x12, 0x15, 0x2e, 0x6f, 0x6a,
	0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x42, 0x6c, 0x6f, 0x67, 0x52,
	0x65, 0x71, 0x1a, 0x15, 0x2e, 0x6f, 0x6a, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x2e, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x42, 0x6c, 0x6f, 0x67, 0x52, 0x73, 0x70, 0x22, 0x00, 0x12, 0x4e, 0x0a, 0x0c, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x41, 0x6c, 0x6c, 0x42, 0x6c, 0x6f, 0x67, 0x12, 0x1d, 0x2e, 0x6f, 0x6a,
	0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x42, 0x6c, 0x6f, 0x67, 0x50,
	0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x1d, 0x2e, 0x6f, 0x6a, 0x2e,
	0x62, 0x6c, 0x6f, 0x67, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x42, 0x6c, 0x6f, 0x67, 0x50, 0x61,
	0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x52, 0x73, 0x70, 0x22, 0x00, 0x42, 0x25, 0x5a, 0x23, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x73, 0x75, 0x4a, 0x75, 0x64,
	0x67, 0x65, 0x2f, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x6c,
	0x6f, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_blog_proto_rawDescOnce sync.Once
	file_blog_proto_rawDescData = file_blog_proto_rawDesc
)

func file_blog_proto_rawDescGZIP() []byte {
	file_blog_proto_rawDescOnce.Do(func() {
		file_blog_proto_rawDescData = protoimpl.X.CompressGZIP(file_blog_proto_rawDescData)
	})
	return file_blog_proto_rawDescData
}

var file_blog_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_blog_proto_goTypes = []interface{}{
	(*QueryBlogPageSizeRsp)(nil), // 0: oj.blog.QueryBlogPageSizeRsp
	(*QueryBlogReq)(nil),         // 1: oj.blog.QueryBlogReq
	(*QueryBlogRsp)(nil),         // 2: oj.blog.QueryBlogRsp
	(*Blog)(nil),                 // 3: oj.blog.Blog
	(*QueryBlogPageSizeReq)(nil), // 4: oj.blog.QueryBlogPageSizeReq
	(*UpdateBlogStatusReq)(nil),  // 5: oj.blog.UpdateBlogStatusReq
	(*AddBlogReq)(nil),           // 6: oj.blog.AddBlogReq
	(*UpdateBlogReq)(nil),        // 7: oj.blog.UpdateBlogReq
	(*CommonRsp)(nil),            // 8: oj.blog.CommonRsp
}
var file_blog_proto_depIdxs = []int32{
	3, // 0: oj.blog.QueryBlogPageSizeRsp.blogs:type_name -> oj.blog.Blog
	3, // 1: oj.blog.QueryBlogRsp.blog:type_name -> oj.blog.Blog
	6, // 2: oj.blog.BlogServer.AddBlog:input_type -> oj.blog.AddBlogReq
	7, // 3: oj.blog.BlogServer.UpdateBlog:input_type -> oj.blog.UpdateBlogReq
	5, // 4: oj.blog.BlogServer.UpdateBlogStatus:input_type -> oj.blog.UpdateBlogStatusReq
	4, // 5: oj.blog.BlogServer.QueryMyBlog:input_type -> oj.blog.QueryBlogPageSizeReq
	4, // 6: oj.blog.BlogServer.QueryBlogPageSize:input_type -> oj.blog.QueryBlogPageSizeReq
	1, // 7: oj.blog.BlogServer.QueryBlog:input_type -> oj.blog.QueryBlogReq
	4, // 8: oj.blog.BlogServer.QueryAllBlog:input_type -> oj.blog.QueryBlogPageSizeReq
	8, // 9: oj.blog.BlogServer.AddBlog:output_type -> oj.blog.CommonRsp
	8, // 10: oj.blog.BlogServer.UpdateBlog:output_type -> oj.blog.CommonRsp
	8, // 11: oj.blog.BlogServer.UpdateBlogStatus:output_type -> oj.blog.CommonRsp
	0, // 12: oj.blog.BlogServer.QueryMyBlog:output_type -> oj.blog.QueryBlogPageSizeRsp
	0, // 13: oj.blog.BlogServer.QueryBlogPageSize:output_type -> oj.blog.QueryBlogPageSizeRsp
	2, // 14: oj.blog.BlogServer.QueryBlog:output_type -> oj.blog.QueryBlogRsp
	0, // 15: oj.blog.BlogServer.QueryAllBlog:output_type -> oj.blog.QueryBlogPageSizeRsp
	9, // [9:16] is the sub-list for method output_type
	2, // [2:9] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_blog_proto_init() }
func file_blog_proto_init() {
	if File_blog_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_blog_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryBlogPageSizeRsp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_blog_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryBlogReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_blog_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryBlogRsp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_blog_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Blog); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_blog_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryBlogPageSizeReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_blog_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateBlogStatusReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_blog_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddBlogReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_blog_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateBlogReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_blog_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommonRsp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_blog_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_blog_proto_goTypes,
		DependencyIndexes: file_blog_proto_depIdxs,
		MessageInfos:      file_blog_proto_msgTypes,
	}.Build()
	File_blog_proto = out.File
	file_blog_proto_rawDesc = nil
	file_blog_proto_goTypes = nil
	file_blog_proto_depIdxs = nil
}
