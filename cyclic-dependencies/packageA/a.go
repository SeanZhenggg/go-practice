package packagea

import (
	"fmt"
	B "practice/cyclic-dependencies/packageB"
)

var (
	varA = 456
)

func Test() {
	fmt.Println(B.VarB)
}
