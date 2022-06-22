package main

import "fmt"

type person struct {
	name string
	age  int
}

func NewPerson() *person {
	p := new(person)
	p.name = "1111"
	p.age = 20
	return p
}

func main() {
	var sp *string
	sp = new(string)
	*sp = "1111"
	fmt.Println(*sp)
	pp := NewPerson()
	fmt.Println(pp.name, pp.age)
}
