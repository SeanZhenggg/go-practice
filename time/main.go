package main

import (
	"github.com/panjf2000/ants/v2"
	"log"
	"sync"
	"time"
)

func main() {
	//example.FailedResetTimer()

	//example.TickerExample()

	now := time.Now()
	for i := 0; i < 100; i++ {
		doIntensiveTask()
	}
	used := time.Since(now)
	log.Printf("time used : %v", used)

	now2 := time.Now()
	pool, err := ants.NewPool(50)
	if err != nil {
		return
	}
	wg := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		pool.Submit(func() {
			doIntensiveTaskByGoroutine(&wg)
		})
	}
	wg.Wait()
	used2 := time.Since(now2)
	log.Printf("time used : %v", used2)
}

func doIntensiveTask() {
	var sum int
	for i := 0; i < 10000000; i++ {
		sum += i
	}
}

func doIntensiveTaskByGoroutine(wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()
	var sum int
	for i := 0; i < 10000000; i++ {
		sum += i
	}
}
