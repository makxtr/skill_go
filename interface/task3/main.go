package main

import "fmt"

type ABC interface {
	A()
	B()
	C()
}

type AB interface {
	A()
	B()
}

type BC interface {
	B()
	C()
}

type abc struct{}

func (a abc) A() {}
func (a abc) B() {}
func (a abc) C() {}

type ab struct{}

func (a ab) A() {}
func (a ab) B() {}

func main() {
	var a interface{}
	a = ab{}
	a1, ok := a.(ABC)
	if ok {
		fmt.Println(a1)
	}

	var b interface{}
	b = abc{}

	ab := b.(BC)
	//ab.A()

	bc := ab.(AB)

	bc.A()

}
