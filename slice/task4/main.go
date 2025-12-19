package main

import "fmt"

func main() {
	a := []int{}

	for i := range 3 {
		a = append(a, i+1)
	}
	println("cap(a)", cap(a))

	b := append(a, 4)
	c := append(b, 5)

	c[1] = 0

	fmt.Println("a =", a)
	fmt.Println("b =", b)
	fmt.Println("c =", c)
}
