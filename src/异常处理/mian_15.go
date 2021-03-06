package main

import "fmt"

/**
panic异常，不可恢复错误。如：数组越界，空指针异常
*/
func TestA() {
	fmt.Println("func TestA()")
}
func TestB(i int) {
	var arr [10]int
	arr[i] = 666 //panic: runtime error: index out of range [20] with length 10
}
func TestC() {
	fmt.Println("func TestC()")
}
func TestD(i int) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()
	var arr [10]int
	arr[i] = 666
}
func TestE(i []int) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()
	i = append(i, 666)
}
func main() {
	TestA()
	TestD(20)
	var i []int
	TestE(i)
	TestC()
}
