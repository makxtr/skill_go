package main

import "fmt"

func readFromChan(ch chan string) {
	val := <-ch
	fmt.Println(val)
}

func main() {
	ch := make(chan string)

	go func() {
		ch <- " Bob"
	}()

	readFromChan(ch)
}
