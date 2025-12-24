package main

import (
	"context"
	"fmt"
	"strconv"
	"sync"
	"time"
)

type Request struct {
	Payload string
}

type Client interface {
	SendRequest(ctx context.Context, request Request) error
	WithLimiter(ctx context.Context, requests []Request)
}

type client struct {
}

func (c client) SendRequest(ctx context.Context, request Request) error {
	time.Sleep(100 * time.Millisecond)
	fmt.Println("sending request", request.Payload)
	return nil
}

// limitation connection count
// limitation connection goroutines
// limitation rps X ticker X
var rps = 100

func (c client) WithLimiter(ctx context.Context, reqs []Request) {
	ticker := time.Tick(time.Second / time.Duration(rps))

	wg := sync.WaitGroup{}

	wg.Add(len(reqs))
	for _, req := range reqs {
		<-ticker
		go func() {
			defer wg.Done()
			c.SendRequest(ctx, req)
		}()
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
}
