package main

import "fmt"

func main() {
	var m1 = make(map[int]int)
	fmt.Printf("m1 : %v, %p\n", m1, &m1)

	test1(&m1)
	fmt.Printf("m1 : %v, %p\n\n", m1, &m1)

	var m2 = make(map[int]int)
	fmt.Printf("m2 : %v, %p\n", m2, &m2)
	test2(m2)
	fmt.Printf("m2 : %v, %p\n\n", m2, &m2)

	var m3 = make(map[int]int)
	fmt.Printf("m3 : %v, %p\n", m3, &m3)
	test3(&m3)
	fmt.Printf("m3 : %v, %p\n\n", m3, &m3)

	var m4 = make(map[int]int)
	fmt.Printf("m4 : %v, %p\n", m4, &m4)
	test4(m4)
	fmt.Printf("m4 : %v, %p", m4, &m4)
}

func test1(m *map[int]int) {
	fmt.Printf("m : %v, %p\n", m, m)
	(*m)[1] = 1
}

func test2(m map[int]int) {
	m[1] = 1
	fmt.Printf("m : %v, %p\n", m, &m)
}

func test3(m *map[int]int) {
	*m = make(map[int]int)
	(*m)[1] = 1
	fmt.Printf("m : %v, %p\n", m, m)
}

func test4(m map[int]int) {
	m = make(map[int]int)
	m[1] = 1
	fmt.Printf("m : %v, %p\n", m, &m)
}
