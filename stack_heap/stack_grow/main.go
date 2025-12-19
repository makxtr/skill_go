package main

//go:noinline
func foo() {
	bar()
}

//go:noinline
func bar() {
	arr := [2100]byte{}
	println(&arr)
}

func main() {
	a := 1
	println("before", &a)
	foo()
	println("after", &a)
}
