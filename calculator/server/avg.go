package main

import (
	"io"
	"log"

	pb "github.com/beslanshapiaev/grpc-go-course/calculator/proto"
)

func (s *Server) Avg(stream pb.CalculatorSerivce_AvgServer) error {
	log.Println("Avg function was invoked")

	var sum int32 = 0
	count := 0

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&pb.AvgResponse{
				Result: float64(sum) / float64(count),
			})
		}

		if err != nil {
			log.Fatalf("error while reading stream %v\n", err)
		}

		log.Printf("Reveiving number: %d\n", req.Number)
		sum += req.Number
		count++
	}
}