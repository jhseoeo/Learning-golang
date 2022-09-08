package main

import "fmt"

func main() {
	samples := []string{"hello", "안녕hello"}

outer: // label of outer for-statements
	for _, sample := range samples {
		for i, r := range sample {
			fmt.Println(i, r, string(r))
			if r == 'l' {
				continue outer // this leads to continue nested loop
			}
		}
		fmt.Println() // this line is not reached because all the strings contain 'l', and the process continued outer loop
	}
}
