package main

import "fmt"

func main() {
	fmt.Println("Greetings and \n\"Salutations\"")
	fmt.Println(
	`Hello There!
I'm Here`)
    fmt.Println('\141', '\x61', '\u0061')

    fmt.Println("---------------------------")

	var s1 string = "qwe";
	var s2 string = "asd";

	fmt.Println(s1 == s2);
	fmt.Println(s1 > s2);
	fmt.Println(s1 + s2);

    fmt.Println("---------------------------")

	var character rune = '\u0061';
	fmt.Println(character)
}