package server

import (
	"golang.org/x/net/context"
	"github.com/albertwilliamsxyz/technical-exercise/api"
)

type Server struct{}

func (server *Server) Add(ctx context.Context, in *api.BinaryOperationRequest) (*api.BinaryOperationResponse, error) {
	result := Add(in.A, in.B)
	return &api.BinaryOperationResponse{Result: result}, nil
}

func (server *Server) Substract(ctx context.Context, in *api.BinaryOperationRequest) (*api.BinaryOperationResponse, error) {
	result := Substract(in.A, in.B)
	return &api.BinaryOperationResponse{Result: result}, nil
}

func (server *Server) Multiply(ctx context.Context, in *api.BinaryOperationRequest) (*api.BinaryOperationResponse, error) {
	result := Multiply(in.A, in.B)
	return &api.BinaryOperationResponse{Result: result}, nil
}

func (server *Server) Divide(ctx context.Context, in *api.BinaryOperationRequest) (*api.BinaryOperationResponse, error) {
	result := Divide(in.A, in.B)
	return &api.BinaryOperationResponse{Result: result}, nil
}
