package main

import (
	"fmt"
	"time"
)

func exampleForChannel1() {
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

func exampleForChannel2() {
	var stop = make(chan struct{})

	go func() {
		for {
			select {
			case <-stop:
				fmt.Println("channel get stop signal!!!")
				return
			default:
				fmt.Println("still working")
				time.Sleep(1 * time.Second)
			}
		}
	}()

	fmt.Println("worker start")
	if v, ok := <-time.After(5 * time.Second); ok {
		fmt.Printf("v: %v, ok : %v\n", v, ok)
		stop <- struct{}{}
	}
	fmt.Println("worker stop")

}

func exampleForChannel3() {
	var ch = make(chan int, 3)
	//for i := 0; i < 3; i++ {
	//	fmt.Printf("send data into channel : %v\n", i)
	//	ch <- i
	//}
	//time.Sleep(time.Second)
	close(ch)

	for {
		v, ok := <-ch
		time.Sleep(time.Second)
		fmt.Printf("v : %v, ok: %v\n", v, ok)
	}
}

func exampleForSelect1() {
	defer fmt.Println("func returned...")
	fmt.Println("func working...")

	select {}

}

func main() {
	//exampleForChannel1()
	//exampleForChannel2()
	//exampleForChannel3()
	go exampleForSelect1()
}
