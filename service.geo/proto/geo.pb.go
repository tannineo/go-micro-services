// Code generated by protoc-gen-go.
// source: service.geo/proto/geo.proto
// DO NOT EDIT!

/*
Package geo is a generated protocol buffer package.

It is generated from these files:
	service.geo/proto/geo.proto

It has these top-level messages:
	Args
	Rectangle
	Point
	Reply
*/
package geo

import proto "github.com/golang/protobuf/proto"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal

// A latitude-longitude bounding box, represented as two diagonally opposite
// points "lo" and "hi".
type Args struct {
	TraceId string     `protobuf:"bytes,1,opt,name=traceId" json:"traceId,omitempty"`
	From    string     `protobuf:"bytes,2,opt" json:"From,omitempty"`
	Rect    *Rectangle `protobuf:"bytes,3,opt,name=rect" json:"rect,omitempty"`
}

func (m *Args) Reset()         { *m = Args{} }
func (m *Args) String() string { return proto.CompactTextString(m) }
func (*Args) ProtoMessage()    {}

func (m *Args) GetRect() *Rectangle {
	if m != nil {
		return m.Rect
	}
	return nil
}

type Rectangle struct {
	Lo *Point `protobuf:"bytes,2,opt,name=lo" json:"lo,omitempty"`
	Hi *Point `protobuf:"bytes,3,opt,name=hi" json:"hi,omitempty"`
}

func (m *Rectangle) Reset()         { *m = Rectangle{} }
func (m *Rectangle) String() string { return proto.CompactTextString(m) }
func (*Rectangle) ProtoMessage()    {}

func (m *Rectangle) GetLo() *Point {
	if m != nil {
		return m.Lo
	}
	return nil
}

func (m *Rectangle) GetHi() *Point {
	if m != nil {
		return m.Hi
	}
	return nil
}

// Points are represented as latitude-longitude pairs in the E7 representation
// (degrees multiplied by 10**7 and rounded to the nearest integer).
// Latitudes should be in the range +/- 90 degrees and longitude should be in
// the range +/- 180 degrees (inclusive).
type Point struct {
	Latitude  int32 `protobuf:"varint,1,opt,name=latitude" json:"latitude,omitempty"`
	Longitude int32 `protobuf:"varint,2,opt,name=longitude" json:"longitude,omitempty"`
}

func (m *Point) Reset()         { *m = Point{} }
func (m *Point) String() string { return proto.CompactTextString(m) }
func (*Point) ProtoMessage()    {}

type Reply struct {
	HotelIds []int32 `protobuf:"varint,1,rep,name=hotelIds" json:"hotelIds,omitempty"`
}

func (m *Reply) Reset()         { *m = Reply{} }
func (m *Reply) String() string { return proto.CompactTextString(m) }
func (*Reply) ProtoMessage()    {}

func init() {
}

// Client API for Geo service

type GeoClient interface {
	// Obtains the Locations contained within the given Rectangle.
	BoundedBox(ctx context.Context, in *Args, opts ...grpc.CallOption) (*Reply, error)
}

type geoClient struct {
	cc *grpc.ClientConn
}

func NewGeoClient(cc *grpc.ClientConn) GeoClient {
	return &geoClient{cc}
}

func (c *geoClient) BoundedBox(ctx context.Context, in *Args, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := grpc.Invoke(ctx, "/geo.Geo/BoundedBox", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Geo service

type GeoServer interface {
	// Obtains the Locations contained within the given Rectangle.
	BoundedBox(context.Context, *Args) (*Reply, error)
}

func RegisterGeoServer(s *grpc.Server, srv GeoServer) {
	s.RegisterService(&_Geo_serviceDesc, srv)
}

func _Geo_BoundedBox_Handler(srv interface{}, ctx context.Context, buf []byte) (interface{}, error) {
	in := new(Args)
	if err := proto.Unmarshal(buf, in); err != nil {
		return nil, err
	}
	out, err := srv.(GeoServer).BoundedBox(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _Geo_serviceDesc = grpc.ServiceDesc{
	ServiceName: "geo.Geo",
	HandlerType: (*GeoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "BoundedBox",
			Handler:    _Geo_BoundedBox_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}
