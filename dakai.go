package main

import (
	"fmt"
	"io/ioutil"
)

func main() {

	//file, _ := os.Open("C:/Users/Administrator/Desktop/纵千里.txt")
	//打开文件
	file := "C:/Users/Administrator/Desktop/纵千里.txt"
	c, _ := ioutil.ReadFile(file) //读取文件内容
	fmt.Printf(":%s\n", c)
	//读取文件
	//buf := make([]byte, 1024)
	//for {
	//	n, _ := file.Read(buf)
	//	if n == 0 {
	//		break
	//	}
	//	fmt.Printf("%s", buf[:n])
	//}

	//reader := bufio.NewReader(file)
	//for {
	//	line, _, _ := reader.ReadLine()
	//	if line == nil {
	//		break
	//	}
	//	fmt.Println(string(line))
	//}

}
