package server

import (
	"context"
	"github.com/VasanthakumarV/goabm/apis/grpc/proto/agentpb"
)


type server struct{}

func (*server) Status(ctx context.Context, req ) {

}

// Status(ctx context.Context, in *AgentStatusRequest, opts ...grpc.CallOption) (AgentStatusService_StatusClient, error)
