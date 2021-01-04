package main

import "fmt"

/*
继承（匿名组合）
*/
type Person struct {
	name string
	sex  byte
	age  int
}
type Student struct {
	Person
	id      int
	address string
}

func main() {
	//顺序初始化
	student := Student{Person{"张三", 'm', 22}, 01, "礼顿道89号"}
	fmt.Printf("student = %+v\n", student)

	//给成员赋值
	var stu Student
	stu.name = "李四"
	stu.sex = 'f'
	stu.age = 18
	stu.id = 2
	stu.address = "礼顿道99号"
	fmt.Printf("stu = %+v\n", stu)

	//给成员赋值
	stu_ := new(Student)
	stu_.name = "王五"
	stu_.sex = 'f'
	stu_.age = 20
	stu_.id = 3
	stu_.address = "礼顿道109号"
	fmt.Printf("stu_ = %+v\n", stu_)
}
