package echoserver

import (
	"context"

	echopb "github.com/troydai/grpc-reconnect/protos"
)

func New() echopb.EchoServer {
	return &impl{}
}

type impl struct {
	echopb.UnimplementedEchoServer
}

func (s *impl) Echo(_ context.Context, req *echopb.EchoRequest) (*echopb.EchoResponse, error) {
	return &echopb.EchoResponse{
		Message: req.Message,
	}, nil
}

func (s *impl) Status(context.Context, *echopb.StatusRequest) (*echopb.StatusResponse, error) {
	return &echopb.StatusResponse{
		Status: "running",
	}, nil
}
