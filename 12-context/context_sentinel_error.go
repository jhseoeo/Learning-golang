package main

import (
	"context"
	"fmt"
	"time"
)

func longRunningThing(ctx context.Context, data int) (int, error) {
	ticker := time.NewTicker(time.Second * 1)
	defer ticker.Stop()

	ctx, cancel := context.WithTimeout(ctx, time.Duration(data)*time.Second)
	defer cancel()

	var result, i int
	for {
		select {
		case <-ticker.C:
			i++
			result += i
		case <-ctx.Done():
			return result, ctx.Err()
		}
	}

}

func longRunningThingManager(ctx context.Context, data int) (int, error) {
	type wrapper struct {
		result int
		err    error
	}

	ch := make(chan wrapper, 1)
	go func() {
		result, err := longRunningThing(ctx, data)
		ch <- wrapper{result, err}
	}()

	select {
	case data := <-ch:
		return data.result, data.err
	case <-ctx.Done():
		return 0, ctx.Err()
	}
}

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	time.AfterFunc(time.Second*2, cancel)
	num, err := longRunningThingManager(ctx, 3)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(num)
}
