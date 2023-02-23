package main

import (
	"fmt"
	"reflect"
	"time"
)

func MakeTimedFunction(f interface{}) interface{} {
	ft := reflect.TypeOf(f)
	fv := reflect.ValueOf(f)

	wrapperF := reflect.MakeFunc(ft, func(args []reflect.Value) []reflect.Value {
		start := time.Now()
		out := fv.Call(args)
		end := time.Now()
		fmt.Println(end.Sub(start))
		return out
	})
	return wrapperF.Interface()
}

func timeMe(a int) int {
	var result int = 0
	for i := 0; i < a*10000; i++ {
		result += 1
	}
	return result
}

func main() {
	timed := MakeTimedFunction(timeMe).(func(int) int)

	fmt.Println(timed(100000))
}
