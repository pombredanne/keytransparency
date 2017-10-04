// Code generated by protoc-gen-go. DO NOT EDIT.
// source: monitor_v1.proto

/*
Package monitor_v1 is a generated protocol buffer package.

Key Transparency Monitor Service


It is generated from these files:
	monitor_v1.proto

It has these top-level messages:
	GetMonitoringRequest
	GetMonitoringResponse
*/
package monitor_v1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import keytransparency_v1 "github.com/google/keytransparency/core/proto/keytransparency_v1"
import trillian "github.com/google/trillian"
import _ "google.golang.org/genproto/googleapis/api/annotations"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// GetMonitoringRequest contains the input parameters of the GetMonitoring APIs.
type GetMonitoringRequest struct {
	// epoch specifies the for which the monitoring results will
	// be returned (epochs start at 1).
	Epoch int64 `protobuf:"varint,1,opt,name=epoch" json:"epoch,omitempty"`
	// kt_URL is the URL of the keytransparency server for which the monitoring
	// result will be returned.
	Kt_URL string `protobuf:"bytes,2,opt,name=kt_URL,json=ktURL" json:"kt_URL,omitempty"`
}

func (m *GetMonitoringRequest) Reset()                    { *m = GetMonitoringRequest{} }
func (m *GetMonitoringRequest) String() string            { return proto.CompactTextString(m) }
func (*GetMonitoringRequest) ProtoMessage()               {}
func (*GetMonitoringRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *GetMonitoringRequest) GetEpoch() int64 {
	if m != nil {
		return m.Epoch
	}
	return 0
}

func (m *GetMonitoringRequest) GetKt_URL() string {
	if m != nil {
		return m.Kt_URL
	}
	return ""
}

type GetMonitoringResponse struct {
	// smr contains the map root for the sparse Merkle Tree signed with the
	// monitor's key on success. If the checks were not successful the
	// smr will be empty. The epochs are encoded into the smr map_revision.
	Smr *trillian.SignedMapRoot `protobuf:"bytes,1,opt,name=smr" json:"smr,omitempty"`
	// seen_timestamp_nanos contains the time in nanoseconds where this particular
	// signed map root was retrieved and processed. The actual timestamp of the
	// smr returned by the server is contained in above smr field.
	SeenTimestampNanos int64 `protobuf:"varint,2,opt,name=seen_timestamp_nanos,json=seenTimestampNanos" json:"seen_timestamp_nanos,omitempty"`
	// errors contains a list of error string representing the verification checks
	// that failed while monitoring the key-transparency server.
	Errors []string `protobuf:"bytes,3,rep,name=errors" json:"errors,omitempty"`
	// error_data contains the original response from the mutations API if and
	// only if at least one verification step failed. It can be used to re-run the
	// verification steps.
	ErrorData *keytransparency_v1.GetMutationsResponse `protobuf:"bytes,4,opt,name=error_data,json=errorData" json:"error_data,omitempty"`
}

func (m *GetMonitoringResponse) Reset()                    { *m = GetMonitoringResponse{} }
func (m *GetMonitoringResponse) String() string            { return proto.CompactTextString(m) }
func (*GetMonitoringResponse) ProtoMessage()               {}
func (*GetMonitoringResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *GetMonitoringResponse) GetSmr() *trillian.SignedMapRoot {
	if m != nil {
		return m.Smr
	}
	return nil
}

func (m *GetMonitoringResponse) GetSeenTimestampNanos() int64 {
	if m != nil {
		return m.SeenTimestampNanos
	}
	return 0
}

func (m *GetMonitoringResponse) GetErrors() []string {
	if m != nil {
		return m.Errors
	}
	return nil
}

func (m *GetMonitoringResponse) GetErrorData() *keytransparency_v1.GetMutationsResponse {
	if m != nil {
		return m.ErrorData
	}
	return nil
}

