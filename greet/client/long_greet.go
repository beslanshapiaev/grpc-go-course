package main

import (
	"context"
	"log"
	"time"

	pb "github.com/beslanshapiaev/grpc-go-course/greet/proto"
)

func doLongGreet(c pb.GreetServiceClient) {
	log.Println("doLongGreet waws invoked")

	reqs := []*pb.GreetRequest {
		{ FirstName: "Adam" },
		{ FirstName: "Eva"},
		{ FirstName: "Viktor"},
		{ FirstName: "Vladislav"},
		{ FirstName: "Anna"},
	}
	stream, err := c.LongGreet(context.Background())

	if err != nil {
		log.Fatalf("error while calling L ongGreet %v\n", err)
	}

	for _, req := range reqs {
		log.Printf("Sending req: %v\n", req)
		stream.Send(req)
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("Error while receiving response from LongGreet: %v\n", err)
	}
	
	log.Printf("LongGreet: %s\n", res.Result)
}