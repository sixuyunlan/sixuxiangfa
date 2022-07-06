package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	key := ""
	loop := true
	balance := 10000.0
	money := 0.0
	note := ""
	details := "时间\t收支\t金额\t余额\t说明"
	for {
		fmt.Println("-----------------家庭收支记账软件-----------------")
		fmt.Println("	1.收支明细")
		fmt.Println("	2.登记收入")
		fmt.Println("	3.登记支出")
		fmt.Println("	4.退出软件")
		fmt.Print("请选择(1-4):")
		fmt.Scanln(&key)
		switch key {
		case "1":
			fmt.Println("收支明细")
			fmt.Println(details)
		case "2":
			fmt.Println("登记收入")
			fmt.Println("请输入金额:")
			fmt.Scanln(&money)
			balance = balance + money
			fmt.Println("请输入说明:")
			fmt.Scanln(&note)
			details += fmt.Sprintf("\n%s\t收入\t%f\t%f\t%s", time.Now().Format("2006-01-02 15:04:05"), money, balance, note)
			fmt.Println("登记成功")
		case "3":
			fmt.Println("登记支出")
			fmt.Println("请输入金额:")
			fmt.Scanln(&money)
			if money > balance {
				fmt.Println("余额不足")
				break //跳出当前循环,中断上{一个 的执行
			}
			balance -= money

			fmt.Println("请输入说明:")
			fmt.Scanln(&note)
			details += fmt.Sprintf("\n%s\t支出\t%f\t%f\t%s", time.Now().Format("2006-01-02 15:04:05"), money, balance, note)
			fmt.Println("登记成功")
		case "4":
			fmt.Println("确定要退出吗？(y/n)")
			for {
				fmt.Scanln(&key)
				if key == "y" || key == "n" {
					break
				}
				fmt.Println("请输入正确的选项")
			}
			if key == "y" {
				loop = false
			}
		default:
			fmt.Println("请输入正确的选项")
		}
		if !loop {
			fmt.Println("程序结束")
			fmt.Println("欢迎下次使用") //如果退出软件，则跳出循环
			os.Exit(0)            //退出循环
		}
	}
}
