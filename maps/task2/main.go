package main

import (
	"fmt"
	"sync"
)

func main() {
	m := make(map[int]int)

	wg := &sync.WaitGroup{}
	wg.Add(100)

	for i := range 100 {
		go func() {
			defer wg.Done()
			m[i] = i
		}()
	}

	wg.Wait()

	fmt.Println(m)
}
