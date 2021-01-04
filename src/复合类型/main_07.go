package main

import "fmt"

/**
map的基本操作
*/
func main() {
	//先声明后赋值
	user := make(map[string]string)
	user["username"] = "admin"
	user["password"] = "123456"
	fmt.Printf("user = %v,username = %v,password = %v\n", user, user["username"], user["password"])

	user_ := map[string]string{
		"username": "admin",
		"password": "123456",
	}
	fmt.Printf("user = %v,username = %v,password = %v\n", user_, user_["username"], user_["password"])

	//遍历
	for k, v := range user_ {
		fmt.Printf("k = %v,v = %v\n", k, v)
	}

	//查看键值是否存在
	_, flag := user_["phone"]
	fmt.Printf("键值存在吗? %v\n", flag)

	//删除元素
	delete(user_, "username")
	fmt.Printf("user = %v,username = %v,password = %v\n", user_, user_["username"], user_["password"])
}
