package main

import "fmt"

//type Stack[T any] struct {
type Stack[T comparable] struct {
	vals []T
}

func (s *Stack[T]) Push(val T) {
	s.vals = append(s.vals, val)
}

func (s *Stack[T]) Pop() (T, bool) {
	if len(s.vals) == 0 {
		var zero T
		return zero, false
	} else {
		top := s.vals[len(s.vals)-1]
		s.vals = s.vals[:len(s.vals)-1]
		return top, true
	}
}

func (s *Stack[T]) Contains(val T) bool {
	for _, v := range s.vals {
		if v == val {
			return true
		}
	}
	return false
}

func main() {
	var intStack Stack[int]
	intStack.Push(10)
	intStack.Push(12)
	intStack.Push(14)
	//intStack.Push("nope") // this occurs error
	fmt.Println(intStack.Contains(10))
	fmt.Println(intStack.Contains(20))
}
