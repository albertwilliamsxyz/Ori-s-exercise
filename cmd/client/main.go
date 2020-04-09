package main

import (
	"os"
	"fmt"
	"log"
	"strconv"
	"github.com/albertwilliamsxyz/technical-exercise/internal/client"
)

const HELP string = `Technical Exercise CLI
Usage:
	client [METHOD] [ARGS]

Methods:
	add
		description: add two numbers
		example input: client add 1 1
		example output: 2
	substract
		description: substract two numbers
		example input: client substract 1 1
		example output: 0
	multiply
		description: multiply two numbers
		example input: client multiply 2 2
		example output: 4
	divide
		description: divide two numbers
		example input: client divide 2 2
		example output: 1
	sum
		description: sum a list of numbers
		example input: client sum 1 2 3 4
		example output: 10
`

func main() {
	client := client.Client{}
	err := client.Connect(os.Getenv("TECHNICAL_EXERCISE_SERVER_ADDRESS"))
	if err != nil {
		log.Fatalf("Something went wrong: %v", err)
	}
	defer client.Disconnect()

	if len(os.Args) < 2 {
		log.Fatalf("You must pass a method\n%s", HELP)
	}
	method := os.Args[1]
	arguments := os.Args[2:len(os.Args)]
	switch method {
	case "help":
		fmt.Printf(HELP)
	case "add":
		add(&client, arguments)
	case "substract":
		substract(&client, arguments)
	case "multiply":
		multiply(&client, arguments)
	case "divide":
		divide(&client, arguments)
	case "sum":
		sum(&client, arguments)
	default:
		fmt.Println(HELP)
	}
}

func add(client *client.Client, arguments []string) {
	if len(arguments) != 2 {
		log.Fatalf("Invalid number of arguments\n%s", HELP)
	}
	a, errA := strconv.Atoi(arguments[0])
	b, errB := strconv.Atoi(arguments[1])
	if errA != nil || errB != nil {
		log.Fatalf("The Arguments must be integers\n%s", HELP)
	}
	result, err := client.Add(int32(a), int32(b))
	if err != nil {
		log.Fatalf("Something went wrong: %v", err)
	}
	fmt.Printf("Result: %d\n", result)
}

func substract(client *client.Client, arguments []string) {
	if len(arguments) != 2 {
		log.Fatalf("Invalid number of arguments\n%s", HELP)
	}
	a, errA := strconv.Atoi(arguments[0])
	b, errB := strconv.Atoi(arguments[1])
	if errA != nil || errB != nil {
		log.Fatalf("The Arguments must be integers\n%s", HELP)
	}
	result, err := client.Substract(int32(a), int32(b))
	if err != nil {
		log.Fatalf("Something went wrong: %v", err)
	}
	fmt.Printf("Result: %d\n", result)
}

func multiply(client *client.Client, arguments []string) {
	if len(arguments) != 2 {
		log.Fatalf("Invalid number of arguments\n%s", HELP)
	}
	a, errA := strconv.Atoi(arguments[0])
	b, errB := strconv.Atoi(arguments[1])
	if errA != nil || errB != nil {
		log.Fatalf("The Arguments must be integers\n%s", HELP)
	}
	result, err := client.Multiply(int32(a), int32(b))
	if err != nil {
		log.Fatalf("Something went wrong: %v", err)
	}
	fmt.Printf("Result: %d\n", result)
}

func divide(client *client.Client, arguments []string) {
	if len(arguments) != 2 {
		log.Fatalf("Invalid number of arguments\n%s", HELP)
	}
	a, errA := strconv.Atoi(arguments[0])
	b, errB := strconv.Atoi(arguments[1])
	if errA != nil || errB != nil {
		log.Fatalf("The Arguments must be integers\n%s", HELP)
	}
	result, err := client.Divide(int32(a), int32(b))
	if err != nil {
		log.Fatalf("Something went wrong: %v", err)
	}
	fmt.Printf("Result: %d\n", result)
}

func sum(client *client.Client, arguments []string) {
	if len(arguments) < 1 {
		log.Fatalf("Invalid number of arguments\n%s", HELP)
	}
	var numbers []int32
	for _, argument := range arguments {
		number, err := strconv.Atoi(argument)
		if err != nil {
			log.Fatalf("The Arguments must be integers\n%s", HELP)
		}
		numbers = append(numbers, int32(number))
	}
	result, err := client.Sum(numbers)
	if err != nil {
		log.Fatalf("Something went wrong: %v", err)
		return
	}
	fmt.Printf("Result: %d\n", result)
}
