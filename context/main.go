package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("got the stop channel")
				return
			default:
				fmt.Println("still working")
				time.Sleep(1 * time.Second)
			}
		}
	}()

	go childProcess(ctx)

	time.Sleep(5 * time.Second)
	fmt.Println("stop the goroutine")
	cancel()
	time.Sleep(5 * time.Second)
}

func childProcess(ctx context.Context) {
	childCtx, cancel := context.WithTimeout(ctx, 8*time.Second)
	defer cancel()

	for {
		select {
		case <-childCtx.Done():
			fmt.Println("childCtx call Done()!!!")
			return
		default:
			fmt.Println("childProcess is running...")
			time.Sleep(time.Second)
		}
	}
}
