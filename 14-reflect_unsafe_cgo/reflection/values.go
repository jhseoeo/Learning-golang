package main

import (
	"fmt"
	"reflect"
)

func changeInt(i *int) {
	*i = 20
}

func changeIntReflect(i *int) {
	reflect.ValueOf(i).Elem().SetInt(20)
}

func main() {
	s := []string{"a", "b", "c"}
	sv := reflect.ValueOf(s)
	s2 := sv.Interface().([]string)
	fmt.Println(s2)

	// change value of a variable using reflection
	i := 10
	iv := reflect.ValueOf(&i)
	ivv := iv.Elem()
	ivv.SetInt(15)
	fmt.Println(i)

	i1, i2 := 10, 10

	changeInt(&i1)
	changeIntReflect(&i2)

	fmt.Println(i1, i2)

	i3 := 10
	reflect.ValueOf(i3).SetInt(1) // this calls will panic
}
