package main

import "fmt"

func main() {
	sum := add(10, 32)
	fmt.Println("sum:", sum)

	quotient, remainder, err := divide(100, 7)
	if err != nil {
		fmt.Println("divide error:", err)
		return
	}

	fmt.Printf("quotient: %d remainder: %d\n", quotient, remainder)

	total := variadicAdd(1, 2, 3, 4)
	fmt.Println("variadic total:", total)
}

func add(a, b int) int {
	return a + b
}

func divide(dividend, divisor int) (int, int, error) {
	if divisor == 0 {
		return 0, 0, fmt.Errorf("cannot divide by zero")
	}
	return dividend / divisor, dividend % divisor, nil
}

func variadicAdd(values ...int) int {
	total := 0
	for _, v := range values {
		total += v
	}
	return total
}
