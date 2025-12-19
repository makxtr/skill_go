package main

import "io"

func readAll(a []io.Reader) {

}

func convert(a []io.ReadWriter) []io.Reader {
	r := make([]io.Reader, len(a))

	for _, v := range a {
		r = append(r, v)
	}

	return r
}

func main() {
	var a = []io.ReadWriter{}

	readAll(convert(a))
}
