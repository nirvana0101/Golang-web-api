package main

import "fmt"

/**
匿名函数
*/
func main() {
	func() {
		fmt.Println("匿名函数的使用")
	}()
}
