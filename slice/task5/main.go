package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5}
	// a = [1,2,[3,4],5] len=5, cap=5
	//
	b := a[2:4] // b =[1,2|3,4|,5] len=2, cap=3

	c := append(b, 10) // b=[1,2|3,4|,10] len=2,cap=3, a=[1,2,3,4,10] len=5, cap=5
	//       \
	// c=[3,4,10] = len3 cap3

	c[1] = 55 // c = [3,55,10] , b = [3,55], a = 1,2,3,55,10

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
