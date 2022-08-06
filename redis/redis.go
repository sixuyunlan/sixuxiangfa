package main

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

//编写一个函数，运用redis的list命令，记录客户最近浏览的10次url，并且每次浏览的时候，将url添加到list的头部，如果list的长度超过10，则将list的尾部的url删除

func main() {
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("连接redis失败，err:", err)
		return
	}
	fmt.Println("连接成功。。。。。。。。")
	defer conn.Close()                               // 关闭conn b
	conn.Do("set", "name", "tom")                    // 向redis写入数据
	sting, _ := redis.String(conn.Do("get", "name")) // 从redis读取数据
	fmt.Println("读取到的数据是：", sting)
	conn.Do("hset", "user1", "age", "18", "name", "tom", "addr", "beijing") // 向redis写入数据
	hash, _ := redis.StringMap(conn.Do("hgetall", "user1"))
	fmt.Println("读取到的数据是：", hash)
	conn.Do("del", "list")
	conn.Do("lpush", "list", "a", "b", "c") // 向redis写入数据
	conn.Do("lpush", "list", "d", "e", "f")
	conn.Do("rpush", "list", "0")
	list, _ := redis.Strings(conn.Do("lrange", "list", "0", "-1"))
	fmt.Println("读取到的数据是：", list)
	conn.Do("sadd", "emails", "asd@sohu.com", "jack@qq.com") // 向redis写入数据
	set, _ := redis.Strings(conn.Do("smembers", "emails"))
	fmt.Println("读取到的数据是：", set)
	conn.Do("hset", "Monster", "name", "Monster", "age", "18", "skill", "eat")
	Monster, _ := redis.Strings(conn.Do("hgetall", "Monster"))
	for i, v := range Monster {
		fmt.Println(i, v)
	}

}
