package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type myMutex struct {
	locked int64
}

func (m *myMutex) Lock() {
	for {
		if atomic.CompareAndSwapInt64(&m.locked, 0, 1) {
			return
		}
	}
}

func (m *myMutex) Unlock() {
	atomic.StoreInt64(&m.locked, 0)
}

func main() {

	wg := &sync.WaitGroup{}
	wg.Add(1000)

	mu := &myMutex{}

	c := 0

	for range 1000 {
		go func() {

			wg.Done()

			mu.Lock()
			c++
			mu.Unlock()
		}()
	}

	wg.Wait()

	fmt.Println(c)
}
