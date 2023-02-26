package main

import (
	"bufio"
	"fmt"
	"net"
)

func process(conn net.Conn) {
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {

		}
	}(conn) // 关闭连接
	for {
		reader := bufio.NewReader(conn)
		var buf [128]byte
		n, err := reader.Read(buf[:]) // 读取数据
		if err != nil {
			fmt.Println("read from client failed, err:", err)
			break
		}
		recvStr := string(buf[:n])
		fmt.Println("收到client端发来的数据：", recvStr)
		_, err = conn.Write([]byte(recvStr))
		if err != nil {
			return
		} // 发送数据
	}
}
func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func(lis net.Listener) {
		err := lis.Close()
		if err != nil {

		}
	}(lis)
	for {
		conn, err := lis.Accept()
		if err != nil {
			fmt.Println(err)
			break
		}
		go process(conn)

	}
}
