// Code generated by protoc-gen-go.
// source: pull_request_review_comment_event.proto
// DO NOT EDIT!

package github

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type PullRequestReviewCommentLink struct {
	Html string `protobuf:"bytes,1,opt,name=html" json:"html,omitempty"`
}

func (m *PullRequestReviewCommentLink) Reset()                    { *m = PullRequestReviewCommentLink{} }
func (m *PullRequestReviewCommentLink) String() string            { return proto.CompactTextString(m) }
func (*PullRequestReviewCommentLink) ProtoMessage()               {}
func (*PullRequestReviewCommentLink) Descriptor() ([]byte, []int) { return fileDescriptor24, []int{0} }

type PullRequestReviewCommentLinks struct {
	Self        *PullRequestReviewCommentLink `protobuf:"bytes,1,opt,name=self" json:"self,omitempty"`
	Html        *PullRequestReviewCommentLink `protobuf:"bytes,2,opt,name=html" json:"html,omitempty"`
	PullRequest *PullRequestReviewCommentLink `protobuf:"bytes,3,opt,name=pull_request,json=pullRequest" json:"pull_request,omitempty"`
}

func (m *PullRequestReviewCommentLinks) Reset()                    { *m = PullRequestReviewCommentLinks{} }
func (m *PullRequestReviewCommentLinks) String() string            { return proto.CompactTextString(m) }
func (*PullRequestReviewCommentLinks) ProtoMessage()               {}
func (*PullRequestReviewCommentLinks) Descriptor() ([]byte, []int) { return fileDescriptor24, []int{1} }

func (m *PullRequestReviewCommentLinks) GetSelf() *PullRequestReviewCommentLink {
	if m != nil {
		return m.Self
	}
	return nil
}

func (m *PullRequestReviewCommentLinks) GetHtml() *PullRequestReviewCommentLink {
	if m != nil {
		return m.Html
	}
	return nil
}

func (m *PullRequestReviewCommentLinks) GetPullRequest() *PullRequestReviewCommentLink {
	if m != nil {
		return m.PullRequest
	}
	return nil
}

type PullRequestReviewComment struct {
	Url              string                         `protobuf:"bytes,1,opt,name=url" json:"url,omitempty"`
	Id               int32                          `protobuf:"varint,2,opt,name=id" json:"id,omitempty"`
	DiffHunk         string                         `protobuf:"bytes,3,opt,name=diff_hunk,json=diffHunk" json:"diff_hunk,omitempty"`
	Path             string                         `protobuf:"bytes,4,opt,name=path" json:"path,omitempty"`
	Position         int32                          `protobuf:"varint,5,opt,name=position" json:"position,omitempty"`
	OriginalPosition int32                          `protobuf:"varint,6,opt,name=original_position,json=originalPosition" json:"original_position,omitempty"`
	CommitId         string                         `protobuf:"bytes,7,opt,name=commit_id,json=commitId" json:"commit_id,omitempty"`
	OriginalCommitId string                         `protobuf:"bytes,8,opt,name=original_commit_id,json=originalCommitId" json:"original_commit_id,omitempty"`
	User             *User                          `protobuf:"bytes,9,opt,name=user" json:"user,omitempty"`
	Body             string                         `protobuf:"bytes,10,opt,name=body" json:"body,omitempty"`
	CreatedAt        string                         `protobuf:"bytes,11,opt,name=created_at,json=createdAt" json:"created_at,omitempty"`
	UpdatedAt        string                         `protobuf:"bytes,12,opt,name=updated_at,json=updatedAt" json:"updated_at,omitempty"`
	HtmlUrl          string                         `protobuf:"bytes,13,opt,name=html_url,json=htmlUrl" json:"html_url,omitempty"`
	PullRequestUrl   string                         `protobuf:"bytes,14,opt,name=pull_request_url,json=pullRequestUrl" json:"pull_request_url,omitempty"`
	XLinks           *PullRequestReviewCommentLinks `protobuf:"bytes,15,opt,name=_links,json=Links" json:"_links,omitempty"`
}

func (m *PullRequestReviewComment) Reset()                    { *m = PullRequestReviewComment{} }
func (m *PullRequestReviewComment) String() string            { return proto.CompactTextString(m) }
func (*PullRequestReviewComment) ProtoMessage()               {}
func (*PullRequestReviewComment) Descriptor() ([]byte, []int) { return fileDescriptor24, []int{2} }

func (m *PullRequestReviewComment) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *PullRequestReviewComment) GetXLinks() *PullRequestReviewCommentLinks {
	if m != nil {
		return m.XLinks
	}
	return nil
}

