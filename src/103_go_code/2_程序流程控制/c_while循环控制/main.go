package main

import (
	"fmt"
)

/* while和do...while的实现说明
Go语言没有while 和do..while语法，如果需要使用类似其它语言(比如java/c 的while 和do..while),可以通过for循环来实现效果

1）for循环实现while效果
	循环变量初始化
	for {
		if循环条件表达式{
			break //跳出for循环
		}
		循环操作(语句)
		循环变量迭代
	}
2）for循环实现do...while效果
	for {
		循环操作(语句)
		循环变量迭代
		if循环条件表达式{
			break //跳出for循环
		}

	}
*/
func main() {
	//1）while效果实现
	var i int = 1
	for {
		if i > 5 {
			break
		}
		fmt.Println("hello,world", i)
		i++
	}

	//2）do...while效果实现
	var j int = 1
	for {
		fmt.Println("hello,world~", j)
		j++
		if j > 5 {
			break
		}

	}

}
