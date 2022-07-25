package main

import (
	"fmt"
	"math/rand"
)

func main() {
	n1 := rand.Intn(10) // default(non-seed) rand number is always 1 (cuz it is hard-coded)
	if n1 == 0 {        // there is no parenthesis around the condition
		fmt.Println("too low")
	} else if n1 > 5 {
		fmt.Println("too big :", n1)
	} else {
		fmt.Println("good :", n1)
	}

	fmt.Println("---------------------------")

	if n2 := rand.Intn(10); n2 == 0 { // both declaring variable and checking condition
		fmt.Println("too low")
	} else if n2 > 5 {
		fmt.Println("too big :", n2)
	} else {
		fmt.Println("good :", n2)
	}

	// fmt.Println(n2) // it causes error. once a if/else statement ends, n1 is not accessable.

}
