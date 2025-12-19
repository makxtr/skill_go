package main

import "fmt"

type Person struct {
	name string
	age  int
}

func changeName(p *Person, name string) {
	p.name = name
}

func getPerson() (p Person) {
	//var p Person

	defer changeName(&p, "Alice") // will be executed after return

	p = Person{
		name: "Bob",
		age:  52,
	}

	return p
}

func main() {
	p := getPerson()
	fmt.Println(p)
}
