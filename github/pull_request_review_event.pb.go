// Code generated by protoc-gen-gogo.
// source: pull_request_review_event.proto
// DO NOT EDIT!

package github

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type PullRequestReviewLink struct {
	Html string `protobuf:"bytes,1,opt,name=html,proto3" json:"html,omitempty"`
}

func (m *PullRequestReviewLink) Reset()         { *m = PullRequestReviewLink{} }
func (m *PullRequestReviewLink) String() string { return proto.CompactTextString(m) }
func (*PullRequestReviewLink) ProtoMessage()    {}
func (*PullRequestReviewLink) Descriptor() ([]byte, []int) {
	return fileDescriptorPullRequestReviewEvent, []int{0}
}

type PullRequestReviewLinks struct {
	Html        *PullRequestReviewLink `protobuf:"bytes,1,opt,name=html" json:"html,omitempty"`
	PullRequest *PullRequestReviewLink `protobuf:"bytes,2,opt,name=pull_request,json=pullRequest" json:"pull_request,omitempty"`
}

func (m *PullRequestReviewLinks) Reset()         { *m = PullRequestReviewLinks{} }
func (m *PullRequestReviewLinks) String() string { return proto.CompactTextString(m) }
func (*PullRequestReviewLinks) ProtoMessage()    {}
func (*PullRequestReviewLinks) Descriptor() ([]byte, []int) {
	return fileDescriptorPullRequestReviewEvent, []int{1}
}

type PullRequestReview struct {
	Id             int32                   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	User           *User                   `protobuf:"bytes,2,opt,name=user" json:"user,omitempty"`
	Body           string                  `protobuf:"bytes,3,opt,name=body,proto3" json:"body,omitempty"`
	SubmittedAt    string                  `protobuf:"bytes,4,opt,name=submitted_at,json=submittedAt,proto3" json:"submitted_at,omitempty"`
	State          string                  `protobuf:"bytes,5,opt,name=state,proto3" json:"state,omitempty"`
	HtmlUrl        string                  `protobuf:"bytes,6,opt,name=html_url,json=htmlUrl,proto3" json:"html_url,omitempty"`
	PullRequestUrl string                  `protobuf:"bytes,7,opt,name=pull_request_url,json=pullRequestUrl,proto3" json:"pull_request_url,omitempty"`
	XLinks         *PullRequestReviewLinks `protobuf:"bytes,8,opt,name=_links,json=Links" json:"_links,omitempty"`
}

func (m *PullRequestReview) Reset()         { *m = PullRequestReview{} }
func (m *PullRequestReview) String() string { return proto.CompactTextString(m) }
func (*PullRequestReview) ProtoMessage()    {}
func (*PullRequestReview) Descriptor() ([]byte, []int) {
	return fileDescriptorPullRequestReviewEvent, []int{2}
}

type PullRequestReviewEvent struct {
	Action       string             `protobuf:"bytes,1,opt,name=action,proto3" json:"action,omitempty"`
	Review       *PullRequestReview `protobuf:"bytes,2,opt,name=review" json:"review,omitempty"`
	PullRequest  *PullRequest       `protobuf:"bytes,3,opt,name=pull_request,json=pullRequest" json:"pull_request,omitempty"`
	Repository   *Repository        `protobuf:"bytes,4,opt,name=repository" json:"repository,omitempty"`
	Sender       *User              `protobuf:"bytes,5,opt,name=sender" json:"sender,omitempty"`
	Installation *Installation      `protobuf:"bytes,6,opt,name=installation" json:"installation,omitempty"`
	Organization *User              `protobuf:"bytes,7,opt,name=organization" json:"organization,omitempty"`
}

func (m *PullRequestReviewEvent) Reset()         { *m = PullRequestReviewEvent{} }
func (m *PullRequestReviewEvent) String() string { return proto.CompactTextString(m) }
func (*PullRequestReviewEvent) ProtoMessage()    {}
func (*PullRequestReviewEvent) Descriptor() ([]byte, []int) {
	return fileDescriptorPullRequestReviewEvent, []int{3}
}

