package main

import (
	"errors"
	"fmt"
)

type Stack struct {
	MaxTop int    //最大的数据
	Top    int    //栈顶
	array  [5]int //数组

}

func (this *Stack) Push(value int) (err error) { //把数据压入栈中
	if this.Top == this.MaxTop-1 {
		fmt.Println("栈满")
		return errors.New("栈满")
	}
	this.Top++                   //栈顶加1
	this.array[this.Top] = value //把数据压入栈中

	return
}

func (this *Stack) List() { //遍历栈
	if this.Top == -1 {
		fmt.Println("栈空")
		return
	}
	for i := this.Top; i >= 0; i-- { //从栈顶开始遍历
		fmt.Printf("arr[%d]=%d\n", i, this.array[i])
	}
}

func (this *Stack) Pop() (value int, err error) { //弹出栈顶的数据
	if this.Top == -1 { //栈空
		fmt.Println("栈空")
		return 0, errors.New("栈空")
	}
	value = this.array[this.Top] //取出栈顶的数据
	this.Top--                   //栈顶减1
	return value, nil            //返回数据
}

func main() {
	stack := &Stack{
		MaxTop: 5,
		Top:    -1,
	}
	stack.Push(1) //把数据压入栈 中
	stack.Push(2) //把数据压入栈中
	stack.Push(3) //把数据压入栈中
	stack.Push(4) //把数据压入栈中
	stack.Push(5) //把数据压入栈中

	stack.List() //遍历栈

	value, _ := stack.Pop() //弹出栈顶的数据
	fmt.Println("出栈value=", value)
	stack.List() //遍历栈

	value, _ = stack.Pop() //弹出栈顶的数据
	fmt.Println("出栈value=", value)
	stack.List() //遍历栈

	value, _ = stack.Pop() //弹出栈顶的数据
	fmt.Println("出栈value=", value)
	stack.List() //遍历栈

	value, _ = stack.Pop() //弹出栈顶的数据
	fmt.Println("出栈value=", value)
	stack.List() //遍历栈

	value, _ = stack.Pop() //弹出栈顶的数据
	fmt.Println("出栈value=", value)
	stack.List() //遍历栈

}
