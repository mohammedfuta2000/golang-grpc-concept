package main

import (
	"context"
	"log"
	"time"

	pb "github.com/mohammedfuta2000/golang-grpc-concept/proto"
)
// send one request, recieve one respnse 
func callSayHello(client pb.GreetServiceClient)  {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res,err:= client.SayHello(ctx, &pb.NoParam{})
	if err!=nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Println(res.Message)
}