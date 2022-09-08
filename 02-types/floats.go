package main

import (
	"fmt"
	"math/cmplx"
)

func main() {
	var f32 float32 = 0.5   // f32 : 32-bits floating point number
	var f64 float32 = 0.524 // f64 : 64-bits floating point number
	fmt.Println(f32, f64)

	fmt.Println("---------------------------")

	var c64 complex64 = complex(2, 3) // c64 : 복소수(float32 + float32)
	c128 := complex(3, 4)             // c128 : 복소수(float64 + float64)
	fmt.Println(c64, c128)
	fmt.Println(real(c128))      // 실수부
	fmt.Println(imag(c128))      // 허수부
	fmt.Println(cmplx.Abs(c128)) // 절댓값..?

}
