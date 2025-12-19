package main

import "fmt"

type Address struct {
	city   string
	street string
	house  int
}

func (a Address) setCity(city string) {
	a = Address{
		city: city,
	}
}

func (a *Address) setStreet(street string) {
	a.street = street
}

func setHouse(addr *Address, house int) {
	addr = &Address{
		house: house,
	}
}

func main() {
	addr := Address{
		city:   "New York",
		street: "Broadway",
		house:  10,
	}

	addr.setCity("London")
	addr.setStreet("Piccadilly")
	setHouse(&addr, 5)

	fmt.Println(addr)
}
