package main

import "fmt"

func main() {

	pow := func(num int) int { // using keyword `func`, we can declare an anonyymous function
		// if we put a function name on anonymous function, it will occur a compile-time error
		return num * num
	}

	for i := 0; i < 5; i++ {
		func(j int) {
			fmt.Println("Printing", pow(j), "from inside of an anonymous function")
		}(i) // anonymous function are declared and called immediately
	}
}
