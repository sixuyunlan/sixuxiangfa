package main

import (
	"fmt"
	"sixuxiangfa/customerManager/service"
)

type customerView struct {
	key             string
	loop            bool
	customerService *service.CustomerService
}

//添加客户
func (c *customerView) add() {
	fmt.Println("-----------------添加客户-----------------")
	fmt.Print("姓名:")
	name := ""
	fmt.Scanln(&name)
	fmt.Print("性别:")
	gener := ""
	fmt.Scanln(&gener)
	fmt.Print("年龄:")
	age := 0
	fmt.Scanln(&age)
	fmt.Print("电话:")
	phone := ""
	fmt.Scanln(&phone)
	fmt.Print("邮箱:")
	email := ""
	fmt.Scanln(&email)
	c.customerService.Add(name, gener, age, phone, email)
	fmt.Println("-----------------添加客户-----------------")
}

//修改客户
func (c *customerView) update() {
	fmt.Println("-----------------修改客户-----------------")
	fmt.Print("请输入客户编号:")
	id := 0
	fmt.Scanln(&id)
	i := c.customerService.FindById(id).(int)
	if i == -1 {
		fmt.Println("没有找到该客户")
		return
	}
	fmt.Print("姓名:")
	name := ""
	fmt.Scanln(&name)
	fmt.Print("性别:")
	gener := ""
	fmt.Scanln(&gener)
	fmt.Print("年龄:")
	age := 0
	fmt.Scanln(&age)
	fmt.Print("电话:")
	phone := ""
	fmt.Scanln(&phone)
	fmt.Print("邮箱:")
	email := ""
	fmt.Scanln(&email)

	c.customerService.Update(id, name, gener, age, phone, email)
}

//删除客户
func (c *customerView) delete() {
	fmt.Println("-----------------删除客户-----------------")
	fmt.Print("请输入客户编号:")
	id := 0
	fmt.Scanln(&id)
	c.customerService.Delete(id)
}

//显示所有的客户信息
func (c *customerView) list() {

	fmt.Println("-----------------客户列表-----------------")
	fmt.Println("编号\t姓名\t性别\t年龄\t电话\t\t邮箱")
	for _, v := range c.customerService.List() {
		fmt.Printf("%d\t%s\t%s\t%d\t%s\t%s\n", v.Id, v.Name, v.Gender, v.Age, v.Phone, v.Email)
	}
	fmt.Println("-----------------客户列表-----------------")
}

//显示主菜单
func (c *customerView) mainMenu() {
	for {
		fmt.Println("-----------------客户信息管理软件-----------------")
		fmt.Println("	1.添加客户")
		fmt.Println("	2.修改客户")
		fmt.Println("	3.删除客户")
		fmt.Println("	4.客户列表")
		fmt.Println("	5.退出")
		fmt.Print("请选择(1-5):")
		fmt.Scanln(&c.key)

		switch c.key {
		case "1":
			c.add()
		case "2":
			c.update()
		case "3":
			c.delete()
		case "4":
			c.list()
		case "5":
			fmt.Println("确定要退出么？请输入y/n")
			for {
				fmt.Scanln(&c.key)
				if c.key == "y" || c.key == "n" {
					break
				}
				fmt.Println("请输入正确的选项,y or n")

			}
			if c.key == "y" {
				c.loop = false
			}

		default:
			fmt.Println("输入错误，请重新输入")
		}
		if c.loop == false {
			break
		}

	}
}
func main() {
	c := customerView{
		"",
		true,
		service.NewCustomerService(),
	}
	c.mainMenu()
}
