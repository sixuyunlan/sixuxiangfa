package main

import (
	"fmt"
	"reflect"
)

func reflectTest01(b interface{}) {
	fmt.Println(reflect.TypeOf(b))
	fmt.Println(reflect.ValueOf(b))
	fmt.Println(reflect.ValueOf(b).Type())
	fmt.Println(reflect.ValueOf(b).Kind())
	fmt.Println(reflect.ValueOf(b).Interface())
	fmt.Println(reflect.ValueOf(b).Interface().(int))

}
func reflectTest02(b interface{}) {
	fmt.Println(reflect.TypeOf(b))
	fmt.Println(reflect.ValueOf(b))
	fmt.Println(reflect.ValueOf(b).Interface())

}

type Student struct {
	Name string
	Age  int
}

type Monster struct {
	Name string
	Age  int
}

func main() {
	num := 100
	reflectTest01(num)

	student := Student{Name: "张三", Age: 18}

	reflectTest02(student)
}
