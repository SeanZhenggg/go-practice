package main

import (
	"fmt"
	"time"
)

func main() {
	current := time.Now()
	twoSecAfter := current.Add(time.Second * 2)
	if twoSecAfter.After(current) {
		fmt.Println("🍎🍎🍎🍎🍎🍎 two seconds after!!!")
	}

	twoSecBefore := current.Add(-time.Second * 2)

	if twoSecBefore.Before(current) {
		fmt.Println("🍎🍎🍎🍎🍎🍎 two seconds before!!!")
	}
	fmt.Println("🍎🍎🍎🍎🍎🍎 time now", time.Now().Format("2006-01-02 15:04:05"))
	if v, ok := <-time.After(time.Second * 2); ok {
		fmt.Printf("🍎🍎🍎🍎🍎🍎 val : %v, ok : %v\n", v.Format("2006-01-02 15:04:05"), ok)
	}
}
