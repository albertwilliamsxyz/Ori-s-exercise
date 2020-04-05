package main

import (
	"fmt"
	"net"
	"log"
	pb "github.com/albertwilliamsxyz/technical-exercise/api"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const grpcPort = ":50051"

type Server struct {}

func (server *Server) Add(ctx context.Context, in *pb.AddRequest) (*pb.AddResponse, error) {
	return &pb.AddResponse{ Result: in.FirstOperand + in.SecondOperand }, nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 50051))
	if err != nil {
		log.Fatalf("Something went wrong: %v", err)
	}

	server := Server{}
	grpcServer := grpc.NewServer()
	pb.RegisterTechnicalExerciseServiceServer(grpcServer, &server)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Something went wrong: %v", err)
	}
}