func init() {
	proto.RegisterType((*GetMonitoringRequest)(nil), "monitor.v1.GetMonitoringRequest")
	proto.RegisterType((*GetMonitoringResponse)(nil), "monitor.v1.GetMonitoringResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for MonitorService service

type MonitorServiceClient interface {
	// GetSignedMapRoot returns the latest valid signed map root the monitor
	// observed. Additionally, the response contains additional data necessary to
	// reproduce errors on failure.
	//
	// Returns the signed map root for the latest epoch the monitor observed. If
	// the monitor could not reconstruct the map root given the set of mutations
	// from the previous to the current epoch it won't sign the map root and
	// additional data will be provided to reproduce the failure.
	GetSignedMapRoot(ctx context.Context, in *GetMonitoringRequest, opts ...grpc.CallOption) (*GetMonitoringResponse, error)
	// GetSignedMapRootByRevision works similar to GetSignedMapRoot but returns
	// the monitor's result for a specific map revision.
	//
	// Returns the signed map root for the specified epoch the monitor observed.
	// If the monitor could not reconstruct the map root given the set of
	// mutations from the previous to the current epoch it won't sign the map root
	// and additional data will be provided to reproduce the failure.
	GetSignedMapRootByRevision(ctx context.Context, in *GetMonitoringRequest, opts ...grpc.CallOption) (*GetMonitoringResponse, error)
}

type monitorServiceClient struct {
	cc *grpc.ClientConn
}

func NewMonitorServiceClient(cc *grpc.ClientConn) MonitorServiceClient {
	return &monitorServiceClient{cc}
}

func (c *monitorServiceClient) GetSignedMapRoot(ctx context.Context, in *GetMonitoringRequest, opts ...grpc.CallOption) (*GetMonitoringResponse, error) {
	out := new(GetMonitoringResponse)
	err := grpc.Invoke(ctx, "/monitor.v1.MonitorService/GetSignedMapRoot", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *monitorServiceClient) GetSignedMapRootByRevision(ctx context.Context, in *GetMonitoringRequest, opts ...grpc.CallOption) (*GetMonitoringResponse, error) {
	out := new(GetMonitoringResponse)
	err := grpc.Invoke(ctx, "/monitor.v1.MonitorService/GetSignedMapRootByRevision", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for MonitorService service

type MonitorServiceServer interface {
	// GetSignedMapRoot returns the latest valid signed map root the monitor
	// observed. Additionally, the response contains additional data necessary to
	// reproduce errors on failure.
	//
	// Returns the signed map root for the latest epoch the monitor observed. If
	// the monitor could not reconstruct the map root given the set of mutations
	// from the previous to the current epoch it won't sign the map root and
	// additional data will be provided to reproduce the failure.
	GetSignedMapRoot(context.Context, *GetMonitoringRequest) (*GetMonitoringResponse, error)
	// GetSignedMapRootByRevision works similar to GetSignedMapRoot but returns
	// the monitor's result for a specific map revision.
	//
	// Returns the signed map root for the specified epoch the monitor observed.
	// If the monitor could not reconstruct the map root given the set of
	// mutations from the previous to the current epoch it won't sign the map root
	// and additional data will be provided to reproduce the failure.
	GetSignedMapRootByRevision(context.Context, *GetMonitoringRequest) (*GetMonitoringResponse, error)
}

func RegisterMonitorServiceServer(s *grpc.Server, srv MonitorServiceServer) {
	s.RegisterService(&_MonitorService_serviceDesc, srv)
}

func _MonitorService_GetSignedMapRoot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMonitoringRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MonitorServiceServer).GetSignedMapRoot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/monitor.v1.MonitorService/GetSignedMapRoot",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MonitorServiceServer).GetSignedMapRoot(ctx, req.(*GetMonitoringRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MonitorService_GetSignedMapRootByRevision_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMonitoringRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MonitorServiceServer).GetSignedMapRootByRevision(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/monitor.v1.MonitorService/GetSignedMapRootByRevision",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MonitorServiceServer).GetSignedMapRootByRevision(ctx, req.(*GetMonitoringRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MonitorService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "monitor.v1.MonitorService",
	HandlerType: (*MonitorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSignedMapRoot",
			Handler:    _MonitorService_GetSignedMapRoot_Handler,
		},
		{
			MethodName: "GetSignedMapRootByRevision",
			Handler:    _MonitorService_GetSignedMapRootByRevision_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "monitor_v1.proto",
}

func init() { proto.RegisterFile("monitor_v1.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 423 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0xc1, 0x6e, 0xd4, 0x30,
	0x10, 0x86, 0x95, 0x0d, 0x5d, 0x69, 0x8d, 0x54, 0x55, 0xd6, 0x16, 0x56, 0x2b, 0x0e, 0xcb, 0x5e,
	0x48, 0x39, 0xd8, 0x4d, 0xb9, 0x71, 0x04, 0xa4, 0xbd, 0xb4, 0x1c, 0x5c, 0x2a, 0x71, 0x8b, 0xdc,
	0x74, 0x94, 0x5a, 0x9b, 0x78, 0x8c, 0x3d, 0x1b, 0x69, 0x55, 0xf5, 0xc2, 0x89, 0x7b, 0x9f, 0x84,
	0x67, 0x41, 0xbc, 0x01, 0x0f, 0x82, 0xe2, 0xa4, 0x20, 0x56, 0x15, 0x5c, 0xb8, 0x65, 0xfc, 0x8f,
	0xe7, 0xff, 0xf2, 0x7b, 0xd8, 0x41, 0x83, 0xd6, 0x10, 0xfa, 0xa2, 0xcd, 0x85, 0xf3, 0x48, 0xc8,
	0xd9, 0x70, 0x22, 0xda, 0x7c, 0xfe, 0xb1, 0x32, 0x74, 0xbd, 0xb9, 0x14, 0x25, 0x36, 0xb2, 0x42,
	0xac, 0x6a, 0x90, 0x6b, 0xd8, 0x92, 0xd7, 0x36, 0x38, 0xed, 0xc1, 0x96, 0x5b, 0x59, 0xa2, 0x07,
	0x19, 0x6f, 0xee, 0x4a, 0x45, 0x9b, 0x3f, 0x70, 0xd4, 0xbb, 0xcc, 0xf7, 0xc9, 0x9b, 0xba, 0x36,
	0xda, 0x0e, 0xf5, 0xb3, 0x61, 0xbc, 0x76, 0x46, 0x6a, 0x6b, 0x91, 0x34, 0x19, 0xb4, 0xa1, 0x57,
	0x97, 0x6f, 0xd9, 0x74, 0x05, 0x74, 0xd6, 0x83, 0x19, 0x5b, 0x29, 0xf8, 0xb4, 0x81, 0x40, 0x7c,
	0xca, 0xf6, 0xc0, 0x61, 0x79, 0x3d, 0x4b, 0x16, 0x49, 0x96, 0xaa, 0xbe, 0xe0, 0x87, 0x6c, 0xbc,
	0xa6, 0xe2, 0x42, 0x9d, 0xce, 0x46, 0x8b, 0x24, 0x9b, 0xa8, 0xbd, 0x35, 0x5d, 0xa8, 0xd3, 0xe5,
	0xf7, 0x84, 0x1d, 0xee, 0x4c, 0x09, 0x0e, 0x6d, 0x00, 0x7e, 0xc4, 0xd2, 0xd0, 0xf8, 0x38, 0xe4,
	0xf1, 0xc9, 0x53, 0xf1, 0x0b, 0xed, 0xdc, 0x54, 0x16, 0xae, 0xce, 0xb4, 0x53, 0x88, 0xa4, 0xba,
	0x1e, 0x7e, 0xcc, 0xa6, 0x01, 0xc0, 0x16, 0x64, 0x1a, 0x08, 0xa4, 0x1b, 0x57, 0x58, 0x6d, 0x31,
	0x44, 0xa7, 0x54, 0xf1, 0x4e, 0xfb, 0x70, 0x2f, 0xbd, 0xef, 0x14, 0xfe, 0x84, 0x8d, 0xc1, 0x7b,
	0xf4, 0x61, 0x96, 0x2e, 0xd2, 0x6c, 0xa2, 0x86, 0x8a, 0xaf, 0x18, 0x8b, 0x5f, 0xc5, 0x95, 0x26,
	0x3d, 0x7b, 0x14, 0xbd, 0x33, 0xb1, 0x13, 0x98, 0x68, 0x73, 0xd1, 0x31, 0x6f, 0x86, 0x3c, 0xee,
	0x91, 0xd5, 0x24, 0xde, 0x7d, 0xa7, 0x49, 0x9f, 0x7c, 0x1d, 0xb1, 0xfd, 0xe1, 0xa7, 0xce, 0xc1,
	0xb7, 0xa6, 0x04, 0xfe, 0x25, 0x61, 0x07, 0x2b, 0xa0, 0x3f, 0xf8, 0xf9, 0x42, 0xfc, 0x7e, 0x59,
	0xf1, 0x50, 0x9c, 0xf3, 0xe7, 0x7f, 0xe9, 0xe8, 0x7d, 0x97, 0xf2, 0xf3, 0xb7, 0x1f, 0x77, 0xa3,
	0x23, 0xfe, 0x42, 0xb6, 0xb9, 0x1c, 0xba, 0xe5, 0x4d, 0x1f, 0xf7, 0xad, 0x6c, 0xb4, 0x93, 0x1e,
	0xc2, 0xa6, 0xa6, 0xf0, 0xba, 0xd6, 0xd4, 0x3d, 0xd1, 0x5d, 0xc2, 0xe6, 0xbb, 0x28, 0x6f, 0xb6,
	0x0a, 0x5a, 0x13, 0x0c, 0xda, 0xff, 0x03, 0x75, 0x1c, 0xa1, 0x5e, 0xf2, 0xec, 0x5f, 0x50, 0xf2,
	0x26, 0x6e, 0xc8, 0xed, 0xe5, 0x38, 0xee, 0xd5, 0xab, 0x9f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x85,
	0x9d, 0xe0, 0xdc, 0xff, 0x02, 0x00, 0x00,
}
