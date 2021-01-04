package main

import "fmt"

func TestFunc(a int) {
	if a == 1 {
		return
	}
	a--
	TestFunc(a)
	fmt.Printf("a = %v\n", a)
}

/**
递归计算过程：10 + 9 + 8 + 7
*/
func sum(a int) (c int) {
	if a == 10 {
		return a
	}
	a += sum(a + 1) //9+10,8+19 ……
	return a
}

/**
func main() {
	TestFunc(10)
	fmt.Println("main")
}
*/

func main() {
	fmt.Println("sum(1)=", sum(1))
}
