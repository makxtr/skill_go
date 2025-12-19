package main

// go build -gcflags "-m -m"

//go:noinline
func foo() {
	bar()
	baz()
}

//go:noinline
func bar() *int {
	b := 10
	return &b
}

//go:noinline
func baz() int {
	c := 20
	return c
}

func main() {
	foo()
}
