package example

import (
	"fmt"
	"time"
)

func BlockingTimerWhenStop() {
	c := make(chan bool)

	go func() {
		for i := 0; i < 3; i++ {
			time.Sleep(time.Second * 7)
			c <- false
		}

		time.Sleep(time.Second * 7)
		c <- true
	}()
	// use variable to flag the channel has received value
	// so that no need for drained channel after timer.Stop
	// var read = false
	go func() {
		timer := time.NewTimer(time.Second * 5)
		for {
			// when timer expired fires, will receive channel value below
			// then drain channel in if body will block forever cuz timer.C retrieved value before
			// need to add another flag to prevent the blocking
			// otherwise use select to drain channel
			//if !timer.Stop() && !read {
			//}
			if !timer.Stop() {
				<-timer.C
			}
			timer.Reset(time.Second * 5)
			select {
			case b := <-c:
				if b == false {
					fmt.Println(time.Now(), ":recv false. continue")
					continue
				}
				fmt.Println(time.Now(), ":recv true. return")
				return
			case <-timer.C:
				//read = true
				fmt.Println(time.Now(), ":timer expired")
				continue
			}
		}
	}()

	//to avoid that all goroutine blocks.
	var s string
	fmt.Scanln(&s)
}
