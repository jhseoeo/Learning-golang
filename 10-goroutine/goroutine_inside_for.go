package main

import "fmt"

func main() {
	a := []int{2, 4, 6, 8, 10}
	ch1 := make(chan int, len(a))
	for _, v := range a {
		go func() {
			ch1 <- v * 2
		}()
	}

	for i := 0; i < len(a); i++ {
		fmt.Println(<-ch1)
	}

	fmt.Println("---------------------------")

	ch2 := make(chan int, len(a))
	for _, v := range a {
		v := v
		go func() {
			ch2 <- v * 2
		}()
	}

	for i := 0; i < len(a); i++ {
		fmt.Println(<-ch2)
	}

	fmt.Println("---------------------------")

	ch3 := make(chan int, len(a))
	for _, v := range a {
		go func(val int) {
			ch3 <- val * 2
		}(v)
	}

	for i := 0; i < len(a); i++ {
		fmt.Println(<-ch3)
	}
}
