package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var status int64

func main() {
	c := sync.NewCond(&sync.Mutex{})
	for i := 0; i < 11; i++ {
		go listen(c)

	}
	go broadcast(c)
	time.Sleep(time.Second * 1)

}
func broadcast(c *sync.Cond) {
	atomic.StoreInt64(&status, 1)
	c.Broadcast()
}
func listen(c *sync.Cond) {
	c.L.Lock()
	for atomic.LoadInt64(&status) != 1 {
		c.Wait()
	}
	fmt.Println("listen")

	c.L.Unlock()
}
