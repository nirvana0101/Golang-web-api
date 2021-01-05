package main

import (
	"fmt"
	"net"
)

/**
客户端
*/
func main() {
	conn, e := net.Dial("tcp", "localhost:8060")
	if e != nil {
		fmt.Println("err =", e)
		return
	}
	defer conn.Close()

	conn.Write([]byte("are you ok?"))
}
