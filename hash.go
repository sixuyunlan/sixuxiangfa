package main

import (
	"fmt"
	"os"
)

type Emp struct {
	Id   int
	Name string
	Next *Emp
}

func (e *Emp) ShowMe() {
	fmt.Printf("链表%d 找到该雇员%d\n", e.Id, e.Name)
}

type Emplink struct {
	Head *Emp
}

//添加员工的方法，保证添加时，编号从小到大排列
func (e *Emplink) Inert(emp *Emp) { //添加员工
	if e.Head == nil { //如果链表为空，则直接添加
		e.Head = emp //添加到链表头
		return       //返回
	}
	cur := e.Head         //设置当前指针
	for cur.Next != nil { //如果当前指针不为空，则继续向后遍历
		cur = cur.Next //继续向后遍历
	}
	cur.Next = emp //添加到链表尾
}

//显示链表中的所有员工
func (e *Emplink) ShowLink(index int) {
	if e.Head == nil {
		fmt.Printf("链表%d 为空\n", index)
		return
	}
	for cur := e.Head; cur != nil; cur = cur.Next {
		fmt.Printf("链表%d 员工id=%d 员工%s", index, cur.Id, cur.Name)

	}
	fmt.Println()

}

//根据id查找对应的员工，如果没有找到，则返回nil
func (e *Emplink) FindById(id int) *Emp {
	if e.Head == nil {
		return nil
	}
	cur := e.Head
	for cur != nil && cur.Id == id {

		return cur
	}
	cur = cur.Next
	return nil
}

type HashTable struct {
	LinkArr [7]Emplink
}

//给HashTable 编写Insert雇员的方法
func (e *HashTable) Insert(emp *Emp) {
	index := e.Hash(emp.Id)
	e.LinkArr[index].Inert(emp)
}

//编写方法，显示hashtable中的所有员工
func (e *HashTable) ShowAllEmp() {
	for i := 0; i < len(e.LinkArr); i++ {
		e.LinkArr[i].ShowLink(i)
	}
}

//编写一个散列方法
func (e *HashTable) Hash(id int) int {
	return id % 7
}

//编写方法，完成查找
func (e *HashTable) FindById(id int) *Emp {
	linkNo := e.Hash(id)
	return e.LinkArr[linkNo].FindById(id)
}

func main() {
	key := ""
	id := 0
	name := ""
	var hashtable HashTable
	for {
		fmt.Println("==============雇员系统菜单===============")
		fmt.Println("input 表示添加雇员")
		fmt.Println("show  表示显示雇员")
		fmt.Println("find  表示查找雇员")
		fmt.Println("exit  表示退出系统")
		fmt.Println("=============================================")
		fmt.Scanln(&key)
		switch key {
		case "input":
			fmt.Println("请输入雇员编号：")
			fmt.Scanln(&id)
			fmt.Println("请输入雇员姓名：")
			fmt.Scanln(&name)
			emp := &Emp{id, name, nil}
			hashtable.Insert(emp)
		case "show":
			hashtable.ShowAllEmp()
		case "find":
			fmt.Println("请输入要查找的雇员编号：")
			fmt.Scanln(&id)
			emp := hashtable.FindById(id)
			if emp != nil {
				fmt.Println("找到了！")
				fmt.Println(emp)
			} else {
				fmt.Println("没有找到！")
			}
		case "exit":
			fmt.Println("退出系统")
			os.Exit(0)

		default:
			fmt.Println("输入错误，请重新输入")

		}
	}
}
