package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	results := make(chan string)
	go fetchService(ctx, "payments", results)
	go fetchService(ctx, "analytics", results)
	go fetchService(ctx, "notifications", results)

	done := make(chan struct{})
	go func() {
		for res := range results {
			fmt.Println("result:", res)
		}
		close(done)
	}()

	select {
	case <-done:
		fmt.Println("all services responded")
	case <-ctx.Done():
		fmt.Println("timeout waiting for services:", ctx.Err())
	}

	pipeline := make(chan int)
	squares := make(chan int)

	go producer(pipeline, 5)
	go squareWorker(pipeline, squares)
	go consumer(squares)

	time.Sleep(3 * time.Second)
}

func fetchService(ctx context.Context, name string, out chan<- string) {
	delay := time.Duration(500+rand.Intn(1200)) * time.Millisecond

	select {
	case <-time.After(delay):
		out <- fmt.Sprintf("%s responded in %s", name, delay)
	case <-ctx.Done():
		fmt.Printf("%s canceled: %v\n", name, ctx.Err())
		return
	}
}

func producer(out chan<- int, count int) {
	defer close(out)
	for i := 1; i <= count; i++ {
		out <- i
		time.Sleep(200 * time.Millisecond)
	}
}

func squareWorker(in <-chan int, out chan<- int) {
	defer close(out)
	for n := range in {
		out <- n * n
	}
}

func consumer(in <-chan int) {
	for v := range in {
		fmt.Println("square:", v)
	}
}
