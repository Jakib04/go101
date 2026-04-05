package main

import (
	"fmt"
	"sort"
)

// ============================================================
// TOPIC: Data Structures - Arrays, Slices, Maps
// ============================================================

// --- EXAMPLE ---

type Developer struct {
	Name  string
	Stack []string
}

func main() {
	// --- ARRAYS (fixed size) ---
	languages := [3]string{"Go", "Rust", "Zig"}
	fmt.Println("Array:", languages)

	// --- SLICES (dynamic size) ---
	topics := []string{"goroutines", "channels"}
	topics = append(topics, "testing", "profiling")
	fmt.Println("Slice:", topics)
	fmt.Println("Length:", len(topics), "Capacity:", cap(topics))

	// Slice a slice
	first2 := topics[:2]
	fmt.Println("First two:", first2)

	// --- MAPS ---
	cfg := map[string]int{
		"maxWorkers":     4,
		"port":           8080,
		"timeoutSeconds": 30,
	}

	// Check if key exists (comma-ok idiom)
	if max, ok := cfg["maxWorkers"]; ok {
		fmt.Println("Workers:", max)
	}

	// Delete a key
	delete(cfg, "timeoutSeconds")

	// Iterate and sort keys
	keys := make([]string, 0, len(cfg))
	for k := range cfg {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	fmt.Println("Config keys:", keys)

	// --- SLICE OF STRUCTS ---
	registry := []Developer{
		{Name: "Linus", Stack: []string{"Go", "C"}},
		{Name: "Ken", Stack: []string{"Go", "Plan 9"}},
	}

	for i, dev := range registry {
		fmt.Printf("%d: %s knows %v\n", i, dev.Name, dev.Stack)
	}
}

// ============================================================
// PRACTICE QUESTIONS
// ============================================================
//
// Q1: Create a slice of integers. Append 5 numbers, then print
//     the slice, its length, and its capacity.
//
// Q2: Write a function `contains(slice []string, target string) bool`
//     that checks if a string exists in a slice.
//
// Q3: Create a map[string][]string where keys are languages
//     and values are frameworks. Example:
//       "Go" -> ["Gin", "Echo", "Fiber"]
//     Add entries and loop through them.
//
// Q4: Write a function `wordCount(text string) map[string]int`
//     that counts how many times each word appears.
//     Hint: use strings.Fields() to split text into words.
//
// Q5: What's the difference between:
//       var s []int         // nil slice
//       s := []int{}        // empty slice
//       s := make([]int, 5) // slice with length 5
//     Try printing len() and cap() for each.
//
// Q6: Copy a slice without sharing the underlying array:
//       original := []int{1, 2, 3}
//       copied := make([]int, len(original))
//       copy(copied, original)
//     Modify `copied` and verify `original` is unchanged.
//
// ============================================================
