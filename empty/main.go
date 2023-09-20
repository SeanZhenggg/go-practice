package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var a struct{}
	var b interface{}
	var c bool

	fmt.Printf("🍎🍎🍎🍎🍎🍎 a value : %v, mem: %v\n", a, unsafe.Sizeof(a))
	fmt.Printf("🍎🍎🍎🍎🍎🍎 b value : %v, mem: %v\n", b, unsafe.Sizeof(b))
	fmt.Printf("🍎🍎🍎🍎🍎🍎 c value : %v, mem: %v\n", c, unsafe.Sizeof(c))
}
