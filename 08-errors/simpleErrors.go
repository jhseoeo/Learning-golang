package main

import (
	"errors"
	"fmt"
)

func doubleEven1(i int) (int, error) {
	if i%2 != 0 {
		return 0, errors.New("only even numbers are processed")
	}
	return i * 2, nil
}

func doubleEven2(i int) (int, error) {
	if i%2 != 0 {
		return 0, fmt.Errorf("%d isn't an even number", i)
	}
	return i * 2, nil
}

func main() {
	i1, err := doubleEven1(3)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(i1)

	i2, err := doubleEven2(3)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(i2)
}
