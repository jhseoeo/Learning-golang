package main

import "fmt"

func main() {
	var i interface{} // empty interfece variable can store a value of any type

	i = 20
	fmt.Println(i)

	i = "hello"
	fmt.Println(i)

	i = struct {
		FirstName string
		LastName  string
	}{"Fred", "Fredson"}
	fmt.Println(i)
}
