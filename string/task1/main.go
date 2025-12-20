package main

import "fmt"

func main() {
	greetings := "привет World!"

	fmt.Println(len(greetings))

	fmt.Printf("%v %b %c \n", greetings[1], greetings[1], greetings[1])

	r := rune(greetings[1])
	fmt.Println(r)
}
