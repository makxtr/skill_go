package main

import (
	"fmt"
	"sync"
)

func fanin(chans ...<-chan int) <-chan int {

}

func main() {
	ch1 := generator(3, 4)
	ch2 := generator(5, 6)

	combined := fanin(ch1, ch2)

	for v := range combined {
		fmt.Println("Получено:", v)
	}
}

func generator(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}
