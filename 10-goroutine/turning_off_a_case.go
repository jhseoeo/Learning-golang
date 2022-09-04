package main

import "fmt"

func main() {
	in := make(chan int)
	in2 := make(chan int)
	done := make(chan struct{})

	for {
		select {
		case v, ok := <-in:
			if !ok {
				in = nil // this case will never succeed again
				continue
			}
			fmt.Println(v)

		case v, ok := <-in2:
			if !ok {
				in2 = nil // this case will never succeed again
				continue
			}
			fmt.Println(v)

		case <-done:
			return
		}
	}
}
