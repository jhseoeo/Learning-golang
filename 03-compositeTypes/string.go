package main

import "fmt"

func main() {
	var s string = "Hello there"
	var b byte = s[5] // indexing like slice. byte can represent each 1-byte character of string
	fmt.Println(s, b, len(s))
	// slicing string
	fmt.Println(s[4:7])
	fmt.Println(s[5:])
	fmt.Println(s[:6])

	fmt.Println("---------------------------")

	var h string = "한글조아"
	fmt.Println(h, len(h))
	// each character of korean letter takes 3 bytes, slicing like below would not be done properly
	fmt.Println(h[2:])
	fmt.Println(h[:5])
	fmt.Println(h[2:7])

	fmt.Println("---------------------------")

	// rune, byte, int can represent each character of string. but int is not recommended
	var a1 rune = 'x'
	var a2 byte = 'y'
	var a3 int = 65
	var s1 string = string(a1)
	var s2 string = string(a2)
	var s3 string = string(a3)
	fmt.Println(s1, s2, s3)

	fmt.Println("---------------------------")

	// difference of rune and byte
	var ss string = "Hello 안녕"
	var bs []byte = []byte(ss) // []byte splits UTF-8 characters. usually use this.
	var rs []rune = []rune(ss) // []rune doesn't split
	fmt.Println(bs, rs)

}
