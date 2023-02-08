package main

import (
	"fmt"
)

/* const说明
介绍：
	＞ 常量使用const 修改
	＞ 常量在定义的时候，必须初始化
	＞ 常量不能修改
	＞ 常量只能修饰bool 数值类型(int, float系列）、string类型
	＞ 语法：const identifier ltype] = value
注意细节：
	1）Golang 中没有常量名必须字母大写的规范 比如 TAX_RATE
	2）仍然通过首字母的大小写来控制常量的访问范围
*/

func main() {

	var num int       //变量声明，可以不赋初值
	num = 9           //变量可修改
	const tax int = 0 //常量声明的时候，必须赋值
	//tax = 10        //常量是不能修改
	fmt.Println(num, tax)

	const name = "tom"
	const score float64 = 0.8
	const div = 9 / 3

	// 简洁写法
	const (
		x = 1
		y = 2
	)
	fmt.Println(x, y)

	// 专业写法
	const (
		a = iota //标识给a赋0,b在a的基础上+1,c在b的基础上+1,d在c的基础上+1
		b
		c
		d
	)
	fmt.Println(a, b, c, d)
}
