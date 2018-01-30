// This file was automatically generated by "microgen 0.7.0b" utility.
// Please, do not edit.
package transportgrpc

import (
	generated "github.com/devimteam/microgen/example/generated"
	protobuf "github.com/devimteam/microgen/example/generated/transport/converter/protobuf"
	protobuf1 "github.com/devimteam/microgen/example/protobuf"
	grpc1 "github.com/go-kit/kit/transport/grpc"
	grpc "google.golang.org/grpc"
)

func NewGRPCClient(conn *grpc.ClientConn, opts ...grpc1.ClientOption) generated.StringService {
	return &generated.Endpoints{
		CountEndpoint: grpc1.NewClient(
			conn,
			"service.string",
			"Count",
			protobuf.EncodeCountRequest,
			protobuf.DecodeCountResponse,
			protobuf1.CountResponse{},
			opts...,
		).Endpoint(),
		TestCaseEndpoint: grpc1.NewClient(
			conn,
			"service.string",
			"TestCase",
			protobuf.EncodeTestCaseRequest,
			protobuf.DecodeTestCaseResponse,
			protobuf1.TestCaseResponse{},
			opts...,
		).Endpoint(),
		UppercaseEndpoint: grpc1.NewClient(
			conn,
			"service.string",
			"Uppercase",
			protobuf.EncodeUppercaseRequest,
			protobuf.DecodeUppercaseResponse,
			protobuf1.UppercaseResponse{},
			opts...,
		).Endpoint(),
	}
}
