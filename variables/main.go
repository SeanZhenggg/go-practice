package main

import "fmt"

var (
	a = 100
	b = 30
)

func main() {
	fmt.Printf("[main] a, b val : %v, %v, a, b mem : %p, %p\n", a, b, &a, &b)
	createVars(b)
}

func createVars(b int) (c int) {
	fmt.Printf("[createVars] a, b val : %v, %v, a, b mem : %p, %p\n", a, b, &a, &b)

	b = 49
	a = 50
	fmt.Printf("[createVars] a, b val : %v, %v, a, b mem : %p, %p\n", a, b, &a, &b)

	//var a = 49
	//fmt.Printf("a val : %v, a mem : %p\n", a, &a)

	a, c := 78, 87
	fmt.Printf("[createVars] a, c val : %v, %v, a, c mem : %p, %p\n", a, c, &a, &c)

	return c
}
