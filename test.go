package main

import (
	"fmt"
)

type Chinese struct {
	Name any
	Age  any
}

func (p Chinese) sayhello() {
	fmt.Println("你好")

}
func (p Chinese) niuyangge() {
	fmt.Println("东北的 喜欢扭秧歌")
}

type Amarican struct {
	Name any
	Age  any
}

func (p Amarican) sayhello() {
	fmt.Println("hi")

}
func (p Amarican) disco() {
	fmt.Println("love disco")
}

type SayHelloer interface {
	sayhello()
}

func greet(s SayHelloer) {
	s.sayhello()
	switch s.(type) {
	case Chinese:
		ch := s.(Chinese)
		ch.niuyangge()
	case Amarican:
		am := s.(Amarican)
		am.disco()
	}
}
func main() {
	c := Chinese{"王强", 21}
	fmt.Println(c)
	greet(c)
	a := Amarican{"Lisa", 18}
	fmt.Println(a)
	greet(a)
	var arr = []SayHelloer{Amarican{"rose", 21}, Chinese{"丽丽", 20}, Chinese{"鹿鹿", 19}}
	fmt.Println(arr)

}
