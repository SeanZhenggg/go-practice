package main

import (
	"fmt"
)

func main() {
	sliceCopy()

	sliceDelete()

	sliceExample()
}

func sliceCopy() {
	numbers := make([]int, 0)

	for i := 0; i < 10; i++ {
		numbers = append(numbers, i)
	}

	copyA := make([]int, len(numbers))
	fmt.Println("copy cnt:", copy(copyA, numbers))
	fmt.Println("copied data:", copyA)

	copyB := make([]int, 3)
	fmt.Println("copy cnt:", copy(copyB, numbers[2:5]))
	fmt.Println("copied data:", copyB)

	copyC := make([]int, 3)
	fmt.Println("copy cnt:", copy(copyC, numbers))
	fmt.Println("copied data:", copyC)
}

func sliceDelete() {
	numbers := make([]int, 0)

	for i := 0; i < 10; i++ {
		numbers = append(numbers, i)
	}

	fmt.Printf("🍎🍎🍎🍎🍎🍎 numbers : memory : %p, value : %v, len : %v, cap : %v \n", numbers, numbers, len(numbers), cap(numbers))

	deleted := append(numbers[:3], numbers[4:]...)

	fmt.Printf("🍎🍎🍎🍎🍎🍎 deleted : memory : %p, value : %v, len : %v, cap : %v \n", deleted, deleted, len(deleted), cap(deleted))
}

func sliceExample() {
	a := make([]int, 20)
	fmt.Printf("🍎🍎🍎🍎🍎🍎 a : memory : %p, value : %v, len : %v, cap : %v \n", a, a, len(a), cap(a))
	a = []int{7, 8, 9, 10}
	fmt.Printf("🍎🍎🍎🍎🍎🍎 a : memory : %p, value : %v, len : %v, cap : %v \n", a, a, len(a), cap(a))
	b := a[15:16]
	fmt.Println(b)
}
