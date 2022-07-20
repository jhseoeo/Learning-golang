package main

import "fmt"

var asd = "aasd";

func main() {
    var a int = 1
	var b string
	var c, d = 123, "456"
	var (
		e int
		f = 20
		g, h = 1.23, "4.56"
		i, j string
	)

	k := 10
	k, l := 20, 20


	var 한글개꿀ㅋㅋ = 1

	fmt.Println(a, b, c, d, e, f, g, h, i, j, k, l)
	fmt.Println(asd);
	fmt.Println(한글개꿀ㅋㅋ)
}