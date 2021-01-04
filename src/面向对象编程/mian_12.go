package main

import "fmt"

/**
接口的继承与转换
*/
type Humaner_ interface {
	sayHi()
}
type Personer_ interface {
	Humaner_
	sing()
}
type _Student struct {
	id   int
	name string
}

func (stu *_Student) sayHi() {
	fmt.Printf("Student %+v say hi\n", *stu)
}
func (stu *_Student) sing() {
	fmt.Printf("Student %+v sing\n", *stu)
}
func main() {
	//接口的继承
	var per Personer_
	per = &_Student{9527, "mike"}
	per.sayHi()
	per.sing()

	//接口的转换（子接口可转父接口）
	var hum Humaner_
	hum = per
	hum.sayHi() //转过来也只接了sayHi那部分

	//接口的转换（父接口转子接口失败,因为父没有子对应sing那块内存空间）
	hum = &_Student{9528, "nike"}
	per = hum //cannot use hum (type Humaner_) as type Personer_ in assignment
	hum.sayHi()
}
