package main

import "fmt"

/**
切片的再回顾
*/
func main() {
	var sli []byte
	fmt.Println(sli, len(sli), cap(sli)) //空数组 [] 0 0

	bytes := make([]byte, 1024)
	fmt.Println(bytes, len(bytes), cap(bytes)) //数组初始长度为1024，默认值为0
}
