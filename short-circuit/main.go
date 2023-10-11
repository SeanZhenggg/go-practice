package main

import (
	"fmt"
)

func main() {
	var a *int
	var b = 0
	//a, b := 10, 0

	if a == nil && (*a)/b == 0 {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
