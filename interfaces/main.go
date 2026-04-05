package main

import "fmt"

// ============================================================
// TOPIC: Interfaces - Implicit Implementation, Polymorphism
// ============================================================

// --- EXAMPLE ---

// Interface definition - any type with a Speak() method satisfies it
type Speaker interface {
	Speak() string
}

type Person struct {
	Name string
	Age  int
}

type User struct {
	Person
	ID string
}

type Robot struct {
	ID string
}

// User implements Speaker (implicitly - no "implements" keyword!)
func (u User) Speak() string {
	return fmt.Sprintf("Hi, I'm %s and I love Go!", u.Name)
}

// Robot also implements Speaker
func (r Robot) Speak() string {
	return fmt.Sprintf("Robot %s reporting for duty.", r.ID)
}

// This function accepts ANY type that implements Speaker
func present(s Speaker) {
	fmt.Println("Speaker says:", s.Speak())
}

func main() {
	user := User{
		Person: Person{Name: "Alice", Age: 30},
		ID:     "U123",
	}
	robot := Robot{ID: "XJ-9"}

	// Polymorphism - same function, different types
	present(user)
	present(robot)

	// Interface slice - mix different types
	speakers := []Speaker{user, robot}
	for _, s := range speakers {
		fmt.Println("-", s.Speak())
	}

	// Type assertion - get the concrete type back
	var s Speaker = user
	if u, ok := s.(User); ok {
		fmt.Println("It's a user:", u.Name)
	}

	// Type switch
	checkType(user)
	checkType(robot)
}

func checkType(s Speaker) {
	switch v := s.(type) {
	case User:
		fmt.Printf("User %s (age %d)\n", v.Name, v.Age)
	case Robot:
		fmt.Printf("Robot %s\n", v.ID)
	default:
		fmt.Println("Unknown speaker")
	}
}

// ============================================================
// PRACTICE QUESTIONS
// ============================================================
//
// Q1: Create an `Animal` struct and make it implement Speaker.
//     Pass it to the `present()` function.
//
// Q2: Create a `Shape` interface with methods:
//       Area() float64
//       Perimeter() float64
//     Implement it for Circle and Rectangle structs.
//     Write a function `printShape(s Shape)` that prints both.
//
// Q3: Create a `Stringer` interface with `String() string`.
//     (This is actually built into Go as fmt.Stringer!)
//     Implement it on Person so fmt.Println prints nicely.
//
// Q4: What happens if Robot is missing the Speak() method?
//     Try commenting it out - you'll get a compile error when
//     passing Robot to present(). This is Go's type safety!
//
// Q5: The empty interface `interface{}` (or `any` in Go 1.18+)
//     accepts any type. Write a function:
//       func describe(i any) string
//     that uses a type switch to describe the value.
//
// Q6: What is the difference between these two?
//       var s Speaker        // nil interface
//       var s Speaker = user // interface holding a User
//     (Hint: nil interface vs nil concrete value)
//
// ============================================================
