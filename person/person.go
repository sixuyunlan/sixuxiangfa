package person

import "fmt"

type person struct {
	Name string
	age  int
}

func Nemperson(name string) *person {
	return &person{
		Name: name,
	}
}
func (p *person) SetAge(age int) {
	if age > 0 && age < 150 {
		p.age = age
	} else {
		fmt.Println("error")
	}
}
func (p *person) GetAge() int {
	return p.age
}
