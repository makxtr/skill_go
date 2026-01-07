package main

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type res struct {
	err error
	val int
}

var errTimeout = errors.New("timed out")

func processData(ctx context.Context, v int) res {
	timer := time.NewTimer(time.Duration(rand.Intn(10)) * time.Second)
	defer timer.Stop()

	select {
	case <-ctx.Done():
		return res{err: errTimeout, val: 0}
	case <-timer.C:
		return res{err: nil, val: v << 1}
	}
}

func main() {
	in := make(chan int)
	out := make(chan int)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	go func() {
		for i := range 10 {
			select {
			case in <- i + 1:
			case <-ctx.Done():
				return
			}
		}
		close(in)
	}()
	now := time.Now()
	processParallel(ctx, in, out, 5)
	for val := range out {
		fmt.Println(val)
	}
	fmt.Println(time.Since(now))
}

func processParallel(ctx context.Context, in <-chan int, out chan<- int, numWorkers int) {
	wg := &sync.WaitGroup{}

	for range numWorkers {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					return
				case v, ok := <-in:
					if !ok {
						return
					}
					r := processData(ctx, v)
					if r.err == nil {
						out <- r.val
					}
				}
			}
		}()
	}

	go func() {
		wg.Wait()
		close(out)
	}()
}
