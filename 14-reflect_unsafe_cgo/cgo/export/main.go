package main

/*
   extern int add(int a, int b);
*/
import "C"
import "fmt"

//export doubler
func doubler(i int) int {
	return i * 2
}

func main() {
	fmt.Println(C.add(1, 2))
}
