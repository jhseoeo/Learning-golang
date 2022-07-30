package main

import "fmt"

type person struct {
	name string
	age  int
}

func modifyFails(i int, s string, p person) { // all the parameters are passed by value(copied), not referenced or aliased
	i *= 2
	s = "goodbye"
	p.name = "Bob" // even for the struct, cannot change the origin value by modifying parameters. 
}

// ----------------------------------------------------------------------------------

// slices and maps are passed passed by pointers
func modifyMap(m map[int]string) {
	// changing map parameters are reflected in the variables passed into the function
	m[2] = "hello"
	m[3] = "goodbye"
}

func modifySlice(s []int){
	// we can modify any element in the slice, but can't lengthen the slice
	for k, v := range s {
		s[k] = v * 2
	}
	s = append(s, 10) // this line actually did not append a value to origin variable
}

// ----------------------------------------------------------------------------------

func main() {
	p, i, s := person{}, 2, "hello"
	modifyFails(i, s, p) // this invocation of function can't make any change on variables above
	fmt.Println(i, s, p)

	fmt.Println("---------------------------")

	m := map[int]string {
		1: "first",
		2: "second",
	}

	sl := []int{1, 2, 3}

	modifyMap(m)
	modifySlice(sl)

	fmt.Println(m)
	fmt.Println(sl)
}