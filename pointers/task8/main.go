package main

import "fmt"

func named() (a, b int) {
	a, b = 1, 2

	defer func() {
		a = 10
		b = 20
	}()

	return a, b
}

func unnamed() (int, int) {
	a, b := 1, 2

	defer func() {
		a = 10
		b = 20
	}()

	return a, b
}

func main() {
	a, b := named()
	fmt.Println(a, b) // 10,20
	a, b = unnamed()
	fmt.Println(a, b) // 1, 2
}
