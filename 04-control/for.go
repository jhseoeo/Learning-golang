package main

import (
	"fmt"
)

func main() {

	// c-style for statement.
	for i := 0; i < 10; i++ { // no parenthesis as like if-statements
		// initialization : should use := to initialize variables. `var` is not legal here
		// comparison : must be an expression that evaluates to a bool
		// increment : similar to other languages
		fmt.Print(i)
	}
	fmt.Println()

	fmt.Println("---------------------------")

	// condition-only statements
	i := 1
	for i < 100 { // it is like while statements in other languages
		fmt.Print(i, " ")
		i = i * 2
	}
	fmt.Println()

	fmt.Println("---------------------------")

	// infinite loop and break, continue statements
	j := 0
	for { // no expression after for keyword
		if j++; j > 10 {
			break // if there is no break statement, loop will last until a keyboard interrupt(ctrl-c) occurs
		} else if j%2 == 0 {
			continue //
		}
		fmt.Print(i, " ")
	}
	fmt.Println()

	fmt.Println("---------------------------")

	// for-range loop (array, slice, string)
	evenVals := []int{2, 4, 6, 8, 10, 12}
	for i, v := range evenVals {
		fmt.Println(i, v) // i is an key(index) of the data, v is value of the data
	}

	// ignoring the key in a for-range loop
	for _, v := range evenVals { // using underscore(_), we can only access the value
		fmt.Print(v, " ") // if you want, it is also possible to ignore the value by using underscore
	}

	fmt.Println("---------------------------")

	// for-range loop (map)
	names := map[string]bool{
		"Fred": true,
		"Raul": true,
		"Will": false,
	}

	for k, v := range names {
		fmt.Println(k, v)
	}

	for k := range names { // by leave off second variable, it is possible to get key only.
		fmt.Println(k, names[k])
	}

	fmt.Println("---------------------------")

	// everytime you run for-range loop over a map, order of the key/values varies.
	for i := 0; i < 3; i++ {
		for k, v := range names {
			fmt.Println(k, v)
		}
		fmt.Println()
	}
	fmt.Println(names) // exception of above rule. this always prints map in ascending order to make it easier to debug and log

	fmt.Println("---------------------------")

	// for-range loop (string)
	samples := []string{"hello", "안녕하세요"}

	for _, sample := range samples {
		for i, r := range sample { // it iterates over the runes, not bytes
			fmt.Println(i, r, string(r)) // key is the number of byte from the beginning of the string, type of value is rune
		}
	}

	fmt.Println("---------------------------")

	for _, v := range evenVals {
		v = v * 2
	}
	fmt.Println(evenVals) // modifying value doesn't change the original data..
}
