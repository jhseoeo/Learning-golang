package main

import "fmt"

type MyInt int

func main() {
	var i interface{}
	var mine MyInt = 20

	i = mine
	i2 := i.(MyInt) // using type assertion, we can confine the type of concrete type that the interface indicates
	fmt.Println(i2 + 1)

	// i3 := i.(string) // this line occurs a panic
	// fmt.Println(i3)

	// i4 := i.(int) // this line also occurs a panic
	// fmt.Println(i4)

	// ok is set to true if the type conversion was successful.
	// if it was not, ok is set to false and the other value is set to its zero value
	i5, ok := i.(int) // i5 is set to 0, ok is set to false
	if !ok {
		msg := fmt.Errorf("unexpected type for %v", i5)
		fmt.Println(msg)
	}

	i6, ok := i.(MyInt)
	if !ok { // i6 is set to 20, ok is set to true
		// this block is not reached
		msg := fmt.Errorf("unexpected type for %v", i6)
		fmt.Println(msg)
	}
}
