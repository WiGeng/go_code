package main

import (
	"fmt"
	_ "unsafe"
)

/* 介绍
在编程中，需要接收用户输入的数据，就可以使用键益输入语句来获取。
InputDemo.go
	步骤：
	1) 导入fmt包
	2）週用fmt包的fmt.Scanln()或者fmt.Scanf()
		函数说明：
		func Scanln
		func Scanln(a ...interface{}) (n int, err error)
		Scanln类似Scan，但会在换行时才停止扫描。最后一个条目后必须有换行或者到达结束位置。

		func Sscanf
		func Sscanf(str string, format string, a ...interface{}) (n int, err error)
		Sscanf从字符串str扫描文本，根据format 参数指定的格式将成功读取的空白分隔的值保存进成功传递给本函数的参数。
		返回成功扫描的条目个数和遇到的任何错误。
*/

func main() {
	/*案例演示：
	要求：可以从控制台接收用户信息，【姓名，年龄，是否在职，薪水】
	1) 使用 fmt.Scanln()获取
	2) 使用 fmt.Scanf() 获取
	*/
	var name string
	var age int
	var status bool
	var salary float32

	//使用 fmt.Scanln()获取
	fmt.Println("请输入姓名[woko]")
	fmt.Scanln(&name)
	fmt.Println("请输入年龄[26]")
	fmt.Scanln(&age)
	fmt.Println("请输入在职状态[true/false]")
	fmt.Scanln(&status)
	fmt.Println("请输入薪水[12000.00]")
	fmt.Scanln(&salary)
	fmt.Printf(" 姓名：%v\n 年龄：%v\n 在职状态：%v\n 薪水: %v\n", name, age, status, salary)

	//使用 fmt.Scanf() 获取
	fmt.Println("请输入姓名 年龄 在职状态 薪水 用空格分隔")
	fmt.Scanf("%s %d %t %f", &name, &age, &status, &salary)
	fmt.Printf(" 姓名：%v\n 年龄：%v\n 在职状态：%v\n 薪水: %v", name, age, status, salary)
}
