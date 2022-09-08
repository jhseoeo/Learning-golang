package main

import (
	"errors"
	"fmt"
	"os"
	"reflect"
)

func fileChecker(name string) error {
	f, err := os.Open(name)
	if err != nil {
		return fmt.Errorf("in fileChecker: %w", err)
	}
	defer f.Close()
	return nil
}

// ----------------------------------------------------------------------------------

type MyErr struct {
	Codes []int
}

func (me MyErr) Error() string {
	return fmt.Sprintf("codes: %v", me.Codes)
}

func (me MyErr) Is(target error) bool {
	if me2, ok := target.(MyErr); ok {
		return reflect.DeepEqual(me, me2)
	}
	return false
}

// ----------------------------------------------------------------------------------

type ResourceErr struct {
	Resource string
	Code     int
}

func (re ResourceErr) Error() string {
	return fmt.Sprintf("%s: %d", re.Resource, re.Code)
}

func (re ResourceErr) Is(target error) bool {
	if other, ok := target.(ResourceErr); ok {
		ignoreResource := other.Resource == ""
		ignoreCode := other.Code == 0
		matchResource := other.Resource == re.Resource
		matchCode := other.Code == re.Code

		return matchCode && matchResource ||
			matchCode && ignoreResource ||
			ignoreCode && matchResource
	}
	return false
}

// ----------------------------------------------------------------------------------

func main() {
	err := fileChecker("nonExistFile.txt")
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			fmt.Println(err)
		}
	}

	fmt.Println("---------------------------")

	err1 := ResourceErr{
		Resource: "Database",
		Code:     123,
	}

	err2 := ResourceErr{
		Resource: "Network",
		Code:     456,
	}

	if errors.Is(err1, ResourceErr{Resource: "Database"}) {
		fmt.Println("The database is broken:", err1)
	}

	if errors.Is(err2, ResourceErr{Resource: "Database"}) {
		fmt.Println("The database is broken:", err2)
	}

	if errors.Is(err1, ResourceErr{Code: 123}) {
		fmt.Println("Code 123 is triggered:", err1)
	}

	if errors.Is(err2, ResourceErr{Code: 123}) {
		fmt.Println("Code 123 is triggered:", err2)
	}

	if errors.Is(err1, ResourceErr{Resource: "Database", Code: 123}) {
		fmt.Println("Database is broken and Code 123 is triggered", err1)
	}

	if errors.Is(err1, ResourceErr{Resource: "Network", Code: 123}) {
		fmt.Println("Network is broken and Code 123 is triggered", err2)
	}
}
