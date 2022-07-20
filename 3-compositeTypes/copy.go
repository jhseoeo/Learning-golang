package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4}

	fmt.Println("---------------------------")

	y := make([]int, 4)
	num := copy(y, x) // copy(destination, source) => the number of elements copied (decided by length of slices)
	fmt.Println(y, num)
	y[2] = 1
	fmt.Println(y, x)

	fmt.Println("---------------------------")

	z := make([]int, 2, 4)
	num = copy(z, x)
	fmt.Println(z, num)

	fmt.Println("---------------------------")

	w := make([]int, 4, 4)
	num = copy(w, x[:2])
	fmt.Println(w, num)

	fmt.Println("---------------------------")

	q := []int{1, 2, 3, 4}
	num = copy(q[:3], q[1:])
	fmt.Println(q, num)

	fmt.Println("---------------------------")

	ex := []int{1, 2, 3, 4}
	ed := [...]int{5, 6, 7, 8}
	ey := make([]int, 2)
	copy(ey, ed[:])
	fmt.Println(ey)
	copy(ed[:], ex)
	fmt.Println(ed)
}
