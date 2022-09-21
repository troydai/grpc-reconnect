package echoserver

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	echopb "github.com/troydai/grpc-reconnect/protos"
)

func New() echopb.EchoServer {
	return &impl{
		id: uuid.New(),
	}
}

type impl struct {
	echopb.UnimplementedEchoServer

	id uuid.UUID
}

func (s *impl) Echo(_ context.Context, req *echopb.EchoRequest) (*echopb.EchoResponse, error) {
	return &echopb.EchoResponse{
		Message: fmt.Sprintf("[%v] %s", s.id, req.Message),
	}, nil
}

func (s *impl) Status(context.Context, *echopb.StatusRequest) (*echopb.StatusResponse, error) {
	return &echopb.StatusResponse{
		Status: fmt.Sprintf("[%v] running", s.id),
	}, nil
}
