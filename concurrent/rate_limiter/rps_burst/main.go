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
	fmt.Println("sending request", request.Payload)
	time.Sleep(1 * time.Second)
	return nil
}

// limitation connection count
// limitation connection goroutines
// limitation rps X ticker X
var rps = 1
var burst = 10

func (c client) WithLimiter(ctx context.Context, reqs []Request) {
	ticker := time.NewTicker(time.Second / time.Duration(rps))
	tickets := make(chan struct{}, burst)

	wg := sync.WaitGroup{}

	go func() {
		for range burst {
			tickets <- struct{}{}
		}
	}()

	go func() {
		for {
			select {
			case <-ticker.C:
				tickets <- struct{}{}
			case <-ctx.Done():
				return
			}
		}
	}()

	wg.Add(len(reqs))
	for _, req := range reqs {
		<-tickets
		go func() {
			defer wg.Done()

			ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
			defer cancel()

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
