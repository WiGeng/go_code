package main

import (
	"awesomeProject/src/103_go_code/8_object-oriented/D_Encapsulation/b_exercise/model"
	"fmt"
)

func main() {
	//创建一个account变量
	account := model.NewAccount("6230270101", "888888", 4000000)
	if account != nil {
		fmt.Println("创建成功=", account)
	} else {
		fmt.Println("创建失败")
	}
}
