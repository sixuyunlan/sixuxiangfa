package main

import (
	"fmt"
	"net"
)

func process(conn net.Conn) { // 客户端连接成功后，调用此函数
	defer conn.Close() // 关闭conn
	for {              //循环读取客户端发送的数据
		buf := make([]byte, 1024) // 创建一个1024字节的缓冲区
		n, err := conn.Read(buf)  //读取客户端发送的数据
		if err != nil {
			fmt.Println("读取客户端发送的数据失败，err:", err)
			return
		}
		fmt.Println("客户端发送了：", string(buf[0:n])) // 把读取到的数据打印出来
	}
}

func main() {
	fmt.Println("服务器开始监听。。。。。。。。")
	listen, err := net.Listen("tcp", "0.0.0.0:8888")
	if err != nil {
		fmt.Println("监听失败，err:", err)
		return
	}
	defer listen.Close()
	for {
		fmt.Println("等待客户端连接。。。。。。。。")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("客户端连接失败，err:", err)

		} else {
			fmt.Printf("客户端连接成功。。。。。。。。 con = %v\n", conn)

		}
		go process(conn) // 启动一个协程，用来处理客户端发送的数据
	}
}
