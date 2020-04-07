package server

import (
	"fmt"
	"testing"
	"reflect"
	"github.com/albertwilliamsxyz/technical-exercise/api"
)

type binaryOperationRequestTestItem struct {
	in api.BinaryOperationRequest
	want api.BinaryOperationResponse
}

type aggregateOperationRequestTestItem struct {
	in api.AggregateOperationRequest
	want api.AggregateOperationResponse
}

func TestServerAdd(t *testing.T) {
	server := Server{}
	tests := []binaryOperationRequestTestItem {
		{
			api.BinaryOperationRequest{A: 1, B: 1},
			api.BinaryOperationResponse{Result: 2},
		},
		{
			api.BinaryOperationRequest{A: 2, B: 2},
			api.BinaryOperationResponse{Result: 4},
		},
		{
			api.BinaryOperationRequest{A: -1, B: 1},
			api.BinaryOperationResponse{Result: 0},
		},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("%v", test.in), func(t *testing.T) {
			result, err := server.Add(nil, &test.in)
			if err != nil {
				t.Errorf("Error: %v", err)
			}
			if !reflect.DeepEqual(*result, test.want) {
				t.Errorf("Got %v, want %v", *result, test.want)
			}
		})
	}
}

func TestServerSubstract(t *testing.T) {
	server := Server{}
	tests := []binaryOperationRequestTestItem {
		{
			api.BinaryOperationRequest{A: 1, B: 1},
			api.BinaryOperationResponse{Result: 0},
		},
		{
			api.BinaryOperationRequest{A: 2, B: 2},
			api.BinaryOperationResponse{Result: 0},
		},
		{
			api.BinaryOperationRequest{A: -1, B: 1},
			api.BinaryOperationResponse{Result: -2},
		},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("%v", test.in), func(t *testing.T) {
			result, err := server.Substract(nil, &test.in)
			if err != nil {
				t.Errorf("Error: %v", err)
			}
			if !reflect.DeepEqual(*result, test.want) {
				t.Errorf("Got %v, want %v", *result, test.want)
			}
		})
	}
}

func TestServerMultiply(t *testing.T) {
	server := Server{}
	tests := []binaryOperationRequestTestItem {
		{
			api.BinaryOperationRequest{A: 1, B: 1},
			api.BinaryOperationResponse{Result: 1},
		},
		{
			api.BinaryOperationRequest{A: 2, B: 2},
			api.BinaryOperationResponse{Result: 4},
		},
		{
			api.BinaryOperationRequest{A: -1, B: 1},
			api.BinaryOperationResponse{Result: -1},
		},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("%v", test.in), func(t *testing.T) {
			result, err := server.Multiply(nil, &test.in)
			if err != nil {
				t.Errorf("Error: %v", err)
			}
			if !reflect.DeepEqual(*result, test.want) {
				t.Errorf("Got %v, want %v", *result, test.want)
			}
		})
	}
}

func TestServerDivide(t *testing.T) {
	server := Server{}
	tests := []binaryOperationRequestTestItem {
		{
			api.BinaryOperationRequest{A: 1, B: 1},
			api.BinaryOperationResponse{Result: 1},
		},
		{
			api.BinaryOperationRequest{A: 2, B: 2},
			api.BinaryOperationResponse{Result: 1},
		},
		{
			api.BinaryOperationRequest{A: -1, B: 1},
			api.BinaryOperationResponse{Result: -1},
		},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("%v", test.in), func(t *testing.T) {
			result, err := server.Divide(nil, &test.in)
			if err != nil {
				t.Errorf("Error: %v", err)
			}
			if !reflect.DeepEqual(*result, test.want) {
				t.Errorf("Got %v, want %v", *result, test.want)
			}
		})
	}
}

func TestServerSum(t *testing.T) {
	server := Server{}
	tests := []aggregateOperationRequestTestItem {
		{
			api.AggregateOperationRequest{Numbers: []int32{1, 1}},
			api.AggregateOperationResponse{Result: 2},
		},
		{
			api.AggregateOperationRequest{Numbers: []int32{1, 2, 1}},
			api.AggregateOperationResponse{Result: 4},
		},
		{
			api.AggregateOperationRequest{Numbers: []int32{-1, 1}},
			api.AggregateOperationResponse{Result: 0},
		},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("%v", test.in), func(t *testing.T) {
			result, err := server.Sum(nil, &test.in)
			if err != nil {
				t.Errorf("Error: %v", err)
			}
			if !reflect.DeepEqual(*result, test.want) {
				t.Errorf("Got %v, want %v", *result, test.want)
			}
		})
	}
}
