package main

import "fmt"

func main() {
	fmt.Println(1_2_3_4, 123_456_789_123123_456) // numbers can distinguished by underbar(_)
	fmt.Println(0x1234, 0o1234, 0b1101)          // hexadecimal, octal, binary representation

	fmt.Println("---------------------------")

	var i8 int8 = -128                  // 8-bit signed integer
	var i64 int64 = 9223372036854775807 // 64-bit signed integer
	var u32 uint32 = 4294967295         // 32-bit unsigned integer
	fmt.Println(i8, i64, u32)

	fmt.Println("---------------------------")

	var b byte = 123 // byte: 8-bit unsigned integer
	var uint8_ uint8 = 234
	fmt.Println(b == uint8_) // doesn't occur error

	fmt.Println("---------------------------")

	var i int = 9223372036854775807 // signed int. its size(32 or 64bits) is decided at compile time, determined by its hardware
	var ui uint = 0                 // unsigned int
	// fmt.Println(i == i64) occurs error
	fmt.Println(i, ui)
}
