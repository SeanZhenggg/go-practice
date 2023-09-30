package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.TrimLeft("aabbccdd", "abc"))
	fmt.Println(strings.TrimRight("aabbccdd/r/n", ""))
}
