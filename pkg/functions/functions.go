package functions

func Add(a, b int32) int32 {
	return a + b
}

func Substract(a, b int32) int32 {
	return a - b
}

func Multiply(a, b int32) int32 {
	return a * b
}

func Divide(a, b int32) int32 {
	return a / b
}

func Sum(numbers []int32) (result int32) {
	for _, number := range numbers {
		result += number
	}
	return
}
