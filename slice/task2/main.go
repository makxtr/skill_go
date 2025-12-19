package main

import "fmt"

func main() {
	a := []int{}

	a = append(a, []int{1, 2, 3, 4, 5, 6, 7}...)

	println("cap(a)", cap(a))
	
	b := append(a, 8)
	c := append(a, 9)

	c[1] = 0

	fmt.Println("a =", a)
	fmt.Println("b =", b)
	fmt.Println("c =", c)
}
