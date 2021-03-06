package proto

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/golang/protobuf/ptypes/timestamp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Link describes a link in the linkgraph.
type Link struct {
	Uuid                 []byte               `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Url                  string               `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	RetrievedAt          *timestamp.Timestamp `protobuf:"bytes,3,opt,name=retrieved_at,json=retrievedAt,proto3" json:"retrieved_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Link) Reset()         { *m = Link{} }
func (m *Link) String() string { return proto.CompactTextString(m) }
func (*Link) ProtoMessage()    {}
func (*Link) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{0}
}

func (m *Link) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Link.Unmarshal(m, b)
}
func (m *Link) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Link.Marshal(b, m, deterministic)
}
func (m *Link) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Link.Merge(m, src)
}
func (m *Link) XXX_Size() int {
	return xxx_messageInfo_Link.Size(m)
}
func (m *Link) XXX_DiscardUnknown() {
	xxx_messageInfo_Link.DiscardUnknown(m)
}

var xxx_messageInfo_Link proto.InternalMessageInfo

func (m *Link) GetUuid() []byte {
	if m != nil {
		return m.Uuid
	}
	return nil
}

func (m *Link) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *Link) GetRetrievedAt() *timestamp.Timestamp {
	if m != nil {
		return m.RetrievedAt
	}
	return nil
}

// Edge describes an edge in the linkgraph.
type Edge struct {
	Uuid                 []byte               `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	SrcUuid              []byte               `protobuf:"bytes,2,opt,name=src_uuid,json=srcUuid,proto3" json:"src_uuid,omitempty"`
	DstUuid              []byte               `protobuf:"bytes,3,opt,name=dst_uuid,json=dstUuid,proto3" json:"dst_uuid,omitempty"`
	UpdatedAt            *timestamp.Timestamp `protobuf:"bytes,4,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Edge) Reset()         { *m = Edge{} }
func (m *Edge) String() string { return proto.CompactTextString(m) }
func (*Edge) ProtoMessage()    {}
func (*Edge) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{1}
}

func (m *Edge) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Edge.Unmarshal(m, b)
}
func (m *Edge) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Edge.Marshal(b, m, deterministic)
}
func (m *Edge) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Edge.Merge(m, src)
}
func (m *Edge) XXX_Size() int {
	return xxx_messageInfo_Edge.Size(m)
}
func (m *Edge) XXX_DiscardUnknown() {
	xxx_messageInfo_Edge.DiscardUnknown(m)
}

var xxx_messageInfo_Edge proto.InternalMessageInfo

func (m *Edge) GetUuid() []byte {
	if m != nil {
		return m.Uuid
	}
	return nil
}

func (m *Edge) GetSrcUuid() []byte {
	if m != nil {
		return m.SrcUuid
	}
	return nil
}

func (m *Edge) GetDstUuid() []byte {
	if m != nil {
		return m.DstUuid
	}
	return nil
}

func (m *Edge) GetUpdatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

// RemoveStaleEdgesQuery describes a query for removing stale edges from the
// graph.
type RemoveStaleEdgesQuery struct {
	FromUuid             []byte               `protobuf:"bytes,1,opt,name=from_uuid,json=fromUuid,proto3" json:"from_uuid,omitempty"`
	UpdatedBefore        *timestamp.Timestamp `protobuf:"bytes,2,opt,name=updated_before,json=updatedBefore,proto3" json:"updated_before,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *RemoveStaleEdgesQuery) Reset()         { *m = RemoveStaleEdgesQuery{} }
func (m *RemoveStaleEdgesQuery) String() string { return proto.CompactTextString(m) }
func (*RemoveStaleEdgesQuery) ProtoMessage()    {}
func (*RemoveStaleEdgesQuery) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{2}
}

