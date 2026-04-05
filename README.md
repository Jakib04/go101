# Go 101 - Learn Go by Doing

A hands-on Go practice repo. Each folder covers a core topic with **working examples** and **practice questions** you can solve right in the code.

## Getting Started

```bash
# Make sure Go is installed
go version

# Run any module
cd basics && go run main.go
cd functions && go run main.go
# ... etc
```

## Topics

| # | Folder | What You'll Learn |
|---|--------|-------------------|
| 1 | [`basics/`](basics/main.go) | Variables, types, control flow, strings, defer |
| 2 | [`functions/`](functions/main.go) | Multiple returns, variadic args, closures |
| 3 | [`structs/`](structs/main.go) | Struct composition, embedding, nested structs |
| 4 | [`methods/`](methods/main.go) | Value vs pointer receivers |
| 5 | [`interfaces/`](interfaces/main.go) | Implicit interfaces, polymorphism, type assertions |
| 6 | [`datastructures/`](datastructures/main.go) | Arrays, slices, maps |
| 7 | [`concurrency/`](concurrency/main.go) | Goroutines, channels, context, pipelines |

## How to Practice

1. **Read** the example code at the top of each file
2. **Run** it to see the output
3. **Solve** the practice questions at the bottom of each file
4. Write your solutions directly in the same file and run again

## Quick Reference

```go
// Variables
name := "Go"                    // short declaration
var age int = 10                // explicit type

// Functions
func add(a, b int) (int, error) { return a + b, nil }

// Structs & Methods
type Dog struct { Name string }
func (d Dog) Bark() string { return "Woof!" }

// Interfaces (implicit!)
type Animal interface { Bark() string }

// Goroutines & Channels
ch := make(chan string)
go func() { ch <- "hello" }()
msg := <-ch

// Error handling
val, err := doSomething()
if err != nil { log.Fatal(err) }
```

## Prerequisites

- [Go 1.22+](https://go.dev/dl/)
- A code editor (VS Code with Go extension recommended)
