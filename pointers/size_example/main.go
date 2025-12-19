package main

import (
	"fmt"
	"unsafe"
)

// go run -gcflags="-m -l" main.go

// 0xc00009000a     0xc000092038
// 	  1 byte			8 byte
// 0xc00009000b     0xc000092040

func main() {
	a := false
	b := true
	c := &a
	d := &b

	fmt.Printf("value = %v, address = %p\n", a, &a)
	fmt.Printf("value = %v, address = %p\n", b, &b)
	fmt.Printf("value = %v, address = %p\n", c, &c)
	fmt.Printf("value = %v, address = %p\n", d, &d)

	diff := int(uintptr(unsafe.Pointer(&c)) - uintptr(unsafe.Pointer(&d)))
	fmt.Printf("Diff: %d byte\n", diff)
}
