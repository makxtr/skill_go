package main

import "fmt"

func spawnMessages(n int) chan string {
	ch := make(chan string, 1)

	for i := 0; i < n; i++ {
		ch <- fmt.Sprintf("msg %d", i+1)
	}

	return ch
}

func main() {
	n := 10

	for msg := range spawnMessages(n) {
		fmt.Println("received:", msg)
	}
}
