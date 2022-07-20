package main

import (
	"fmt"
	"math/cmplx"
)

func main() {
	var f32 float32 = 0.5;
	var f64 float32 = 0.524;
	fmt.Println(f32, f64);

    fmt.Println("---------------------------")

	var c64 complex64 = complex(2, 3)
	c128 := complex(3, 4)
	fmt.Println(c64, c128)
	fmt.Println(real(c128))
	fmt.Println(imag(c128))
	fmt.Println(cmplx.Abs(c128))
	
}