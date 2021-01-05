package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

/**
创建文件,并以字符串方式写入内容
*/
func WriteFile(path string) {
	//新建文件
	file, e := os.Create(path)
	if e != nil {
		fmt.Println("err = ", e)
		return
	}
	//推荐写法，函数的最后执行，无论如何都会关闭文件
	defer file.Close()

	var i int64
	for i = 0; i < 10; i++ {
		_, e := file.WriteString(strconv.FormatInt(i, 10) + "\n")
		if e != nil {
			fmt.Println("err = ", e)
		}
	}
}

/*
读取文件
*/
func ReadFile(path string) {
	file, e := os.Open(path)
	if e != nil {
		fmt.Println("err = ", e)
		return
	}
	//推荐写法，在函数的最后执行，无论如何都会关闭文件
	defer file.Close()
	bytes := make([]byte, 1024) //创建切片并指定初始容量
	n, e := file.Read(bytes)
	if e != nil && e != io.EOF { //不是io.EOF才打印
		fmt.Println("err = ", e)
		return
	}
	fmt.Println("bytes = ", string(bytes[:n]))

}

/*
按行读取
*/
func ReadFile_(path string) {
	file, e := os.Open(path)
	if e != nil {
		fmt.Println("err = ", e)
		return
	}
	//推荐写法，在函数的最后执行，无论如何都会关闭文件
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		line, _, e := reader.ReadLine()
		if e != nil {
			fmt.Println(e)
			if e == io.EOF {
				break
			}
		}
		fmt.Println(string(line))
	}
	file.Close()
}
func main() {
	WriteFile("demo.txt")              //项目根路径
	WriteFile("./src/文件读写操作/demo.txt") //使用相对路径
	ReadFile("demo.txt")
	ReadFile_("demo.txt")
}
