package main

import (
	"log"
	"net"
	"time"

	"github.com/VasanthakumarV/goabm/apis/grpc/proto"
	"github.com/VasanthakumarV/goabm/internal/shopper"
	"github.com/VasanthakumarV/goabm/pkgs/abm"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) Status(req *agent.AgentStatusRequest, stream agent.AgentStatusService_StatusServer) error {

	// Create a world with arg which specify
	// the buffer size for the channel and the max number of
	// goroutines
	w := abm.NewWorld(1, 2)

	// Create slice of agents to send in
	ss := []abm.Agent{&shopper.Agent{}, &shopper.Agent{}}

	// Run the simulation for 5 ticks and
	in := w.Start(5, ss)

	// Read in and print the changed status of agents
	// from the channel returned after start
	for s := range in {
		status := s.(*shopper.Agent)
		res := &agent.AgentStatusResponse{
			Status: status.Status,
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
