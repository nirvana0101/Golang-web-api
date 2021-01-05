package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(4) //指定以多少核运算
	for {
		go fmt.Print(0)
		fmt.Print(1)
	}
}
