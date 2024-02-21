package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	pb "github.com/mohammedfuta2000/golang-grpc-concept/proto"
)

const (
	port = ":8080"
)

type helloServer struct{
	pb.GreetServiceServer

}

func main()  {
	log.Println("Starting server")
	// what is a listener and how is it differnet from http.listen()
	lis, err := net.Listen("tcp",port)
	if err!=nil {
		log.Fatalf("failed to start the server: %v",err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterGreetServiceServer(grpcServer, &helloServer{})
	log.Printf("server started at %v", lis.Addr())
	if err:= grpcServer.Serve(lis); err!=nil{
		log.Fatalf("failed to start the server: %v",err)
	}
}