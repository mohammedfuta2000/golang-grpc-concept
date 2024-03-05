package main

import (
	"context"

	pb "github.com/mohammedfuta2000/golang-grpc-concept/proto"
)

func (s *helloServer) SayHello(ctx context.Context, req *pb.NoParam)(*pb.HelloResponse, error)  {
	return &pb.HelloResponse{
		Message: "Hello",
	},nil
}
