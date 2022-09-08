/*
As you know, goto is NOT recommended. Consider carefully before using it.
*/
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	a := 10
	// goto skip // cannot jump over variable declaration
	b := 20
	goto skip
skip:
	c := 30

	fmt.Println(a, b, c)
	if c > a {
		// goto inner // cannot jump into block
	}

	if a < b {
		goto inner
	inner:
		fmt.Println("a is less than b")
	}

	x := rand.Intn(10)

	for x < 100 {
		if a%5 == 0 {
			goto done // in this case, instead of using boolean flag, using goto makes code clearer and readable
		}
		a = a*2 + 1
	}

	fmt.Println("do something when the loop completes normally")
done:
	fmt.Println("do complicated stuff no matter why we left the loop")
	fmt.Println(a)
}