func (m *RemoveStaleEdgesQuery) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemoveStaleEdgesQuery.Unmarshal(m, b)
}
func (m *RemoveStaleEdgesQuery) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemoveStaleEdgesQuery.Marshal(b, m, deterministic)
}
func (m *RemoveStaleEdgesQuery) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoveStaleEdgesQuery.Merge(m, src)
}
func (m *RemoveStaleEdgesQuery) XXX_Size() int {
	return xxx_messageInfo_RemoveStaleEdgesQuery.Size(m)
}
func (m *RemoveStaleEdgesQuery) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoveStaleEdgesQuery.DiscardUnknown(m)
}

var xxx_messageInfo_RemoveStaleEdgesQuery proto.InternalMessageInfo

func (m *RemoveStaleEdgesQuery) GetFromUuid() []byte {
	if m != nil {
		return m.FromUuid
	}
	return nil
}

func (m *RemoveStaleEdgesQuery) GetUpdatedBefore() *timestamp.Timestamp {
	if m != nil {
		return m.UpdatedBefore
	}
	return nil
}

// Range specifies the [fromID, toID) range to use when streaming Links or Edges.
type Range struct {
	FromUuid []byte `protobuf:"bytes,1,opt,name=from_uuid,json=fromUuid,proto3" json:"from_uuid,omitempty"`
	ToUuid   []byte `protobuf:"bytes,2,opt,name=to_uuid,json=toUuid,proto3" json:"to_uuid,omitempty"`
	// Return results before this filter timestamp.
	Filter               *timestamp.Timestamp `protobuf:"bytes,3,opt,name=filter,proto3" json:"filter,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Range) Reset()         { *m = Range{} }
func (m *Range) String() string { return proto.CompactTextString(m) }
func (*Range) ProtoMessage()    {}
func (*Range) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{3}
}

func (m *Range) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Range.Unmarshal(m, b)
}
func (m *Range) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Range.Marshal(b, m, deterministic)
}
func (m *Range) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Range.Merge(m, src)
}
func (m *Range) XXX_Size() int {
	return xxx_messageInfo_Range.Size(m)
}
func (m *Range) XXX_DiscardUnknown() {
	xxx_messageInfo_Range.DiscardUnknown(m)
}

var xxx_messageInfo_Range proto.InternalMessageInfo

func (m *Range) GetFromUuid() []byte {
	if m != nil {
		return m.FromUuid
	}
	return nil
}

func (m *Range) GetToUuid() []byte {
	if m != nil {
		return m.ToUuid
	}
	return nil
}

func (m *Range) GetFilter() *timestamp.Timestamp {
	if m != nil {
		return m.Filter
	}
	return nil
}

func init() {
	proto.RegisterType((*Link)(nil), "proto.Link")
	proto.RegisterType((*Edge)(nil), "proto.Edge")
	proto.RegisterType((*RemoveStaleEdgesQuery)(nil), "proto.RemoveStaleEdgesQuery")
	proto.RegisterType((*Range)(nil), "proto.Range")
}

func init() { proto.RegisterFile("api.proto", fileDescriptor_00212fb1f9d3bf1c) }

var fileDescriptor_00212fb1f9d3bf1c = []byte{
	// 391 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0x4f, 0xef, 0xd2, 0x40,
	0x10, 0xcd, 0x42, 0xf9, 0xd3, 0x01, 0x0d, 0xd9, 0x44, 0xc5, 0x62, 0x22, 0x69, 0x8c, 0xe1, 0x54,
	0x08, 0x9e, 0x3c, 0x78, 0xc0, 0x84, 0xe8, 0xc1, 0x8b, 0x55, 0xce, 0xa4, 0xd0, 0x69, 0x6d, 0x68,
	0xd9, 0xba, 0x3b, 0xc5, 0xf0, 0x19, 0xfc, 0xb4, 0x7e, 0x03, 0xb3, 0xdb, 0x56, 0x2b, 0xf2, 0xe3,
	0xf7, 0x3b, 0x75, 0xdf, 0xbc, 0x37, 0x7d, 0x6f, 0x67, 0x16, 0xec, 0x20, 0x4f, 0xbc, 0x5c, 0x0a,
	0x12, 0xbc, 0x63, 0x3e, 0xce, 0xcb, 0x58, 0x88, 0x38, 0xc5, 0xb9, 0x41, 0xbb, 0x22, 0x9a, 0x53,
	0x92, 0xa1, 0xa2, 0x20, 0xcb, 0x4b, 0x9d, 0x33, 0xb9, 0x14, 0x60, 0x96, 0xd3, 0xb9, 0x24, 0xdd,
	0x03, 0x58, 0x9f, 0x92, 0xe3, 0x81, 0x73, 0xb0, 0x8a, 0x22, 0x09, 0xc7, 0x6c, 0xca, 0x66, 0x43,
	0xdf, 0x9c, 0xf9, 0x08, 0xda, 0x85, 0x4c, 0xc7, 0xad, 0x29, 0x9b, 0xd9, 0xbe, 0x3e, 0xf2, 0x77,
	0x30, 0x94, 0x48, 0x32, 0xc1, 0x13, 0x86, 0xdb, 0x80, 0xc6, 0xed, 0x29, 0x9b, 0x0d, 0x96, 0x8e,
	0x57, 0x3a, 0x78, 0xb5, 0x83, 0xf7, 0xb5, 0x8e, 0xe0, 0x0f, 0xfe, 0xe8, 0x57, 0xe4, 0xfe, 0x64,
	0x60, 0xad, 0xc3, 0x18, 0xaf, 0xba, 0x3d, 0x87, 0xbe, 0x92, 0xfb, 0xad, 0xa9, 0xb7, 0x4c, 0xbd,
	0xa7, 0xe4, 0x7e, 0x53, 0x51, 0xa1, 0xa2, 0x92, 0x6a, 0x97, 0x54, 0xa8, 0xc8, 0x50, 0x6f, 0x01,
	0x8a, 0x3c, 0x0c, 0xa8, 0xcc, 0x63, 0xdd, 0x9b, 0xc7, 0xae, 0xd4, 0x2b, 0x72, 0x7f, 0xc0, 0x13,
	0x1f, 0x33, 0x71, 0xc2, 0x2f, 0x14, 0xa4, 0xa8, 0x73, 0xa9, 0xcf, 0x05, 0xca, 0x33, 0x9f, 0x80,
	0x1d, 0x49, 0x91, 0x6d, 0x1b, 0x11, 0xfb, 0xba, 0x60, 0x0c, 0x57, 0xf0, 0xb8, 0x36, 0xdc, 0x61,
	0x24, 0x24, 0x9a, 0xb0, 0xb7, 0x4d, 0x1f, 0x55, 0x1d, 0xef, 0x4d, 0x83, 0xfb, 0x1d, 0x3a, 0x7e,
	0x70, 0x8c, 0xf1, 0xb6, 0xd1, 0x33, 0xe8, 0x91, 0x68, 0x8e, 0xa3, 0x4b, 0xc2, 0x10, 0x4b, 0xe8,
	0x46, 0x49, 0x4a, 0x28, 0x1f, 0x30, 0xfe, 0x4a, 0xb9, 0xfc, 0xc5, 0xc0, 0xd6, 0x7b, 0xfe, 0x20,
	0x83, 0xfc, 0x1b, 0x7f, 0x0d, 0xb0, 0xc9, 0x15, 0x4a, 0x32, 0xab, 0x1f, 0x94, 0x8d, 0x9e, 0x06,
	0x4e, 0x13, 0xfc, 0xd5, 0x99, 0xa5, 0xd5, 0x94, 0x06, 0x4e, 0x13, 0xf0, 0x57, 0xd0, 0xd1, 0x7a,
	0xc5, 0x87, 0x55, 0xd5, 0x5c, 0xef, 0x9f, 0x7f, 0x2d, 0x98, 0x56, 0x99, 0x21, 0xdf, 0xa1, 0xd2,
	0xdc, 0x82, 0xf1, 0x8f, 0x30, 0xba, 0xdc, 0x0a, 0x7f, 0x51, 0x37, 0x5c, 0x5b, 0x97, 0xf3, 0xf4,
	0xbf, 0xfb, 0xaf, 0xf5, 0x03, 0xdf, 0x75, 0x0d, 0x7e, 0xf3, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x40,
	0x5e, 0xe7, 0x01, 0x33, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// LinkGraphClient is the client API for LinkGraph service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LinkGraphClient interface {
	// InsertLink inserts or updates a link.
	InsertLink(ctx context.Context, in *Link, opts ...grpc.CallOption) (*Link, error)
	// InsertEdge inserts or updates an edge.
	InsertEdge(ctx context.Context, in *Edge, opts ...grpc.CallOption) (*Edge, error)
	// Links streams the set of links in the specified ID range.
	Links(ctx context.Context, in *Range, opts ...grpc.CallOption) (LinkGraph_LinksClient, error)
	// Edges streams the set of edges in the specified ID range.
	Edges(ctx context.Context, in *Range, opts ...grpc.CallOption) (LinkGraph_EdgesClient, error)
	// RemoveStaleEdges removes any edge that originates from the specified
	// link ID and was updated before the specified timestamp.
	RemoveStaleEdges(ctx context.Context, in *RemoveStaleEdgesQuery, opts ...grpc.CallOption) (*empty.Empty, error)
}

type linkGraphClient struct {
	cc *grpc.ClientConn
}

func NewLinkGraphClient(cc *grpc.ClientConn) LinkGraphClient {
	return &linkGraphClient{cc}
}

func (c *linkGraphClient) InsertLink(ctx context.Context, in *Link, opts ...grpc.CallOption) (*Link, error) {
	out := new(Link)
	err := c.cc.Invoke(ctx, "/proto.LinkGraph/InsertLink", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *linkGraphClient) InsertEdge(ctx context.Context, in *Edge, opts ...grpc.CallOption) (*Edge, error) {
	out := new(Edge)
	err := c.cc.Invoke(ctx, "/proto.LinkGraph/InsertEdge", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *linkGraphClient) Links(ctx context.Context, in *Range, opts ...grpc.CallOption) (LinkGraph_LinksClient, error) {
	stream, err := c.cc.NewStream(ctx, &_LinkGraph_serviceDesc.Streams[0], "/proto.LinkGraph/Links", opts...)
	if err != nil {
		return nil, err
	}
	x := &linkGraphLinksClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type LinkGraph_LinksClient interface {
	Recv() (*Link, error)
	grpc.ClientStream
}

type linkGraphLinksClient struct {
	grpc.ClientStream
}

func (x *linkGraphLinksClient) Recv() (*Link, error) {
	m := new(Link)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *linkGraphClient) Edges(ctx context.Context, in *Range, opts ...grpc.CallOption) (LinkGraph_EdgesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_LinkGraph_serviceDesc.Streams[1], "/proto.LinkGraph/Edges", opts...)
	if err != nil {
		return nil, err
	}
	x := &linkGraphEdgesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type LinkGraph_EdgesClient interface {
	Recv() (*Edge, error)
	grpc.ClientStream
}

type linkGraphEdgesClient struct {
	grpc.ClientStream
}

func (x *linkGraphEdgesClient) Recv() (*Edge, error) {
	m := new(Edge)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *linkGraphClient) RemoveStaleEdges(ctx context.Context, in *RemoveStaleEdgesQuery, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/proto.LinkGraph/RemoveStaleEdges", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LinkGraphServer is the server API for LinkGraph service.
type LinkGraphServer interface {
	// InsertLink inserts or updates a link.
	InsertLink(context.Context, *Link) (*Link, error)
	// InsertEdge inserts or updates an edge.
	InsertEdge(context.Context, *Edge) (*Edge, error)
	// Links streams the set of links in the specified ID range.
	Links(*Range, LinkGraph_LinksServer) error
	// Edges streams the set of edges in the specified ID range.
	Edges(*Range, LinkGraph_EdgesServer) error
	// RemoveStaleEdges removes any edge that originates from the specified
	// link ID and was updated before the specified timestamp.
	RemoveStaleEdges(context.Context, *RemoveStaleEdgesQuery) (*empty.Empty, error)
}

// UnimplementedLinkGraphServer can be embedded to have forward compatible implementations.
type UnimplementedLinkGraphServer struct {
}

func (*UnimplementedLinkGraphServer) InsertLink(ctx context.Context, req *Link) (*Link, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InsertLink not implemented")
}
func (*UnimplementedLinkGraphServer) InsertEdge(ctx context.Context, req *Edge) (*Edge, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InsertEdge not implemented")
}
func (*UnimplementedLinkGraphServer) Links(req *Range, srv LinkGraph_LinksServer) error {
	return status.Errorf(codes.Unimplemented, "method Links not implemented")
}
func (*UnimplementedLinkGraphServer) Edges(req *Range, srv LinkGraph_EdgesServer) error {
	return status.Errorf(codes.Unimplemented, "method Edges not implemented")
}
func (*UnimplementedLinkGraphServer) RemoveStaleEdges(ctx context.Context, req *RemoveStaleEdgesQuery) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveStaleEdges not implemented")
}

func RegisterLinkGraphServer(s *grpc.Server, srv LinkGraphServer) {
	s.RegisterService(&_LinkGraph_serviceDesc, srv)
}

func _LinkGraph_InsertLink_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Link)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LinkGraphServer).InsertLink(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.LinkGraph/InsertLink",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LinkGraphServer).InsertLink(ctx, req.(*Link))
	}
	return interceptor(ctx, in, info, handler)
}

func _LinkGraph_InsertEdge_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Edge)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LinkGraphServer).InsertEdge(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.LinkGraph/InsertEdge",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LinkGraphServer).InsertEdge(ctx, req.(*Edge))
	}
	return interceptor(ctx, in, info, handler)
}

func _LinkGraph_Links_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Range)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(LinkGraphServer).Links(m, &linkGraphLinksServer{stream})
}

type LinkGraph_LinksServer interface {
	Send(*Link) error
	grpc.ServerStream
}

type linkGraphLinksServer struct {
	grpc.ServerStream
}

func (x *linkGraphLinksServer) Send(m *Link) error {
	return x.ServerStream.SendMsg(m)
}

func _LinkGraph_Edges_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Range)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(LinkGraphServer).Edges(m, &linkGraphEdgesServer{stream})
}

type LinkGraph_EdgesServer interface {
	Send(*Edge) error
	grpc.ServerStream
}

type linkGraphEdgesServer struct {
	grpc.ServerStream
}

func (x *linkGraphEdgesServer) Send(m *Edge) error {
	return x.ServerStream.SendMsg(m)
}

func _LinkGraph_RemoveStaleEdges_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveStaleEdgesQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LinkGraphServer).RemoveStaleEdges(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.LinkGraph/RemoveStaleEdges",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LinkGraphServer).RemoveStaleEdges(ctx, req.(*RemoveStaleEdgesQuery))
	}
	return interceptor(ctx, in, info, handler)
}

var _LinkGraph_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.LinkGraph",
	HandlerType: (*LinkGraphServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "InsertLink",
			Handler:    _LinkGraph_InsertLink_Handler,
		},
		{
			MethodName: "InsertEdge",
			Handler:    _LinkGraph_InsertEdge_Handler,
		},
		{
			MethodName: "RemoveStaleEdges",
			Handler:    _LinkGraph_RemoveStaleEdges_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Links",
			Handler:       _LinkGraph_Links_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Edges",
			Handler:       _LinkGraph_Edges_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "api.proto",
}
