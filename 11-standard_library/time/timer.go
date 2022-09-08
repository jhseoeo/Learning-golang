package main

import (
	"fmt"
	"time"
)

func main() {
	dura := time.Second * 2
	timer := time.NewTicker(dura)
	defer timer.Stop() // shutdown ticker
	after := time.After(dura * 3)

	time.AfterFunc(dura*1, func() {
		fmt.Println("응애")
	})

loop1:
	for {
		select {
		case <-timer.C: // channel that listens ticking
			fmt.Println("야옹")
			timer.Reset(dura / 2) // reconfirguration tick interval
		case <-after:
			fmt.Println("끝")
			break loop1
		}
	}
}
