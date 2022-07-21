package main

import "fmt"

func main() {
	var arr [3]int // array declaration
	fmt.Println(arr)

	fmt.Println("---------------------------")

	// declare with literal
	var x = [3]int{1, 2, 3}                  // [1, 2, 3]
	var y = [12]int{1, 5: 4, 6, 10: 100, 15} // [1, 0, 0, 0, 0, 4, 6, 0, 0, 0, 100, 15]
	var z = [...]int{4, 5, 6, 7, 8}          // [4, 5, 6, 7, 8]

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)

	fmt.Println("---------------------------")

	x[2] = 10
	fmt.Println(x[2]) // indexing by bracket

	fmt.Println("---------------------------")

	var multidimentional = [2][3]int{{1, 2, 3}, {1, 2, 3}} // multidimetional array
	fmt.Println(multidimentional)

	fmt.Println("---------------------------")

	fmt.Println(len(z)) // get size(length) of array
	fmt.Println(len(multidimentional[1]))
	fmt.Println(len([...]int{123, 456, 789}))

	fmt.Println("---------------------------")

	// var size int = 5				// can not specify size of array with variables.
	// var arr = [size]int{1, 2, 3} // it occurs an error.
}
