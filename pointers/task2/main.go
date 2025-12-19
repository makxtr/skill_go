package main

import "fmt"

func changeSlice(s []int) {
	s[1] = 10
}

func main() {
	s := []int{1, 2, 3}
	changeSlice(s)

	fmt.Println(s)

	// 1 2 3
}
