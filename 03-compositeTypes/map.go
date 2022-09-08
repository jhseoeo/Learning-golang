package main

import "fmt"

func main() {
	var nilMap map[string]int
	mamap := map[string]int{}
	fmt.Println(nilMap, mamap, nilMap == nil, mamap == nil)

	fmt.Println("---------------------------")

	var teams = map[string][]string{
		"Orcas":   []string{"Fred", "Ralph"},
		"Lions":   []string{"Sarah", "Peter"},
		"Kittens": {"Waldo", "Raul"}, // 이렇게 해도 됨
	}
	fmt.Println(teams)

	fmt.Println("---------------------------")

	ages := make(map[int][]string, 10)
	fmt.Println(ages, len(ages)) // len returns the number of key-value pairs of map

	fmt.Println("---------------------------")

	totalWins := map[string]int{}
	totalWins["Orcas"] = 1 // assigning value into specific key with =, not :=
	totalWins["Lions"] = 2
	fmt.Println(totalWins["Orcas"])
	fmt.Println(totalWins["Kittens"]) // an initial value of int is zero.
	totalWins["Kittens"]++            // value of "Kittens" would be 1
	totalWins["Lions"] = 3
	fmt.Println(totalWins["Kittens"])
	fmt.Println(totalWins["Lions"])

	fmt.Println("---------------------------")

	m := map[string]int{
		"Hello": 5,
		"world": 0,
	}

	// comma idiom
	v, ok := m["Hello"] // v gets value of the given key
	fmt.Println(v, ok)  // ok gets if the key exists in map

	v, ok = m["world"] // though v gets zero value,
	fmt.Println(v, ok) // ok gets true because "world" exists in map

	v, ok = m["nono"]  // "nono" doesn't exist in map, so ok gets false
	fmt.Println(v, ok) // v gets zero value,

	delete(m, "Hello") // deletes key "Hello" from the map
	v, ok = m["Hello"] // the key doesn't exist in map as it is deleted
	fmt.Println(v, ok)

	fmt.Println("---------------------------")

	intSet := map[int]bool{} // using a map as a set
	vals := []int{5, 10, 2, 5, 8, 7, 3, 9, 1, 2, 10}

	for _, v := range vals {
		intSet[v] = true
	}
	fmt.Println(len(vals), len(intSet))
	fmt.Println(intSet[5])
	fmt.Println(intSet[500])
	if intSet[100] {
		fmt.Println("100 is in the set")
	}
	// if you want more functions like union, intersection, and subtraction, use third-party libraries
}
