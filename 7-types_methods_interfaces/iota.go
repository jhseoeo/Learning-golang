package main

import "fmt"

func main() {
	type MailCategory int
	const (
		Uncategorized MailCategory = iota // 0
		Personal                          // 1
		Spam                              // 2
		Social                            // 3
		Advertisement                     // 4
	)

	fmt.Println(Uncategorized, Personal, Spam, Social, Advertisement)

	fmt.Println("---------------------------")

	type BitField int
	const (
		Field1 BitField = 1 << iota // assigned 1
		Field2                      // assigned 2
		Field3                      // assigned 4
		Field4                      // assigned 8
		_                           // passed 16
		Field6                      // assigned 32
	)

	fmt.Println(Field1, Field2, Field3, Field4, Field6)
}
