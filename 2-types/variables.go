package main

import "fmt"

var asd = "aasd"

func main() {
	var a int = 1         // variable declaration
	var b string          // declare string variable. it's value would be zero value of string ("")
	var c, d = 123, "456" // multiple declaration. their type would be int, string
	var (                 // inside parentheses, variables can be declared line by line
		e    int            //
		f    = 20           // it's type would be int
		g, h = 1.23, "4.56" // multiple declaration, only specified their values
		i, j string         // only specified their types
	)

	k := 10        // variable declaration by := operator
	k, l := 20, 20 // multiple declaration

	var 한글개꿀ㅋㅋ = 1 // 한글도 됩니다

	fmt.Println(a, b, c, d, e, f, g, h, i, j, k, l)
	fmt.Println(asd)
	fmt.Println(한글개꿀ㅋㅋ)
}
