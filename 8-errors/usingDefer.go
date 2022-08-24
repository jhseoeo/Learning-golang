package main

import (
	"errors"
	"fmt"
)

func doThing1(val int) (int, error) {
	if val%2 == 0 {
		return val, nil
	} else {
		return 0, errors.New("hi")
	}
}

func doThing2(val int) (int, error) {
	if val%2 == 0 {
		return val, nil
	} else {
		return 0, errors.New("bye")
	}
}

func doThing3(val1, val2 int) (int, error) {
	if (val1+val2)%4 == 0 {
		return val1 + val2, nil
	} else {
		return 0, errors.New("no")
	}
}

func doSomeThings_before(val1, val2 int) (int, error) {
	val3, err := doThing1(val1)
	if err != nil {
		return 0, fmt.Errorf("in DoSomeThings: %w", err)
	}

	val4, err := doThing2(val2)
	if err != nil {
		return 0, fmt.Errorf("in DoSomeThings: %w", err)
	}

	res, err := doThing3(val3, val4)
	if err != nil {
		return 0, fmt.Errorf("in DoSomeThings: %w", err)
	}
	return res, nil
}

func doSomeThings_after(val1, val2 int) (_ int, err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("in DoSomeThings: %w", err)
		}
	}()

	val3, err := doThing1(val1)
	if err != nil {
		return 0, err
	}

	val4, err := doThing2(val2)
	if err != nil {
		return 0, err
	}

	return doThing3(val3, val4)
}

func main() {
	val, err := doSomeThings_after(2, 4)
	fmt.Println(val, err)
}
