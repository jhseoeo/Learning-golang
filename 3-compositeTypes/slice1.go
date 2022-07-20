package main

import "fmt"

func main() {
	var slice []int // declaration

    fmt.Println("---------------------------")

	// declare  with literal
	var x = []int{1, 2, 3}
	var y = []int{1, 5: 4, 6, 10: 100, 15} // [1, 0, 0, 0, 0, 4, 6, 0, 0, 0, 100, 15]

	fmt.Println(x)
	fmt.Println(y)

	fmt.Println("---------------------------")

	x[2] = 10
	fmt.Println(x[2]) // indexing by bracket

    fmt.Println("---------------------------")
		
	var multidimentional = [][]int{{1, 2, 3}, {1, 2, 3}} // multidimetional slice
	fmt.Println(multidimentional)

    fmt.Println("---------------------------")

	fmt.Println(slice) // empty slice is equal to  
	fmt.Println(slice == nil) // comparation between two slices occurs error; only possible comparation is the one between slice and nil

    fmt.Println("---------------------------")

	// capacity grows as it gets appended
	slice = append(slice, 10)
	fmt.Println(slice, len(slice), cap(slice))
	slice = append(slice, 20, 30, 40, 50)
	fmt.Println(slice, len(slice), cap(slice))
	slice = append(slice, x...)
	fmt.Println(slice, len(slice), cap(slice))
	// append(slice, x...) it occurs an error

    fmt.Println("---------------------------")

	initialized_capacity := make([]int, 0, 5) // type, length, capacity
	// initialized_capacity := make([]int, 6, 5) // it occurs an error (length > capacity)
	initialized_capacity = append(initialized_capacity, 5, 6, 7, 8)
	fmt.Println(initialized_capacity, len(initialized_capacity), cap(initialized_capacity))
}