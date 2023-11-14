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

	//loc, _ := time.LoadLocation("Asia/Taipei")
	//fmt.Printf("loc : %v\n", loc)
	//localTime, _ := time.ParseInLocation(time.RFC3339Nano, "2023-09-27T15:45:00.0000000-04:00", loc)
	//fmt.Printf("the_time : %v\n", localTime.Unix())
	//time, err := time.ParseInLocation("20060102", "20231017", loc)
	//if err == nil {
	//	fmt.Printf("time: %v", time.String())
	//}
	//startTime := time.Now()
	//time.Sleep(1 * time.Second)
	//fmt.Printf("formatTimeDuration : %s\n", formatTimeDuration(time.Since(startTime)))

	t, err := time.Parse(time.RFC3339, "2023-11-06T03:00:00.660+00:00")
	if err != nil {
		fmt.Printf("err : %v\n", err)
	}

	fmt.Printf(t.Local().Format(time.DateTime))
}

func formatTimeDuration(duration time.Duration) string {
	fmt.Printf("duration : %d\n", duration)
	sec := duration / 1000000000
	fmt.Printf("sec : %d\n", sec)
	nsec := duration % 1000000000
	fmt.Printf("nsec : %d\n", nsec)
	fmt.Printf("duration.Seconds() : %f\n", float64(sec)+float64(nsec)/1e9)

	durationSecs := time.Duration(duration.Seconds())
	fmt.Printf("durationSecs : %d\n", durationSecs)
	hours := durationSecs / 3600
	durationSecs %= 3600
	minutes := durationSecs / 60
	durationSecs %= 60

	return fmt.Sprintf("%02d:%02d:%02d", hours, minutes, durationSecs)
}
