package service

import (
	"fmt"
	"sixuxiangfa/customerManager/model"
)

type CustomerService struct { //定义一个客户服务结构体
	customers   []model.Customer
	customerNum int
}

//编写一个方法，可以返回 *CustomerService
func NewCustomerService() *CustomerService {
	customerService := &CustomerService{}
	customerService.customerNum = 1
	customer := model.NewCustomer(1, "张三", "男", 20, "123456789", "zs@sohu.com")
	customerService.customers = append(customerService.customers, *customer)
	return customerService

}

//返回客户切片
func (cs *CustomerService) List() []model.Customer {
	return cs.customers //返回客户切片

}
func (cs *CustomerService) Add(name string, gener string, age int, phone string, email string) { //添加客户
	cs.customerNum++                                                              //客户编号自增
	customer := model.NewCustomer(cs.customerNum, name, gener, age, phone, email) //创建一个客户
	cs.customers = append(cs.customers, *customer)                                //将客户添加到客户切片

}

func (cs *CustomerService) Delete(id int) {
	var index int = cs.FindById(id).(int)
	for {
		//获取客户编号
		if index != -1 { //如果客户编号存在
			cs.customers = append(cs.customers[:index], cs.customers[index+1:]...) //切片删除
			fmt.Println("-----------------删除成功-----------------")

		} else {
			fmt.Println("Error:客户不存在")

		}
		break
	}

}

func (cs *CustomerService) FindById(id int) interface{} {
	for i, v := range cs.customers { //遍历客户切片
		if v.Id == id { //如果客户编号相等
			return i //返回客户编号
		} //如果客户编号不相等
	}
	return -1 //返回-1

}

func (cs *CustomerService) Update(id int, name string, gender string, age int, phone string, email string) {
	index := cs.FindById(id).(int)
	for {
		if index != -1 {
			cs.customers[index].Name = name
			cs.customers[index].Gender = gender
			cs.customers[index].Age = age
			cs.customers[index].Phone = phone
			cs.customers[index].Email = email

			fmt.Println("-----------------修改成功-----------------")

		} else {
			fmt.Println("Error:客户不存在")

		}
		break

	}

}
