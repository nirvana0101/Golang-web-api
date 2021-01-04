package main

import (
	"fmt"
)

/**
接口类型
*/
type Humaner interface {
	sayHi()
}
type Student__ struct {
	id   int
	name string
}

func (stu *Student__) sayHi() {
	fmt.Printf("Student %+v say hi\n", *stu)
}

type Teacher struct {
	id      int
	subject string
}

func (tea *Teacher) sayHi() {
	fmt.Printf("Teacher %+v say hi\n", *tea)
}

/**
多态的另外一种书写形式，通过函数传入接口类型的参数实现
*/
func PrintSayHi(hum Humaner) {
	hum.sayHi()
}

func main() {
	var hum Humaner

	hum = &Student__{1, "小明"}

	hum.sayHi()

	PrintSayHi(hum)

	hum = &Teacher{1, "生物"}

	hum.sayHi()

	PrintSayHi(hum)
}
