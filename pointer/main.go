package main

import (
	"fmt"
	"time"
	"unsafe"
)

// 定義一個名為 Person 的結構（struct）
type Person struct {
	Name string // 名稱
	Age  int    // 年齡
}

func modifyValue(x *int) {
	*x = 100
}

func main() {
	// 創建一個 Person 變數，名稱為 Alice，年齡為 30 歲
	alice := Person{Name: "Alice", Age: 30}

	// 計算 Person 變數 alice 的大小
	personSize := unsafe.Sizeof(alice)

	// 創建一個指向 Person 變數的Pointer，名稱為 Bob，年齡為 24 歲
	bob := &Person{Name: "Bob", Age: 24}

	// 計算指向 Person 變數 bob 的Pointer的大小
	pointerSize := unsafe.Sizeof(bob)

	fmt.Printf("alice 變數 %v 的大小： %d 個位元組\n", alice, personSize)
	fmt.Printf("bob 指向 Person 變數 %v 的Pointer的大小： %d 個位元組\n", bob, pointerSize)
	fmt.Printf("alice 變數的記憶體位置 : %v\n", &alice) // 列印 alice 變數的記憶體位置
	fmt.Printf("bob 變數的記憶體位置 : %v\n", &bob)     // 列印 bob 變數（Pointer）的記憶體位置

	y := 10
	modifyValue(&y) // 通過Pointer修改y的值
	fmt.Println("y =", y)

	// 創建一個包含大量資料的整數陣列
	var arr [N]int
	for i := 0; i < N; i++ {
		arr[i] = i
	}

	// 測試不使用指標的情況，複製陣列
	startTime := time.Now()
	result1 := sumValues(arr)
	duration1 := time.Since(startTime)

	// 測試使用指標的情況，避免複製陣列
	startTime = time.Now()
	result2 := sumPointers(&arr)
	duration2 := time.Since(startTime)

	fmt.Printf("不使用指標的結果：%d，執行時間：%v\n", result1, duration1)
	fmt.Printf("使用指標的結果：%d，執行時間：%v\n", result2, duration2)
}

const N = 1000000 // 陣列大小

// 使用值傳遞的函數，不使用指標
func sumValues(arr [N]int) int {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	return sum
}

// 使用指標傳遞的函數，避免陣列複製
func sumPointers(arr *[N]int) int {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	return sum
}
