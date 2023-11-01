package main

import "fmt"

type Test struct {
	name string
}

func main() {
	deferUsageExample()
}

func deferUsageExample() {
	a := &Test{name: "sean"}
	defer func(a *Test) {
		fmt.Printf("in defer, a.name = %s\n", a.name)
	}(a)

	a.name = "tammy"
	fmt.Printf("in main, a.name = %s\n", a.name)
}
