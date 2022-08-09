package main

import "fmt"

type Foo struct {
	foo int
	bar string
}

type person struct {
	FirstName  string
	MiddleName *string
	LastName   string
}

func stringp(s string) *string {
	// helper function that returns address of parameter variable
	return &s
}

func main() {
	x := &Foo{}
	// you cannot use an & before primitive literal(numbers, booleans, strings, etc..) or constants because they don't have memory addressess
	// z := &"string" // this statement occurs an error
	var y string
	z := &y // to point to a primitive type, declare a variable first

	fmt.Println(x, z)

	fmt.Println("---------------------------")

	p := person{
		FirstName: "Pat",
		// MiddleName: "Perry",  // or
		// MiddleName: &"Perry", // this lines won't compile
		MiddleName: stringp("Perry"), // helper function turned a constant value into a pointer
		LastName:   "Peterson",
	}

	fmt.Println(p)
}
