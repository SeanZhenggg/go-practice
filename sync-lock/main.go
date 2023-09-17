package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var lock = sync.Mutex{}

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

func test3(ctx context.Context) {
	for {
		fmt.Printf("🍎🍎🍎🍎🍎🍎-----test 3-----\n")
		select {
		case <-ctx.Done():
			fmt.Printf(" <-ctx.Done()")
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

// func test4() {
// 	for {
// 		fmt.Printf("🍎🍎🍎🍎🍎🍎-----test 4-----\n")
// 		select {

// 		case a := <-testCh:
// 			fmt.Printf("🍎🍎🍎🍎🍎🍎 Print a %v\n", a)
// 		case b := <-testCh2:
// 			fmt.Printf("🍎🍎🍎🍎🍎🍎 Print a %v\n", b)
// 		}
// 		fmt.Printf("🍎🍎🍎🍎🍎🍎-----test 4 end-----\n")
// 	}
// }

func main() {

	// go printer("hello")
	// go printer("world")

	// go producer()
	// go consumer()

	// go producer()
	// go consumer()

	//         但是取决于谁想执行加锁操作, 所以不完美
	// for {

	// }

	// channel()

	// channel2()

	// channel3()

	// go producer2()
	// go producer3()
	// go consumer2()

	// for {

	// }

	// 定义一个数组模拟缓冲区
	// var buff = make(chan int, 5)
	// var exitCh = make(chan int)
	// go producer4(buff)
	// go consumer4(buff, exitCh)

	// <-exitCh
	// fmt.Println("程序结束了")

	// test1("")
	// test2("")
	// ctx := context.Background()
	// ctx, _ = context.WithTimeout(ctx, time.Second)

	// go test3(ctx)
	// go test4()

	// for {
	// }
}
