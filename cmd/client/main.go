package main

import (
	"log"

	pb "github.com/albertwilliamsxyz/technical-exercise/api"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Something went wrong: %v", err)
	}
	defer conn.Close()

	c := pb.NewTechnicalExerciseServiceClient(conn)
	request := &pb.BinaryOperationRequest{
		A:  1,
		B: 2,
	}
	response, err := c.Add(context.Background(), request)
	if err != nil {
		log.Fatalf("Something went wrong: %v", err)
	}
	log.Printf("Result: %v", response.Result)

	request = &pb.BinaryOperationRequest{
		A:  1,
		B: 2,
	}
	response, err = c.Substract(context.Background(), request)
	if err != nil {
		log.Fatalf("Something went wrong: %v", err)
	}
	log.Printf("Result: %v", response.Result)

	request = &pb.BinaryOperationRequest{
		A:  1,
		B: 2,
	}
	response, err = c.Multiply(context.Background(), request)
	if err != nil {
		log.Fatalf("Something went wrong: %v", err)
	}
	log.Printf("Result: %v", response.Result)

	request = &pb.BinaryOperationRequest{
		A:  1,
		B: 2,
	}
	response, err = c.Divide(context.Background(), request)
	if err != nil {
		log.Fatalf("Something went wrong: %v", err)
	}
	log.Printf("Result: %v", response.Result)
}
