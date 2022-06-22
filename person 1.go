package main

import (
	"fmt"
	"sixuxiangfa/person"
)

func main() {
	p := person.Nemperson("lili")
	p.SetAge(50)
	fmt.Println(p.Name)
	fmt.Println(p.GetAge())
	fmt.Println(*p)
	fmt.Println(p)
}
