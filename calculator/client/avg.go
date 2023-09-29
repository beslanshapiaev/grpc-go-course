package main

import (
	"context"
	"log"

	pb "github.com/beslanshapiaev/grpc-go-course/calculator/proto"
)

func doAvg(c pb.CalculatorSerivceClient) {
	log.Println("doAvg was invoced")

	stream, err := c.Avg(context.Background())

	if err != nil {
		log.Fatalf("Error while opening the stream: %v\n", err)
	}

	numbers := []int32{1,2,3,4,5,6,7,8,3,2,4,53,10,3333,56,7,8,8}

	for _, number := range numbers {
		log.Printf("Sending number: %d\n", number)
		stream.Send(&pb.AvgRequest{
			Number: number,
		})
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("Error while reveiving response: %v\n", err)
	}

	log.Printf("Avg: %f\n", res.Result)
}