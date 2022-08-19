package main

import (
	"fmt"
	"time"
)

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

// defining methods for user-defined type
func (p Person) String() string { // The receiver appears between the keyword func and the name of the method
	return fmt.Sprintf("%s %s, age %d", p.FirstName, p.LastName, p.Age)
}

// ----------------------------------------------------------------------------------

type Counter struct {
	total       int
	lastUpdated time.Time
}

// pointer receiver should be used when the method modifies the receiver or handles nil instances
func (c *Counter) Increment() {
	c.total++
	c.lastUpdated = time.Now()
}

// value receiver can be used when the method doesn't modify the receiver
func (c Counter) String() string {
	return fmt.Sprintf("total: %d, last updated: %v", c.total, c.lastUpdated)
}

func doUpdateWrong(c Counter) {
	c.Increment()
	fmt.Println("in doUpdateWrong:", c.String())
}

func doUpdateRight(c *Counter) {
	c.Increment()
	fmt.Println("in doUpdateRight:", c.String())
}

// ----------------------------------------------------------------------------------

// an implementation of a binary tree that takes advantage of nil values for the receiver
type IntTree struct {
	val         int
	left, right *IntTree
}

func (it *IntTree) Insert(val int) *IntTree {
	if it == nil { // case that handles when the receiver is a nil instance
		return &IntTree{val: val} // cannot assign its address into receiver directly
	}
	if val < it.val {
		it.left = it.left.Insert(val)
	} else if val > it.val {
		it.right = it.right.Insert(val)
	}
	return it
}

func (it *IntTree) Contains(val int) bool {
	switch {
	case it == nil: // case that handles when the receiver is a nil instance
		return false
	case val < it.val:
		return it.left.Contains(val)
	case val > it.val:
		return it.right.Contains(val)
	default:
		return true
	}
}

// ----------------------------------------------------------------------------------

func main() {
	p := Person{
		FirstName: "Fred",
		LastName:  "Fredson",
		Age:       52,
	}
	output := p.String()
	fmt.Println(output)

	fmt.Println("---------------------------")

	var c Counter
	fmt.Println(c.String())
	// Go automatically converts it to a pointer type.
	c.Increment() //c.Increment() is converted to (&c).Increment()
	fmt.Println(c.String())

	var cc Counter
	doUpdateWrong(cc)
	fmt.Println("in main:", cc.String())
	doUpdateRight(&cc)
	fmt.Println("in main:", cc.String())

	fmt.Println("---------------------------")

	var it *IntTree
	it = it.Insert(5)
	it = it.Insert(3)
	it = it.Insert(10)
	it = it.Insert(2)
	fmt.Println(it.Contains(2))  // true
	fmt.Println(it.Contains(12)) // false
}
