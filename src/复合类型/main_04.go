package main

import "fmt"

/**
值传递与指针传递
*/
func main_() {
	a := 10
	b := 20
	swap(a, b) //新建内存空间，存值,即10 ，20
	fmt.Printf("main: a = %d,b= %d\n", a, b)
}

/**
传值：交换a,b的值
*/
func swap(a, b int) {
	a, b = b, a
	fmt.Printf("swap: a = %d,b= %d\n", a, b)
}

func main() {
	a := 10
	b := 20
	fmt.Println("main:指针&a指向的值为,", *&a)
	fmt.Printf("main: &a = %v,&b= %v\n", &a, &b)
	swap_(&a, &b) //新建内存空间，存地址，即：0xc00000a0a8，0xc00000a0c0
	fmt.Printf("main: a = %d,b= %d\n", a, b)
}

/**
传地址：交换a,b的值
*/
func swap_(pointer_a, pointer_b *int) {

	fmt.Println("swap_:指针pointer_a指向的值为,", *pointer_a) //取地址的值

	fmt.Printf("swap_: pointer_a = %v,pointer_b= %v\n", pointer_a, pointer_b) //打印指针

	*pointer_a, *pointer_b = *pointer_b, *pointer_a //把指针b指向的值，赋予指针a指向的内存,
	// 把指针a指向的值，赋予指针b指向的内存

	fmt.Printf("swap_: a = %d,b= %d\n", *pointer_a, *pointer_b)
}
