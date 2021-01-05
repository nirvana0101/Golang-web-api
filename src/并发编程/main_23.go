package main

import (
	"fmt"
	"time"
)

/**
多任务资源竞争
*/
func Printer(s string) {
	for _, data := range s {
		fmt.Printf("%c", data)
		time.Sleep(time.Second)
	}
}
func PerSon1() {
	Printer("hello")
}
func PerSon2() {
	Printer("world")
}
func main() {
	go PerSon1()
	go PerSon2()
	for {
	}
}
