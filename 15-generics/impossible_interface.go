package main

import "fmt"

type ImpossiblePrintableInt interface {
	int
	String() string
}

type ImpossibleStruct[T ImpossiblePrintableInt] struct {
	val T
}

type PrintableInt interface {
	~int
	String() string
}

type MyInt int

func (mi MyInt) String() string {
	return fmt.Sprint(mi)
}

func main() {
	s := ImpossibleStruct[int]{10}
	s2 := ImpossibleStruct[MyInt]{10}
	fmt.Println(s.val, s2.val)
}
