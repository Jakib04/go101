package main

import (
	"fmt"
	"strings"
)

const course = "Go 101"

func main() {
	fmt.Println("Welcome to", course)

	name := "Ada"
	greeting := fmt.Sprintf("Hello, %s!", strings.ToUpper(name))
	fmt.Println(greeting)

	result, err := add(21, 21)
	if err != nil {
		fmt.Println("add error:", err)
		return
	}

	switch {
	case result > 50:
		fmt.Println("Big number:", result)
	case result == 42:
		fmt.Println("The answer:", result)
	default:
		fmt.Println("Result:", result)
	}

	for i := 0; i < 3; i++ {
		defer fmt.Println("deferred:", i)
		fmt.Println("loop iteration:", i)
	}

	user := User{Name: "Grace", Age: 37}
	user.Greet()
	introduce(user)
}

func add(a, b int) (int, error) {
	sum := a + b
	if sum < 0 {
		return 0, fmt.Errorf("sum overflowed: %d", sum)
	}
	return sum, nil
}

type User struct {
	Name string
	Age  int
}

func (u User) Greet() {
	fmt.Printf("Hi, I'm %s and I'm %d.\n", u.Name, u.Age)
}

func (u User) Speak() string {
	return fmt.Sprintf("I'm %s and I love Go.", u.Name)
}

type Speaker interface {
	Speak() string
}

func introduce(s Speaker) {
	fmt.Println("Speaker says:", s.Speak())
}
