// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: proto/cashier/cashier.proto

package cashier

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type Cashier struct {
	UserId   string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Code     string `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
	Password string `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	Name     string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
}

func (m *Cashier) Reset()         { *m = Cashier{} }
func (m *Cashier) String() string { return proto.CompactTextString(m) }
func (*Cashier) ProtoMessage()    {}
func (*Cashier) Descriptor() ([]byte, []int) {
	return fileDescriptor_73f06902101431d3, []int{0}
}
func (m *Cashier) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Cashier) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Cashier.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Cashier) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Cashier.Merge(m, src)
}
func (m *Cashier) XXX_Size() int {
	return m.Size()
}
func (m *Cashier) XXX_DiscardUnknown() {
	xxx_messageInfo_Cashier.DiscardUnknown(m)
}

var xxx_messageInfo_Cashier proto.InternalMessageInfo

func (m *Cashier) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *Cashier) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *Cashier) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *Cashier) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type ListQuery struct {
	Limit int64  `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Page  int64  `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
	Sort  string `protobuf:"bytes,3,opt,name=sort,proto3" json:"sort,omitempty"`
	Where string `protobuf:"bytes,4,opt,name=where,proto3" json:"where,omitempty"`
}

func (m *ListQuery) Reset()         { *m = ListQuery{} }
func (m *ListQuery) String() string { return proto.CompactTextString(m) }
func (*ListQuery) ProtoMessage()    {}
func (*ListQuery) Descriptor() ([]byte, []int) {
	return fileDescriptor_73f06902101431d3, []int{1}
}
func (m *ListQuery) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ListQuery) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ListQuery.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ListQuery) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListQuery.Merge(m, src)
}
func (m *ListQuery) XXX_Size() int {
	return m.Size()
}
func (m *ListQuery) XXX_DiscardUnknown() {
	xxx_messageInfo_ListQuery.DiscardUnknown(m)
}

var xxx_messageInfo_ListQuery proto.InternalMessageInfo

func (m *ListQuery) GetLimit() int64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *ListQuery) GetPage() int64 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *ListQuery) GetSort() string {
	if m != nil {
		return m.Sort
	}
	return ""
}

func (m *ListQuery) GetWhere() string {
	if m != nil {
		return m.Where
	}
	return ""
}

