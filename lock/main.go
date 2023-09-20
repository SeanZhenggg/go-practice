package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var (
	count int
	mu    sync.Mutex
	muA   sync.Mutex
	muB   sync.Mutex
	rw    sync.RWMutex
	wg    sync.WaitGroup
	lock  sync.Mutex
)

func printer(str string) {
	lock.Lock() //添加锁,如果不添加那么可能执行输出hello也可能执行输出world,那么就是无序的
	for _, val := range str {
		fmt.Printf("%c", val)
		time.Sleep(time.Millisecond * 300)
	}
	lock.Unlock()
}

var buff [10]int

func producer() {
	lock.Lock()
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		num := rand.Intn(100)
		fmt.Println("生产者生产了", num)
		buff[i] = num
		time.Sleep(time.Millisecond * 1000)
	}
	lock.Unlock()
}

func consumer() {
	lock.Lock()
	for i := 0; i < 10; i++ {
		num := buff[i]
		fmt.Println("消费者消费到了", num)
		time.Sleep(time.Millisecond * 1000)
	}
	lock.Unlock()
}

func channel() {
	var myCh = make(chan int)
	go func() {
		if val, ok := <-myCh; ok {
			fmt.Printf("val : %v", val)
		}

		fmt.Println("done")
	}()

	myCh <- 666 // 报错
	fmt.Printf("val : %v", <-myCh)
	time.Sleep(time.Second * 1)
	fmt.Println("done in main goroutine")
}

func channel2() {
	myCh := make(chan int, 3)
	myCh <- 1   //写入了一个数据1
	myCh <- 2   //写入了一个数据2
	myCh <- 3   //写入了一个数据3
	close(myCh) //管道必须关闭,否则报错
	for {
		if value, ok := <-myCh; ok {
			fmt.Println(ok, value) //先后输出 1 2 3
		} else {
			break
		}
	}
}

func channel3() {
	var myCh chan int
	myCh = make(chan int, 3)
	fmt.Println(<-myCh) //报错
}

var myChan = make(chan int, 10)

func producer2() {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		num := rand.Intn(100)
		myChan <- num
		fmt.Println("生产者生产了", num)
		time.Sleep(time.Millisecond * 300)
	}
}

func producer3() {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		num := rand.Intn(100)
		myChan <- num
		fmt.Println("生产者生产了", num)
		time.Sleep(time.Millisecond * 300)
	}
}

func consumer2() {
	for i := 0; i < 10; i++ {
		num := <-myChan
		fmt.Println("消费者消费到了", num)
	}
}

// 定义一个函数模拟生产者
func producer4(buff chan<- int) {
	rand.Seed(time.Now().UnixNano()) // 种随机种子
	for i := 0; i < 5; i++ {
		// 产生随机数
		num := rand.Intn(100)
		fmt.Println("生产者生产了", num)
		// 将生产好的数据放入缓冲区
		buff <- num
		//time.Sleep(time.Millisecond * 300)
	}
}

// 定义一个函数模拟消费者
func consumer4(buff <-chan int, exitCh chan<- int) {
	for i := 0; i < 5; i++ {
		num := <-buff
		fmt.Println("-------消费者消费到了", num)
	}
	exitCh <- 666
}

func test1(vars string) {
	fmt.Printf("🍎🍎🍎🍎🍎🍎-----test 1-----\n")
	switch vars {
	case "A":
		fmt.Printf("🍎🍎🍎🍎🍎🍎 Print A\n")
	case "B":
		fmt.Printf("🍎🍎🍎🍎🍎🍎 Print B\n")
	default:
		fmt.Printf("🍎🍎🍎🍎🍎🍎 default !!!\n")
	}
	fmt.Printf("🍎🍎🍎🍎🍎🍎-----test 1 end-----\n")
}

func test2(vars string) {
	fmt.Printf("🍎🍎🍎🍎🍎🍎-----test 2-----\n")
	switch vars {
	case "A":
		fmt.Printf("🍎🍎🍎🍎🍎🍎 Print A\n")
	case "B":
		fmt.Printf("🍎🍎🍎🍎🍎🍎 Print B\n")
	}

	fmt.Printf("🍎🍎🍎🍎🍎🍎-----test 2 end-----\n")
}

var testCh = make(chan int, 3)
var testCh2 = make(chan int, 3)

func test3(wg *sync.WaitGroup, ctx context.Context) {

	for {
		fmt.Printf("🍎🍎🍎🍎🍎🍎-----test 3-----\n")
		select {
		case <-ctx.Done():
			fmt.Printf(" <-ctx.Done()")
			wg.Done()
			return
		case a := <-testCh:
			fmt.Printf("🍎🍎🍎🍎🍎🍎 Print a %v\n", a)
		case b := <-testCh2:
			fmt.Printf("🍎🍎🍎🍎🍎🍎 Print a %v\n", b)
		default:
			fmt.Printf("🍎🍎🍎🍎🍎🍎 default !!!\n")
		}
		fmt.Printf("🍎🍎🍎🍎🍎🍎-----test 3 end-----\n")
	}
}

func test4(wg *sync.WaitGroup, ctx context.Context) {
	for {
		fmt.Printf("🍎🍎🍎🍎🍎🍎-----test 4-----\n")
		select {
		case <-ctx.Done():
			fmt.Printf(" <-ctx.Done()")
			wg.Done()
			return
		case a := <-testCh:
			fmt.Printf("🍎🍎🍎🍎🍎🍎 Print a %v\n", a)
		case b := <-testCh2:
			fmt.Printf("🍎🍎🍎🍎🍎🍎 Print a %v\n", b)
		}
		fmt.Printf("🍎🍎🍎🍎🍎🍎-----test 4 end-----\n")
	}
}

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

	// go printer("hello")
	// go printer("world")

	// go producer()
	// go consumer()

	// go producer()
	// go consumer()

	// channel()

	// channel2()

	// channel3()

	// go producer2()
	// go producer3()
	// go consumer2()

	// 定义一个数组模拟缓冲区
	// var buff = make(chan int, 5)
	// var exitCh = make(chan int)
	// go producer4(buff)
	// go consumer4(buff, exitCh)

	// <-exitCh
	// fmt.Println("程序结束了")

	// test1("")
	// test2("")

	//ctx := context.Background()
	//ctx, _ = context.WithTimeout(ctx, 5*time.Second)
	//wg := &sync.WaitGroup{}
	//wg.Add(1)
	////go test3(wg, ctx)
	//go test4(wg, ctx)
	//wg.Wait()

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
