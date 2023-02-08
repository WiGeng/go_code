package main

import (
	"awesomeProject/src/103_go_code/3_func-pkg/A_function/e_function_init/utils"
	"fmt"
)

/* init函数：
基本介绍： 每一个源文件都可以包含一个init函数，该函数会在mian函数执行前，被go运行框架调用，也就是说会在main函数前被调用
init函数注意事项和使用细节：
	1）如果一个文件同时包含全局变量定义，init函数和main函数，则执行的流程全局变量定义 init函数 -> main函数
	2）init的主要工作就是完成一些初始化操作
*/

var Age = age() //全局变量先被初始化

func age() int {
	fmt.Println("age()...")
	return 90
}

func init() {
	fmt.Println("init~")
}
func main() {
	fmt.Println("main()...Age=", Age)
	fmt.Printf("Car是%v,Color为%v", utils.Car, utils.Color)
}
