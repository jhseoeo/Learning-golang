package main

import "fmt"

type Foo struct {
	x int
	S string
}

func (f Foo) Hello() string {
	return "Hello"
}

func (f Foo) goodbye() string {
	return "goodbye"
}

type Bar = Foo

func main() {
	bar := Bar{
		x: 20,
		S: "Hello",
	}
	var f Foo = bar

	fmt.Println(f.Hello(), f.goodbye())
}
