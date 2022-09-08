package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4}
	y := make([]int, 4)

	num := copy(y, x)   // copy(destination, source). x is copied into y
	fmt.Println(y, num) // num: the number of elements copied (decided by length of slices)
	y[2] = 1            // if we change any value of y,
	fmt.Println(y, x)   // values of x still unchanged (doesn't share memory spaces)

	fmt.Println("---------------------------")

	z := make([]int, 2, 4) // because length of z is 2,
	num = copy(z, x)       // when it is copied, only two of x are copied.
	fmt.Println(z, num)    // [1, 2] 2

	fmt.Println("---------------------------")

	w := make([]int, 4, 4) //
	num = copy(w, x[:2])   // because x[:2] has only two element (length is 2),
	fmt.Println(w, num)    // x[:2] is copied into first two elements

	fmt.Println("---------------------------")

	q := []int{1, 2, 3, 4}
	num = copy(q[:3], q[1:]) // [2, 3, 4] is copied into [1, 2, 3]
	fmt.Println(q, num)      // [2, 3, 4, 4]

	fmt.Println("---------------------------")

	ex := []int{1, 2, 3, 4}
	ed := [...]int{5, 6, 7, 8}
	ey := make([]int, 2)
	copy(ey, ed[:]) // whether use [:] to get slice, same result
	fmt.Println(ey)
	copy(ed[:], ex) // whether use [:] to get slice, same result. (ex is successfully copied into ed)
	fmt.Println(ed)
}
