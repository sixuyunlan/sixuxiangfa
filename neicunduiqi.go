package main

import (
	"fmt"
	"unsafe"
)

type DemoA struct {
	A int8
	B int64
	C int16
}
type DemoB struct {
	A int8
	C int16
	B int64
}

func main() {
	fmt.Println(unsafe.Sizeof(DemoA{})) //内存补齐
	fmt.Println(unsafe.Sizeof(DemoB{}))

}
