package main

import "fmt"

func main() {
	var x int32 = 10
	var y bool = true

	// pointer is a variable that holds the location in memory where a value is stored
	pointerX := &x       // pointerX stores first byte address of x. if x stored at location 10 to 13, pointerX will point to 10
	pointerY := &y       // & is the address operator. it returns the address of the memory location where the value is stored
	var pointerZ *string // because it doesn't point to anything, pointerZ has zero value(nil)

	fmt.Println(pointerX, pointerY, pointerZ) // addresses of variables
	fmt.Println(pointerZ == nil)              // referencing *pointerZ occurs an runtime error(panics), because pointerZ is nil
	fmt.Println(*pointerX, *pointerY)         // * is indirection operator.
	*pointerX, *pointerY = 8, false
	fmt.Println(x, y) // by changing pointer's values, we can change origin value

	fmt.Println("---------------------------")

	var a = new(int)      // built-in function `new` creates a pointer variable
	fmt.Println(a == nil) // false
	fmt.Println(*a)       // it points zero value of given type
	*a = 2
	fmt.Println(*a)
}

// if we need to distinguish variables that contains zero value or is not assigned any value, use a pointer and make it direct to nil
// this pattern is useful when you are handling JSON
// remember that pointers also indicate mutability, so be careful when using this pattern
