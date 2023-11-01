package main

import "fmt"

func main() {
	var m1 = make(map[int]int)
	fmt.Printf("m1 : %v, %p\n", m1, &m1)

	passMapAsPointer(&m1)
	fmt.Printf("m1 : %v, %p\n\n", m1, &m1)

	var m2 = make(map[int]int)
	fmt.Printf("m2 : %v, %p\n", m2, &m2)
	passMapAsValue(m2)
	fmt.Printf("m2 : %v, %p\n\n", m2, &m2)

	var m3 map[int]int
	fmt.Printf("m3 : %v, %p\n", m3, &m3)
	passMapAsPointerAndRedefine(&m3)
	fmt.Printf("m3 : %v, %p\n\n", m3, &m3)

	var m4 map[int]int
	fmt.Printf("m4 : %v, %p\n", m4, &m4)
	passMapAsValueAndRedefine(m4)
	fmt.Printf("m4 : %v, %p", m4, &m4)

	nilMapCheck()
}

func passMapAsPointer(m *map[int]int) {
	(*m)[1] = 1
	fmt.Printf("m : %v, %p\n", m, m)
}

func passMapAsValue(m map[int]int) {
	m[1] = 1
	fmt.Printf("m : %v, %p\n", m, &m)
}

func passMapAsPointerAndRedefine(m *map[int]int) {
	*m = make(map[int]int)
	(*m)[1] = 1
	fmt.Printf("m : %v, %p\n", m, m)
}

func passMapAsValueAndRedefine(m map[int]int) {
	m = make(map[int]int)
	m[1] = 1
	fmt.Printf("m : %v, %p\n", m, &m)
}

func nilMapCheck() {
	var m map[string]int

	if m == nil {
		fmt.Println("is nil map...")
	} else {
		fmt.Println("is not nil map...")
	}

	if len(m) == 0 {
		fmt.Println("is nil map...")
	} else {
		fmt.Println("is not nil map...")
	}
}
