package main

import (
	"context"
	"log"

	pb "github.com/beslanshapiaev/grpc-go-course/calculator/proto"
)

func doSum(c pb.CalculatorSerivceClient) {
	log.Println("doSum invoked")

	res, err := c.Sum(context.Background(), &pb.SumRequest{
		FirstNumber: 1,
		SecondNumber: 2,
	})

	if err != nil {
		log.Fatalf("Could not sum: %v", err)
	}

	log.Printf("Sum: %d\n", res.Result)
}