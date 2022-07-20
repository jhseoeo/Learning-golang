package main

import "fmt"

func main() {
	var slice1 []int     // nil
	var slice2 = []int{} // zero-length slice, not nil
	fmt.Println(slice1 == nil, slice2 == nil)

	fmt.Println("---------------------------")

	x := []int{1, 2, 3, 4}
	y := x[:2]
	z := x[1:]
	d := x[1:3]
	e := x[:]
	fmt.Println(x, y, z, d, e)

	fmt.Println("---------------------------")

	e[2] = 1 // Slicing overwraps storage
	fmt.Println(x, y, z, d, e)

	fmt.Println(cap(x), cap(y), cap(z), cap(d), cap(e))
	y = append(y, 30)
	// z = append(z, 30) // try this also

	fmt.Println(x, y, z, d, e) // it's too confusing :(
	fmt.Println(cap(x), cap(y), cap(z), cap(d), cap(e))

	fmt.Println("---------------------------")

	xx := make([]int, 0, 5)
	xx = append(xx, 1, 2, 3, 4)
	yy := xx[:2]
	zz := xx[2:]
	fmt.Println(cap(xx), cap(yy), cap(zz))
	yy = append(yy, 30, 40, 50)
	xx = append(xx, 60)
	zz = append(zz, 70)
	fmt.Println("xx:", xx)
	fmt.Println("yy:", yy)
	fmt.Println("zz:", zz)

	fmt.Println("---------------------------")

	xxx := []int{1, 2, 3, 4, 5}
	yyy := xxx[:2:2] // limits capacity of slice, they not share additional capacity
	zzz := xxx[2:4:4]
	fmt.Println(cap(xxx), cap(yyy), cap(zzz))
	yyy = append(yyy, 30, 40, 50) // this appending never interacts with other slices
	xxx = append(xxx, 60)
	zzz = append(zzz, 70)
	fmt.Println(xxx, yyy, zzz)

	fmt.Println("---------------------------")

	ax := [...]int{1, 2, 3, 4}
	ay := ax[:2]
	az := ax[2:]
	ax[0] = 10
	fmt.Println(ax, ay, az)

}
