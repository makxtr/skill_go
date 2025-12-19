package main

import "fmt"

func main() {
	// nil map
	var m map[string]int

	m["A"] = 1
	m["B"] = 2
	m["C"] = 3

	fmt.Println(m)

	// nil chan
	var ch chan int
	go func() {
		for i := range 5 {
			ch <- i
		}
	}()

	for val := range ch {
		fmt.Println(val)
	}

	// nil func
	var fn func()
	fn()
}
