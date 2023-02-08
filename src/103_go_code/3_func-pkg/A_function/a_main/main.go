package main

import (
	utils "awesomeProject/src/103_go_code/3_func-pkg/A_function/b_utils"
	"fmt"
)

/* 函数
介绍：
	为完成某一功能的程序指令（语句）的集合，成为函数，有自定义函数和系统函数之分
1. 基本语法
	func 函数名（形参列表）（返回值类型列表）{
		执行语句…
		return 返回值列表
	}
	1） 形参列表：表示函数的输入
	2） 函数中的语句：表示为了实现某一功能代码块·
	3) 函数可以有返回值,也可以没有
2. 函数-调用机制
	1）为了让大家更好的理解西数调用过程，看两个案例，并西出示意图，这个很重要
			a) 传入一个数+1 test
			b）计算两个数,并返回 getSum
	2）return句
	＞基本语法
		Go函数支持返回多个值，这一点是其它编程语言没有的。
		func 函数名（形表列表）（返回值类型列表）{
				语句…
				return 返回値列表
		}
	>  如果返回多个值时，在接收时，希望忽略某个返回值，则使用_符号表示占位忽略
	>  如果返回値只有一个，（返回值类型列表）可以不写()
3. 函数-递归调用
	基本介绍： 一个函数在函数体内有调用了本身，称之为递归调用
	函数递归需要遵守的重要原则：
	1）执行一个函数时，就创建一个新的受保护的独立空间(新函数栈)
	2）函数的局部变量是独立的，不会相互影响
	3）递归必须向退出递归的条件逼近，否则就是无限递归
	4）当一个函数执行完毕，或者遇到return，就会返回；遵守谁调用，就将结果返回给谁；同时当函数执行完毕或者返回时，该函数本身也会被销毁

4. 函数注意事项和细节讨论
	1） 函数的形参列表可以是多个，返回值列表也可以是多个
	2） 形参列表和返回值列表的数据类型可以是值类型和引用类型
	3） 函数的命名遵循标识符命名规范，首字母不能是数宇，首字母大写该函数可以被本包文件和其它包文件使用，类似public，首字母小写，只能被本包文件使用，其它包文件不能使用，类似private
	4） 函数中的变量是局部的，函数外不生效
	5） 基本数据类型和数组默认都是值传递的，即进行值拷贝。在函数内修改，不会影响到原来的值
	6） 如果希望函数内的变量能修改函数外的变量，可以传入变量的地址&，函数内以指针的方式操作变量。从效果上看类似引用
	7） Go函数不支持重载
	8)  在Go中，函数也是一种数据类型，可以赋值给一个变量，则该变量就是一个函数类型的变量了，通过该变量可以对函数调用。
	9） 函数既然是一种数据类型，因此在Go中，函数可以作为形参来调用
	10）为了简化数据类型定义，Go支持自定义数据类型
		基本语法，type 自定义数据类型名 数据类型 //理解：相当于一个别名
		案例：type mylnt int 	//这时myint 就等价int 来使用了。
		案例：type mySum func(int,int) int  	// 这时mySum 就等价一个函数类型func(int,int) int
	11) 支持对函数返回值命名
	12）在go中支持可变参数
		a. 支持0到多个参数
			func sum(args... int) sum int {
			}
		b. 支持1到多个参数
			func sum(n1 int, args... int) sum int {
			}
		说明：args是slice，通过args[indesx]可以访问到各个值。
*/

//快速入门

//在Go中，函数也是一种数据类型
func getSum(n1 int, n2 int) int {
	return n1 + n2
}

func main() {
	var n1 float64 = 12.43
	var n2 float64 = 18.68
	var operator byte = '-'
	result := utils.Cal(n1, n2, operator)
	fmt.Println("result=", result)

	//输出函数的数据类型
	a := getSum
	fmt.Printf("a的数据类型是%T,getSum的数据类型是%T\n", a, getSum)
	res := a(10, 20)
	fmt.Println("res=", res)
}
