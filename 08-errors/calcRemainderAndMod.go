package main

import (
	"errors"
	"fmt"
	"os"
)

func calcRemainderAndMod(numerator, denomiator int) (int, int, error) {
	if denomiator == 0 {
		return 0, 0, errors.New("denimiator is 0")
	}
	return numerator / denomiator, numerator % denomiator, nil
}

func main() {
	numerator, denomiator := 20, 3
	remainder, mod, err := calcRemainderAndMod(numerator, denomiator)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(remainder, mod)
}
