package main

import (
	"log"
	"unsafe"
)

type Builder struct {
	addr *Builder
	buf  []byte
}

func noescape(p unsafe.Pointer) unsafe.Pointer {
	x := uintptr(p)
	return unsafe.Pointer(x ^ 0)
}

func (b *Builder) copyCheck() {
	log.Printf("b.addr: %v", b.addr)
	log.Printf("b: %v", b)
	if b.addr == nil {
		b.addr = (*Builder)(noescape(unsafe.Pointer(b)))
	} else if b.addr != b {
		panic("strings: illegal use of non-zero Builder copied by value")
	}
	log.Printf("b.addr: %v", b.addr)
	log.Printf("b: %v", b)
}
func (b *Builder) Write(p []byte) (int, error) {
	b.copyCheck()
	return 0, nil
}

func main() {
	// test case
	var a Builder
	a.Write([]byte("testa"))
	var b = a
	b.Write([]byte("testb")) // will panic here
}
