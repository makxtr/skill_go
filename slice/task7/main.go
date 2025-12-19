package main

import "fmt"

func main() {
	s := make([]int, 0, 5)
	s = append(s, 1, 2, 3)

	subSlice := s[1:3]

	subSlice[0] = 99
	subSlice = append(subSlice, 4)

	s = append(s, 5, 6, 7)

	subSlice[1] = 100

	fmt.Println("s:", s)
	fmt.Println("subSlice", subSlice)
}
