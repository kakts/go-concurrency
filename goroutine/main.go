package main

import (
	"fmt"
)

func sayHello() {
	fmt.Println("Hello")
}

func main() {
	fmt.Println("Before")
	// goroutineの呼び出し。 go キーワード関数呼び出しの前に置く
	// sayHello()の実行後に何も処理がないので、sayHelloの呼び出しをホストするゴルーチンが終了する前に終了してしまうことがほぼ確実
	// 結果として、Helloという文字列が表示されない
	go sayHello()
}