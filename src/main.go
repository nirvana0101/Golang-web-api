package main

import (
	"fmt"
	"hello"
)

func main() {
	var (
		a = 11
		b = 3.14
		c = "helle world !"
	)
	fmt.Printf("a = %v, b =%v, c = %v\n", a, b, c)

	hello.Hello_02(81)
}
