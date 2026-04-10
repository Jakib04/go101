package main

import (
	"fmt"
)

// ============================================================
// TOPIC: Go Basics - Variables, Types, Control Flow, Strings
// ============================================================

// --- EXAMPLE ---

var course = "Go 101" // package-level variable

const maxRetries = 5 // constant declaration 

func main() {
	fmt.Println("Welcome to", course)
	fmt.Println("Max retries:", maxRetries)

	// Test the isEven function
	fmt.Println("Is 19 even?", isEven(19))
	fmt.Println("Is 12 even?", isEven(12))

	// Test the greet function
	fmt.Println(greet("Jahed"))
	fmt.Println(greet("Mustafiz"))

}

func isEven(n int) bool {
	return n%2 == 0
}

func greet(name string) string {
	return fmt.Sprintf("Hello, %s! Welcome to Go 101.", name)
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
