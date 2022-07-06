package main

import (
	"errors"
	"fmt"
)

func main() {
	e := errors.New("原始错e")            //创建一个错误
	w := fmt.Errorf("Wrap了一个错误:%w", e) //创建一个错误
	fmt.Println(w)                     //输出错误

}
