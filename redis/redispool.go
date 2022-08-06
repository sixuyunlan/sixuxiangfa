package main

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

var pool *redis.Pool

func init() {
	pool = &redis.Pool{
		MaxIdle:     8,
		MaxActive:   0,
		IdleTimeout: 100,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "localhost:6379")
		},
	}

}

func main() {
	conn := pool.Get()
	defer conn.Close()
	conn.Do("set", "name", "tom")
	r, _ := redis.String(conn.Do("get", "name"))
	conn.Do("set", "name2", "tom2")
	r2, _ := redis.String(conn.Do("get", "name2"))
	fmt.Println(r, r2)
}
