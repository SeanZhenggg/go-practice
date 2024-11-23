package example

import (
	"log"
	"sync"
	"time"
)

func TickerExample() {
	ticker := time.NewTicker(1 * time.Second)
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer func() {
			wg.Done()
		}()
		for {
			select {
			case v, ok := <-ticker.C:
				log.Printf("ticker is running... v: %s, running: %t", v, ok)
			default:
				time.Sleep(990 * time.Millisecond)
				log.Printf("not running...")
			}

		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(5100 * time.Millisecond)
		ticker.Stop()
	}()

	wg.Wait()

	log.Printf("close program...")
}
