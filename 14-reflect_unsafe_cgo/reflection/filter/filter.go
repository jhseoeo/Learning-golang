package filter

import (
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

func FilterInt(slice []int, filter func(int) bool) []int {
	res := []int{}
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
