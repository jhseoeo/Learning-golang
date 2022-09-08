package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func div(numerator, denominator int) int { // parameters and each type of this, and the return type specified here
	// func div(numerator int, denominator int) int { same with above
	// if there is no return type specified (as like main function), no return statement is needed in the function body.
	if denominator == 0 {
		return 0
	}
	return numerator / denominator
}

// ----------------------------------------------------------------------------------
// In Go, we can't pass the parameters as default value (optional parameter) and named parameters
// you should supply all the parameters for a function

// to emulate optional parameters, using struct which fields are matched with desired parameters, and pass it as parameter.

type MyFuncOpts struct {
	FirstName string
	LastName  string
	Age       int
}

// Named and optional parameters are useful when a function has too many parameters and we have to manage this.
// However, functions in a good program have just a few parameters.
func MyFunc(opts MyFuncOpts) string {
	return opts.FirstName + " " + opts.LastName + "/" + strconv.Itoa(opts.Age)
}

// ----------------------------------------------------------------------------------

// variadic parameter
func addTo(base int, vals ...int) int { // put three dots(...) before type to declare a parameter as variadic
	var res int
	for _, v := range vals {
		res += v
	}
	return res
}

// ----------------------------------------------------------------------------------

// Go allows multiple return values
// types of multiple return values are listed in parantheses, seperated by commas
func divAndRemainder(numerator int, denomiator int) (int, int, error) {
	if denomiator == 0 {
		return 0, 0, errors.New("cannot divide by zero")
	} else {
		return numerator / denomiator, numerator % denomiator, nil // must return all of return values, without parantheses
		// if there is no error, just return nil for error.
	}
}

// ----------------------------------------------------------------------------------

// we can specify names for return values
func divAndRemainder_Named(numerator int, denomiator int) (result int, remainder int, err error) {
	// named return values are initialized by its zero value.
	// even if there is a single named return value, we must surround with parentheses.
	if denomiator == 0 {
		err = errors.New("cannot divide by zero")
		return result, remainder, err
		// return 0, 0, errors.New("cannot divide by zero") // this statement is also legal. it is not essential to use name of return value
	} else {
		result = numerator / denomiator
		remainder = numerator % denomiator
		return result, remainder, err
	}
	// the name of variables that is allocated by this return values can differ from name of return values
	// it can clearify your code, but also can be shadowed
}

// ----------------------------------------------------------------------------------

// when we specified names for return values, we can write only `return` without specifying values to return
func divAndRemainder_Blank(numerator int, denomiator int) (result int, remainder int, err error) {
	if denomiator == 0 {
		err = errors.New("cannot divide by zero")
		return // variables of named return value are returned directlu
	} else {
		result = numerator / denomiator
		remainder = numerator % denomiator
		return // blank return can reduce amount of typing, but it is less readable.
	}
	// this blank return makes it harder to understand data flow.
}

// ----------------------------------------------------------------------------------

func main() {
	res := div(8, 5)
	fmt.Println(res)

	fmt.Println("---------------------------")

	person1 := MyFunc(MyFuncOpts{ // pass the (anonymous) struct which fields are matched with desired parameters
		LastName: "Johnson",
		Age:      24,
	})
	person2 := MyFunc(MyFuncOpts{
		FirstName: "Junhyuk",
		LastName:  "Seo",
	})

	fmt.Println(person1, person2)

	fmt.Println("---------------------------")

	addVal1 := addTo(1, 2, 3, 4, 5)            // we can pass parameters as multiple parameters
	addVal2 := addTo(2, []int{4, 6, 8, 10}...) // we can pass parameters as slice, but must put three dots(...) after slice.
	fmt.Println(addVal1, addVal2)

	fmt.Println("---------------------------")

	// we should assign multiple return values of function into multiples variables
	result1, remainder1, err1 := divAndRemainder(5, 2) // if we try assigning multiple return values into a single variable, there will be a compile-time error
	result2, _, err2 := divAndRemainder_Named(5, 2)    // if we don't need to get remainder as variable, just using _, we can ignore it
	result3, remainder3, err3 := divAndRemainder_Blank(5, 2)

	if err1 != nil {
		fmt.Println(err1)
		os.Exit(1)
	} else {
		fmt.Println(result1, remainder1)
	}
	fmt.Println(err2, result2)
	fmt.Println(result3, remainder3, err3)

	fmt.Println("---------------------------")

}