func init() {
	proto.RegisterType((*PullRequestReviewLink)(nil), "github.PullRequestReviewLink")
	proto.RegisterType((*PullRequestReviewLinks)(nil), "github.PullRequestReviewLinks")
	proto.RegisterType((*PullRequestReview)(nil), "github.PullRequestReview")
	proto.RegisterType((*PullRequestReviewEvent)(nil), "github.PullRequestReviewEvent")
}
func (m *PullRequestReviewLink) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PullRequestReviewLink) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Html) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintPullRequestReviewEvent(dAtA, i, uint64(len(m.Html)))
		i += copy(dAtA[i:], m.Html)
	}
	return i, nil
}

func (m *PullRequestReviewLinks) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PullRequestReviewLinks) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Html != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintPullRequestReviewEvent(dAtA, i, uint64(m.Html.Size()))
		n1, err := m.Html.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.PullRequest != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintPullRequestReviewEvent(dAtA, i, uint64(m.PullRequest.Size()))
		n2, err := m.PullRequest.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	return i, nil
}

func (m *PullRequestReview) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PullRequestReview) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Id != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintPullRequestReviewEvent(dAtA, i, uint64(m.Id))
	}
	if m.User != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintPullRequestReviewEvent(dAtA, i, uint64(m.User.Size()))
		n3, err := m.User.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	if len(m.Body) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintPullRequestReviewEvent(dAtA, i, uint64(len(m.Body)))
		i += copy(dAtA[i:], m.Body)
	}
	if len(m.SubmittedAt) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintPullRequestReviewEvent(dAtA, i, uint64(len(m.SubmittedAt)))
		i += copy(dAtA[i:], m.SubmittedAt)
	}
	if len(m.State) > 0 {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintPullRequestReviewEvent(dAtA, i, uint64(len(m.State)))
		i += copy(dAtA[i:], m.State)
	}
	if len(m.HtmlUrl) > 0 {
		dAtA[i] = 0x32
		i++
		i = encodeVarintPullRequestReviewEvent(dAtA, i, uint64(len(m.HtmlUrl)))
		i += copy(dAtA[i:], m.HtmlUrl)
	}
	if len(m.PullRequestUrl) > 0 {
		dAtA[i] = 0x3a
		i++
		i = encodeVarintPullRequestReviewEvent(dAtA, i, uint64(len(m.PullRequestUrl)))
		i += copy(dAtA[i:], m.PullRequestUrl)
	}
	if m.XLinks != nil {
		dAtA[i] = 0x42
		i++
		i = encodeVarintPullRequestReviewEvent(dAtA, i, uint64(m.XLinks.Size()))
		n4, err := m.XLinks.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n4
	}
	return i, nil
}

func (m *PullRequestReviewEvent) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PullRequestReviewEvent) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Action) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintPullRequestReviewEvent(dAtA, i, uint64(len(m.Action)))
		i += copy(dAtA[i:], m.Action)
	}
	if m.Review != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintPullRequestReviewEvent(dAtA, i, uint64(m.Review.Size()))
		n5, err := m.Review.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n5
	}
	if m.PullRequest != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintPullRequestReviewEvent(dAtA, i, uint64(m.PullRequest.Size()))
		n6, err := m.PullRequest.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n6
	}
	if m.Repository != nil {
		dAtA[i] = 0x22
		i++
		i = encodeVarintPullRequestReviewEvent(dAtA, i, uint64(m.Repository.Size()))
		n7, err := m.Repository.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n7
	}
	if m.Sender != nil {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintPullRequestReviewEvent(dAtA, i, uint64(m.Sender.Size()))
		n8, err := m.Sender.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n8
	}
	if m.Installation != nil {
		dAtA[i] = 0x32
		i++
		i = encodeVarintPullRequestReviewEvent(dAtA, i, uint64(m.Installation.Size()))
		n9, err := m.Installation.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n9
	}
	if m.Organization != nil {
		dAtA[i] = 0x3a
		i++
		i = encodeVarintPullRequestReviewEvent(dAtA, i, uint64(m.Organization.Size()))
		n10, err := m.Organization.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n10
	}
	return i, nil
}

