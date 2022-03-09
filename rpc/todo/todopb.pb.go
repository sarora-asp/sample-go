// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.6.1
// source: rpc/todo/todopb.proto

package todo

import (
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Todo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TaskId      int32                `protobuf:"varint,6,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
	Title       string               `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Desc        string               `protobuf:"bytes,7,opt,name=desc,proto3" json:"desc,omitempty"`
	UserId      int32                `protobuf:"varint,8,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Completed   bool                 `protobuf:"varint,2,opt,name=completed,proto3" json:"completed,omitempty"`
	CreatedAt   *timestamp.Timestamp `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt   *timestamp.Timestamp `protobuf:"bytes,9,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	CompletedAt *timestamp.Timestamp `protobuf:"bytes,4,opt,name=completed_at,json=completedAt,proto3" json:"completed_at,omitempty"`
	DueDate     *timestamp.Timestamp `protobuf:"bytes,5,opt,name=due_date,json=dueDate,proto3" json:"due_date,omitempty"`
}

func (x *Todo) Reset() {
	*x = Todo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_todo_todopb_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Todo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Todo) ProtoMessage() {}

func (x *Todo) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_todo_todopb_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Todo.ProtoReflect.Descriptor instead.
func (*Todo) Descriptor() ([]byte, []int) {
	return file_rpc_todo_todopb_proto_rawDescGZIP(), []int{0}
}

func (x *Todo) GetTaskId() int32 {
	if x != nil {
		return x.TaskId
	}
	return 0
}

func (x *Todo) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Todo) GetDesc() string {
	if x != nil {
		return x.Desc
	}
	return ""
}

func (x *Todo) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *Todo) GetCompleted() bool {
	if x != nil {
		return x.Completed
	}
	return false
}

func (x *Todo) GetCreatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Todo) GetUpdatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *Todo) GetCompletedAt() *timestamp.Timestamp {
	if x != nil {
		return x.CompletedAt
	}
	return nil
}

func (x *Todo) GetDueDate() *timestamp.Timestamp {
	if x != nil {
		return x.DueDate
	}
	return nil
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_todo_todopb_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_todo_todopb_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_rpc_todo_todopb_proto_rawDescGZIP(), []int{1}
}

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TaskId    int32                `protobuf:"varint,6,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
	Title     string               `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Desc      string               `protobuf:"bytes,7,opt,name=desc,proto3" json:"desc,omitempty"`
	Completed bool                 `protobuf:"varint,2,opt,name=completed,proto3" json:"completed,omitempty"`
	DueDate   *timestamp.Timestamp `protobuf:"bytes,5,opt,name=due_date,json=dueDate,proto3" json:"due_date,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_todo_todopb_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_todo_todopb_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_rpc_todo_todopb_proto_rawDescGZIP(), []int{2}
}

func (x *Request) GetTaskId() int32 {
	if x != nil {
		return x.TaskId
	}
	return 0
}

func (x *Request) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Request) GetDesc() string {
	if x != nil {
		return x.Desc
	}
	return ""
}

func (x *Request) GetCompleted() bool {
	if x != nil {
		return x.Completed
	}
	return false
}

func (x *Request) GetDueDate() *timestamp.Timestamp {
	if x != nil {
		return x.DueDate
	}
	return nil
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int32   `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Success bool    `protobuf:"varint,2,opt,name=success,proto3" json:"success,omitempty"`
	Msg     string  `protobuf:"bytes,3,opt,name=msg,proto3" json:"msg,omitempty"`
	Todo    *Todo   `protobuf:"bytes,4,opt,name=todo,proto3" json:"todo,omitempty"`
	Todos   []*Todo `protobuf:"bytes,5,rep,name=todos,proto3" json:"todos,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_todo_todopb_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_todo_todopb_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_rpc_todo_todopb_proto_rawDescGZIP(), []int{3}
}

func (x *Response) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *Response) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *Response) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *Response) GetTodo() *Todo {
	if x != nil {
		return x.Todo
	}
	return nil
}

func (x *Response) GetTodos() []*Todo {
	if x != nil {
		return x.Todos
	}
	return nil
}

var File_rpc_todo_todopb_proto protoreflect.FileDescriptor

