package main

import (
	"fmt"
	"strings"
)

// ============================================================
// TOPIC: Methods - Value vs Pointer Receivers
// ============================================================

// --- EXAMPLE ---

type Notebook struct {
	Title  string
	Notes  []string
	Closed bool
}

// Value receiver - does NOT modify the original
func (n Notebook) Count() int {
	return len(n.Notes)
}

// Value receiver - reads but does not modify
func (n Notebook) Summary() string {
	return fmt.Sprintf("%s: %s", n.Title, strings.Join(n.Notes, ", "))
}

// Pointer receiver - DOES modify the original
func (n *Notebook) Add(note string) {
	if n.Closed {
		fmt.Println("notebook is closed; cannot add note")
		return
	}
	n.Notes = append(n.Notes, note)
}

// Pointer receiver - modifies state
func (n *Notebook) Close() {
	n.Closed = true
}

func main() {
	nb := &Notebook{Title: "Go Webinar"}

	nb.Add("functions return multiple values")
	nb.Add("struct embedding promotes fields")

	fmt.Println("notes count:", nb.Count())
	fmt.Println(nb.Summary())

	nb.Close()
	nb.Add("this will not be added") // closed!
}

// ============================================================
// PRACTICE QUESTIONS
// ============================================================
//
// Q1: Create a `Counter` struct with a `Value int` field.
//     Add methods:
//       - Increment() - pointer receiver, adds 1
//       - Reset()     - pointer receiver, sets to 0
//       - Get() int   - value receiver, returns Value
//
// Q2: Create a `Rectangle` struct with Width and Height.
//     Add methods:
//       - Area() float64
//       - Perimeter() float64
//       - Scale(factor float64) - pointer receiver
//
// Q3: Why does Add() use a pointer receiver (*Notebook) but
//     Count() uses a value receiver (Notebook)?
//     (Answer: Add modifies the struct, Count only reads it)
//
// Q4: What happens if you change Add() to a value receiver?
//     Try it and see - the notes won't actually be saved!
//
// Q5: Add a method `Remove(index int)` to Notebook that
//     removes a note by index (use pointer receiver).
//     Hint: use append(n.Notes[:i], n.Notes[i+1:]...)
//
// ============================================================
