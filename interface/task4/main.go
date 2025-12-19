package main

import "fmt"

func main() {
	var x = func() int { return 1 }
	x = nil

	test(x)
}

func test(x interface{}) {
	switch x.(type) {
	case int:
		fmt.Println("int", x)
	case string:
		fmt.Println("string", x)
	case nil:
		fmt.Println("nil", x)
	case func() int:
		fmt.Println("func() int", x())
	default:
		fmt.Println("unknown type", x)
	}
}
