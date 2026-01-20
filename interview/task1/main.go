package main

import (
	"context"
	"fmt"
	"math/rand"
	"os/signal"
	"syscall"
	"time"
)

const (
	maxDepth      = 10 // maximum recursion depth of the flow
	maxGoroutines = 20 // soft goroutine limit
)

var currentGoroutines int

func main() {
	// Main goroutine waits for SIGINT / SIGTERM to terminate.
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	results := make(chan string)

	// Single goroutine which is waiting on the result channel.
	go resultLogger(results)

	// Start initial "flow" goroutine that can spawn others.
	go spawnFlow(ctx, results, 0, 0)

	fmt.Println("app started; send SIGTERM (kill -TERM) or Ctrl+C to stop")

	// Block main goroutine until we get a signal.
	<-ctx.Done()
	fmt.Println("signal received, shutting down...")

	close(results)

	time.Sleep(500 * time.Millisecond)
	fmt.Println("main exit")
}

// resultLogger is a single goroutine consuming from the results channel.
func resultLogger(results <-chan string) {
	for msg := range results {
		fmt.Println("[result]", msg)
	}
	fmt.Println("resultLogger done")
}

// spawnFlow represents a "flow" goroutine.
// Each such goroutine:
//  1. waits for some timeout,
//  2. writes to the result channel,
//  3. may spawn more goroutines of the same function up to some limit.
func spawnFlow(ctx context.Context, results chan<- string, id int, depth int) {
	fmt.Println("spawnFlow started", id, depth)

	if depth > maxDepth {
		return
	}

	currentGoroutines++
	defer func() {
		currentGoroutines--
	}()

	delay := time.Duration(rand.Intn(8)) * time.Second

	select {
	case <-ctx.Done():
		return
	case <-time.After(delay):
	}

	results <- fmt.Sprintf("goroutine id=%d depth=%d delay=%v active=%d",
		id, depth, delay, currentGoroutines)

	children := rand.Intn(6)

	for i := 0; i < children; i++ {
		if currentGoroutines >= maxGoroutines {
			return
		}

		childID := rand.Int()
		go spawnFlow(ctx, results, childID, depth+1)
	}
}
