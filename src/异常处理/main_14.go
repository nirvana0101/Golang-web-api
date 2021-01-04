package main

import (
	"errors"
	"fmt"
)

/**
error类型错误的处理(go的error类型相当于java的RuntimeException,不会导致程序中断)
*/
func Div(a, b float64) (result float64, err error) {
	if b == 0 {
		err = errors.New("runtime error: divide by zero")
		return
	}
	result = a / b
	return
}
func main() {
	r, err := Div(1, 0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(r)
	}
}
