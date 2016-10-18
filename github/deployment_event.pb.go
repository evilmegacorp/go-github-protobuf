// Code generated by protoc-gen-go.
// source: deployment_event.proto
// DO NOT EDIT!

package github

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type DeploymentEvent struct {
	Deployment *Deployment `protobuf:"bytes,1,opt,name=deployment" json:"deployment,omitempty"`
	Repository *Repository `protobuf:"bytes,2,opt,name=repository" json:"repository,omitempty"`
	Sender     *User       `protobuf:"bytes,3,opt,name=sender" json:"sender,omitempty"`
}

func (m *DeploymentEvent) Reset()                    { *m = DeploymentEvent{} }
func (m *DeploymentEvent) String() string            { return proto.CompactTextString(m) }
func (*DeploymentEvent) ProtoMessage()               {}
func (*DeploymentEvent) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{0} }

func (m *DeploymentEvent) GetDeployment() *Deployment {
	if m != nil {
		return m.Deployment
	}
	return nil
}

func (m *DeploymentEvent) GetRepository() *Repository {
	if m != nil {
		return m.Repository
	}
	return nil
}

func (m *DeploymentEvent) GetSender() *User {
	if m != nil {
		return m.Sender
	}
	return nil
}

func init() {
	proto.RegisterType((*DeploymentEvent)(nil), "github.DeploymentEvent")
}

func init() { proto.RegisterFile("deployment_event.proto", fileDescriptor7) }

var fileDescriptor7 = []byte{
	// 157 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0x4b, 0x49, 0x2d, 0xc8,
	0xc9, 0xaf, 0xcc, 0x4d, 0xcd, 0x2b, 0x89, 0x4f, 0x2d, 0x03, 0x92, 0x7a, 0x05, 0x45, 0xf9, 0x25,
	0xf9, 0x42, 0x6c, 0xe9, 0x99, 0x25, 0x19, 0xa5, 0x49, 0x52, 0x5c, 0xa5, 0xc5, 0xa9, 0x45, 0x10,
	0x31, 0x29, 0x81, 0xa2, 0xd4, 0x82, 0xfc, 0xe2, 0xcc, 0x92, 0xfc, 0xa2, 0x4a, 0x98, 0x08, 0x42,
	0x37, 0x44, 0x44, 0x69, 0x3e, 0x23, 0x17, 0xbf, 0x0b, 0x5c, 0xd0, 0x15, 0x64, 0xa2, 0x90, 0x11,
	0x17, 0x17, 0x42, 0x9d, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0xb7, 0x91, 0x90, 0x1e, 0xc4, 0x02, 0x3d,
	0x84, 0xe2, 0x20, 0x24, 0x55, 0x20, 0x3d, 0x08, 0xdb, 0x24, 0x98, 0x50, 0xf5, 0x04, 0xc1, 0x65,
	0x82, 0x90, 0x54, 0x09, 0xa9, 0x70, 0xb1, 0x15, 0xa7, 0xe6, 0xa5, 0xa4, 0x16, 0x49, 0x30, 0x83,
	0xd5, 0xf3, 0xc0, 0xd4, 0x87, 0x02, 0xfd, 0x10, 0x04, 0x95, 0x4b, 0x62, 0x03, 0x3b, 0xd4, 0x18,
	0x10, 0x00, 0x00, 0xff, 0xff, 0xbe, 0xb6, 0xed, 0xc7, 0xfa, 0x00, 0x00, 0x00,
}
