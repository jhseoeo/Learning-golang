package main

import (
	"fmt"
	"reflect"
)

func hasNoValue(i interface{}) bool {
	iv := reflect.ValueOf(i)
	if !iv.IsValid() {
		return true
	}

	switch iv.Kind() {
	case reflect.Ptr, reflect.Slice, reflect.Map, reflect.Func, reflect.Interface:
		return iv.IsNil()
	default:
		return false
	}
}

func main() {
	var s interface{} = "asd"
	var p *string

	fmt.Println(hasNoValue(s), hasNoValue(p), hasNoValue(nil))
}
