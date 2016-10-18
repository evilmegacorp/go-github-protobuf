// Code generated by protoc-gen-go.
// source: issue.proto
// DO NOT EDIT!

package github

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Label struct {
	Url   string `protobuf:"bytes,1,opt,name=url" json:"url,omitempty"`
	Name  string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Color string `protobuf:"bytes,3,opt,name=color" json:"color,omitempty"`
}

func (m *Label) Reset()                    { *m = Label{} }
func (m *Label) String() string            { return proto.CompactTextString(m) }
func (*Label) ProtoMessage()               {}
func (*Label) Descriptor() ([]byte, []int) { return fileDescriptor15, []int{0} }

type Issue struct {
	Id            int32      `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Url           string     `protobuf:"bytes,2,opt,name=url" json:"url,omitempty"`
	RepositoryUrl string     `protobuf:"bytes,3,opt,name=repository_url,json=repositoryUrl" json:"repository_url,omitempty"`
	LabelsUrl     string     `protobuf:"bytes,4,opt,name=labels_url,json=labelsUrl" json:"labels_url,omitempty"`
	CommentsUrl   string     `protobuf:"bytes,5,opt,name=comments_url,json=commentsUrl" json:"comments_url,omitempty"`
	EventsUrl     string     `protobuf:"bytes,6,opt,name=events_url,json=eventsUrl" json:"events_url,omitempty"`
	HtmlUrl       string     `protobuf:"bytes,7,opt,name=html_url,json=htmlUrl" json:"html_url,omitempty"`
	Number        int32      `protobuf:"varint,8,opt,name=number" json:"number,omitempty"`
	State         string     `protobuf:"bytes,9,opt,name=state" json:"state,omitempty"`
	Title         string     `protobuf:"bytes,10,opt,name=title" json:"title,omitempty"`
	Body          string     `protobuf:"bytes,11,opt,name=body" json:"body,omitempty"`
	User          *User      `protobuf:"bytes,12,opt,name=user" json:"user,omitempty"`
	Labels        []*Label   `protobuf:"bytes,13,rep,name=labels" json:"labels,omitempty"`
	Assignee      *User      `protobuf:"bytes,14,opt,name=assignee" json:"assignee,omitempty"`
	Milestone     *Milestone `protobuf:"bytes,15,opt,name=milestone" json:"milestone,omitempty"`
	Locked        bool       `protobuf:"varint,16,opt,name=locked" json:"locked,omitempty"`
	Comments      int32      `protobuf:"varint,17,opt,name=comments" json:"comments,omitempty"`
	ClosedAt      string     `protobuf:"bytes,18,opt,name=closed_at,json=closedAt" json:"closed_at,omitempty"`
	CreatedAt     string     `protobuf:"bytes,19,opt,name=created_at,json=createdAt" json:"created_at,omitempty"`
	UpdatedAt     string     `protobuf:"bytes,20,opt,name=updated_at,json=updatedAt" json:"updated_at,omitempty"`
	ClosedBy      *User      `protobuf:"bytes,21,opt,name=closed_by,json=closedBy" json:"closed_by,omitempty"`
}

func (m *Issue) Reset()                    { *m = Issue{} }
func (m *Issue) String() string            { return proto.CompactTextString(m) }
func (*Issue) ProtoMessage()               {}
func (*Issue) Descriptor() ([]byte, []int) { return fileDescriptor15, []int{1} }

func (m *Issue) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *Issue) GetLabels() []*Label {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *Issue) GetAssignee() *User {
	if m != nil {
		return m.Assignee
	}
	return nil
}

func (m *Issue) GetMilestone() *Milestone {
	if m != nil {
		return m.Milestone
	}
	return nil
}

func (m *Issue) GetClosedBy() *User {
	if m != nil {
		return m.ClosedBy
	}
	return nil
}

func init() {
	proto.RegisterType((*Label)(nil), "github.Label")
	proto.RegisterType((*Issue)(nil), "github.Issue")
}

func init() { proto.RegisterFile("issue.proto", fileDescriptor15) }

var fileDescriptor15 = []byte{
	// 414 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0x92, 0x4f, 0x6f, 0xd3, 0x40,
	0x10, 0xc5, 0xd5, 0x24, 0x76, 0xed, 0x71, 0x92, 0xb6, 0x4b, 0x41, 0x4b, 0x10, 0x52, 0xa8, 0x84,
	0x54, 0x2e, 0x41, 0x82, 0x4f, 0x00, 0x9c, 0x90, 0xe0, 0x62, 0xa9, 0xe7, 0xca, 0x7f, 0x46, 0x8d,
	0xc5, 0xda, 0x1b, 0xed, 0xae, 0x91, 0x7c, 0xe5, 0x93, 0xb3, 0x3b, 0x63, 0x3b, 0x97, 0xdc, 0x76,
	0xde, 0xef, 0xe9, 0x25, 0x6f, 0xc6, 0x90, 0x35, 0xd6, 0xf6, 0x78, 0x38, 0x19, 0xed, 0xb4, 0x88,
	0x5f, 0x1a, 0x77, 0xec, 0xcb, 0x1d, 0xf4, 0x16, 0x0d, 0x6b, 0xbb, 0x9b, 0xb6, 0x51, 0x68, 0x9d,
	0xee, 0x46, 0xd3, 0xc3, 0x0f, 0x88, 0x7e, 0x15, 0x25, 0x2a, 0x71, 0x0b, 0xcb, 0xde, 0x28, 0x79,
	0xb5, 0xbf, 0x7a, 0x4c, 0xf3, 0xf0, 0x14, 0x02, 0x56, 0x5d, 0xd1, 0xa2, 0x5c, 0x90, 0x44, 0x6f,
	0x71, 0x0f, 0x51, 0xa5, 0x95, 0x36, 0x72, 0x49, 0x22, 0x0f, 0x0f, 0xff, 0x22, 0x88, 0x7e, 0x86,
	0x5f, 0x16, 0x5b, 0x58, 0x34, 0x35, 0x85, 0x44, 0xb9, 0x7f, 0x4d, 0xa9, 0x8b, 0x73, 0xea, 0x47,
	0xd8, 0x1a, 0x3c, 0x69, 0xdb, 0x38, 0x6d, 0x86, 0xe7, 0x00, 0x39, 0x6a, 0x73, 0x56, 0x9f, 0xbc,
	0xed, 0x3d, 0x80, 0x0a, 0xff, 0xcb, 0x92, 0x65, 0x45, 0x96, 0x94, 0x95, 0x80, 0x3f, 0xc0, 0xba,
	0xd2, 0x6d, 0x8b, 0x9d, 0x63, 0x43, 0x44, 0x86, 0x6c, 0xd2, 0xc6, 0x04, 0xfc, 0x3b, 0x1b, 0x62,
	0x4e, 0x60, 0x25, 0xe0, 0xb7, 0x90, 0x1c, 0x5d, 0xab, 0x08, 0x5e, 0x13, 0xbc, 0x0e, 0x73, 0x40,
	0x6f, 0x20, 0xee, 0xfa, 0xb6, 0x44, 0x23, 0x13, 0x2a, 0x32, 0x4e, 0xa1, 0xbc, 0x75, 0x85, 0x43,
	0x99, 0x72, 0x79, 0x1a, 0x82, 0xea, 0x1a, 0xa7, 0x50, 0x02, 0xab, 0x34, 0x84, 0xe5, 0x95, 0xba,
	0x1e, 0x64, 0xc6, 0xcb, 0x0b, 0x6f, 0xb1, 0x87, 0x55, 0x38, 0x85, 0x5c, 0x7b, 0x2d, 0xfb, 0xb2,
	0x3e, 0xf0, 0x7d, 0x0e, 0x4f, 0x5e, 0xcb, 0x89, 0xf8, 0xe5, 0xc4, 0xdc, 0x51, 0x6e, 0xf6, 0x4b,
	0xef, 0xd9, 0x4c, 0x1e, 0xba, 0x51, 0x3e, 0x42, 0xf1, 0x08, 0x49, 0x61, 0x6d, 0xf3, 0xd2, 0x21,
	0xca, 0xed, 0x85, 0xb0, 0x99, 0x8a, 0xcf, 0x90, 0xce, 0x17, 0x97, 0x37, 0x64, 0xbd, 0x9b, 0xac,
	0xbf, 0x27, 0x90, 0x9f, 0x3d, 0xa1, 0xbb, 0xd2, 0xd5, 0x1f, 0xac, 0xe5, 0xad, 0x77, 0x27, 0xf9,
	0x38, 0x89, 0x1d, 0x24, 0xd3, 0x72, 0xe5, 0x1d, 0x6d, 0x65, 0x9e, 0xc5, 0x3b, 0x48, 0x2b, 0xa5,
	0x2d, 0xd6, 0xcf, 0x85, 0x93, 0x82, 0x0a, 0x27, 0x2c, 0x7c, 0x73, 0xe1, 0x0c, 0x95, 0x41, 0xbf,
	0x28, 0xa2, 0xaf, 0xf8, 0x0c, 0xa3, 0xc2, 0xb8, 0x3f, 0xd5, 0x13, 0xbe, 0x67, 0x3c, 0x2a, 0x1e,
	0x7f, 0x9a, 0xa3, 0xcb, 0x41, 0xbe, 0xbe, 0x54, 0x95, 0xf1, 0xf7, 0xa1, 0x8c, 0xe9, 0x83, 0xfe,
	0xfa, 0x3f, 0x00, 0x00, 0xff, 0xff, 0x8d, 0x6d, 0xb5, 0x04, 0x04, 0x03, 0x00, 0x00,
}