var file_rpc_todo_todopb_proto_rawDesc = []byte{
	0x0a, 0x15, 0x72, 0x70, 0x63, 0x2f, 0x74, 0x6f, 0x64, 0x6f, 0x2f, 0x74, 0x6f, 0x64, 0x6f, 0x70,
	0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x74, 0x6f, 0x64, 0x6f, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xec,
	0x02, 0x0a, 0x04, 0x54, 0x6f, 0x64, 0x6f, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x61, 0x73, 0x6b, 0x5f,
	0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x74, 0x61, 0x73, 0x6b, 0x49, 0x64,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x65, 0x73, 0x63, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x65, 0x73, 0x63, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65,
	0x64, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x3d, 0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x70, 0x6c,
	0x65, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0b, 0x63, 0x6f, 0x6d, 0x70, 0x6c,
	0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x35, 0x0a, 0x08, 0x64, 0x75, 0x65, 0x5f, 0x64, 0x61,
	0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x64, 0x75, 0x65, 0x44, 0x61, 0x74, 0x65, 0x22, 0x07, 0x0a,
	0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0xa1, 0x01, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x06, 0x74, 0x61, 0x73, 0x6b, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x65, 0x73, 0x63, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x64, 0x65, 0x73, 0x63, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74,
	0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65,
	0x74, 0x65, 0x64, 0x12, 0x35, 0x0a, 0x08, 0x64, 0x75, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x07, 0x64, 0x75, 0x65, 0x44, 0x61, 0x74, 0x65, 0x22, 0x8c, 0x01, 0x0a, 0x08, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x1e, 0x0a, 0x04, 0x74, 0x6f, 0x64, 0x6f, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x54, 0x6f, 0x64,
	0x6f, 0x52, 0x04, 0x74, 0x6f, 0x64, 0x6f, 0x12, 0x20, 0x0a, 0x05, 0x74, 0x6f, 0x64, 0x6f, 0x73,
	0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x54, 0x6f,
	0x64, 0x6f, 0x52, 0x05, 0x74, 0x6f, 0x64, 0x6f, 0x73, 0x32, 0xee, 0x01, 0x0a, 0x0b, 0x54, 0x4f,
	0x64, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2a, 0x0a, 0x0a, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x54, 0x6f, 0x64, 0x6f, 0x12, 0x0a, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x54,
	0x6f, 0x64, 0x6f, 0x1a, 0x0e, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x29, 0x0a, 0x08, 0x67, 0x65, 0x74, 0x54, 0x6f, 0x64, 0x6f,
	0x73, 0x12, 0x0b, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0e,
	0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x2a, 0x0a, 0x07, 0x67, 0x65, 0x74, 0x54, 0x6f, 0x64, 0x6f, 0x12, 0x0d, 0x2e, 0x74, 0x6f,
	0x64, 0x6f, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x74, 0x6f, 0x64,
	0x6f, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x2d, 0x0a, 0x0a,
	0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x6f, 0x64, 0x6f, 0x12, 0x0d, 0x2e, 0x74, 0x6f, 0x64,
	0x6f, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x74, 0x6f, 0x64, 0x6f,
	0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x2d, 0x0a, 0x0a, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x64, 0x6f, 0x12, 0x0d, 0x2e, 0x74, 0x6f, 0x64, 0x6f,
	0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x2e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x17, 0x5a, 0x15, 0x73, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x74, 0x77, 0x69, 0x72, 0x70, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x74,
	0x6f, 0x64, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rpc_todo_todopb_proto_rawDescOnce sync.Once
	file_rpc_todo_todopb_proto_rawDescData = file_rpc_todo_todopb_proto_rawDesc
)

func file_rpc_todo_todopb_proto_rawDescGZIP() []byte {
	file_rpc_todo_todopb_proto_rawDescOnce.Do(func() {
		file_rpc_todo_todopb_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpc_todo_todopb_proto_rawDescData)
	})
	return file_rpc_todo_todopb_proto_rawDescData
}

var file_rpc_todo_todopb_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_rpc_todo_todopb_proto_goTypes = []interface{}{
	(*Todo)(nil),                // 0: todo.Todo
	(*Empty)(nil),               // 1: todo.Empty
	(*Request)(nil),             // 2: todo.Request
	(*Response)(nil),            // 3: todo.Response
	(*timestamp.Timestamp)(nil), // 4: google.protobuf.Timestamp
}
var file_rpc_todo_todopb_proto_depIdxs = []int32{
	4,  // 0: todo.Todo.created_at:type_name -> google.protobuf.Timestamp
	4,  // 1: todo.Todo.updated_at:type_name -> google.protobuf.Timestamp
	4,  // 2: todo.Todo.completed_at:type_name -> google.protobuf.Timestamp
	4,  // 3: todo.Todo.due_date:type_name -> google.protobuf.Timestamp
	4,  // 4: todo.Request.due_date:type_name -> google.protobuf.Timestamp
	0,  // 5: todo.Response.todo:type_name -> todo.Todo
	0,  // 6: todo.Response.todos:type_name -> todo.Todo
	0,  // 7: todo.TOdoService.createTodo:input_type -> todo.Todo
	1,  // 8: todo.TOdoService.getTodos:input_type -> todo.Empty
	2,  // 9: todo.TOdoService.getTodo:input_type -> todo.Request
	2,  // 10: todo.TOdoService.deleteTodo:input_type -> todo.Request
	2,  // 11: todo.TOdoService.updateTodo:input_type -> todo.Request
	3,  // 12: todo.TOdoService.createTodo:output_type -> todo.Response
	3,  // 13: todo.TOdoService.getTodos:output_type -> todo.Response
	3,  // 14: todo.TOdoService.getTodo:output_type -> todo.Response
	3,  // 15: todo.TOdoService.deleteTodo:output_type -> todo.Response
	3,  // 16: todo.TOdoService.updateTodo:output_type -> todo.Response
	12, // [12:17] is the sub-list for method output_type
	7,  // [7:12] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_rpc_todo_todopb_proto_init() }
func file_rpc_todo_todopb_proto_init() {
	if File_rpc_todo_todopb_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rpc_todo_todopb_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Todo); i {
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
		file_rpc_todo_todopb_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_rpc_todo_todopb_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
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
		file_rpc_todo_todopb_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
			RawDescriptor: file_rpc_todo_todopb_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_rpc_todo_todopb_proto_goTypes,
		DependencyIndexes: file_rpc_todo_todopb_proto_depIdxs,
		MessageInfos:      file_rpc_todo_todopb_proto_msgTypes,
	}.Build()
	File_rpc_todo_todopb_proto = out.File
	file_rpc_todo_todopb_proto_rawDesc = nil
	file_rpc_todo_todopb_proto_goTypes = nil
	file_rpc_todo_todopb_proto_depIdxs = nil
}