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
	go sayHello()
}