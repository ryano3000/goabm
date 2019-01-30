package server

import (
	"context"
	"github.com/Vasanthakumar.V/goabm/apis/grpc/proto/agentpb"
)


type server struct{}

func (*server) Status(ctx context.Context, req *agentpb.) {

}

// Status(ctx context.Context, in *AgentStatusRequest, opts ...grpc.CallOption) (AgentStatusService_StatusClient, error)
