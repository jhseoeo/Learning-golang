package main

import "fmt"

type Status int

const (
	InvalidLogin Status = iota + 1
	NotFound
)

type StatusErr struct {
	Status  Status
	Message string
}

func (se StatusErr) Error() string {
	return fmt.Sprintf("%d - %s", se.Status, se.Message)
}

// ----------------------------------------------------------------------------------

func GenerateWrongError(flag bool) error {
	var genErr StatusErr
	if flag {
		genErr = StatusErr{
			Status: NotFound,
		}
	}
	return genErr
}

func GenerateError1(flag bool) error {
	if flag {
		return StatusErr{
			Status: NotFound,
		}
	}
	return nil
}

func GenerateError2(flag bool) error {
	var genErr error
	if flag {
		genErr = StatusErr{
			Status: NotFound,
		}
	}
	return genErr
}

func main() {
	err := GenerateWrongError(true)
	fmt.Println(err != nil)
	err = GenerateWrongError(false)
	fmt.Println(err != nil)
}
