package main

import (
	"os"
	"log"
	"net"
	"github.com/albertwilliamsxyz/technical-exercise/api"
	"github.com/albertwilliamsxyz/technical-exercise/internal/server"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", os.Getenv("TECHNICAL_EXERCISE_SERVER_ADDRESS"))
	if err != nil {
		log.Fatalf("Something went wrong: %v", err)
	}
	server := server.Server{}
	grpcServer := grpc.NewServer()
	api.RegisterTechnicalExerciseServiceServer(grpcServer, &server)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Something went wrong: %v", err)
	}
}
