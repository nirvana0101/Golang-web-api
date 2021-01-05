package main

import (
	"fmt"
	"net"
)

func main() {
	conn, e := net.Dial("tcp", "localhost:8086")
	if e != nil {
		fmt.Println("err = ", e)
		return
	}
	defer conn.Close()

	bytes := make([]byte, 1024)

	for {
		fmt.Println("请输入：……")
		fmt.Scan(&bytes)
		fmt.Println("发送：", string(bytes))
		conn.Write(bytes)

		n, e := conn.Read(bytes)
		if e != nil {
			fmt.Println("err = ", e)
			return
		}
		fmt.Println("接收到来自", conn.RemoteAddr().String(), string(bytes[:n]))
	}
}
