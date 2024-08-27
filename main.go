package main

import (
	"fmt"
	"ticketget/cpp"
)

func main() {
	var a int
	c := cpp.NewCppCrawler()
	for {
		fmt.Println("----------本项目仅供学习交流使用，禁止用于商业用途-----------------")
		fmt.Println("-------CPP 抢票助手------")
		fmt.Println("1.本地信息重置")
		fmt.Println("2.登录")
		fmt.Println("3.票种选择")
		fmt.Println("4.购票人选择")
		fmt.Println("5.直接抢票")
		fmt.Println("6.定时抢票")
		fmt.Println("7.时间间隔设置")
		fmt.Println("8.版本重置")
		fmt.Println("9.退出")
		fmt.Println("-------------------------")
		fmt.Println("请选择:")
		fmt.Scanln(&a)
		switch a {
		case 1:
			c.InfoClear()
		case 2:
			c.Login()
		case 3:
			c.GetTicketInfo()
			c.ChoseTicket()
		case 4:
			c.GetPersonInfo()
			c.ChosePerson()
		case 5:
			c.GrapTicket()
		case 6:
			c.CronTicket()
		case 7:
			c.SetTimeSleep()
		case 8:
			c.SetVersion()
		case 9:
			goto breakFlag
		default:
			fmt.Println("输入非法！")
		}
	}
breakFlag:
	fmt.Println("程序结束!")
}
