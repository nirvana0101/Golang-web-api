package main

import (
	"fmt"
)

type User struct {
	username string
	password string
}

/**
空接口与类型断言
*/
func main() {
	//空接口是万能类型
	var i interface{}
	i = 1
	var j interface{}
	j = "mail"
	fmt.Printf("i = %v,j = %v\n", i, j)
	var user interface{}
	user = User{"mike", "1w3456"}
	fmt.Printf("user = %v\n", user)

	//类型断言
	var sli []interface{}
	sli = append(sli, i)
	sli = append(sli, j)
	sli = append(sli, user)
	for _, value := range sli {

		switch t := value.(type) {
		case int:
			fmt.Printf("%T ,%v\n", t, value)
		case string:
			fmt.Printf("%T ,%v\n", t, value)
		case User:
			fmt.Printf("%T ,%v\n", t, value)
		default:
			fmt.Printf("%T ,%v\n", t, value)
		}
	}
}
