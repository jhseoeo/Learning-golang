package main

import "fmt"

func main() {
    const a int = 1 // typed
	const b, c = 123, "456" // untyped
	const (
		d = 20
		e, f = 1.23, "4.56"
	)

	fmt.Println(a, b, c, d, e, f)

	var i int = a
	// var f float64 = a // occurs error

}