type Request struct {
	ListQuery *ListQuery `protobuf:"bytes,1,opt,name=list_query,json=listQuery,proto3" json:"list_query,omitempty"`
	Cashier   *Cashier   `protobuf:"bytes,2,opt,name=cashier,proto3" json:"cashier,omitempty"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_73f06902101431d3, []int{2}
}
func (m *Request) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Request.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return m.Size()
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetListQuery() *ListQuery {
	if m != nil {
		return m.ListQuery
	}
	return nil
}

func (m *Request) GetCashier() *Cashier {
	if m != nil {
		return m.Cashier
	}
	return nil
}

type Response struct {
	Valid    bool       `protobuf:"varint,1,opt,name=valid,proto3" json:"valid,omitempty"`
	Total    int64      `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
	Cashier  *Cashier   `protobuf:"bytes,3,opt,name=cashier,proto3" json:"cashier,omitempty"`
	Cashiers []*Cashier `protobuf:"bytes,4,rep,name=cashiers,proto3" json:"cashiers,omitempty"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_73f06902101431d3, []int{3}
}
func (m *Response) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Response.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return m.Size()
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetValid() bool {
	if m != nil {
		return m.Valid
	}
	return false
}

func (m *Response) GetTotal() int64 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *Response) GetCashier() *Cashier {
	if m != nil {
		return m.Cashier
	}
	return nil
}

func (m *Response) GetCashiers() []*Cashier {
	if m != nil {
		return m.Cashiers
	}
	return nil
}

func init() {
	proto.RegisterType((*Cashier)(nil), "cashier.Cashier")
	proto.RegisterType((*ListQuery)(nil), "cashier.ListQuery")
	proto.RegisterType((*Request)(nil), "cashier.Request")
	proto.RegisterType((*Response)(nil), "cashier.Response")
}

func init() { proto.RegisterFile("proto/cashier/cashier.proto", fileDescriptor_73f06902101431d3) }

var fileDescriptor_73f06902101431d3 = []byte{
	// 373 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0xcd, 0x6a, 0xe2, 0x50,
	0x18, 0x4d, 0x4c, 0xc6, 0xc4, 0xcf, 0xcd, 0xcc, 0x65, 0x60, 0x82, 0x03, 0x41, 0xb2, 0x1a, 0x06,
	0xab, 0xd4, 0x3e, 0x41, 0x6b, 0xa1, 0x14, 0xba, 0xe9, 0x85, 0xae, 0x25, 0x35, 0x5f, 0x6b, 0xe0,
	0x6a, 0xe2, 0xbd, 0xd7, 0x4a, 0xdf, 0xc2, 0xc7, 0xe8, 0xa3, 0x74, 0xe9, 0xb2, 0xcb, 0xa2, 0x2f,
	0x52, 0xee, 0x4f, 0x52, 0x17, 0x85, 0xea, 0x2a, 0xe7, 0x1c, 0xce, 0xc9, 0xf9, 0x72, 0x20, 0xf0,
	0xb7, 0xe4, 0x85, 0x2c, 0x06, 0x93, 0x54, 0x4c, 0x73, 0xe4, 0xd5, 0xb3, 0xaf, 0x55, 0x12, 0x58,
	0x9a, 0x3c, 0x40, 0x30, 0x32, 0x90, 0xfc, 0x81, 0x60, 0x29, 0x90, 0x8f, 0xf3, 0x2c, 0x72, 0xbb,
	0xee, 0xbf, 0x16, 0x6d, 0x2a, 0x7a, 0x9d, 0x11, 0x02, 0xfe, 0xa4, 0xc8, 0x30, 0x6a, 0x68, 0x55,
	0x63, 0xd2, 0x81, 0xb0, 0x4c, 0x85, 0x58, 0x15, 0x3c, 0x8b, 0x3c, 0xad, 0xd7, 0x5c, 0xf9, 0xe7,
	0xe9, 0x0c, 0x23, 0xdf, 0xf8, 0x15, 0x4e, 0xc6, 0xd0, 0xba, 0xc9, 0x85, 0xbc, 0x5d, 0x22, 0x7f,
	0x26, 0xbf, 0xe1, 0x07, 0xcb, 0x67, 0xb9, 0xd4, 0x3d, 0x1e, 0x35, 0x44, 0xc5, 0xca, 0xf4, 0xd1,
	0xd4, 0x78, 0x54, 0x63, 0xa5, 0x89, 0x82, 0x4b, 0x5b, 0xa1, 0xb1, 0x4a, 0xaf, 0xa6, 0xc8, 0xab,
	0xf7, 0x1b, 0x92, 0x4c, 0x21, 0xa0, 0xb8, 0x58, 0xa2, 0x90, 0xe4, 0x14, 0x80, 0xe5, 0x42, 0x8e,
	0x17, 0xaa, 0x4c, 0x77, 0xb4, 0x87, 0xa4, 0x5f, 0x0d, 0x50, 0x9f, 0x41, 0x5b, 0xac, 0xbe, 0xe8,
	0x3f, 0x54, 0x8b, 0xe8, 0xfa, 0xf6, 0xf0, 0x67, 0xed, 0xb7, 0xf3, 0xd0, 0x7a, 0xb2, 0xb5, 0x0b,
	0x21, 0x45, 0x51, 0x16, 0x73, 0x81, 0xea, 0x98, 0xa7, 0x94, 0xd9, 0xc9, 0x42, 0x6a, 0x88, 0x52,
	0x65, 0x21, 0x53, 0x66, 0xbf, 0xc5, 0x90, 0xfd, 0x12, 0xef, 0x9b, 0x12, 0xd2, 0x83, 0xd0, 0x42,
	0x11, 0xf9, 0x5d, 0xef, 0x4b, 0x73, 0xed, 0x18, 0xbe, 0x34, 0x20, 0xb4, 0xaa, 0x20, 0x3d, 0xf0,
	0xce, 0x19, 0x23, 0x9f, 0x7e, 0xbb, 0x4b, 0xe7, 0xd7, 0x9e, 0x62, 0xce, 0x4f, 0x1c, 0x72, 0x02,
	0xbe, 0x5a, 0xe4, 0x50, 0x7b, 0x0f, 0xbc, 0x2b, 0x3c, 0xd8, 0x3d, 0x80, 0xe6, 0x88, 0x63, 0x2a,
	0xf1, 0x88, 0xc0, 0x5d, 0x99, 0x1d, 0x17, 0xb8, 0x44, 0x86, 0x07, 0x07, 0x2e, 0xa2, 0xd7, 0x6d,
	0xec, 0x6e, 0xb6, 0xb1, 0xfb, 0xbe, 0x8d, 0xdd, 0xf5, 0x2e, 0x76, 0x36, 0xbb, 0xd8, 0x79, 0xdb,
	0xc5, 0xce, 0x7d, 0x53, 0xff, 0x1a, 0x67, 0x1f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xbc, 0xa9, 0x73,
	0x18, 0x39, 0x03, 0x00, 0x00,
}

func (m *Cashier) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Cashier) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Cashier) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintCashier(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Password) > 0 {
		i -= len(m.Password)
		copy(dAtA[i:], m.Password)
		i = encodeVarintCashier(dAtA, i, uint64(len(m.Password)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Code) > 0 {
		i -= len(m.Code)
		copy(dAtA[i:], m.Code)
		i = encodeVarintCashier(dAtA, i, uint64(len(m.Code)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.UserId) > 0 {
		i -= len(m.UserId)
		copy(dAtA[i:], m.UserId)
		i = encodeVarintCashier(dAtA, i, uint64(len(m.UserId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ListQuery) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ListQuery) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ListQuery) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Where) > 0 {
		i -= len(m.Where)
		copy(dAtA[i:], m.Where)
		i = encodeVarintCashier(dAtA, i, uint64(len(m.Where)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Sort) > 0 {
		i -= len(m.Sort)
		copy(dAtA[i:], m.Sort)
		i = encodeVarintCashier(dAtA, i, uint64(len(m.Sort)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Page != 0 {
		i = encodeVarintCashier(dAtA, i, uint64(m.Page))
		i--
		dAtA[i] = 0x10
	}
	if m.Limit != 0 {
		i = encodeVarintCashier(dAtA, i, uint64(m.Limit))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Request) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Request) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Request) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Cashier != nil {
		{
			size, err := m.Cashier.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintCashier(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.ListQuery != nil {
		{
			size, err := m.ListQuery.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintCashier(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Response) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Response) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Response) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Cashiers) > 0 {
		for iNdEx := len(m.Cashiers) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Cashiers[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintCashier(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if m.Cashier != nil {
		{
			size, err := m.Cashier.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintCashier(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.Total != 0 {
		i = encodeVarintCashier(dAtA, i, uint64(m.Total))
		i--
		dAtA[i] = 0x10
	}
	if m.Valid {
		i--
		if m.Valid {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintCashier(dAtA []byte, offset int, v uint64) int {
	offset -= sovCashier(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Cashier) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.UserId)
	if l > 0 {
		n += 1 + l + sovCashier(uint64(l))
	}
	l = len(m.Code)
	if l > 0 {
		n += 1 + l + sovCashier(uint64(l))
	}
	l = len(m.Password)
	if l > 0 {
		n += 1 + l + sovCashier(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovCashier(uint64(l))
	}
	return n
}

func (m *ListQuery) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Limit != 0 {
		n += 1 + sovCashier(uint64(m.Limit))
	}
	if m.Page != 0 {
		n += 1 + sovCashier(uint64(m.Page))
	}
	l = len(m.Sort)
	if l > 0 {
		n += 1 + l + sovCashier(uint64(l))
	}
	l = len(m.Where)
	if l > 0 {
		n += 1 + l + sovCashier(uint64(l))
	}
	return n
}

func (m *Request) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ListQuery != nil {
		l = m.ListQuery.Size()
		n += 1 + l + sovCashier(uint64(l))
	}
	if m.Cashier != nil {
		l = m.Cashier.Size()
		n += 1 + l + sovCashier(uint64(l))
	}
	return n
}

func (m *Response) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Valid {
		n += 2
	}
	if m.Total != 0 {
		n += 1 + sovCashier(uint64(m.Total))
	}
	if m.Cashier != nil {
		l = m.Cashier.Size()
		n += 1 + l + sovCashier(uint64(l))
	}
	if len(m.Cashiers) > 0 {
		for _, e := range m.Cashiers {
			l = e.Size()
			n += 1 + l + sovCashier(uint64(l))
		}
	}
	return n
}

func sovCashier(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozCashier(x uint64) (n int) {
	return sovCashier(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Cashier) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCashier
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Cashier: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Cashier: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UserId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCashier
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCashier
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCashier
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UserId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Code", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCashier
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCashier
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCashier
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Code = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Password", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCashier
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCashier
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCashier
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Password = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCashier
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCashier
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCashier
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCashier(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCashier
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthCashier
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
func (m *ListQuery) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCashier
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ListQuery: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ListQuery: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Limit", wireType)
			}
			m.Limit = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCashier
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Limit |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Page", wireType)
			}
			m.Page = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCashier
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Page |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sort", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCashier
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCashier
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCashier
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sort = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Where", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCashier
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCashier
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCashier
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Where = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCashier(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCashier
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthCashier
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
func (m *Request) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCashier
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Request: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Request: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ListQuery", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCashier
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthCashier
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCashier
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ListQuery == nil {
				m.ListQuery = &ListQuery{}
			}
			if err := m.ListQuery.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Cashier", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCashier
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthCashier
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCashier
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Cashier == nil {
				m.Cashier = &Cashier{}
			}
			if err := m.Cashier.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCashier(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCashier
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthCashier
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
func (m *Response) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCashier
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Response: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Response: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Valid", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCashier
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Valid = bool(v != 0)
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Total", wireType)
			}
			m.Total = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCashier
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Total |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Cashier", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCashier
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthCashier
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCashier
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Cashier == nil {
				m.Cashier = &Cashier{}
			}
			if err := m.Cashier.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Cashiers", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCashier
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthCashier
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCashier
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Cashiers = append(m.Cashiers, &Cashier{})
			if err := m.Cashiers[len(m.Cashiers)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCashier(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCashier
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthCashier
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
func skipCashier(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCashier
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
					return 0, ErrIntOverflowCashier
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowCashier
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
			if length < 0 {
				return 0, ErrInvalidLengthCashier
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupCashier
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthCashier
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthCashier        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCashier          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupCashier = fmt.Errorf("proto: unexpected end of group")
)
