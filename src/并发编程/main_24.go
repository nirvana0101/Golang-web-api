package main

import (
	"fmt"
	"time"
)

var ch = make(chan int)

/**
使用channel解决多任务资源竞争，实现同步
*/
func Printer_(s string) {
	for _, data := range s {
		fmt.Printf("%c", data)
		time.Sleep(time.Second)
	}
}
func PerSon1_() {
	Printer_("hello")
	ch <- 666 //写数据
}
func PerSon2_() {
	//取数据
	<-ch
	Printer_("world")
}
func main() {
	go PerSon1_()
	go PerSon2_()
	for {
	}
}
