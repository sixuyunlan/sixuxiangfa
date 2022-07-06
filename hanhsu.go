package main

import "fmt"

func sum(a, b int) (sum any, err any) {
	if a < 0 || b < 0 {
		err = fmt.Errorf("a or b is negative")
		return
	}
	return a + b, nil
}

type person1 struct {
	name string
	age  int
	addr address
}
type address struct {
	city string
	area string
}

func main() {
	fmt.Println(sum(1, 2))
	fmt.Println(sum(-1, -2))

	sum2 := func(a, b int) (sum any) {
		return a + b
	}
	fmt.Println(sum2(2, 3))

	cl := colsure()
	fmt.Println(cl())
	fmt.Println(cl())
	fmt.Println(cl())
	p := person1{"牛逼", 18, address{"北京", "朝阳"}}
	fmt.Println(p)
	fmt.Println(p.addr.city)
	p.Walk()

}
func colsure() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

type WalkRun interface {
	Walk()
	Run()
}

func (p person1) Walk() {
	fmt.Println("walk")
}
func (p person1) Run() {
	fmt.Println("run")
}
