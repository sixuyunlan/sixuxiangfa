package main

import (
	"errors"
	"fmt"
)

type Stack struct {
	MaxTop int     //最大的数据
	Top    int     //栈顶
	array  [20]int //数组

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

//判断一个字符是不是运算符
func (this *Stack) isOperator(value byte) bool {
	if value == '+' || value == '-' || value == '*' || value == '/' {
		return true
	}
	return false
}

//运算的方法
func (this *Stack) calc(num1, num2 int, oper byte) int {
	res := 0
	switch oper {
	case '+':
		res = num1 + num2
	case '-':
		res = num1 - num2
	case '*':
		res = num1 * num2
	case '/':
		res = num1 / num2
	default:
		fmt.Println("错误的运算符")
	}
	return res
}

//编写一个方法，返回某个运算符的优先级
func (this *Stack) priority(oper byte) int {
	var res = 0
	switch oper {
	case '+':
		return 0
	case '-':
		return 0 //
	case '*':
		return 1
	case '/':
		return 1
	default:
		fmt.Println("错误的运算符")
	}
	return res
}

func main() {
	numStack := &Stack{

		MaxTop: 20,
		Top:    -1,
	}
	operStack := &Stack{ //操作符栈
		MaxTop: 20,
		Top:    -1,
	}
	exp := "3+2*6-2"
	index := 0
	num1 := 0
	num2 := 0
	oper := '+'
	res
	for {
		ch := exp[index : index+1]      //取出一个字符
		temp := []byte(ch)[0]           //转换成字节
		if operStack.isOperator(temp) { //判断是不是运算符

			if operStack.Top == -1 { //如果操作符栈为空

			//如果是运算符，则把运算符压入操作符栈
			operStack.Push(temp)
			} else {
				if operStack.priority(temp) > operStack.priority(operStack.array[operStack.Top]) { //如果当前运算符优先级大于栈顶运算符优先级
					operStack.Push(temp)
				} 


		} else { //如果不是运算符，则把数字压入数字栈
			num1, _ := numStack.Pop()              //弹出数字栈的数据
			num2, _ := numStack.Pop()              //弹出数字栈的数据
			oper, _ := operStack.Pop()             //弹出操作符栈的数据
			res := numStack.calc(num1, num2, oper) //计算结果
			numStack.Push(res)                     //把结果压入数字栈
			numStack.Push(temp)                    //把数字压入数字栈

		}
	}
}
