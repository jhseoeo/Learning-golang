package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func child_loop(ctx context.Context) {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			fmt.Println("signal from child context")
		case <-ctx.Done():
			fmt.Println("child context cancelled")
			return
		}
	}
}

func parent_loop(ctx context.Context) {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	go func() {
		ctx, cancel := context.WithTimeout(ctx, time.Second*3)
		defer cancel()
		child_loop(ctx)
	}()

	for {
		select {
		case <-ticker.C:
			fmt.Println("signal from parent context")
		case <-ctx.Done():
			fmt.Println("parent context cancelled")
			return
		}
	}
}

func main() {
	ctx := context.Background()
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		ctx, cancel := context.WithTimeout(ctx, time.Second*2)
		defer cancel()
		parent_loop(ctx)
	}()

	wg.Wait()
}
