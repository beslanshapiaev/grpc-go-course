package main

import (
	"context"
	"log"

	pb "github.com/beslanshapiaev/grpc-go-course/calculator/proto"
	"google.golang.org/grpc/status"
)

func doSqrt(c pb.CalculatorSerivceClient, n int32) {
	log.Println("doSqrt was invoked")
	res, err := c.Sqrt(context.Background(), &pb.SqrtRequest{Number: n})

	if err != nil {
		e, ok := status.FromError(err)
		
		if ok {
			log.Printf("Error message from server; %s\n", e.Message())
			log.Printf("Error code from server; %s\n", e.Code())
		}else {
			log.Fatalf("A non gRPC error: %v\n", err)
		}
	}

	log.Printf("Sqrt: %f\n", res.Result)
}