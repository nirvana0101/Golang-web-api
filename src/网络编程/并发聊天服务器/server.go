package main

import (
	"fmt"
	"net"
	"strings"
)

func handler(conn net.Conn) {
	defer conn.Close()
	addr := conn.RemoteAddr().String()
	fmt.Println(addr, "connect seccess")
	bytes := make([]byte, 1024)
	for {
		n, err := conn.Read(bytes)
		if err != nil {
			fmt.Println("err = ", err)
			return
		}
		if "exit" == string(bytes[:n]) {
			fmt.Println(addr, "退出")
			return
		} else {
			fmt.Println("receive: ", addr, string(bytes[:n]))
		}
		conn.Write([]byte(strings.ToUpper(string(bytes[:n]))))
	}
}

/**
并发聊天服务器服务端
*/
func main() {
	listener, e := net.Listen("tcp", "localhost:8086")
	if e != nil {
		fmt.Println("err = ", e)
		return
	}
	fmt.Println("server started")
	defer listener.Close()
	for {
		conn, e := listener.Accept()
		if e != nil {
			fmt.Println("err = ", e)
			return
		}
		go handler(conn)
	}
}
