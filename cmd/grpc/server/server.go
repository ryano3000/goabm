package main

import (
	"log"
	"net"
	"time"

	"github.com/VasanthakumarV/goabm/apis/grpc/proto"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) Status(req *agent.AgentStatusRequest, stream agent.AgentStatusService_StatusServer) error {
	for i := 0; i < 10; i++ {
		res := &agent.AgentStatusResponse{
			Status: "s",
		}
		stream.Send(res)
		time.Sleep(time.Second)
	}
	return nil
}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatal("Failed to listen")
	}
	s := grpc.NewServer()
	agent.RegisterAgentStatusServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatal("Failed to serve")
	}

}
