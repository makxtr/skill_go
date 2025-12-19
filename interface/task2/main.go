package main

import (
	"errors"
	"fmt"
)

type SomeError struct{}

func (s SomeError) Error() string {
	return "some error"
}

var SomeErrorN = errors.New("Something went wrong")

func foo() error {
	var result *SomeError
	
	//result = SomeError

	return result
}

func main() {
	result := foo()

	if result != nil {
		fmt.Println("Error occurred!!!", result)
	}
}
