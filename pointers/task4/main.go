package main

import "fmt"

func changeMap(m map[string]int) {
	m["Bob"] = 52
}

func main() {
	s := make(map[string]int)
	s["Bob"] = 25

	changeMap(s)

	fmt.Println(s)
}
