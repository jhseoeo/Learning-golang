package main

import "fmt"

func Map[T1, T2 any](s []T1, f func(T1) T2) []T2 {
	r := make([]T2, len(s))
	for i, v := range s {
		r[i] = f(v)
	}
	return r
}

func Filter[T any](s []T, f func(T) bool) []T {
	var r []T
	for _, v := range s {
		if f(v) {
			r = append(r, v)
		}
	}
	return r
}

func Reduce[T1, T2 any](s []T1, f func(T2, T1) T2, initial T2) T2 {
	r := initial
	for _, v := range s {
		r = f(r, v)
	}
	return r
}

// WE CANNOT IMPLEMENT GENERIC METHOD
//type FunctionalSlice[T any] []T
//
//func (fs FunctionalSlice[T]) Map[E any](f func(T) E) FunctionalSlice[E] {
//	out := make(FunctionalSlice[E], len(fs))
//	for i, v := range fs {
//		out[i] = f(v)
//	}
//	return out
//}
//
//func (fs FunctionalSlice[T]) Filter(f func(T) bool) FunctionalSlice[T] {
//	var out []T
//	for _, v := range fs {
//		if f(v) {
//			out = append(out, v)
//		}
//	}
//	return out
//}
//
//func (fs FunctionalSlice[T]) Reduce[E any](f func(E, T) E, start E) E {
//	out := start
//	for _, v := range fs {
//		out = f(out, v)
//	}
//	return out
//}

func main() {
	words := []string{"One", "Potato", "Two", "Potato"}

	filtered := Filter(words, func(s string) bool {
		return s != "Potato"
	})
	fmt.Println(filtered)

	lengths := Map(filtered, func(s string) int {
		return len(s)
	})
	fmt.Println(lengths)

	sum := Reduce(lengths, func(acc int, cur int) int {
		return acc + cur
	}, 0)
	fmt.Println(sum)

	//words = FunctionalSlice[string]{"One", "Potato", "Two", "Potato"}
	//sum := words.Filter(func(s string) bool {
	//	return s != "Potato"
	//}).Map(func(s string) int {
	//	return len(s)
	//}).Reduce(func(acc, cur int) int {
	//	return acc + cur
	//})
}
