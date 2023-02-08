package main

import (
	"fmt"
	_ "strconv"
	_ "unsafe"
)

/* 指针基本介绍
1） 基本数据类型，变量存的就是值，也叫值类型
2） 获取变量的地址，用&，比如： var num int,获取num的地址：&num
3） 指针类型，指针变量存的是一个地址，这个地址指向的空间存的才是值。比如：var ptr *int = &num
4） 获取指针类型所指向的值，使用：*，比如：var ptr *int,使用*ptr获取ptr指向的值
*/
func main() {
	//基本数据类型在内存中的布局
	var i int32 = 10
	//取i的地址，&i
	fmt.Println("i的指针地址是：", &i)

	var ptr *int32 = &i
	//1. ptr是一个指针变量
	//2. Ptr的类型是*int
	//3. ptr的值是&i
	fmt.Printf("ptr的值是：%v\n", ptr)
	fmt.Println("ptr的地址是：", &ptr)
	fmt.Println("ptr指向的值是：", *ptr)
	//通过ptr修改i的值
	*ptr = 100
	fmt.Println("i修改后的值是：", i)

	/* 指针细节说明
	1) 值类型：都有对应的指针类型,形式为*数据类型,比如int的对应的指针就是*int,tloat32对应的指针类型就是*float32,依次类推。
	2) 值类型包括：基本数据类型 int系列，float系列, bool, string, 数组和结构体struct
	*/

	/* 值类型和引用类型
		1） 值类型：基本数据类型int系列, float系列, bool, string、数组和结构体struct
		2） 引用类型：指针、slice切片、 map、 管道chan、interface 等都是引用类型

	   值类型和引用类型的使用特点
		1）值类型，变量直接存储值，内存通常在栈中分配
		2）引用类型：变量存储的是一个地址，这个地址对应的空间才真正存储数据(值)。内存通常在堆上分配，当没有任何变量引用这个
		地址时，该地址对应的数据空间就成为一个垃圾，由GC来回收
	*/

}
