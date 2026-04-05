package main

import "fmt"

// ============================================================
// TOPIC: Structs - Composition, Embedding, Nested Structs
// ============================================================

// --- EXAMPLE ---

type Address struct {
	City    string
	Country string
}

type Person struct {
	Name    string
	Age     int
	Address Address // nested struct
}

type Employee struct {
	Person       // embedded struct (field promotion)
	Role   string
	Salary int
}

func main() {
	// Struct literal
	person := Person{
		Name: "Rob",
		Age:  45,
		Address: Address{
			City:    "San Francisco",
			Country: "USA",
		},
	}
	fmt.Printf("Person: %+v\n", person)
	fmt.Println("City:", person.Address.City)

	// Struct embedding - fields are promoted
	employee := Employee{
		Person: person,
		Role:   "Engineer",
		Salary: 160000,
	}

	// Access promoted fields directly
	fmt.Println("Name:", employee.Name)           // promoted from Person
	fmt.Println("City:", employee.Address.City)    // promoted through Person
	fmt.Printf("Full: %s, %s, $%d\n", employee.Name, employee.Role, employee.Salary)

	// Structs are value types (copied on assignment)
	p2 := person
	p2.Name = "Alice"
	fmt.Println("Original:", person.Name)  // still "Rob"
	fmt.Println("Copy:", p2.Name)          // "Alice"
}

// ============================================================
// PRACTICE QUESTIONS
// ============================================================
//
// Q1: Create a `Book` struct with Title, Author, and Pages.
//     Create an instance and print it.
//
// Q2: Create a `Library` struct that has a Name and a slice
//     of Books ([]Book). Add 3 books and loop through them.
//
// Q3: Create a `Student` struct that embeds `Person` and adds
//     a `GPA` field. Show that you can access Name directly.
//
// Q4: Write a function `newPerson(name string, age int) *Person`
//     that returns a pointer to a new Person. Why might you
//     return a pointer instead of a value?
//
// Q5: What happens when you compare two structs with ==?
//     Try comparing two Person values. Does it work?
//     (Hint: it works if all fields are comparable)
//
// Q6: Create an anonymous struct inline:
//     config := struct{ Host string; Port int }{"localhost", 8080}
//
// ============================================================
