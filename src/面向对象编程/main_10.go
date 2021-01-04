package main

import "fmt"

/**
方法(带有接收者的函数)
*/
type Person_ struct {
	name string
	sex  byte
	age  int
}
type Student_ struct {
	Person_
	id      int
	address string
}

func (stu Student_) PrintInfo() {
	fmt.Printf("stu = %+v\n", stu)
}

func (stu *Student_) setInfo(name string) {
	stu.name = name
}

func main() {
	stu := Student_{Person_{"张三", 'm', 22}, 01, "礼顿道89号"}
	stu.PrintInfo()
	stu.setInfo("王一博")
	stu.PrintInfo()
}
