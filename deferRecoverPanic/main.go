package main

import (
	"fmt"
	"log"
	"time"
)

//func main() {
//	defer fmt.Println("main defer 1")
//	defer fmt.Println("main defer 2")
//	defer fmt.Println("main defer 3")
//
//	fmt.Println("main")
//	return
//}

func panicExample1() {
	defer fmt.Println("in main")
	defer func() {
		defer func() {
			fmt.Println("in defer defer")
			panic("panic again and again")
		}()
		fmt.Println("in defer")
		panic("panic again")
	}()

	panic("panic once")
}

func panicExample2() {
	go callPanicFunc()

	go normalPrintFunc()
}

func callPanicFunc() {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("recover from panic : %v\n", r)
		}
		callPanicFunc()
	}()
	for {
		time.Sleep(3 * time.Second)
		panic("caused panic!!!")
	}
}

func normalPrintFunc() {
	for {
		log.Printf("go routine 2 running...")
		time.Sleep(1 * time.Second)
	}
}

type Test struct {
	name string
}

func deferUsageExample() {
	a := &Test{name: "sean"}
	defer func(a *Test) {
		fmt.Printf("in defer, a.name = %s\n", a.name)
	}(a)

	a.name = "tammy"
	fmt.Printf("in main, a.name = %s\n", a.name)
}

func main() {
	panicExample2()

	select {}
}
