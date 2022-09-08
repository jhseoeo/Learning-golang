package main

import "fmt"

type MyInt int

func typeSwitch(i interface{}) {
	switch i.(type) {
	case nil:
		fmt.Println("nil")
	case int:
		fmt.Println("int")
	case MyInt:
		fmt.Println("MyInt")
	case string:
		fmt.Println("string")
	case bool, rune:
		fmt.Println("bool or rune")
	default:
		fmt.Println("what is this")
	}
}

func main() {
	var a int = 12
	var b MyInt = 23
	var c *string
	d := "asdasd"
	e := map[string]string{
		"hi": "there",
	}
	var f interface{}

	typeSwitch(a)
	typeSwitch(b)
	typeSwitch(c)
	typeSwitch(d)
	typeSwitch(e)
	typeSwitch(f)
}
