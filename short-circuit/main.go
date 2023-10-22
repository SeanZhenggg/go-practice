package main

import (
	"fmt"
)

type c struct {
	name string
}
type b struct {
	c *c
}

type A struct {
	b *b
}

func main() {
	//var a *int
	//var b = 0
	//a, b := 10, 0

	//if a == nil && (*a)/b == 0 {
	//	fmt.Println("true")
	//} else {
	//	fmt.Println("false")
	//}

	aa := &A{}
	if aa.b != nil && aa.b.c != nil {
		fmt.Println("aa.b.c.name : %v\n", aa.b.c.name)
	} else {
		fmt.Printf("wrong : %v\n", aa.b.c)
	}
}