type PullRequestReviewCommentEvent struct {
	Action      string                    `protobuf:"bytes,1,opt,name=action" json:"action,omitempty"`
	Comment     *PullRequestReviewComment `protobuf:"bytes,2,opt,name=comment" json:"comment,omitempty"`
	PullRequest *PullRequest              `protobuf:"bytes,3,opt,name=pull_request,json=pullRequest" json:"pull_request,omitempty"`
	Repository  *Repository               `protobuf:"bytes,4,opt,name=repository" json:"repository,omitempty"`
	Sender      *User                     `protobuf:"bytes,5,opt,name=sender" json:"sender,omitempty"`
}

func (m *PullRequestReviewCommentEvent) Reset()                    { *m = PullRequestReviewCommentEvent{} }
func (m *PullRequestReviewCommentEvent) String() string            { return proto.CompactTextString(m) }
func (*PullRequestReviewCommentEvent) ProtoMessage()               {}
func (*PullRequestReviewCommentEvent) Descriptor() ([]byte, []int) { return fileDescriptor24, []int{3} }

func (m *PullRequestReviewCommentEvent) GetComment() *PullRequestReviewComment {
	if m != nil {
		return m.Comment
	}
	return nil
}

func (m *PullRequestReviewCommentEvent) GetPullRequest() *PullRequest {
	if m != nil {
		return m.PullRequest
	}
	return nil
}

func (m *PullRequestReviewCommentEvent) GetRepository() *Repository {
	if m != nil {
		return m.Repository
	}
	return nil
}

func (m *PullRequestReviewCommentEvent) GetSender() *User {
	if m != nil {
		return m.Sender
	}
	return nil
}

func init() {
	proto.RegisterType((*PullRequestReviewCommentLink)(nil), "github.PullRequestReviewCommentLink")
	proto.RegisterType((*PullRequestReviewCommentLinks)(nil), "github.PullRequestReviewCommentLinks")
	proto.RegisterType((*PullRequestReviewComment)(nil), "github.PullRequestReviewComment")
	proto.RegisterType((*PullRequestReviewCommentEvent)(nil), "github.PullRequestReviewCommentEvent")
}

func init() { proto.RegisterFile("pull_request_review_comment_event.proto", fileDescriptor24) }

