package main

import "fmt"

type Address struct {
	City    string
	Country string
}

type Person struct {
	Name    string
	Age     int
	Address Address
}

type Employee struct {
	Person
	Role   string
	Salary int
}

func main() {
	person := Person{
		Name: "Rob",
		Age:  45,
		Address: Address{
			City:    "San Francisco",
			Country: "USA",
		},
	}

	fmt.Printf("Person: %#v\n", person)

	employee := Employee{
		Person: person,
		Role:   "Engineer",
		Salary: 160000,
	}

	fmt.Printf("Employee name via promotion: %s\n", employee.Name)
	fmt.Printf("Employee city via embedding: %s\n", employee.Address.City)
	fmt.Printf("Full employee struct: %#v\n", employee)
}
