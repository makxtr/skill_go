package main

import "fmt"

func modify(s []int, n int) {
	s = append(s, n)
	s[0] = 999

}

func main() {
	s1 := make([]int, 3, 5)
	s2 := s1[:2]

	s1[0] = 1
	s2[1] = 2

	modify(s1, 55)
	modify(s2, 66)

	fmt.Println("s1:", s1)
	fmt.Println("s2:", s2)
	fmt.Println("s1 cap:", cap(s1), "s2 cap:", cap(s2))
}
