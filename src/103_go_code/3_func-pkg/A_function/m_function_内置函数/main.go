package main

import (
	"fmt"
)

/* 内置函数
介绍：
	golang设计者为了编程方便，提供了一些函数，这些函数可以直接使用，我们称为Go的内置函数
	1) len：用来求长度，比如string、array、slice、map、channel
	2) new：用来分配内存，主要用来分配值类型，比如int、float32、struct..返回的是指针
	3) make：用来分配内存，主要用来分配引用类型，比如chan、map、slice。
*/

func main() {

	num1 := 100
	fmt.Printf("num1的类型%T , num1的值=%v , num1的地址%v\n", num1, num1, &num1)

	num2 := new(int) // *int
	//num2的类型%T => *int
	//num2的值 = 地址 0xc0000b2020 （这个地址是系统分配）
	//num2的地址%v = 地址 0xc0000ac020 	 (这个地址是系统分配)
	//num2指向的值 = 100
	*num2 = 100
	fmt.Printf("num2的类型%T , num2的值=%v , num2的地址=%v\nnum2这个指针，指向的值=%v",
		num2, num2, &num2, *num2)
}
