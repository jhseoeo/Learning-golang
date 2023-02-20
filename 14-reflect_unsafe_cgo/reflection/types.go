package main

import (
	"fmt"
	"reflect"
)

type Foo struct {
	A int    `myTag:"value"`
	B string `myTag:"value2"`
}

func main() {
	x := 1
	xt := reflect.TypeOf(x)
	fmt.Println(xt.Name(), xt.Kind())

	f := Foo{}
	ft := reflect.TypeOf(f)
	fmt.Println(ft.Name(), ft.Kind())
	for i := 0; i < ft.NumField(); i++ { // iterating struct's field
		curField := ft.Field(i)
		fmt.Println(curField.Name, curField.Type.Name(), curField.Tag.Get("myTag"))
	}

	xpt := reflect.TypeOf(&x)
	fmt.Println(xpt.Name(), xpt.Kind(), xpt.Elem().Name(), xpt.Elem().Kind())

	tt := reflect.TypeOf(ft)
	fmt.Println(tt.Name(), tt.Kind(), tt.Elem().Name(), tt.Elem().Kind())

	tnt := reflect.TypeOf(ft.Name())
	fmt.Println(tnt.Name(), tt.Kind())

}
