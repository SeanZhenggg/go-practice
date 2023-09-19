package main

import "fmt"

// 宣告 const 分組時，用 iota 賦予值，籲社會從 0 開始，並且陸續+1
const (
	a = iota
	b = iota
	c //忽略值時，會採用前一個值
	d
)

// 重新宣告一個 const 時，iota 會重置為0
const e = iota

const (
	f = iota
	g
	h
)

const (
	i = 1 << iota
	j
	k
	l
	m
	n
	o
)

const (
	p = iota << 1
	q
	r
	s
	t
	u
	v
)

func main() {
	fmt.Println(a, b, c, d, e, f, g, h)
}
