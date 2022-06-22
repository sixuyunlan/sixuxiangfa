package main

import (
	"fmt"
	"io/ioutil"
	"time"
)

var quary = "test"
var matches int

var workCount = 0
var maxWorkCont = 32

var searchRequest = make(chan string)
var workDone = make(chan bool)
var foundMatch = make(chan bool)

func main() {
	start := time.Now()
	workCount = 1
	go search("/Users/", true)
	waitForWorkers()
	fmt.Println(matches, "matches")
	fmt.Println(time.Since(start))
}
func waitForWorkers() {
	for {
		select {
		case path := <-searchRequest:
			workCount++
			go search(path, true)
		case <-workDone:
			workCount--
			if workCount == 0 {
				return
			}
		case <-foundMatch:
			matches++

		}

	}

}

func search(path string, master bool) {

	files, err := ioutil.ReadDir(path)
	if err == nil {
		for _, file := range files {
			name := file.Name()
			if name == quary {
				foundMatch <- true
			}
			if file.IsDir() {
				if workCount < maxWorkCont {
					searchRequest <- path + name + "/"
				} else {
					search(path+name+"/", false)
				}

			}
		}

	}
	if master {
		workDone <- true
	}

}
