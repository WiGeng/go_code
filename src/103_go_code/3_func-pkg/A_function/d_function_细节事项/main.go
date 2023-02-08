package main

import (
	"fmt"
)

//在Go中，函数也是一种数据类型，可以赋值给一个变量，则该变量就是一个函数类型的变量了，通过该变量可以对函数调用
func getSum(n1 int, n2 int) int {
	return n1 + n2
}

//函数既然是一种数据类型，因此在Go中，函数可以作为形参来调用
func myFun1(funvar func(int, int) int, num1 int, num2 int) int {
	return funvar(num1, num2)
}

//Go支持自定义函数别名，针对函数
type myTypeFun func(int, int) int

func myFun2(funvar myTypeFun, num1 int, num2 int) int {
	return funvar(num1, num2)
}

//支持对函数返回值命名
func cal(n1 int, n2 int) (sum int, sub int) {
	sum = n1 + n2
	sub = n1 - n2
	return
}

//go支持可变参数案例演示：編写一个函数sum，可以求出 1到多个int的和
func chageSum(n1 int, args ...int) int { //‼️可变参数args需要放在形参列表的最后一个
	sum := n1
	for i := 0; i < len(args); i++ {
		sum += args[i]
	}
	return sum
}

func main() {
	//输出函数的数据类型
	a := getSum
	fmt.Printf("a的数据类型是%T,getSum的数据类型是%T\n", a, getSum)
	res1 := a(10, 20)
	fmt.Println("res1=", res1)

	//将函数作为形参
	res2 := myFun1(getSum, 50, 60)
	fmt.Println("res2=", res2)

	//为了简化数据类型定义，Go支持自定义数据类型
	//基本数据类型
	type myInt int // 在go中，myInt和int虽然都是int类型，但是go默认这两个是两种不同的类型
	var num myInt = 40
	var num3 int = 20
	num4 := int(num)
	fmt.Printf("num的数据类型是%T,num的值是%v\n", num, num)
	fmt.Printf("num的数据类型是%T,num的值是%v\n", num3, num3)
	fmt.Printf("num的数据类型是%T,num的值是%v\n", num4, num4)

	//自定义函数别名，针对函数
	res3 := myFun2(getSum, 500, 600)
	fmt.Println("res3=", res3)

	//支持对函数返回值命名
	sum, sub := cal(30, 50)
	fmt.Printf("a1= %v,b1= %v\n", sum, sub)

	//go支持可变参数测试
	res5 := chageSum(10, 0, -1, -5, 96)
	fmt.Println("res5=", res5)

}
