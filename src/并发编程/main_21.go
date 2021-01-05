package main

import (
	"fmt"
	"runtime"
)

/**
终止函数与终止协程
*/
func Print() {
	defer fmt.Println("c")
	return //只会终止当前函数
	fmt.Println("b")
}
func Print_() {
	defer fmt.Println("c") //无论如何都会执行
	runtime.Goexit()
	fmt.Println("b")
}
func main() {
	go func() {
		fmt.Println("a")
		Print_()
		fmt.Println("d") //终止协程后不在打印
	}()
	for {
	}
}