var fileDescriptor24 = []byte{
	// 505 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x54, 0x5f, 0x6b, 0x13, 0x41,
	0x10, 0x27, 0xff, 0x2e, 0xc9, 0x24, 0xc6, 0x74, 0x04, 0x39, 0xa3, 0x85, 0x10, 0x2a, 0x16, 0x94,
	0x3c, 0x44, 0x10, 0x11, 0x5f, 0xa4, 0x88, 0x0a, 0x3e, 0x94, 0x85, 0x3e, 0x2f, 0x97, 0xdc, 0xa6,
	0x59, 0x7a, 0xb9, 0x3b, 0xef, 0xf6, 0x2a, 0xfd, 0x00, 0x7e, 0x4b, 0x3f, 0x88, 0x8f, 0xee, 0xcc,
	0xde, 0x5d, 0x22, 0xa4, 0xb5, 0x7d, 0xc9, 0xcd, 0xfe, 0x7e, 0xbf, 0xf9, 0xb3, 0x33, 0xb3, 0x81,
	0x57, 0x69, 0x11, 0x45, 0x32, 0x53, 0x3f, 0x0a, 0x95, 0x1b, 0xfb, 0xbd, 0xd6, 0xea, 0xa7, 0x5c,
	0x25, 0xdb, 0xad, 0x8a, 0x8d, 0x54, 0xd7, 0xf6, 0x77, 0x9e, 0x66, 0x89, 0x49, 0xd0, 0xbb, 0xd4,
	0x66, 0x53, 0x2c, 0x27, 0x50, 0xe4, 0x2a, 0x73, 0xd8, 0x64, 0x9c, 0xa9, 0x34, 0xc9, 0xb5, 0x49,
	0xb2, 0x9b, 0x12, 0x39, 0xa2, 0x70, 0x65, 0x34, 0x07, 0xcd, 0x16, 0xf0, 0xe2, 0xdc, 0x82, 0xc2,
	0x81, 0x82, 0x33, 0x9c, 0xb9, 0x04, 0xdf, 0x75, 0x7c, 0x85, 0x08, 0xed, 0x8d, 0xd9, 0x46, 0x7e,
	0x63, 0xda, 0x38, 0xed, 0x0b, 0xb6, 0x67, 0xbf, 0x1b, 0x70, 0x7c, 0x97, 0x53, 0x8e, 0xef, 0xa1,
	0x9d, 0xab, 0x68, 0xcd, 0x5e, 0x83, 0xc5, 0xc9, 0xdc, 0x55, 0x37, 0xbf, 0xcb, 0x49, 0xb0, 0x07,
	0x79, 0x72, 0xbe, 0xe6, 0x43, 0x3c, 0xc9, 0x03, 0xbf, 0xc0, 0x70, 0xbf, 0x5b, 0x7e, 0xeb, 0x01,
	0x11, 0x06, 0xe9, 0x8e, 0x9d, 0xfd, 0x69, 0x81, 0x7f, 0x9b, 0x1a, 0xc7, 0xd0, 0x2a, 0xb2, 0xaa,
	0x1d, 0x64, 0xe2, 0x08, 0x9a, 0x3a, 0xe4, 0x7a, 0x3b, 0xc2, 0x5a, 0xf8, 0x1c, 0xfa, 0xa1, 0x5e,
	0xaf, 0xe5, 0xa6, 0x88, 0xaf, 0xb8, 0x88, 0xbe, 0xe8, 0x11, 0xf0, 0xb5, 0x70, 0xed, 0x4c, 0x03,
	0xb3, 0xf1, 0xdb, 0xae, 0x9d, 0x64, 0xe3, 0x04, 0x7a, 0x3c, 0x27, 0x9d, 0xc4, 0x7e, 0x87, 0xc3,
	0xd4, 0x67, 0x7c, 0x0d, 0x47, 0x49, 0xa6, 0x2f, 0x75, 0x1c, 0x44, 0xb2, 0x16, 0x79, 0x2c, 0x1a,
	0x57, 0xc4, 0x79, 0x25, 0xb6, 0x99, 0x69, 0x37, 0xb4, 0x91, 0xb6, 0xa0, 0xae, 0xcb, 0xec, 0x80,
	0x6f, 0x21, 0xbe, 0x01, 0xac, 0x23, 0xed, 0x54, 0x3d, 0x56, 0xd5, 0xa1, 0xce, 0x2a, 0xf5, 0x14,
	0xda, 0xb4, 0x49, 0x7e, 0x9f, 0x9b, 0x38, 0xac, 0x9a, 0x78, 0x61, 0x31, 0xc1, 0x0c, 0xdd, 0x64,
	0x99, 0x84, 0x37, 0x3e, 0xb8, 0x9b, 0x90, 0x8d, 0xc7, 0x00, 0xab, 0x4c, 0x05, 0x46, 0x85, 0x32,
	0x30, 0xfe, 0x80, 0x99, 0x7e, 0x89, 0x7c, 0x32, 0x44, 0x17, 0x69, 0x58, 0xd1, 0x43, 0x47, 0x97,
	0x88, 0xa5, 0x9f, 0x41, 0x8f, 0x06, 0x29, 0xa9, 0xbf, 0x8f, 0x98, 0xec, 0xd2, 0xf9, 0xc2, 0xf6,
	0xf8, 0x14, 0xc6, 0xff, 0xbc, 0x04, 0x92, 0x8c, 0x58, 0x32, 0xda, 0x9b, 0x1c, 0x29, 0x3f, 0x82,
	0x27, 0x23, 0xda, 0x41, 0xff, 0x31, 0x97, 0xfe, 0xf2, 0x3e, 0xf3, 0xcf, 0x45, 0x87, 0x3f, 0xb3,
	0x5f, 0xcd, 0xdb, 0x37, 0xfb, 0x33, 0x3d, 0x37, 0x7c, 0x0a, 0x5e, 0xb0, 0xe2, 0x29, 0xb8, 0x15,
	0x28, 0x4f, 0xf8, 0x01, 0xba, 0xe5, 0xbb, 0x2c, 0x57, 0x77, 0xfa, 0xbf, 0xc4, 0xa2, 0x72, 0xc0,
	0x77, 0x07, 0x37, 0xf7, 0xc9, 0xa1, 0x00, 0xfb, 0x8b, 0x8a, 0x0b, 0x80, 0xdd, 0x13, 0xe7, 0x95,
	0x1a, 0x2c, 0xb0, 0xf2, 0x12, 0x35, 0x23, 0xf6, 0x54, 0x78, 0x02, 0x5e, 0xae, 0xe2, 0xd0, 0x8e,
	0xb6, 0x73, 0x60, 0xb4, 0x25, 0xb7, 0xf4, 0xf8, 0xcf, 0xe1, 0xed, 0xdf, 0x00, 0x00, 0x00, 0xff,
	0xff, 0xb9, 0xd5, 0xef, 0xf5, 0x80, 0x04, 0x00, 0x00,
}
