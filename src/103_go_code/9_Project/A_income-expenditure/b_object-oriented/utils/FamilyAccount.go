package utils

import "fmt"

type FamilyAccount struct {
	key     string
	loop    bool
	balance float64
	money   float64
	note    string
	flag    bool
	details string
}

// 显示收支明细
func (this *FamilyAccount) ShowDetils() {
	fmt.Println("-----------------当前收支明细记录-----------------")
	if this.flag == true {
		fmt.Println(this.details)
	} else {
		fmt.Println("当前没有收支明细... 来一笔吧!")
	}
}

// 显示收入

func (this *FamilyAccount) Incoming() {
	fmt.Println("本次收入金额:")
	fmt.Scanln(&this.money)
	this.balance += this.money // 修改账户余额
	fmt.Println("本次收入说明:")
	fmt.Scanln(&this.note)
	//将这个收入情况，拼接到details变量
	//收入    11000           1000            有人发红包
	this.details += fmt.Sprintf("\n收入\t%v\t\t%v\t\t%v", this.balance, this.money, this.note)
	this.flag = true
}

// 显示支出
func (this *FamilyAccount) Pay() {
	fmt.Println("本次支出金额:")
	fmt.Scanln(&this.money)
	//这里需要做一个必要的判断
	if this.money > this.balance {
		fmt.Println("余额不足")
	}
	this.balance -= this.money
	fmt.Println("本次支出说明:")
	fmt.Scanln(&this.note)
	this.details += fmt.Sprintf("\n支出\t%v\t\t%v\t\t%v", this.balance, this.money, this.note)
	this.flag = true
}

//  退出程序
func (this *FamilyAccount) Exit() {
	fmt.Println("你确定要退出吗? y/n")
	choice := ""
	for {
		fmt.Scanln(&choice)
		if choice == "y" || choice == "n" {
			break
		}
		fmt.Println("你的输入有误，请重新输入 y/n")
	}
	if choice == "y" {
		this.loop = false
	}

}

// 定义工厂模式构造方法，显示一个*FamilyAccount实例
func NewFamilyAccount() *FamilyAccount {
	return &FamilyAccount{
		key:     "",
		loop:    true,
		balance: 100000.0,
		money:   0.0,
		note:    "",
		flag:    true,
		details: "收支\t账户金额\t收支金额\t说明",
	}
}

// 显示主菜单
func (this *FamilyAccount) ShowMenu() {
	for {
		fmt.Println("\n-----------------家庭收支记账软件-----------------")
		fmt.Println("                  1 收支明细")
		fmt.Println("                  2 登记收入")
		fmt.Println("                  3 登记支出")
		fmt.Println("                  4 退出软件")
		fmt.Print("请选择(1-4)：")
		fmt.Scanln(&this.key)

		switch this.key {
		case "1":
			this.ShowDetils()
		case "2":
			this.Incoming()
		case "3":
			this.Pay()
		case "4":
			this.Exit()
		default:
			fmt.Println("请输入正确的选项..")

		}
		if !this.loop {
			break
		}
	}
}
