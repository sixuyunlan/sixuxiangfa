package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {

	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("客户端连接失败，err:", err)
		return
	}
	reader := bufio.NewReader(os.Stdin)

	for {
		line, _ := reader.ReadString('\n')
		//如果客户端输入的是exit，则退出
		line = strings.Trim(line, "\r\n") //去掉换行符
		if line == "exit" {
			fmt.Println("客户端退出。。。。。。。。")
			break
		}

		_, err := conn.Write([]byte(line + "\n"))
		if err != nil {
			fmt.Println("写入失败，err:", err)
		}
		fmt.Println("客户端发送了：", line)

	}
}
