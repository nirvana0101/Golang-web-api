package main

import "fmt"

/**
切片的声明与使用
*/
func main() {
	//先声明后赋值
	var slice []int
	slice = append(slice, 1)
	slice = append(slice, 2)
	fmt.Printf("slice = %d,len = %d,cap = %d\n", slice, len(slice), cap(slice))

	//声明同时赋值
	slice_ := []int{1, 2, 3, 4}
	fmt.Printf("slice_ = %d,len = %d,cap = %d\n", slice_, len(slice_), cap(slice_))

	//切片的截取
	arr := [10]int{0, 1, 2, 3, 4, 5}
	slice_01 := arr[0:3]
	fmt.Printf("slice_01 = %d,len = %d,cap = %d\n", slice_01, len(slice_01), cap(slice_01))
}
