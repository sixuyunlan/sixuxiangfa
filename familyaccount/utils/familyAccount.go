package utils

import (
	"fmt"
	"os"
	"time"
)

type FamilyAccount struct {
	key     string
	loop    bool
	balance float64
	money   float64
	note    string
	flag    any
	details string
}

func NewFamilyAccount() *FamilyAccount { //创建一个家庭收支记账软件
	return &FamilyAccount{
		key:     "",
		loop:    true,
		balance: 10000.0,
		money:   0.0,
		note:    "",
		flag:    false,
		details: "日期\t收支\t金额\t余额\t说明",
	}
}

//再使用该软件前，有一个登入软件的功能，可以让用户输入账号和密码，如果账号和密码正确，则可以进入主菜单，否则，则提示用户账号或密码错误，并且提示用户重新输入。
//如果用户输入了账号和密码，则可以进入主菜单，否则，则提示用户账号或密码错误，并且提示用户重新输入。
func (fa *FamilyAccount) Login() {

	fmt.Println("请输入密码:")
	for {
		fmt.Scanln(&fa.key)
		if fa.key == "123456" {
			fmt.Println("登入成功")
			break
		} else {
			fmt.Println("账号或密码错误，请重新输入")
		}

	}
}

func (this *FamilyAccount) showDetails() { //显示收支明细
	if this.flag == true {
		fmt.Println("收支明细")
		fmt.Println(this.details)
	} else {
		fmt.Println("暂无收支明细")
	}

}
func (fa *FamilyAccount) registerIncome() { //登记收入
	fmt.Println("登记收入")
	fmt.Println("请输入金额:")
	fmt.Scanln(&fa.money)
	fa.balance = fa.balance + fa.money
	fmt.Println("请输入说明:")
	fmt.Scanln(&fa.note)

	fa.details += fmt.Sprintf("\n%s\t收入\t%f\t%f\t%s", time.Now().Format("2006-01-02 15:04:05"), fa.money, fa.balance, fa.note)
	fa.flag = true
	fmt.Println("登记成功")
}

func (fa *FamilyAccount) registerExpense() { //登记支出
	fmt.Println("登记支出")
	fmt.Println("请输入金额:")
	fmt.Scanln(&fa.money)
	if fa.money > fa.balance {
		fmt.Println("余额不足")
		return
	}
	fa.balance -= fa.money

	fmt.Println("请输入说明:")
	fmt.Scanln(&fa.note)
	fa.details += fmt.Sprintf("\n%s\t支出\t%f\t%f\t%s", time.Now().Format("2006-01-02 15:04:05"), fa.money, fa.balance, fa.note)
	fa.flag = true

	fmt.Println("登记成功")
}

func (fa *FamilyAccount) exitProgram() { //退出程序

	fmt.Println("确定要退出吗？(y/n)")
	for {
		fmt.Scanln(&fa.key)
		if fa.key == "y" || fa.key == "n" {
			break
		}
		fmt.Println("请输入正确的选项,y or n")
	}
	if fa.key == "y" {
		fa.loop = false
	}

}

func (fa *FamilyAccount) MainMenu() {
	fa.Login()

	for {
		fmt.Println("-----------------家庭收支记账软件-----------------")
		fmt.Println("	1.收支明细")
		fmt.Println("	2.登记收入")
		fmt.Println("	3.登记支出")
		fmt.Println("	4.退出软件")
		fmt.Print("请选择(1-4):")
		fmt.Scanln(&fa.key)
		switch fa.key {
		case "1":
			fa.showDetails()
		case "2":
			fa.registerIncome()
		case "3":
			fa.registerExpense()
		case "4":
			fa.exitProgram()
		default:
			fmt.Println("请输入正确的选项")
		}
		if !fa.loop {
			fmt.Println("程序结束")
			fmt.Println("欢迎下次使用") //如果退出软件，则跳出循环
			os.Exit(0)            //退出循环
		}

	}
}
