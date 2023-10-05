package main

import "fmt"

var (
	a = 100
	b = 30
)

type Person struct {
	Name  string
	Score *int
}

type School struct {
	Person Person
}

func main() {
	fmt.Printf("[main] a, b val : %v, %v, a, b mem : %p, %p\n", a, b, &a, &b)
	createVars(b)

	createVars2()
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

func createVars2() {
	var a School

	var b = new(map[int]int)
	var c map[int]int

	fmt.Printf("b %+v\n", b == nil)
	fmt.Printf("c %+v\n", c == nil)
	c[3] = 1
	fmt.Printf("a %+v\n", a)
	a.Person = Person{Name: "a"}
	fmt.Printf("a.Person %+v\n", a.Person)
}
