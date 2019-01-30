package main

import (
	"context"
	"io"
	"log"

	"github.com/VasanthakumarV/goabm/apis/grpc/proto"
	"google.golang.org/grpc"
)

func doServerStreaming(c agent.AgentStatusServiceClient) {
	req := &agent.AgentStatusRequest{}
	resStream, _ := c.Status(context.Background(), req)
	for {
		msg, err := resStream.Recv()
		if err == io.EOF {
			break
		}
		log.Printf("Response: %v", msg.GetStatus())
	}
}

func main() {
	cc, err := grpc.Dial("0.0.0.0:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatal("Could not connect")
	}
	defer cc.Close()
	c := agent.NewAgentStatusServiceClient(cc)
	doServerStreaming(c)
}
