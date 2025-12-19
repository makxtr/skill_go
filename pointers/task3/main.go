package main

import "fmt"

func changeSlice(s []int) {
	s = append(s, 4) // len=4, cap=4

	//*s = append(*s, 4)
}

func main() {
	s := make([]int, 3, 4) // len=3, cap=4 [1,2,3],0]
	for i := range 3 {
		s[i] = i + 1
	}

	changeSlice(s) // len=3, cap=4, [1,2,3,],4]

	fmt.Println(s)

	fmt.Println(s[3:4])
}
