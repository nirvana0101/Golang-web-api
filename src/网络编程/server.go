package main

import (
	"fmt"
	"net"
)

/**
Socket服务端
*/
func main() {
	listener, e := net.Listen("tcp", "localhost:8060")
	if e != nil {
		fmt.Println("err = ", e)
		return
	}
	defer listener.Close()

	conn, e := listener.Accept()
	if e != nil {
		fmt.Println("err = ", e)
		return
	}
	defer conn.Close()
	bytes := make([]byte, 1024)
	n, e := conn.Read(bytes)
	if e != nil {
		fmt.Println("err = ", e)
		return
	}
	fmt.Println("receive: " + string(bytes[:n]))
}
