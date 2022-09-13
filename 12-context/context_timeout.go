package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func runLoops(ctx context.Context, label string) {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	deadline, ok := ctx.Deadline()
	if ok {
		fmt.Println("loop runs until deadline:", deadline)
	} else {
		fmt.Println("there is no deadline; process stopped")
		return
	}

	for {
		select {
		case <-ticker.C:
			fmt.Println("signal from", label)
		case <-ctx.Done():
			return
		}
	}
}

func main() {
	ctx := context.Background()
	var wg sync.WaitGroup
	wg.Add(3)

	go func() {
		defer wg.Done()

		ctx, cancel := context.WithTimeout(ctx, time.Second*2)
		defer cancel()

		runLoops(ctx, "timeout")
	}()
	go func() {
		defer wg.Done()

		ctx, cancel := context.WithDeadline(ctx, time.Now().Add(time.Second*4))
		defer cancel()

		runLoops(ctx, "deadline")
	}()
	go func() {
		defer wg.Done()

		ctx, cancel := context.WithCancel(ctx)
		defer cancel()

		runLoops(ctx, "no timeout")
	}()

	wg.Wait()
	fmt.Println("good")
}
