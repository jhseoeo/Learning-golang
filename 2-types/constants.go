package main

import "fmt"

func main() {
	const a int = 1         // typed
	const b, c = 123, "456" // untyped, multiple declaration
	const (                 // decalaration constants using parentheses
		d    = 20           // inside parentheses, constants can be declared line by line
		e, f = 1.23, "4.56" // multiple decalaration
	)

	fmt.Println(a, b, c, d, e, f)

	var i int = a // constants can be value of variables when their types are same
	// var f float64 = a // if types are different, it occurs error
	fmt.Println(i)
}