func encodeFixed64PullRequestReviewEvent(dAtA []byte, offset int, v uint64) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	dAtA[offset+4] = uint8(v >> 32)
	dAtA[offset+5] = uint8(v >> 40)
	dAtA[offset+6] = uint8(v >> 48)
	dAtA[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32PullRequestReviewEvent(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintPullRequestReviewEvent(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *PullRequestReviewLink) Size() (n int) {
	var l int
	_ = l
	l = len(m.Html)
	if l > 0 {
		n += 1 + l + sovPullRequestReviewEvent(uint64(l))
	}
	return n
}

func (m *PullRequestReviewLinks) Size() (n int) {
	var l int
	_ = l
	if m.Html != nil {
		l = m.Html.Size()
		n += 1 + l + sovPullRequestReviewEvent(uint64(l))
	}
	if m.PullRequest != nil {
		l = m.PullRequest.Size()
		n += 1 + l + sovPullRequestReviewEvent(uint64(l))
	}
	return n
}

func (m *PullRequestReview) Size() (n int) {
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovPullRequestReviewEvent(uint64(m.Id))
	}
	if m.User != nil {
		l = m.User.Size()
		n += 1 + l + sovPullRequestReviewEvent(uint64(l))
	}
	l = len(m.Body)
	if l > 0 {
		n += 1 + l + sovPullRequestReviewEvent(uint64(l))
	}
	l = len(m.SubmittedAt)
	if l > 0 {
		n += 1 + l + sovPullRequestReviewEvent(uint64(l))
	}
	l = len(m.State)
	if l > 0 {
		n += 1 + l + sovPullRequestReviewEvent(uint64(l))
	}
	l = len(m.HtmlUrl)
	if l > 0 {
		n += 1 + l + sovPullRequestReviewEvent(uint64(l))
	}
	l = len(m.PullRequestUrl)
	if l > 0 {
		n += 1 + l + sovPullRequestReviewEvent(uint64(l))
	}
	if m.XLinks != nil {
		l = m.XLinks.Size()
		n += 1 + l + sovPullRequestReviewEvent(uint64(l))
	}
	return n
}

func (m *PullRequestReviewEvent) Size() (n int) {
	var l int
	_ = l
	l = len(m.Action)
	if l > 0 {
		n += 1 + l + sovPullRequestReviewEvent(uint64(l))
	}
	if m.Review != nil {
		l = m.Review.Size()
		n += 1 + l + sovPullRequestReviewEvent(uint64(l))
	}
	if m.PullRequest != nil {
		l = m.PullRequest.Size()
		n += 1 + l + sovPullRequestReviewEvent(uint64(l))
	}
	if m.Repository != nil {
		l = m.Repository.Size()
		n += 1 + l + sovPullRequestReviewEvent(uint64(l))
	}
	if m.Sender != nil {
		l = m.Sender.Size()
		n += 1 + l + sovPullRequestReviewEvent(uint64(l))
	}
	if m.Installation != nil {
		l = m.Installation.Size()
		n += 1 + l + sovPullRequestReviewEvent(uint64(l))
	}
	if m.Organization != nil {
		l = m.Organization.Size()
		n += 1 + l + sovPullRequestReviewEvent(uint64(l))
	}
	return n
}

func sovPullRequestReviewEvent(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozPullRequestReviewEvent(x uint64) (n int) {
	return sovPullRequestReviewEvent(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *PullRequestReviewLink) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPullRequestReviewEvent
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: PullRequestReviewLink: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PullRequestReviewLink: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Html", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPullRequestReviewEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPullRequestReviewEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Html = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPullRequestReviewEvent(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPullRequestReviewEvent
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *PullRequestReviewLinks) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPullRequestReviewEvent
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: PullRequestReviewLinks: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PullRequestReviewLinks: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Html", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPullRequestReviewEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthPullRequestReviewEvent
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Html == nil {
				m.Html = &PullRequestReviewLink{}
			}
			if err := m.Html.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PullRequest", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPullRequestReviewEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthPullRequestReviewEvent
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.PullRequest == nil {
				m.PullRequest = &PullRequestReviewLink{}
			}
			if err := m.PullRequest.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPullRequestReviewEvent(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPullRequestReviewEvent
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *PullRequestReview) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPullRequestReviewEvent
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: PullRequestReview: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PullRequestReview: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPullRequestReviewEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field User", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPullRequestReviewEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthPullRequestReviewEvent
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.User == nil {
				m.User = &User{}
			}
			if err := m.User.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Body", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPullRequestReviewEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPullRequestReviewEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Body = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SubmittedAt", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPullRequestReviewEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPullRequestReviewEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SubmittedAt = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field State", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPullRequestReviewEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPullRequestReviewEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.State = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HtmlUrl", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPullRequestReviewEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPullRequestReviewEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.HtmlUrl = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PullRequestUrl", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPullRequestReviewEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPullRequestReviewEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PullRequestUrl = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field XLinks", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPullRequestReviewEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthPullRequestReviewEvent
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.XLinks == nil {
				m.XLinks = &PullRequestReviewLinks{}
			}
			if err := m.XLinks.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPullRequestReviewEvent(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPullRequestReviewEvent
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *PullRequestReviewEvent) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPullRequestReviewEvent
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: PullRequestReviewEvent: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PullRequestReviewEvent: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Action", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPullRequestReviewEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPullRequestReviewEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Action = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Review", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPullRequestReviewEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthPullRequestReviewEvent
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Review == nil {
				m.Review = &PullRequestReview{}
			}
			if err := m.Review.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PullRequest", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPullRequestReviewEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthPullRequestReviewEvent
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.PullRequest == nil {
				m.PullRequest = &PullRequest{}
			}
			if err := m.PullRequest.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Repository", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPullRequestReviewEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthPullRequestReviewEvent
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Repository == nil {
				m.Repository = &Repository{}
			}
			if err := m.Repository.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPullRequestReviewEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthPullRequestReviewEvent
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Sender == nil {
				m.Sender = &User{}
			}
			if err := m.Sender.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Installation", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPullRequestReviewEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthPullRequestReviewEvent
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Installation == nil {
				m.Installation = &Installation{}
			}
			if err := m.Installation.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Organization", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPullRequestReviewEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthPullRequestReviewEvent
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Organization == nil {
				m.Organization = &User{}
			}
			if err := m.Organization.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPullRequestReviewEvent(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPullRequestReviewEvent
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipPullRequestReviewEvent(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPullRequestReviewEvent
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowPullRequestReviewEvent
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowPullRequestReviewEvent
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthPullRequestReviewEvent
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowPullRequestReviewEvent
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipPullRequestReviewEvent(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthPullRequestReviewEvent = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPullRequestReviewEvent   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("pull_request_review_event.proto", fileDescriptorPullRequestReviewEvent)
}

var fileDescriptorPullRequestReviewEvent = []byte{
	// 475 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x53, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x6d, 0xd2, 0xd8, 0x29, 0x63, 0xab, 0x2a, 0x43, 0xa8, 0xdc, 0x4a, 0x84, 0x12, 0x71, 0xa8,
	0x84, 0x48, 0x21, 0x08, 0xc4, 0x11, 0x90, 0x38, 0x20, 0x71, 0x40, 0x2b, 0xf5, 0x6c, 0xd9, 0xcd,
	0x92, 0xae, 0x70, 0xbc, 0xc1, 0x5e, 0x17, 0x95, 0x7f, 0x80, 0x33, 0x9f, 0xd4, 0x23, 0x9f, 0x40,
	0xe1, 0x47, 0xd8, 0x9d, 0xdd, 0x34, 0x1b, 0x30, 0xa8, 0x07, 0x6b, 0x67, 0xe6, 0xbd, 0xd9, 0x9d,
	0x99, 0x37, 0x86, 0xbb, 0x8b, 0xa6, 0x28, 0xd2, 0x8a, 0x7f, 0x6c, 0x78, 0xad, 0xf4, 0x79, 0x26,
	0xf8, 0xa7, 0x94, 0x9f, 0xf1, 0x52, 0x8d, 0x17, 0x95, 0x54, 0x12, 0xc3, 0x99, 0x50, 0xa7, 0x4d,
	0xbe, 0xff, 0xd0, 0x9e, 0xe3, 0x13, 0x39, 0x3f, 0x9a, 0xc9, 0x99, 0x3c, 0x22, 0x38, 0x6f, 0xde,
	0x93, 0x47, 0x0e, 0x59, 0x36, 0x6d, 0x1f, 0x9a, 0x9a, 0x57, 0xce, 0xde, 0xa9, 0xf8, 0x42, 0xd6,
	0x42, 0xc9, 0xea, 0xdc, 0x45, 0xd0, 0x7f, 0x75, 0x19, 0x13, 0x65, 0xad, 0xb2, 0xa2, 0xc8, 0x94,
	0x90, 0xa5, 0x8d, 0x8d, 0x1e, 0xc0, 0xed, 0x77, 0x9a, 0xc9, 0x2c, 0x91, 0x51, 0x75, 0x6f, 0x45,
	0xf9, 0x01, 0x11, 0x7a, 0xa7, 0x6a, 0x5e, 0x24, 0x9d, 0x83, 0xce, 0xe1, 0x0d, 0x46, 0xf6, 0xe8,
	0x4b, 0x07, 0x76, 0x5b, 0xd9, 0x35, 0x3e, 0xf6, 0xe8, 0xd1, 0xe4, 0xce, 0xd8, 0xf5, 0xd2, 0xca,
	0xb6, 0xb7, 0xe1, 0x0b, 0x88, 0xfd, 0x22, 0x93, 0xee, 0x75, 0x52, 0xa3, 0xc5, 0x2a, 0x3c, 0xfa,
	0xda, 0x85, 0x9b, 0x7f, 0xd1, 0x70, 0x1b, 0xba, 0x62, 0x4a, 0x85, 0x04, 0x4c, 0x5b, 0x78, 0x00,
	0x3d, 0x33, 0x2a, 0x77, 0x7f, 0xbc, 0xbc, 0xff, 0x58, 0xc7, 0x18, 0x21, 0xa6, 0xd7, 0x5c, 0x4e,
	0xcf, 0x93, 0x4d, 0xdb, 0xab, 0xb1, 0xf1, 0x1e, 0xc4, 0x75, 0x93, 0xcf, 0x85, 0x52, 0x7c, 0x9a,
	0x66, 0x2a, 0xe9, 0x11, 0x16, 0x5d, 0xc5, 0x5e, 0x2a, 0x1c, 0x40, 0xa0, 0xe7, 0xa9, 0x78, 0x12,
	0x10, 0x66, 0x1d, 0xdc, 0x83, 0x2d, 0xd3, 0x5e, 0xda, 0x54, 0x45, 0x12, 0x12, 0xd0, 0x37, 0xfe,
	0x71, 0x55, 0xe0, 0x21, 0xec, 0xac, 0x2d, 0x83, 0xa1, 0xf4, 0x89, 0xb2, 0xed, 0xb5, 0x65, 0x98,
	0x4f, 0x21, 0x4c, 0x0b, 0x33, 0xd8, 0x64, 0x8b, 0xaa, 0x1e, 0xfe, 0x77, 0x2a, 0x35, 0x0b, 0xe8,
	0x18, 0x5d, 0x76, 0x5b, 0x04, 0x7a, 0x6d, 0x76, 0x0d, 0x77, 0x21, 0xcc, 0x4e, 0x8c, 0xf0, 0x4e,
	0x51, 0xe7, 0x69, 0xe1, 0x42, 0xbb, 0x93, 0x6e, 0x3e, 0x7b, 0xff, 0x7c, 0x89, 0x39, 0x22, 0x3e,
	0xfb, 0x43, 0xb8, 0x4d, 0x4a, 0xbc, 0xd5, 0x96, 0xe8, 0xcb, 0x85, 0x13, 0x80, 0xd5, 0x9e, 0xd2,
	0x40, 0xa3, 0x09, 0x2e, 0xb3, 0xd8, 0x15, 0xc2, 0x3c, 0x16, 0xde, 0x87, 0xb0, 0xe6, 0xe5, 0x54,
	0xcb, 0x17, 0xb4, 0xc8, 0xe7, 0x30, 0x7c, 0x0e, 0xb1, 0xbf, 0xdb, 0x34, 0xf7, 0x68, 0x32, 0x58,
	0x72, 0xdf, 0x78, 0x18, 0x5b, 0x63, 0xe2, 0x23, 0x88, 0x65, 0x35, 0xcb, 0x4a, 0xf1, 0xd9, 0x66,
	0xf6, 0x5b, 0x5e, 0x59, 0x63, 0xbc, 0x1a, 0x5c, 0x5c, 0x0e, 0x37, 0x2e, 0x7e, 0x0e, 0x3b, 0xdf,
	0xf5, 0xf7, 0x43, 0x7f, 0xdf, 0x7e, 0x0d, 0x37, 0xf2, 0x90, 0x7e, 0xa7, 0x27, 0xbf, 0x03, 0x00,
	0x00, 0xff, 0xff, 0xb0, 0xd3, 0xd1, 0x92, 0xee, 0x03, 0x00, 0x00,
}
