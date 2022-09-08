package main

import "fmt"

func main() {
	x1 := 10 // shadowed variable.
	if x1 > 5 {
		fmt.Println(x1) // 10 => x1 is not yet shadowed
		x1 := 5         // shadowing variable. from this line to end of this block
		fmt.Println(x1) // 5 => x1 is shadowed until if-blocks ends
	}
	fmt.Println(x1) // 10 => if-block ends. x1 is not shadowed

	fmt.Println("---------------------------")

	x2 := 10
	if x2 > 5 {
		x2, y2 := 5, 20
		fmt.Println(x2, y2)
	}

	fmt.Println(x2) // y2 is not accessable here

	fmt.Println("---------------------------")

	// NEVER shadow predeclared identifiers in universe blocks
	// true, int := 10, "hi"
	// fmt.Println(true, int) // true, int are not accessable until the main function ends
	// shadow tool can't even detect shadowing universe block identifiers

	// Do noet shadow package name
	// fmt := "oops"    // fmt is shadowed
	// fmt.Println(fmt) // there is no method fmt.Println because fmt is shadowed
	// fmt is not accessable until the main function ends
	// to detect shadowing variables, type `shadow ./...`

}
