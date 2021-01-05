package main

import (
	"fmt"
	"time"
)

/**
多任务同时执行
*/
func newTask() {
	for {
		fmt.Println("this is a newTask")
		time.Sleep(time.Second)
	}
}
func main() {
	go newTask()
	for {
		fmt.Println("this is a main goroutine")
		time.Sleep(time.Second)
	}
}
