package main

import "fmt"

type Speaker interface {
	Speak() string
}

type User struct {
	Name string
}

func (u User) Speak() string {
	return fmt.Sprintf("Hi, I'm %s and I love Go!", u.Name)
}

type Robot struct {
	ID string
}

func (r Robot) Speak() string {
	return fmt.Sprintf("Robot %s reporting for duty.", r.ID)
}

func main() {
	present(User{Name: "Grace"})
	present(Robot{ID: "XJ-9"})
}

func present(s Speaker) {
	fmt.Println("Speaker says:", s.Speak())
}
