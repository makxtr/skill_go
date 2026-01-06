package start

import (
	"fmt"
	"math/rand"
	"time"
)

func processData(v int) int {
	time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
	return v << 1
}

func main() {
	in := make(chan int)
	out := make(chan int)

	go func() {
		for i := range 10 {
			in <- i
		}
		close(in)
	}()
	now := time.Now()
	processParallel(in, out, 5)
	for val := range out {
		fmt.Println(val)
	}
	fmt.Println(time.Since(now))
}

func processParallel(in <-chan int, out chan<- int, numWorkers int) {
}
