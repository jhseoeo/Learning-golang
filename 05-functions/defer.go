package main

import (
	"fmt"
	"os"
)

// `defer` usually used to clean up temporary resources such as files or network connections
func example() int {
	defer func() int {
		return 2 // there is no way to read this value
	}() // this statement is called after the example() ends
	return 555
}

// ----------------------------------------------------------------------------------

// common pattern that allocates a resource to also return a closure that cleans up the resource
func getFile(name string) (*os.File, func(), error) {
	file, err := os.Open(name)
	if err != nil {
		return nil, nil, err
	} else { // it returns resource and a closure that cleans up the resource
		return file, func() { file.Close() }, nil
	}
}

// ----------------------------------------------------------------------------------

func main() {
	n := example()
	fmt.Println(n)

	fmt.Println("---------------------------")

	// _, closer, err := getFile(os.Args[1])
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer closer() // releases the resource by .using defer and closer function

	fmt.Println("---------------------------")

	j := 2

	defer func(i int) {
		fmt.Println(i)
	}(j)

	j++
	fmt.Println(j)
}
