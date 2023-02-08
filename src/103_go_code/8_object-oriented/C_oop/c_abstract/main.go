package main

import (
	"fmt"
)

//定义一个结构体Account
type Account struct {
	AccountNo string
	Pwd       string
	Balance   float64
}

//1. 存款
func (account *Account) Deposite(money float64, pwd string) {
	//看下输入的密码是否正确
	if pwd != account.Pwd {
		fmt.Println("你输入的密码不正确")
		return
	}
	//看看存款金额是否正确
	if money <= 0 {
		fmt.Println("你输入的金额不正确")
		return
	}
	account.Balance += money
	fmt.Println("存款成功~~")

}

//2. 取款
func (account *Account) WithDraw(money float64, pwd string) {
	//看下输入的密码是否正确
	if pwd != account.Pwd {
		fmt.Println("你输入的密码不正确")
		return
	}
	//看看取款金额是否正确
	if money <= 0 || money > account.Balance {
		fmt.Println("你输入的金额不正确")
		return
	}
	account.Balance -= money
	fmt.Println("取款成功~~")

}

//3. 查询余额
func (account *Account) Query(pwd string) {
	//看下输入的密码是否正确
	if pwd != account.Pwd {
		fmt.Println("你输入的密码不正确")
		return
	}
	fmt.Printf("你的账号为=%v 余额=%v \n", account.AccountNo, account.Balance)

}

func main() {

	//测试
	account := Account{
		AccountNo: "6230270100001472636",
		Pwd:       "123456",
		Balance:   1000.0,
	}

	account.Query("666666")
	account.Deposite(200.0, "666666")
	account.Query("666666")
	account.WithDraw(150.0, "666666")
	account.Query("666666")
	for {
		var key int
		var password string
		var money float64
		fmt.Println(" --------------------------------------")
		fmt.Println("|  1. Deposite [1]                     |")
		fmt.Println("|  2. WithDraw [2]                     |")
		fmt.Println("|  3. Query    [3]                     |")
		fmt.Println(" --------------------------------------")
		fmt.Println("请输入需要执行的操作[Deposite[1]|WithDraw[2]|Query[3]]")
		fmt.Scanln(&key)
		switch key {
		case 1:
			fmt.Println("请输入银行卡密码")
			fmt.Scanln(&password)
			fmt.Println("请输入要存入的金额")
			fmt.Scanf("%v", &money)
			account.Deposite(money, password)
		case 2:
			fmt.Println("请输入银行卡密码")
			fmt.Scanln(&password)
			fmt.Println("请输入要取出的金额")
			fmt.Scanf("%v", &money)
			account.WithDraw(money, password)
		case 3:
			fmt.Println("请输入银行卡密码")
			fmt.Scanln(&password)
			account.Query(password)
		}
	}

}
