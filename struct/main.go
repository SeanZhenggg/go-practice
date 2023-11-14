package main

import "fmt"

type embed struct {
	prop1 string
	prop2 int
	prop3 bool
}

type a struct {
	Name    string
	Value   int
	IsAt    bool
	private rune
	Embed   embed
	Arr     []int
}

type b struct {
	Name    string
	Value   int
	IsAt    bool
	private rune
	Embed   embed
	Arr     []int
}

func main() {
	aa := &a{
		Name:    "a",
		Value:   1,
		IsAt:    true,
		private: 'c',
		Embed: embed{
			prop1: "c",
			prop2: 2,
			prop3: false,
		},
		Arr: []int{1, 2, 3},
	}
	var copy1 = (*b)(aa)
	fmt.Printf("aa %v, mem: %p\n", aa, aa)
	fmt.Printf("copy %v, mem: %p\n", copy1, copy1)
}
