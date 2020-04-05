package server

import (
	"fmt"
	"testing"
)

type binaryOperationTestItem struct {
	a, b int32
	want int32
}

func TestAdd(t *testing.T) {
	tests := []binaryOperationTestItem {
		{0, 0, 0},
		{1, 1, 2},
		{1, 0, 1},
		{0, 1, 1},
		{-1, 1, 0},
		{1, -1, 0},
		{-1, -1, -2},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("%d,%d", test.a, test.b), func(t *testing.T) {
			result := Add(test.a, test.b)
			if result != test.want {
				t.Errorf("Got %d, want %d", result, test.want)
			}
		})
	}
}

func TestSubstract(t *testing.T) {
	tests := []binaryOperationTestItem {
		{0, 0, 0},
		{1, 0, 1},
		{1, 1, 0},
		{0, 1, -1},
		{-1, 1, -2},
		{-1, -1, 0},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("%d,%d", test.a, test.b), func(t *testing.T) {
			result := Substract(test.a, test.b)
			if result != test.want {
				t.Errorf("Got %d, want %d", result, test.want)
			}
		})
	}
}

func TestMultiply(t *testing.T) {
	tests := []binaryOperationTestItem {
		{0, 0, 0},
		{1, 0, 0},
		{1, 1, 1},
		{0, 1, 0},
		{-1, 1, -1},
		{-1, -1, 1},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("%d,%d", test.a, test.b), func(t *testing.T) {
			result := Multiply(test.a, test.b)
			if result != test.want {
				t.Errorf("Got %d, want %d", result, test.want)
			}
		})
	}
}

func TestDivide(t *testing.T) {
	tests := []binaryOperationTestItem {
		{1, 1, 1},
		{0, 1, 0},
		{-1, 1, -1},
		{-1, -1, 1},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("%d,%d", test.a, test.b), func(t *testing.T) {
			result := Divide(test.a, test.b)
			if result != test.want {
				t.Errorf("Got %d, want %d", result, test.want)
			}
		})
	}
}