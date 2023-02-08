package main

import (
	"fmt"
)

/*------------------------（1）函数调用机制案例-----------------------------------*/
func test(n int) {
	n = n + 1
	fmt.Println("n的值最终为", n)
}

func getSum(n1 int, n2 int) int {
	sum := n1 + n2
	return sum
}

func getSumSub(n1 int, n2 int) (int, int) {
	sum := n1 + n2
	sub := n1 - n2
	return sum, sub
}

/*------------------------（1）函数递归调用案例-----------------------------------*/
func recursion1(n int) {
	if n > 2 {
		n--
		recursion1(n)
	}
	fmt.Println("n= ", n)
}
func recursion2(n int) {
	if n > 2 {
		n--
		recursion2(n)
	} else {
		fmt.Println("n= ", n)
	}
}

//斐波那契数列求值
/* 题1：斐波那契数,请使用递归的方式,求出斐波那契数1,1,2,3,5,8,13..  给你一个整数n，求出它的斐波那契数是多少？*/
func fbn(n int) int {
	if n == 1 || n == 2 {
		return 1
	} else {
		return fbn(n-1) + fbn(n-2)
	}
}

//题2：求函数值  已知f(1)=3; f(n)=2*f(n-1)+1;  	请使用递归的思想编程，求出f(n)的值？
func f(n int) int {
	if n == 1 {
		return 3
	} else {
		return 2*f(n-1) + 1
	}
}

/* 题3:有一堆桃子，猴子第一天吃了其中的一半，并在多吃了一个。
以后猴子每天都吃其中的一半，然后再多吃一个，当到第10天时，发现只有一个桃子了。
问题：最初共有多少个桃子*/
//思路：
// 第10天只有一个桃子；
// 第9天的桃子数 = （第10天桃子数+1）* 2 ；
// 规律：第n天的桃子数：peach (n) = (peach(n+1) + 1) * 2

func peach(d int) int {
	if d > 10 || d < 1 {
		fmt.Println("输入的天数不对")
		return 0
	} else if d == 10 {
		return 1
	} else {
		return (peach(d+1) + 1) * 2
	}
}

/*------------------------函数练习案例-----------------------------------*/

//题目4:请编写一个函数 swap(n1 *int, n2 *int) 可以交换 n1 和 n2的值
func swap(n1 *int, n2 *int) {
	//定义一个临时变量
	t := *n1
	*n1 = *n2
	*n2 = t
}

func main() {
	//函数调用机制
	n := 10
	test(n)
	fmt.Println("n的初始值是", n)

	sum := getSum(10, 20) //这块的“getSum(10, 20)”相当于“return sum”中的“sum”
	fmt.Println("main中sum为:", sum)

	sum, sub := getSumSub(30, 20)
	fmt.Printf("main中Sub为%v,main中Sum为%v\n", sub, sum)

	res1, _ := getSumSub(30, 20) //使用_ 表示占位符，忽略其中的返回值
	fmt.Printf("main中Sub为%v\n", res1)

	//函数递归调用
	recursion1(4)
	recursion2(4)

	//斐波纳切递归调用测试
	fmt.Println("res=", fbn(3))
	fmt.Println("res=", fbn(4))
	fmt.Println("res=", fbn(5))

	//求函数值测试
	fmt.Println("res=", f(1))
	fmt.Println("res=", f(3))
	fmt.Println("res=", f(5))

	//猴子吃桃测试
	var d int = 1
	fmt.Printf("第%v天的桃子数量是%v\n", d, peach(d))

	//n1、n2值替换
	a := 10
	b := 20
	swap(&a, &b) //传入的地址
	fmt.Printf("a=%v, b=%v", a, b)
}
