package main

import (
	"context"
	"fmt"
	"strconv"
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
	time.Sleep(1 * time.Second)
	fmt.Println("sending request", request.Payload)
	return nil
}

// limitation connection count
// limitation connection goroutines (10000) - X - sem,tokens - X
// limitation rps
var maxGo = 100

func (c client) WithLimiter(ctx context.Context, reqs []Request) {
	tokens := make(chan struct{}, maxGo)
	go func() {
		for range maxGo {
			tokens <- struct{}{}
		}
	}()

	for _, req := range reqs {
		<-tokens
		go func() {
			defer func() { tokens <- struct{}{} }()

			c.SendRequest(ctx, req)
		}()
	}

	for range maxGo {
		<-tokens
	}
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
