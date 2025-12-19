package main

import "fmt"

func changeName(name *string) {
	fmt.Printf("changeName value = %v, address = %p\n", name, &name)

	fmt.Printf("%t\n", name)
	fmt.Printf("%t\n", *name)

	//*name = "Alice"
}

func main() {
	name := "Bob"
	fmt.Printf("main value = %v, address = %p\n", name, &name)
	changeName(&name)
	//println(name)
}
