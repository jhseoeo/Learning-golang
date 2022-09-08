package main

import (
	"fmt"

	"github.com/learning-go/9/circular_dependency/person"
	"github.com/learning-go/9/circular_dependency/pet"
)

func main() {
	var a pet.Pet
	var b person.Person
	fmt.Println(a, b)
}
