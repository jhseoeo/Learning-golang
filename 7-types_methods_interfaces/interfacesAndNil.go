package main

import "fmt"

func main() {
	var s *string
	fmt.Println(s, s == nil) // zero value of pointer, nil
	var i interface{}
	fmt.Println(i, i == nil) // zero value of interface, also nil
	i = s                    // its value is nil, but its type is non-nil  => interface is non-nil
	fmt.Println(i, i == nil)
}
