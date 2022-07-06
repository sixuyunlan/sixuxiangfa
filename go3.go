package main

import (
	"fmt"
	"time"
)

func main() {
	cha := make(chan int)
	chb := make(chan int)

	go func() {
		for {

			select {
			case a := <-cha:
				fmt.Println("channel A get new value", a)
			case b := <-chb:
				fmt.Println("channel B get new value", b)
			default:
				time.Sleep(time.Second * 1)
				fmt.Println("channel A and B both are empty")
			}
		}
	}()
	go func() {
		cha <- 1
		cha <- 2

	}()

	chb <- 3
	chb <- 4
	ma := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}
	ma = append(ma, []int{13, 14, 15, 16})
	//遍历ma
	for _, v := range ma {
		for _, v2 := range v {
			fmt.Printf("%d\t", v2)
		}
		fmt.Println()

	}

}
