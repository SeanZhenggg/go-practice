package example

import (
	"fmt"
	"time"
)

func FailedResetTimer() {
	signal := make(chan int)
	go func() {
		for {
			signal <- 1
		}
	}()

	t := time.NewTimer(time.Second)
	for {
		select {
		case <-signal:
			//time.Sleep(2 * time.Second)
			t.Reset(time.Second)
			fmt.Print("1")
		case <-t.C:
			fmt.Print("!")
		default:
			fmt.Print(".")
		}
	}
}
