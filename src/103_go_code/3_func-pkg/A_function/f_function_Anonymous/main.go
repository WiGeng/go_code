package main

import (
	"fmt"
)

/* 匿名函数
介绍
	Go支持匿名函数（无名），如果我们某个函数数只是希望使用一次，可以考虑使用匿名函数，匿名函数也可以实现多次调用。
匿名函数使用方式1
	在定义匿名函数时就直接调用
匿名函数使用方式2
	将匿名函数赋给一个变量(函数变量)，再通过该变量来调用匿名函数
全局匿名函数
	如果将匿名函数赋给一个全局变量，那么这个匿名函数，就成为一个全局匿名函数，可以在程序有效。
*/

var (
	Fun1 = func(n1 int, n2 int) int { //Fun1就是一个全局匿名函数
		return n1 * n2
	}
)

func main() {
	//在定义匿名函数时就直接调用
	//求两个数的和，匿名函数方式实现
	res := func(n1 int, n2 int) int {
		return n1 + n2
	}(10, 20)
	fmt.Println("res=", res)

	//将匿名函数func (n1 int, n2 int) int赋给a变量
	//则a的数据类型就是函数类型，此时，我们可以通过a完成调用
	a := func(n1 int, n2 int) int {
		return n1 - n2
	}
	res2 := a(10, 30)
	fmt.Println("res2=", res2)
	fmt.Printf("a变量的类型是%T\n", a)

	//全局匿名函数的使用
	res3 := Fun1(4, 9)
	fmt.Println("res3=", res3)
}
