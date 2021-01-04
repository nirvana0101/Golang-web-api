package main

import "fmt"

/**
数组的声明与使用
*/
func main() {
	//先声明后赋值
	var arr [10]int
	arr[0] = 0
	arr[1] = 1
	fmt.Printf("arr = %d,len = %d\n", arr, len(arr))

	//声明同时赋值
	arr_ := [10]int{0, 1, 2, 3}
	fmt.Printf("arr_ = %d,len = %d\n", arr_, len(arr_))
}
