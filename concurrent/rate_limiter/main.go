package main

import (
	"context"
	"fmt"
	"runtime"
	"strconv"
	"sync"
	"time"
)

type Request struct {
	Payload string
}

type Client interface {
	SendRequest(ctx context.Context, request Request) error
	WithLimiter(ctx context.Context, requests []Request) error
}

type client struct {
}

func (c client) SendRequest(ctx context.Context, request Request) error {
	time.Sleep(1 * time.Second)
	fmt.Println("sending request", request.Payload)
	return nil
}

// limitation connection count - X worker pool X
// limitation connection goroutines
// limitation rps

var rps = 100

func (c client) WithLimiter(ctx context.Context, requests []Request) {
	ticker := time.NewTicker(time.Second / time.Duration(rps))
	defer ticker.Stop()

	wg := sync.WaitGroup{}
	for _, req := range requests {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			wg.Add(1)
			go func() {
				defer wg.Done()
				c.SendRequest(ctx, req)
			}()
		}
	}
	wg.Wait()
}

func main() {
	ctx := context.Background()
	c := client{}
	requests := make([]Request, 1000)
	for i := 0; i < 1000; i++ {
		requests[i] = Request{Payload: strconv.Itoa(i)}
	}
	c.WithLimiter(ctx, requests)

	printMemUsage()
}

func generate(reqs []Request) chan Request {
	ch := make(chan Request)

	go func() {
		for _, req := range reqs {
			ch <- req
		}
		close(ch)
	}()

	return ch
}

func printMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("Alloc = %v MiB", m.Alloc/1024/1024)
	fmt.Printf("\tTotalAlloc = %v MiB", m.TotalAlloc/1024/1024)
	fmt.Printf("\tSys = %v MiB", m.Sys/1024/1024)
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}
