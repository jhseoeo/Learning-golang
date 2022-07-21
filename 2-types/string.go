package main

import "fmt"

func main() {
	fmt.Println("Greetings and \n\"Salutations\"") // line alignment by \n
	fmt.Println(                                   // strings that covered by `` can contain carrage return
		`Hello There! 
I'm Here`)

	fmt.Println("---------------------------")

	fmt.Println('\141', '\x61', '\u0061')                         // notations that indicates 97
	fmt.Println(string('\141'), string('\x61'), string('\u0061')) // convert into string, it prints 'a'

	fmt.Println("---------------------------")

	var s1 string = "qwe"
	var s2 string = "asd"

	fmt.Println(s1 == s2) // comparison operators between string
	fmt.Println(s1 > s2)
	fmt.Println(s1 + s2)

	fmt.Println("---------------------------")

	var character rune = '\u0061' // rune represents a single character
	fmt.Println(character)
}
