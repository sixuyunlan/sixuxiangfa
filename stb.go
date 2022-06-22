package main

import (
	"fmt"
	"strconv"
)

func main() {
	str := "1"

	var b bool
	b, _ = strconv.ParseBool(str)
	fmt.Printf("%T,b:%v\n", b, b)
	str1 := "fhjg"
	var n1 int64

	n1, _ = strconv.ParseInt(str1, 10, 0)
	fmt.Printf("%T,b:%v\n", n1, n1)

	var n2 int
	n2, _ = strconv.Atoi(str1)
	fmt.Printf("%T,b:%v\n", n2, n2)
	str2 := "12.4"
	var f1 float64
	f1, _ = strconv.ParseFloat(str2, 0)
	fmt.Printf("%T,b:%v\n", f1, f1)
}
