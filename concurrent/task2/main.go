package main

import (
	"fmt"
)

func main() {
	counter := 0

	for i := 0; i < 100; i++ {
		go func() {
			counter++
		}()
	}

	fmt.Println(counter)
}
