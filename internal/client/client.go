package client

import (
	"github.com/albertwilliamsxyz/technical-exercise/api"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type Client struct {
	connection *grpc.ClientConn
	client api.TechnicalExerciseServiceClient
}

func (client *Client) Connect(address string) (err error) {
	client.connection, err = grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		return err
	}
	client.client = api.NewTechnicalExerciseServiceClient(client.connection)
	return nil
}

func (client *Client) Disconnect() {
	client.connection.Close()
}

func (client *Client) Add(a, b int32) (int32, error) {
	request := &api.BinaryOperationRequest{
		A: a,
		B: b,
	}
	response, err := client.client.Add(context.Background(), request)
	if err != nil {
		return 0, err
	}
	return response.Result, nil
}

func (client *Client) Substract(a, b int32) (int32, error) {
	request := &api.BinaryOperationRequest{
		A: a,
		B: b,
	}
	response, err := client.client.Substract(context.Background(), request)
	if err != nil {
		return 0, err
	}
	return response.Result, nil
}

func (client *Client) Multiply(a, b int32) (int32, error) {
	request := &api.BinaryOperationRequest{
		A: a,
		B: b,
	}
	response, err := client.client.Multiply(context.Background(), request)
	if err != nil {
		return 0, err
	}
	return response.Result, nil
}

func (client *Client) Divide(a, b int32) (int32, error) {
	request := &api.BinaryOperationRequest{
		A: a,
		B: b,
	}
	response, err := client.client.Divide(context.Background(), request)
	if err != nil {
		return 0, err
	}
	return response.Result, nil
}

func (client *Client) Sum(numbers []int32) (int32, error) {
	request := &api.AggregateOperationRequest{
		Numbers: numbers,
	}
	response, err := client.client.Sum(context.Background(), request)
	if err != nil {
		return 0, err
	}
	return response.Result, nil
}
