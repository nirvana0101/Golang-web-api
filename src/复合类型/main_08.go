package main

import "fmt"

/**
结构体的定义与基本操作
*/
type Studenet struct {
	id      int
	name    string
	sex     string
	age     int
	address string
}

/**
传值
*/
func change(student Studenet) {
	student.id = 02
	fmt.Println("change：student = ", student)
}

/**
传地址
*/
func change_(student *Studenet) {
	student.id = 02
	fmt.Println("change_：student = ", student)
}
func main() {
	var student Studenet
	student.id = 1
	student.name = "张三"
	student.sex = "男"
	student.age = 22
	student.address = "礼顿道89号"
	fmt.Println("student = ", student)

	_student := new(Studenet)
	_student.id = 1
	_student.name = "张三"
	_student.sex = "男"
	_student.age = 22
	_student.address = "礼顿道89号"
	fmt.Println("_student = ", _student)

	//测试传值
	change(student)
	fmt.Println("main: student = ", student)

	//测试传地址
	change_(_student)
	fmt.Println("main: student = ", _student)
}
