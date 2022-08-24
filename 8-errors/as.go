package main

import (
	"errors"
	"fmt"
)

type MyErr struct {
	Message string
	ErrCode int
}

func (me MyErr) Error() string {
	return me.Message
}

func (me MyErr) Code() int {
	return me.ErrCode
}

func AFunctionThatReturnsAnError() error {
	return MyErr{Message: "Hi", ErrCode: 123}
}

func main() {
	err1 := AFunctionThatReturnsAnError()
	var myErr MyErr

	if errors.As(err1, &myErr) {
		fmt.Println(err1)
	}

	fmt.Println("---------------------------")

	err2 := AFunctionThatReturnsAnError()
	var coder interface {
		Code() int
	}

	if errors.As(err2, &coder) {
		fmt.Println(coder.Code())
	}
}
