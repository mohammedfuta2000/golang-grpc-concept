package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/mohammedfuta2000/golang-grpc-concept/proto"
)

func callSayHelloBiDirectionalStreaming(client pb.GreetServiceClient, names *pb.NamesList){
	log.Printf("BiDirectional streaming started")
	stream, err:=client.SayHelloBiDirectionalStreaming(context.Background())
	if err!=nil {
		log.Fatalf("could not send names: %v", err)
	}

	waitc := make(chan struct{})

	go func(){
		for{
			message, err:= stream.Recv()
			if err==io.EOF {
				break
			}
			if err !=nil {
				log.Fatalf("Some error while streaming: %v", err)
			}
			log.Println(message)
		}
		close(waitc)
	}()
	for _, name:= range names.Names{
		req:=&pb.HelloRequest{
			Name: name,
		}
		if err:= stream.Send(req); err!=nil{
			log.Fatalf("Error while sending %v", err)
		}
		time.Sleep(2*time.Second)
	}
	stream.CloseSend()
	<-waitc
	log.Printf("Bidirectional streaming finished")
}