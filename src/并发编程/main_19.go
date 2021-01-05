package main

import (
	"fmt"
	"time"
)

/**
主协程中断导致子协程退出
*/
func newTask_() {
	for {
		fmt.Println("this is a newTask")
		time.Sleep(time.Second)
	}
}
func main() {
	go newTask_()
	i := 0
	for {
		i++
		fmt.Println("this is a main goroutine")
		time.Sleep(time.Second)
		if i == 2 {
			break
		}
	}
}
