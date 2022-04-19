package main

import (
	"context"
	"github.com/go-kit/kit/transport/grpc"
	"github.com/pp553933054/micro-go-book/ch12-trace/zipkin-kit/client"
	"github.com/pp553933054/micro-go-book/ch12-trace/zipkin-kit/pb"
	"github.com/pp553933054/micro-go-book/ch12-trace/zipkin-kit/string-service/endpoint"
)

type grpcServer struct {
	diff grpc.Handler
}

func (s *grpcServer) Diff(ctx context.Context, r *pb.StringRequest) (*pb.StringResponse, error) {
	_, resp, err := s.diff.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.StringResponse), nil

}

func NewGRPCServer(ctx context.Context, endpoints endpoint.StringEndpoints, serverTracer grpc.ServerOption) pb.StringServiceServer {
	return &grpcServer{
		diff: grpc.NewServer(
			endpoints.StringEndpoint,
			client.DecodeGRPCStringRequest,
			client.EncodeGRPCStringResponse,
			serverTracer,
		),
	}
}
