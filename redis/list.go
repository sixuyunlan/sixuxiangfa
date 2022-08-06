package main

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

func ListRedis(url string) {
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("连接redis失败，err:", err)
		return
	}
	fmt.Println("连接成功。。。。。。。。")
	defer conn.Close()             // 关闭conn b
	conn.Do("lpush", "list1", url) // 向redis写入数据
	list1, _ := redis.Strings(conn.Do("lrange", "list1", "0", "-1"))
	fmt.Println("读取到的数据是：", list1)
	conn.Do("ltrim", "list1", "0", "9")
	list1, _ = redis.Strings(conn.Do("lrange", "list1", "0", "-1"))
	fmt.Println("读取到的数据是：", list1)

}
func main() {
	ListRedis("1,2,3,4,5,6,7,8,9,10,11,12,13")
}
