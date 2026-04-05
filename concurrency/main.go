package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

// ============================================================
// TOPIC: Concurrency - Goroutines, Channels, Context
// ============================================================

// --- EXAMPLE ---

func main() {
	// --- GOROUTINES + CHANNELS ---
	// Channels let goroutines communicate safely
	results := make(chan string, 3) // buffered channel

	go fetchService("payments", results)
	go fetchService("analytics", results)
	go fetchService("notifications", results)

	// Collect results
	for i := 0; i < 3; i++ {
		fmt.Println("Result:", <-results)
	}

	// --- CONTEXT WITH TIMEOUT ---
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	ch := make(chan string, 1)
	go func() {
		time.Sleep(500 * time.Millisecond)
		ch <- "done"
	}()

	select {
	case msg := <-ch:
		fmt.Println("Got:", msg)
	case <-ctx.Done():
		fmt.Println("Timeout:", ctx.Err())
	}

	// --- PIPELINE PATTERN ---
	pipeline := make(chan int)
	squares := make(chan int)

	go producer(pipeline, 5)
	go squareWorker(pipeline, squares)

	// Consumer runs in main goroutine
	for v := range squares {
		fmt.Println("Square:", v)
	}
}

func fetchService(name string, out chan<- string) {
	delay := time.Duration(100+rand.Intn(300)) * time.Millisecond
	time.Sleep(delay)
	out <- fmt.Sprintf("%s responded in %s", name, delay)
}

func producer(out chan<- int, count int) {
	defer close(out)
	for i := 1; i <= count; i++ {
		out <- i
	}
}

func squareWorker(in <-chan int, out chan<- int) {
	defer close(out)
	for n := range in {
		out <- n * n
	}
}

// ============================================================
// PRACTICE QUESTIONS
// ============================================================
//
// Q1: Launch 5 goroutines that each print their index.
//     Use sync.WaitGroup to wait for all to finish.
//
// Q2: Create a channel of ints. Send numbers 1-10 in a
//     goroutine, close the channel, then range over it
//     in main to print them.
//
// Q3: Write a fan-out pattern: one producer sends work to
//     3 worker goroutines via a shared channel.
//
// Q4: Use `select` with two channels and a timeout:
//       ch1 sends after 1s, ch2 sends after 2s,
//       timeout after 1.5s. What gets printed?
//
// Q5: What happens if you send to an unbuffered channel
//     with no receiver? (Answer: deadlock!)
//     Try: ch := make(chan int); ch <- 1
//
// Q6: Use context.WithCancel() to cancel a long-running
//     goroutine. The goroutine should check ctx.Done()
//     in a select loop.
//
// ============================================================
