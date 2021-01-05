package main

import (
	"fmt"
	"time"
)

/*
实现同步与数据交换
*/
func main() {
	var ch = make(chan string)
	defer fmt.Println("主协程调用结束")

	go func() {
		defer fmt.Println("子协程调用结束")
		for i := 0; i < 2; i++ {
			fmt.Println("子协程调用", i)
			time.Sleep(time.Second)
		}
		ch <- "子协程调用结束"
	}()
	s := <-ch
	fmt.Println("收到", s)
}
