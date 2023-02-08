package main

import "fmt" //提供格式化，输入输出的函数

/*
🚩🚩Golang常用的转义字符(escape char)
1) \t :一个制表位,实现对齐的功能
2) \n :换行符
3) \\ :一个\
4) \" :一个"
5) \r :一个回车fmt.Println("天龙八部雪山飞狐\r张飞");
*/

func main() {
	//演示一个制表位,实现对齐的功能
	fmt.Println("woke\twill succeed")
	//演示换行符
	fmt.Println("woke\nwill succeed")
	//演示字符串中添加\
	fmt.Println("woke\\will\\succeed")
	//演示字符串中添加"
	fmt.Println("我是 \"woke\"")
	//演示回车，从当前行的最前面开始输出，覆盖之前的内容
	fmt.Println("蜗壳二零二二元气满满\r魏更")
}
