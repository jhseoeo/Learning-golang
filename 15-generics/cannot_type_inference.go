package main

import "fmt"

type Integer interface {
	int | int8 | int16 | int32 | int64 |
		uint | uint8 | uint16 | uint32 | uint64
}

func Convert[T1, T2 Integer](in T1) T2 {
	return T2(in)
}

// Invaliid!
func PlusOneThousand[T Integer](in T) T {
	return in + 1000
}

// Valid!
func PlusOneHundred[T Integer](in T) T {
	return in + 100
}

func main() {
	var a int = 10
	//b1 := Convert(a) // occurs error
	b2 := Convert[int, int64](a)
	//var b3 int64 = Convert(a) // occurs error

	fmt.Println(b2)
}
