package main

import (
	"fmt"
	"runtime"
)

func main() {
	test1()
}

func test1() {
	test2()
}

func test2() {
	test3()
}

func test3() {
	var i = 1
	for {
		pc, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		}
		fmt.Printf("%v:%s:%d\n", pc, file, line)
		i++
	}
}
