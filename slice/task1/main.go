package main

import "fmt"

func main() {
	a := []int{1, 2, 3}

	println("cap(a) =", cap(a))

	b := append(a, 4)
	c := append(a, 5)

	c[1] = 0

	fmt.Println("a =", a)
	fmt.Println("b =", b)
	fmt.Println("c =", c)
}
