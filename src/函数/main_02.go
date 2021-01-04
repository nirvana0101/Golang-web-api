package main

import "fmt"

/**
函数类型
*/
type Calculate func(int, int) int

func Add(a, b int) int {
	return a + b
}
func Minus(a, b int) int {
	return a - b
}
func Calculata(a, b int, calculate Calculate) int {
	return calculate(a, b)
}
func main() {
	fmt.Println("Calculata(1, 1, Add)", Calculata(1, 1, Add))
	fmt.Println("Calculata(1, 1, Minus)", Calculata(1, 1, Minus))
}
