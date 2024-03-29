// Code generated by protoc-gen-go.
// source: transport/task.proto
// DO NOT EDIT!

package transport

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Task struct {
	Id      string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Title   string `protobuf:"bytes,2,opt,name=title" json:"title,omitempty"`
	Content string `protobuf:"bytes,3,opt,name=content" json:"content,omitempty"`
	IdList  string `protobuf:"bytes,4,opt,name=idList" json:"idList,omitempty"`
}

func (m *Task) Reset()                    { *m = Task{} }
func (m *Task) String() string            { return proto.CompactTextString(m) }
func (*Task) ProtoMessage()               {}
func (*Task) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func (m *Task) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Task) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Task) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *Task) GetIdList() string {
	if m != nil {
		return m.IdList
	}
	return ""
}

func init() {
	proto.RegisterType((*Task)(nil), "transport.Task")
}

func init() { proto.RegisterFile("transport/task.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 126 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0x29, 0x29, 0x4a, 0xcc,
	0x2b, 0x2e, 0xc8, 0x2f, 0x2a, 0xd1, 0x2f, 0x49, 0x2c, 0xce, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9,
	0x17, 0xe2, 0x84, 0x8b, 0x2a, 0xc5, 0x71, 0xb1, 0x84, 0x24, 0x16, 0x67, 0x0b, 0xf1, 0x71, 0x31,
	0x65, 0xa6, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x31, 0x65, 0xa6, 0x08, 0x89, 0x70, 0xb1,
	0x96, 0x64, 0x96, 0xe4, 0xa4, 0x4a, 0x30, 0x81, 0x85, 0x20, 0x1c, 0x21, 0x09, 0x2e, 0xf6, 0xe4,
	0xfc, 0xbc, 0x92, 0xd4, 0xbc, 0x12, 0x09, 0x66, 0xb0, 0x38, 0x8c, 0x2b, 0x24, 0xc6, 0xc5, 0x96,
	0x99, 0xe2, 0x93, 0x59, 0x5c, 0x22, 0xc1, 0x02, 0x96, 0x80, 0xf2, 0x92, 0xd8, 0xc0, 0x36, 0x1a,
	0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x7e, 0x10, 0x08, 0x5b, 0x89, 0x00, 0x00, 0x00,
}
