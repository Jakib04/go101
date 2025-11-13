package main

import (
	"fmt"
	"sort"
)

type Developer struct {
	Name  string
	Stack []string
}

type Person struct {
	Name string
	Age  int
}

type Employee struct {
	Person
	Role string
}

func main() {
	languages := [3]string{"Go", "Rust", "Zig"}
	fmt.Println("Array:", languages)

	topics := []string{"goroutines", "channels"}
	topics = append(topics, "testing", "profiling")
	fmt.Println("Slice:", topics, "len", len(topics), "cap", cap(topics))

	cfg := map[string]int{
		"maxWorkers":     4,
		"port":           8080,
		"timeoutSeconds": 30,
	}

	if max, ok := cfg["maxWorkers"]; ok {
		fmt.Println("workers:", max)
	}

	keys := make([]string, 0, len(cfg))
	for k := range cfg {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	fmt.Println("Config keys:", keys)

	dev := Developer{
		Name:  "Linus",
		Stack: []string{"Go", "C", "Assembly"},
	}
	fmt.Println("Developer:", dev)

	registry := make([]Developer, 0)
	registry = append(registry,
		dev,
		Developer{Name: "Ken", Stack: []string{"Go", "Plan 9"}},
	)

	for idx, d := range registry {
		fmt.Printf("%d: %s knows %v\n", idx, d.Name, d.Stack)
	}

	manager := Employee{
		Person: Person{Name: "Rob", Age: 45},
		Role:   "Engineering Manager",
	}
	fmt.Printf("Employee: %s (%d) - %s\n", manager.Name, manager.Age, manager.Role)
	fmt.Println("Embedded person:", manager.Person)
}
