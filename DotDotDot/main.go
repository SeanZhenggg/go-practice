package main

import (
	"fmt"
	"reflect"
)

func main() {
	sum()
	//fmt.Println(s1)

	sum(1, 5, 9)
	//fmt.Println(s2)
	slice := []int{2, 3, 5}

	sum1 := sumUnpacking(slice...) // 把slice 解開、剝皮後傳入，同下
	fmt.Println(sum1)
}

func sum(nums ...int) int {
	fmt.Printf("nums type : %v, len : %v, cap: %v \n", reflect.TypeOf(nums).Elem().Kind(), len(nums), cap(nums))
	var total int
	for _, num := range nums {
		total += num
	}
	return total
}

func sumUnpacking(nums ...int) int { // 傳入int但不曉得參數長度為何
	var total int
	for _, num := range nums {
		total += num
	}
	return total
}
