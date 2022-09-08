package main

import (
	"fmt"
	"time"
)

func main() {
	a := time.Now()
	b := time.Now()

	d := a.Sub(b)
	c := a.Add(d)
	e := a.AddDate(1, 2, 3)

	fmt.Println(d)
	fmt.Println(c)
	fmt.Println(e)
}
