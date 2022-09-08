package main

import "fmt"

type Adder struct {
	start int
}

func (a Adder) AddTo(val int) int {
	return a.start + val
}

// ----------------------------------------------------------------------------------

type Person struct {
	name string
	age  int
}
type Score int
type HighScore Score
type Employee Person

// ----------------------------------------------------------------------------------

func main() {
	myAdder := Adder{start: 10}
	fmt.Println(myAdder.AddTo(5)) // prints 15

	f1 := myAdder.AddTo // We can also assign the method to a variable or pass it to a parameter of type func(int)int
	fmt.Println(f1(10)) // prints 20

	f2 := Adder.AddTo
	fmt.Println(f2(myAdder, 15)) // prints 25

	var i int = 300
	var s Score = 100
	var hs HighScore = 200
	// hs = s            // compilation error!
	// s = i             // compilation error!
	s = Score(i)      // ok
	hs = HighScore(s) // ok
	fmt.Println(hs + hs)
}
