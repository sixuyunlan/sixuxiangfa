package main

import (
	"fmt"
	"time"
)

func porc() {
	panic("OK")
}

func main() {
	go func() {
		t := time.NewTicker(time.Second)
		for {
			select {
			case <-t.C:
				go func() {
					defer func() {
						if err := recover(); err != nil {
							fmt.Println(err)
						}
					}()
					porc()
				}()
			}
		}
	}()
	select {}
}
