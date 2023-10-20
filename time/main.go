package main

import (
	"fmt"
	"time"
)

func main() {
	//current := time.Now()
	//twoSecAfter := current.Add(time.Second * 2)
	//if twoSecAfter.After(current) {
	//	fmt.Println("🍎🍎🍎🍎🍎🍎 two seconds after!!!")
	//}
	//
	//twoSecBefore := current.Add(-time.Second * 2)
	//
	//if twoSecBefore.Before(current) {
	//	fmt.Println("🍎🍎🍎🍎🍎🍎 two seconds before!!!")
	//}
	//fmt.Println("🍎🍎🍎🍎🍎🍎 time now", time.Now().Format("2006-01-02 15:04:05"))
	//if v, ok := <-time.After(time.Second * 2); ok {
	//	fmt.Printf("🍎🍎🍎🍎🍎🍎 val : %v, ok : %v\n", v.Format("2006-01-02 15:04:05"), ok)
	//}

	loc, _ := time.LoadLocation("Asia/Taipei")
	//fmt.Printf("loc : %v\n", loc)
	//localTime, _ := time.ParseInLocation(time.RFC3339Nano, "2023-09-27T15:45:00.0000000-04:00", loc)
	//fmt.Printf("the_time : %v\n", localTime.Unix())
	time, err := time.ParseInLocation("20060102", "20231017", loc)
	if err == nil {
		fmt.Printf("time: %v", time.String())
	}
}
