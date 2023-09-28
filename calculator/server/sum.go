package main

import (
	"context"
	"fmt"

	pb "github.com/beslanshapiaev/grpc-go-course/calculator/proto"
)



func (s *Server) Sum(ctx context.Context, in *pb.SumRequest) (*pb.SumResponse, error ){
	fmt.Printf("Sum invoked with %v\n", in)

	return &pb.SumResponse{
		Result: in.FirstNumber + in.SecondNumber,
	}, nil
}