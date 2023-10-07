package main

import "fmt"

//func main() {
//	defer fmt.Println("main defer 1")
//	defer fmt.Println("main defer 2")
//	defer fmt.Println("main defer 3")
//
//	fmt.Println("main")
//	return
//}

func main() {
	defer fmt.Println("in main")
	defer func() {
		defer func() {
			fmt.Println("in defer defer")
			panic("panic again and again")
		}()
		fmt.Println("in defer")
		panic("panic again")
	}()

	panic("panic once")
}
