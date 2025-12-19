package main

import "fmt"

func getBytes(start, end int) []byte {
	arr := [999999999]byte{}
	slice := arr[start:end]

	return slice
}

func main() {
	s := getBytes(10, 20)
	
	fmt.Println(s)
}
