package main

import (
	"fmt"
	"reflect"
)

func Filter(slice []string, filter func(string) bool) []string {
	res := []string{}
	for _, s := range slice {
		if filter(s) {
			res = append(res, s)
		}
	}
	return res
}

func ReflectedFilter(slice interface{}, filter interface{}) interface{} {
	sv := reflect.ValueOf(slice)
	fv := reflect.ValueOf(filter)

	sliceLen := sv.Len()
	out := reflect.MakeSlice(sv.Type(), 0, sliceLen)
	for i := 0; i < sliceLen; i++ {
		curVal := sv.Index(i)
		values := fv.Call([]reflect.Value{curVal})
		if values[0].Bool() {
			out = reflect.Append(out, curVal)
		}
	}
	return out.Interface()
}

func main() {
	names := []string{"Andrew", "Bob", "Clara", "Hortense"}
	longNames := ReflectedFilter(names, func(s string) bool {
		return len(s) > 3
	}).([]string)
	fmt.Println((longNames))

	ages := []int{20, 50, 13}
	adults := ReflectedFilter(ages, func(age int) bool {
		return age >= 18
	}).([]int)
	fmt.Println(adults)
}
