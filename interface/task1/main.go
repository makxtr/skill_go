package main

import "fmt"

type SomeStruct struct{}

func foo() interface{} {
	var result *SomeStruct

	return result
}

func main() {
	res := foo()

	if res != nil {
		fmt.Println("res != nil res =", res)
	}
}
