package main

import (
	"encoding/json"
	"fmt"
)

func failedUpdate(g *int) {
	x := 10
	g = &x // where the pointer is pointing at is changed
}

func update(g *int) {
	*g = 10 // dereferencing => success to change original value
}

// ----------------------------------------------------------------------------------

type Foo struct {
	foo int
	bar string
}

func MakeFoo1(f *Foo) error {
	// not +recommended format
	f.foo = 20
	f.bar = "val"
	return nil
}

// rather than using pointer parameter to pass a value, just return this.
func MakeFoo2() (Foo, error) {
	// recommended format
	f := Foo{
		foo: 20,
		bar: "val",
	}
	return f, nil
}

// ----------------------------------------------------------------------------------

func main() {
	var x *int
	var y int

	failedUpdate(x)
	failedUpdate(&y) // those two invocations cannot changed x and y's values

	fmt.Println(x, y)

	// update(x) // this invocation occurs an error. cannot dereference nil
	update(&y)

	fmt.Println(y)

	fmt.Println("---------------------------")

	a := Foo{}
	MakeFoo1(&a) // rather than using this pattern,

	b, _ := MakeFoo2() // use this pattern

	fmt.Println(a, b)

	fmt.Println("---------------------------")

	// exception for above pattern
	// when parsing JSON, functions like Unmarshal pass a value by pointer parameter
	f := struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}{}
	err := json.Unmarshal([]byte(`{"name": "Bob", "age": 30}`), &f)
	if err == nil {
		fmt.Println(f)
	}

}

// if the size of the data to return is more than a megabyte, it is way faster to pass pointer of data rather than to pass data itself
