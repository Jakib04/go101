package main

import (
	"fmt"
	"strings"
)

// ============================================================
// TOPIC: Go Basics - Variables, Types, Control Flow, Strings
// ============================================================

// --- EXAMPLE ---

var course = "Go 101" // package-level variable

func main() {
	fmt.Println("Welcome to", course)

	// Variables & short declaration
	name := "Ada"
	age := 25
	isStudent := true
	fmt.Printf("name=%s age=%d student=%t\n", name, age, isStudent)

	// String formatting
	greeting := fmt.Sprintf("Hello, %s!", strings.ToUpper(name))
	fmt.Println(greeting)

	// Multiple return values + error handling
	result, err := add(21, 21)
	if err != nil {
		fmt.Println("add error:", err)
		return
	}
	fmt.Println("21 + 21 =", result)

	// Switch statement (no break needed in Go)
	switch {
	case result > 50:
		fmt.Println("Big number!")
	case result == 42:
		fmt.Println("The answer to everything!")
	default:
		fmt.Println("Just a number:", result)
	}

	// For loop + defer (LIFO order)
	for i := 0; i < 3; i++ {
		defer fmt.Println("deferred:", i)
		fmt.Println("loop iteration:", i)
	}
}

func add(a, b int) (int, error) {
	sum := a + b
	if sum < 0 {
		return 0, fmt.Errorf("negative sum: %d", sum)
	}
	return sum, nil
}

// ============================================================
// PRACTICE QUESTIONS
// ============================================================
//
// Q1: Declare a constant called `maxRetries` with value 5.
//     Print it inside main().
//
// Q2: Write a function `isEven(n int) bool` that returns true
//     if n is even. Call it from main() with a few test values.
//
// Q3: Write a function `greet(name string) string` that returns
//     "Hello, <name>! Welcome to Go 101." using fmt.Sprintf.
//
// Q4: Use a for loop to print numbers 1 to 10. Skip even
//     numbers using `continue`.
//
// Q5: Write a switch statement that takes a day string
//     ("Monday", "Saturday", etc.) and prints whether it's
//     a weekday or weekend.
//
// Q6: What will this print? Why?
//     for i := 0; i < 3; i++ {
//         defer fmt.Println(i)
//     }
//     (Answer: 2, 1, 0 - defer is LIFO)
//
// ============================================================
