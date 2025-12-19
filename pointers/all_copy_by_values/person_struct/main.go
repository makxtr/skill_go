package main

import "fmt"

type Person struct {
	name string
	age  int
}

func (p *Person) changeName(name string) {
	fmt.Printf("%v\n", p)
	fmt.Printf("%p\n", p)
	fmt.Printf("%v\n", *p)
	fmt.Printf("%v\n", &p)
	//p = &Person{
	//	name: name,
	//}

	//*p = Person{name: name}

	p.name = name

}

func main() {
	p := Person{"Bob", 20}

	fmt.Printf("%p\n", &p)
	p.changeName("Alice")

	println(p.name)
}
