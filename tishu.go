package main

import (
	"fmt"
	"sync"
	"time"
)

func count(n int, animals any) {
	for i := 0; i < n; i++ {
		fmt.Println(i+1, animals)
		time.Sleep(time.Millisecond * 1)
	}

}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		count(5, "羊")
		wg.Done()
	}()
	go func() {
		count(5, "牛")
		wg.Done()
	}()
	wg.Wait()
}
