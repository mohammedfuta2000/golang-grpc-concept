package main

import (
	"log"

	pb "github.com/mohammedfuta2000/golang-grpc-concept/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	port = ":8080"
)

func main()  {
	// this is akin to http.post
	conn, err := grpc.Dial("localhost"+port,grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err!=nil {
		log.Fatalf("did not connect: %v",err)
	}
	defer conn.Close()

	client:= pb.NewGreetServiceClient(conn)

	
	
	names:= &pb.NamesList{
		Names: []string{"Alice", "Bob"},
	}
	
	// callSayHello(client)
	// callSayHelloServerStream(client, names)
	// callSayHelloCleintStreaming(client,names)
	callSayHelloBiDirectionalStreaming(client,names)
}