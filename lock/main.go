package main

import (
	"fmt"
	"sync"
)

var (
	count int
	mu    sync.Mutex
	muA   sync.Mutex
	muB   sync.Mutex
	rw    sync.RWMutex
	wg    sync.WaitGroup
)

func main() {
	// wg.Add(6)
	// for i := 0; i < 3; i++ {
	// 	go reader(i)
	// }

	// for i := 0; i < 3; i++ {
	// 	go writer(i)
	// }
	// time.Sleep(1 * time.Second)
	// fmt.Println("🍎🍎🍎🍎🍎🍎 final count : ", count)
	// wg.Wait()

	// mu.Lock()
	// defer mu.Unlock()
	// copyTest(&mu)
	wg.Add(2)
	go deadlock1()
	go deadlock2()
	wg.Wait()
}

func reader(id int) {
	fmt.Printf("🍎🍎🍎🍎🍎🍎 reader %v start\n", id)
	rw.RLock()
	defer func() {
		rw.RUnlock()
		wg.Done()
	}()

	fmt.Printf("🍎🍎🍎🍎🍎🍎 reader %v read count : %v\n", id, count)
}

func writer(id int) {
	fmt.Printf("🍎🍎🍎🍎🍎🍎 writer %v start\n", id)
	rw.Lock()
	defer func() {
		rw.Unlock()
		wg.Done()
	}()

	count++
	fmt.Printf("🍎🍎🍎🍎🍎🍎 writer %v write count : %v\n", id, count)
}

// 这里复制了一个锁，造成了死锁
func copyTest(mu *sync.Mutex) {
	mu.Lock()
	defer mu.Unlock()
	fmt.Println("ok")
}

func deadlock1() {
	muA.Lock()
	defer muA.Unlock()
	fmt.Println("🍎🍎🍎🍎🍎🍎 deadlock1")
	muB.Lock()
	defer muB.Unlock()

	wg.Done()
}

func deadlock2() {
	muB.Lock()
	defer muB.Unlock()
	fmt.Println("🍎🍎🍎🍎🍎🍎 deadlock2")
	muA.Lock()
	defer muA.Unlock()
	wg.Done()
}
