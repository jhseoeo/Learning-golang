package main

import (
	"fmt"
	"reflect"
)

func main() {
	var stringType = reflect.TypeOf((*string)(nil)).Elem()
	var stringSliceType = reflect.TypeOf([]string(nil))

	sv := reflect.New(stringType).Elem()
	sv.SetString("hello")

	ssv := reflect.MakeSlice(stringSliceType, 0, 10)
	ssv = reflect.Append(ssv, sv)

	ss := ssv.Interface().([]string)
	fmt.Println(ss)
}
