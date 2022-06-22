package main

import (
	"errors"
	"fmt"
	"os"
)

type Queue struct { //队列结构体
	maxSize int    //队列最大长度
	array   [4]int //队列数组
	front   int    //队列头
	rear    int    //队列尾
}

func (this *Queue) AddQueue(val int) (err error) { //添加队列
	if this.rear == this.maxSize-1 { //队列满了
		return errors.New("queue is full") //队列满了
	}
	this.rear++ //指针后移

	this.array[this.rear] = val //添加数据
	return
}
func (this *Queue) GetQueue() (val int, err error) { //取出队列
	if this.rear == this.front { //队列空了
		return -1, errors.New("queue is empty") //队列空了
	}
	this.front++                       //指针后移
	return this.array[this.front], nil //返回队列数据
}
func (this *Queue) ShowQueue() { //获取队列
	for i := this.front + 1; i < this.rear; i++ { //遍历队列
		fmt.Printf("array[%d]=%d\t", i, this.array[i]) //打印队列数据
	}
}

func main() { //主函数

	queue := &Queue{ //创建队列
		maxSize: 5,  //队列最大长度
		front:   -1, //队列头
		rear:    -1, //队列尾
	}
	var key string //接收用户输入
	var val int    //接收用户输入

	for { //循环
		fmt.Println("1,输入add 表示添加数据到队列") //打印菜单
		fmt.Println("2,输入get 表示从队列获取数据") //打印菜单
		fmt.Println("3,输入show 表示显示队列")   //打印菜单
		fmt.Println("4,输入exit 表示退出程序")   //打印菜单

		fmt.Scanln(&key) //接收用户输入
		switch key {     //判断用户输入
		case "add": //如果用户输入add
			fmt.Println("请输入数据")       //打印提示
			fmt.Scanln(&val)           //接收用户输入
			err := queue.AddQueue(val) //调用添加队列函数
			if err != nil {            //如果添加队列函数返回错误
				fmt.Println("shibai") //打印错误信息
			} else { //如果添加队列函数返回成功
				fmt.Println("添加成功") //打印成功信息
			}

		case "get": //如果用户输入get
			val, err := queue.GetQueue() //调用取出队列函数
			if err != nil {              //如果取出队列函数返回错误
				fmt.Println("shibai") //打印错误信息
			} else { //如果取出队列函数返回成功
				fmt.Println("取出的数据为", val) //打印取出的数据
			}

		case "show": //如果用户输入show
			queue.ShowQueue() //调用显示队列函数

		case "exit": //如果用户输入exit
			os.Exit(0) //退出程序

		} //end switch
	} //end for

} //end main
