package client

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
	request := &pb.AddRequest{
		FirstOperand:  1,
		SecondOperand: 2,
	}
	response, err := c.Add(context.Background(), request)
	if err != nil {
		log.Fatalf("Something went wrong: %v", err)
	}
	log.Printf("Result: %v", response.Result)
}
