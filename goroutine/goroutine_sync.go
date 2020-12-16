package main

import (
	"fmt"
	"sync"
)

func sayHello2() {
	fmt.Println("Hello")
}

func main() {
	fmt.Println("Before")

	// goroutine同期用の合流グループ作成
	var wg sync.WaitGroup
	sayHello := func() {
		// 合流完了
		defer wg.Done()
		fmt.Println("Hello")
	}
	wg.Add(1)
	go sayHello()
	wg.Wait() // 合流ポイント
}