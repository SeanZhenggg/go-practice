package main

import "fmt"

func main() {
	var ch = make(chan int)
	var count int
	go func() {
		for i := 0; i < 50000; i++ {
			count += i
		}
		ch <- 1
		fmt.Printf("count in goroutine : %v\n", count)
	}()

	defer func() {
		for i := 0; i < 50000; i++ {
			count += i
		}
		ch <- 1
		fmt.Printf("count in defer : %v\n", count)
	}()

	v, ok := <-ch
	fmt.Printf("done!!! count : %v, v : %v, ok : %v\n", count, v, ok)
}
