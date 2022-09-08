package main

import "fmt"

func main() {
	words := []string{"a", "cow", "smile", "gopher", "octopus", "영어싫어", ""}

	for _, word := range words {
		switch size := len(word); size { // as like if-statements, there is no parenthesis
		case 1, 2, 3, 4: // using comma(,), we can make multiple matches
			fmt.Println(word, "is a short word!")
		case 5:
			fmt.Println(word, "is exactly the right length:", size)
			break                         // this makes it exits switch earlier
			fmt.Println("not be printed") // because of break above, this line is unreachable
		case 6, 7, 8, 9: // empty case (nothing happened)
			// fallthrough // if keyword `fallthrough` is here, as like other languages, run next case's block
		default:
			fmt.Println(word, "is a wrong word")
		}
	}

	fmt.Println("---------------------------")

	// blank switch, break a loop in switch/case statements
loop: // to break in switch/case statement, attach a label to for-statements
	for i := 0; i < 10; i++ {
		switch { // this is a blank switch. no variable
		case i%2 == 0: // here can be boolean expression
			fmt.Println(i, "is even number")
		case i%3 == 0:
			fmt.Println(i, "is divisible by 3 but not 2")
		case i%7 == 0:
			fmt.Println("exit the loop")
			break loop // if this statement is just `break`, it cannot exit the loop. just proceed until the loop ends
		default:
			fmt.Println(i, "is boring")
		}
	}

	fmt.Println("---------------------------")

	for i := 0; i < 10; i++ {
		switch { // blank switch
		case i < 5: // here can be boolean expression
			fmt.Println(i, "is larger than 5")
		case i == 0:
			fmt.Println(i, "is smaller than 5")
		default:
			fmt.Println(i, "is 5")
		}
	}
}
