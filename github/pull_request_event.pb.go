// Code generated by protoc-gen-go.
// source: pull_request_event.proto
// DO NOT EDIT!

package github

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type PullRequestEvent struct {
	Action      string       `protobuf:"bytes,1,opt,name=action" json:"action,omitempty"`
	Number      int32        `protobuf:"varint,2,opt,name=number" json:"number,omitempty"`
	PullRequest *PullRequest `protobuf:"bytes,3,opt,name=pull_request,json=pullRequest" json:"pull_request,omitempty"`
	Repository  *Repository  `protobuf:"bytes,4,opt,name=repository" json:"repository,omitempty"`
	Sender      *User        `protobuf:"bytes,5,opt,name=sender" json:"sender,omitempty"`
}

func (m *PullRequestEvent) Reset()                    { *m = PullRequestEvent{} }
func (m *PullRequestEvent) String() string            { return proto.CompactTextString(m) }
func (*PullRequestEvent) ProtoMessage()               {}
func (*PullRequestEvent) Descriptor() ([]byte, []int) { return fileDescriptor23, []int{0} }

func (m *PullRequestEvent) GetPullRequest() *PullRequest {
	if m != nil {
		return m.PullRequest
	}
	return nil
}

func (m *PullRequestEvent) GetRepository() *Repository {
	if m != nil {
		return m.Repository
	}
	return nil
}

func (m *PullRequestEvent) GetSender() *User {
	if m != nil {
		return m.Sender
	}
	return nil
}

func init() {
	proto.RegisterType((*PullRequestEvent)(nil), "github.PullRequestEvent")
}

func init() { proto.RegisterFile("pull_request_event.proto", fileDescriptor23) }

var fileDescriptor23 = []byte{
	// 206 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x92, 0x28, 0x28, 0xcd, 0xc9,
	0x89, 0x2f, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x89, 0x4f, 0x2d, 0x4b, 0xcd, 0x2b, 0xd1, 0x2b,
	0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4b, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0x92, 0xe2, 0x2a, 0x2d,
	0x4e, 0x2d, 0x82, 0x88, 0x49, 0x09, 0x14, 0xa5, 0x16, 0xe4, 0x17, 0x67, 0x96, 0xe4, 0x17, 0x55,
	0x42, 0x45, 0x04, 0x41, 0xfa, 0xa1, 0xda, 0x21, 0x42, 0x4a, 0x57, 0x18, 0xb9, 0x04, 0x02, 0x80,
	0xa2, 0x41, 0x10, 0x51, 0x57, 0x90, 0x99, 0x42, 0x62, 0x5c, 0x6c, 0x89, 0xc9, 0x25, 0x99, 0xf9,
	0x79, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x50, 0x1e, 0x48, 0x3c, 0xaf, 0x34, 0x37, 0x29,
	0xb5, 0x48, 0x82, 0x09, 0x28, 0xce, 0x1a, 0x04, 0xe5, 0x09, 0x99, 0x71, 0xf1, 0x20, 0xbb, 0x4c,
	0x82, 0x19, 0x28, 0xcb, 0x6d, 0x24, 0xac, 0x07, 0x71, 0x94, 0x1e, 0x92, 0xf9, 0x41, 0xdc, 0x05,
	0x08, 0x8e, 0x90, 0x11, 0x17, 0x17, 0xc2, 0x8d, 0x12, 0x2c, 0x60, 0x5d, 0x42, 0x30, 0x5d, 0x41,
	0x70, 0x99, 0x20, 0x24, 0x55, 0x42, 0x2a, 0x5c, 0x6c, 0xc5, 0xa9, 0x79, 0x29, 0x40, 0x37, 0xb0,
	0x82, 0xd5, 0xf3, 0xc0, 0xd4, 0x87, 0x02, 0x7d, 0x1e, 0x04, 0x95, 0x4b, 0x62, 0x03, 0xfb, 0xce,
	0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xc2, 0xaf, 0x5f, 0x4c, 0x32, 0x01, 0x00, 0x00,
}
