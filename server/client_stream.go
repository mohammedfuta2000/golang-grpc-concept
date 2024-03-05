package main

import (
	"io"
	"log"

	pb "github.com/mohammedfuta2000/golang-grpc-concept/proto"
)

func(s *helloServer) SayHelloCleintStreaming(stream pb.GreetService_SayHelloCleintStreamingServer) error{
	var message []string
	for {
		req,err:= stream.Recv()
		if err==io.EOF {
			return stream.SendAndClose(&pb.MessagesList{Messages: message})
		}
		if err!=nil {
			return err
		}
		log.Printf("Got request with name: %v", req.Name)
		message = append(message, "Hello", req.Name)
	}
}