package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 100; i++ {
		go func(val int) {
			fmt.Println(val)
		}(i)
	}
}

// ключевой момент в том что main может завершится быстрее
