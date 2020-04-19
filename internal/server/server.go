package server

import (
	"fmt"
	"golang.org/x/net/context"
	"github.com/albertwilliamsxyz/technical-exercise/api"
	"github.com/albertwilliamsxyz/technical-exercise/pkg/functions"
)

type Server struct{}

func (server *Server) Add(ctx context.Context, in *api.BinaryOperationRequest) (*api.BinaryOperationResponse, error) {
	result := functions.Add(in.A, in.B)
	return &api.BinaryOperationResponse{Result: result}, nil
}

func (server *Server) Substract(ctx context.Context, in *api.BinaryOperationRequest) (*api.BinaryOperationResponse, error) {
	result := functions.Substract(in.A, in.B)
	return &api.BinaryOperationResponse{Result: result}, nil
}

func (server *Server) Multiply(ctx context.Context, in *api.BinaryOperationRequest) (*api.BinaryOperationResponse, error) {
	result := functions.Multiply(in.A, in.B)
	return &api.BinaryOperationResponse{Result: result}, nil
}

func (server *Server) Divide(ctx context.Context, in *api.BinaryOperationRequest) (*api.BinaryOperationResponse, error) {
	if in.B == 0 {
		return nil, fmt.Errorf("Division by zero")
	}
	result := functions.Divide(in.A, in.B)
	return &api.BinaryOperationResponse{Result: result}, nil
}

func (server *Server) Sum(ctx context.Context, in *api.AggregateOperationRequest) (*api.AggregateOperationResponse, error) {
	result := functions.Sum(in.Numbers)
	return &api.AggregateOperationResponse{Result: result}, nil
}
