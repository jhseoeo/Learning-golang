package main

import "fmt"

func main() {
	// declaration of channel(unbuffered)
	ch := make(chan int)

	chw := make(chan<- int) // write-only channel
	chr := make(<-chan int) // read-only channel

	ch <- 1   // write the value in 1 to ch
	b := <-ch // reads a value from ch and assigns it to b
	fmt.Println(b)

	fmt.Println("---------------------------")

	chb := make(chan int, 10) // buffered channel

	fmt.Println("---------------------------")

	// for-rage loop using a channel
	for v := range ch {
		fmt.Println(v)
	}

	fmt.Println("---------------------------")

	// closing channel
	close(ch)

	// when reading data from closed channel, using comma ok idiom
	v1, ok := <-ch

	fmt.Println("---------------------------")

	select {
	case v := <-ch:
		fmt.Println(v)
	case v := <-chr:
		fmt.Println(v)
	case chw <- 5:
		fmt.Println("wrote")
	case <-chb:
		fmt.Println("got value on chb, but ignored it")
	}

	fmt.Println("---------------------------")

	select {
	case v := <-ch:
		fmt.Println("read from ch:", v)
	default:
		fmt.Println("no value written to ch")
	}

}
