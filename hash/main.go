package main

import (
	"crypto/md5"
	"crypto/sha256"
	"fmt"
)

func main() {
	data := []byte("hello world")
	hash := sha256.Sum256(data)
	fmt.Printf("hash1 %x\n", string(hash[:]))
	hash2 := md5.Sum(hash[:])
	fmt.Printf("hash2 %x\n", string(hash2[:]))
}
