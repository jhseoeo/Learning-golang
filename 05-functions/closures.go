package main

import (
	"fmt"
	"sort"
)

// this function returns a function
func makeMult(base int) func(int) int {
	return func(factor int) int {
		return base * factor
	}
}

// ----------------------------------------------------------------------------------

func main() {
	type Person struct {
		FirstName string
		LastName  string
		Age       int
	}

	people := []Person{
		{"Kimkim", "Kim", 25},
		{"Junhyuk", "Seo", 24},
		{"Leelee", "Lee", 26},
	}

	// we can pass functions as parameter in Go
	sort.Slice(people, func(i int, j int) bool { // sort.Slice sorts the slice using function that is passed in
		return people[i].Age < people[j].Age // sorting by Age field
	})
	fmt.Println(people)

	sort.Slice(people, func(i int, j int) bool {
		return people[i].LastName < people[j].LastName
	})
	fmt.Println(people)

	fmt.Println("---------------------------")

	// initialize variables by functions
	twoBase := makeMult(2)
	threeBase := makeMult(3)

	for i := 0; i < 3; i++ {
		fmt.Println(twoBase(i), threeBase(i))
	}
}
