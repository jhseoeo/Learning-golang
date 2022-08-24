package main

import (
	"fmt"
)

type Sentinel string

func (s Sentinel) Error() string {
	return string(s)
}

const (
	ErrFoo = Sentinel("foo error")
	ErrBar = Sentinel("bar error")
)

func doubleEven1(i int) (int, error) {
	if i%2 != 0 {
		return 0, ErrFoo
	}
	return i * 2, nil
}

func main() {
	i, err := doubleEven1(21)
	if err == ErrFoo {
		fmt.Println(ErrFoo)
	}
	fmt.Println(i)
}
