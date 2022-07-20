package main

import "fmt"

func main() {
    fmt.Println(1_2_3_4, 123_456_789_123123_456)
    fmt.Println(0x1234, 0o1234, 0b1101)

    fmt.Println("---------------------------")

    var i8 int8 = -128
    var i64 int64 = 9223372036854775807
    var u32 uint32 = 4294967295
    fmt.Println(i8, i64, u32)

    fmt.Println("---------------------------")
    
    var b byte = 123
    var uint8 uint8 = 234
    fmt.Println(b == uint8) // doesn't occur error

    fmt.Println("---------------------------")

    var i int = 9223372036854775807
    var asd uint = 0
    // fmt.Println(i == i64) occurs error
}