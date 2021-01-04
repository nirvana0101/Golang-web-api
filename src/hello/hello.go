package hello

import "fmt"

/**
定义一个私有函数
*/
func hello() {
	fmt.Println("hello")
}

/**
定义一个公有函数
*/
func Hello() {
}

/**
定义一个有返回值的函数
*/
func hello_() (a int) {
	return 1
}

/**
定义一个有多个返回值的函数
*/
func hello_01() (a, b int) {
	return 1, 2
}

/**
定义一个有参数有返回值的函数
*/
func Hello_02(a int) (b int) {
	fmt.Println(a)
	return a
}
