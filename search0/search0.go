package main

import (
	"fmt"
	"io/ioutil"
	"time"
)

var quary = "test"

var matches int

func main() {
	start := time.Now()
	search("/Users/")
	fmt.Println(matches)
	fmt.Println(time.Since(start))

}

func search(path string) {
	files, err := ioutil.ReadDir(path)
	if err == nil {
		for _, file := range files {
			name := file.Name()
			if name == quary {

				matches++
			}
			if file.IsDir() {
				search(path + name + "/")
			}
		}
	}
}
