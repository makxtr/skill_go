package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

// Написать 3 функции
// writer - генерит числа от 1 до 10
// doubler - умножает числа на 2 имитируя работу (500 мс)
// reader - читает и выводит на экран
func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	reader(double(writer(ctx)))
}

func writer(ctx context.Context) <-chan int {
	c := make(chan int)
	go func() {
		for {
			select {
			case <-ctx.Done():
				close(c)
				return
			default:
				c <- rand.Intn(10)
			}
		}
	}()

	return c
}

func double(c <-chan int) <-chan int {
	dbl := func(i int) int {
		time.Sleep(500 * time.Millisecond)
		return i << 1
	}

	r := make(chan int)

	go func() {
		for e := range c {
			r <- dbl(e)
		}
		close(r)
	}()

	return r
}

func reader(r <-chan int) {
	for e := range r {
		fmt.Println(e)
	}
}
