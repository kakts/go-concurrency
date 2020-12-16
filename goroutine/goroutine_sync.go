package main

import (
	"fmt"
	"sync"
)

func sayHello2() {
	// 合流完了
	defer wg.Done()
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
	// goroutine実行
	// 呼び出しのあとに、wg.Wait()で合流ポイントを作ることにより、意図した通りにメインルーチン呼び出しの前にsayHello()が完了する
	go sayHello()
	wg.Wait() // 合流ポイント
}