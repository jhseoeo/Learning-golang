package main

import "fmt"

func main() {
	type person struct { // define struct type
		name string
		age  int
		pet  string
	}

	var john person // struct variable declaration
	var james = person{"James", 24, "cat"}
	kim := person{}                   // there is no difference on empty struct and zero value of struct
	fred := person{"Fred", 22, "dog"} // values are assigned to the fields in the order they were declared in struct definition
	beth := person{                   // using key names (recommended)
		age:  20,
		name: "Beth",
	}
	fred.pet = "parrot" // can use dotted notation

	fmt.Println(john, kim, james, fred, beth)

	fmt.Println("---------------------------")

	var human struct { // anonymous structs
		name string
		age  int
		pet  string
	}
	human.name = "Bob"
	human.age = 24
	human.pet = "dog"

	pet := struct { // directly initializing anonymous structs
		name string
		kind string
	}{
		name: "choco",
		kind: "dog",
	}

	fmt.Println(human, pet)

	fmt.Println("---------------------------")

	type firstPerson struct {
		name string
		age  int
	}
	f1 := firstPerson{"kim", 24}
	f2 := firstPerson{"lee", 25}
	fmt.Println(f1 == f2) // comparing two firstPerson instances is possible when they are composed of comparable types

	type secondPerson struct {
		name string
		age  int
	}
	s1 := secondPerson{"choi", 26}
	// fmt.Println(s1 == f1)     // comparing(==, !=) secondPerson with firstPerson is impossible
	fmt.Println(firstPerson(s1)) // convert secondPerson into firstPerson is possible because they have same fields

	type thirdPerson struct {
		age  int
		name string
	}
	// t1 := thirdPerson{27, "Park"}
	// fmt.Println(firstPerson(t1)) // convert thirdPerson into firstPerson is impossible because their fields have different order.

	type fourthPerson struct {
		firstName string
		age       int
	}
	// f3 := fourthPerson{"Kang", 28}
	// fmt.Println(firstPerson(f3)) // convert fourthPerson into firstPerson is impossible because their fields have different name.

	type fifthPerson struct {
		name  string
		age   int
		hobby string
	}
	// f4 := fifthPerson{"Seo", 24, "Cooking"}
	// fmt.Println(firstPerson(f4)) // convert fourthPerson into firstPerson is impossible because there is an additional field.

	var g struct {
		name string
		age  int
	}
	g = f1 // =, == are possible when anonymous struct have same field
	fmt.Println(g == f1)

	fmt.Println("---------------------------")
	// using a map and struct as a set
	// unlike using bool that takes 1 byte, using struct takes 0 byte
	intSet := map[int]struct{}{}
	vals := []int{5, 10, 2, 5, 8, 7, 3, 9, 1, 2, 10}

	for _, v := range vals {
		intSet[v] = struct{}{}
	}

	if _, ok := intSet[5]; ok { // it look messier than using bool
		fmt.Println("5 is in the set")
	}
	// if you want more functions like union, intersection, and subtraction, use third-party libraries
}
