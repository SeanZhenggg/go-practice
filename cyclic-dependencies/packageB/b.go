package packageb

import (
	"fmt"
	"practice/cyclic-dependencies/packagea"
)

var (
	VarB = 123
)

func Test() {
	fmt.Println(packagea.VarA)
}
