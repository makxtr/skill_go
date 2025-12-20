package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func worker(wg *sync.WaitGroup) {
	defer wg.Done()

	time.Sleep(1 * time.Millisecond)
}

func main() {
	runtime.GOMAXPROCS(2)
	MAX_TASKS := 10_000

	wg := &sync.WaitGroup{}
	wg.Add(MAX_TASKS)

	start := time.Now()
	for range MAX_TASKS {
		go worker(wg)
	}
	wg.Wait()
	fmt.Println(time.Since(start))
}
