package main

import (
	"fmt"
	"runtime"
)

/**
让出时间片，让其他协程执行完再执行
*/
func main() {
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("hello")
		}
	}()
	for i := 0; i < 2; i++ {
		runtime.Gosched()
		fmt.Println("world")
	}
}
