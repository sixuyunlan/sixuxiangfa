package main

import (
	"fmt"
	"unicode/utf8"
)

func main() { //main函数
	array := [...]any{"a", "b", "c", 1} //数组
	for _, v := range array {           // _ 可以忽略
		fmt.Println(v) // a b c 1

	}
	slice := array[:2] //切片
	slice[1] = 0       //修改切片
	fmt.Println(slice) //[a 0]
	b := array         //指针
	b[0] = "d"         //修改指针
	fmt.Println(array) //[d b c 1]

	slice1 := make([]any, 3, 5)           //切片
	fmt.Println(slice1)                   //[0 0 0]
	slice2 := []any{1, 2, 3}              //切片
	fmt.Println(slice2)                   //[1 2 3]
	fmt.Println(len(slice2), cap(slice2)) //3 3

	slice3 := append(slice2, 4, 5, 6)     //添加元素
	fmt.Println(slice3)                   //[1 2 3 4 5 6]
	fmt.Println(len(slice3), cap(slice3)) //7 7

	s := "Hello我是你爹"
	bs := []byte(s)
	fmt.Println(bs)
	fmt.Println(s[0], s[1], s[15])
	fmt.Println(utf8.RuneCountInString(s))

	aa := [][]any{{1, 2, 3}, {4, 5, 6}}
	aa = append(aa, []any{7, 8, 9})
	fmt.Println(aa)

}
