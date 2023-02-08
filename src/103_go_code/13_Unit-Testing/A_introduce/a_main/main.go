package main

import "fmt"

/* 单元测试
基本介绍：
	GO语言中自需有一个轻量级的测试框架testing和自带的go test命令来实现单元测试和性能测试，testing框架和其他语言中的测试框架
	类似，可以基于这个框架写针对相应函数的测试用例，也可以基于该框架写相应的压力测试用例。通过单元测试，可以解决如下问题：
		1）确保每个函数是可运行，并且运行结果是正确的
		2）确保写出来的代码性能是好的
		3）单元测试能及时的发现程序设计或实现的逻辑错误，使问题及早暴露，便于问题的定位解决，而性能测试的重点在于发现程序设计
		  上的一些问题，让程序能够在高并发的情况下还能保持稳定
*/

/* 使用go的单元测试，对addUpper和Sub函数进行测试 */

//一个被测试函数
func addUpper(n int) int {
	res := 0
	for i := 1; i <= n-1; i++ {
		res += i
	}
	return res
}

func addUpper2(n int) int {
	res := 0
	for i := 1; i <= n-1; i++ {
		res += i
	}
	return res
}

func main() {

	// 传统的测试方法，就是在main函数中使用看看结果是否正确
	res := addUpper(10) // 1.+ 10 = 55
	if res != 55 {
		fmt.Printf("addUpper错误 返回值=%v 期望值=%v\n", res, 55)
	} else {
		fmt.Printf("addUpper正确 返回值=%v 期望值=%v\n", res, 55)
	}
}
