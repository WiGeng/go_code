package main

import (
	"fmt"
	"strings"
)

/* 闭包
介绍：闭包就是一个函数与其相关引用环境组合的一个整体（实体）
	1）闭包是类，函数是操作，n是字段。函数和它使用到n构成闭包
	2）当我们反复的调用f函数时，因为n只初始化一次，因此每调用一次就进行累计
	5）要搞清楚闭包的关键，就是要分析出返回的函数它使用(引用)到哪些变量，因为函数和它引用到的变量共同构成闭包
*/

//累加器
func AddUpper() func(int) int { //AddUpper是一个函数，返回的类型是func(int) int
	var n int = 18
	return func(x int) int { //返回的是一个匿名函数，但是这个匿名函数引用到外部的n，因此这个匿名函数和n就组成了一个整体叫闭包
		n = n + x
		return n
	}
}

/*闭包的最佳实践
请编写一个程序，具体要求如下
1）编写一个函数makesuffix(suffix string）可以接收一个文件后缀名(比如.jpg)，并返回一个闭包
2）调用闭包，可以传入一个文件名，如果该文件名没有指定的后缀(比如.jpg），则返回 文件名：jpg，如果己经有.jpg后缀，则返回原文件名。
3）要求使用闭包的方式完成
4) strings.HasSuffix [判断是否有后缀字符串]
*/

func makeSuffix(suffix string) func(string) string {

	return func(name string) string {
		//如果 name 没有指定后缀，则加上，否则就返回原来的名字
		if !strings.HasSuffix(name, suffix) { //判断字符串name是否以suffix结尾
			return name + suffix
		}
		return name
	}
}

func makeSuffix2(suffix string, name string) string {

	//如果 name 没有指定后缀，则加上，否则就返回原来的名字
	if !strings.HasSuffix(name, suffix) {
		return name + suffix
	}

	return name

}

func main() {
	//使用前面的代码
	f := AddUpper()
	fmt.Println(f(1)) // 19
	fmt.Println(f(2)) // 21
	fmt.Println(f(3)) // 24

	//测试makeSuffix的使用
	//返回一个闭包
	f2 := makeSuffix(".jpg")               //如果使用闭包完成，好处是只需要传入一次后缀。
	fmt.Println("文件名处理后=", f2("winter"))   // winter.jgp
	fmt.Println("文件名处理后=", f2("bird.jpg")) // bird.jpg

	fmt.Println("文件名处理后=", makeSuffix2("jpg", "winter"))   // winter.jgp
	fmt.Println("文件名处理后=", makeSuffix2("jpg", "bird.jpg")) // bird.jpg
}
