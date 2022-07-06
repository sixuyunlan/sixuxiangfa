package main

import "fmt"

func main() {
	var x = 0
	fmt.Println(x)

	for i := 0; i < 10; i++ {

		if i%2 != 1 {
			continue

		}
		x += i
	}
	fmt.Println(x)
}
