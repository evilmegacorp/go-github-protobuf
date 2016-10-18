// Code generated by protoc-gen-go.
// source: team_add.proto
// DO NOT EDIT!

package github

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type TeamAddEvent struct {
	Team         *Team       `protobuf:"bytes,1,opt,name=team" json:"team,omitempty"`
	Repository   *Repository `protobuf:"bytes,2,opt,name=repository" json:"repository,omitempty"`
	Organization *User       `protobuf:"bytes,3,opt,name=organization" json:"organization,omitempty"`
	Sender       *User       `protobuf:"bytes,4,opt,name=sender" json:"sender,omitempty"`
}

func (m *TeamAddEvent) Reset()                    { *m = TeamAddEvent{} }
func (m *TeamAddEvent) String() string            { return proto.CompactTextString(m) }
func (*TeamAddEvent) ProtoMessage()               {}
func (*TeamAddEvent) Descriptor() ([]byte, []int) { return fileDescriptor34, []int{0} }

func (m *TeamAddEvent) GetTeam() *Team {
	if m != nil {
		return m.Team
	}
	return nil
}

func (m *TeamAddEvent) GetRepository() *Repository {
	if m != nil {
		return m.Repository
	}
	return nil
}

func (m *TeamAddEvent) GetOrganization() *User {
	if m != nil {
		return m.Organization
	}
	return nil
}

func (m *TeamAddEvent) GetSender() *User {
	if m != nil {
		return m.Sender
	}
	return nil
}

func init() {
	proto.RegisterType((*TeamAddEvent)(nil), "github.TeamAddEvent")
}

func init() { proto.RegisterFile("team_add.proto", fileDescriptor34) }

var fileDescriptor34 = []byte{
	// 181 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x2b, 0x49, 0x4d, 0xcc,
	0x8d, 0x4f, 0x4c, 0x49, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4b, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0x92, 0xe2, 0x2a, 0x2d, 0x4e, 0x2d, 0x82, 0x88, 0x49, 0x71, 0x81, 0xd4, 0x40, 0xd9,
	0x02, 0x45, 0xa9, 0x05, 0xf9, 0xc5, 0x99, 0x25, 0xf9, 0x45, 0x95, 0x10, 0x11, 0xa5, 0x3d, 0x8c,
	0x5c, 0x3c, 0x21, 0x40, 0x05, 0x8e, 0x29, 0x29, 0xae, 0x65, 0xa9, 0x79, 0x25, 0x42, 0x0a, 0x5c,
	0x2c, 0x20, 0x0d, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0xdc, 0x46, 0x3c, 0x7a, 0x10, 0x13, 0xf5, 0x40,
	0x6a, 0x82, 0xc0, 0x32, 0x42, 0x46, 0x5c, 0x5c, 0x08, 0x63, 0x24, 0x98, 0xc0, 0xea, 0x84, 0x60,
	0xea, 0x82, 0xe0, 0x32, 0x41, 0x48, 0xaa, 0x84, 0x0c, 0xb8, 0x78, 0xf2, 0x8b, 0xd2, 0x13, 0xf3,
	0x32, 0xab, 0x12, 0x4b, 0x32, 0xf3, 0xf3, 0x24, 0x98, 0x51, 0x4d, 0x0f, 0x05, 0x3a, 0x37, 0x08,
	0x45, 0x85, 0x90, 0x0a, 0x17, 0x5b, 0x71, 0x6a, 0x5e, 0x4a, 0x6a, 0x91, 0x04, 0x0b, 0x16, 0xb5,
	0x50, 0xb9, 0x24, 0x36, 0xb0, 0x2f, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x55, 0x94, 0x1d,
	0x65, 0x09, 0x01, 0x00, 0x00,
}
