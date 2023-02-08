package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	/* 跳转控制语句-break
	介绍：break语句用于终止某个语句块的执行，用于中断当前for循环或者跳出switch语句
	1) 基本语法:
		{
			...
			break
			...
		}
	2）break语句出现在多层嵌套的语句块中时，可以通过标签指明是要终止那一层语句块
	*/

	/* 练习：
	随机生成一个1-100的数，当生成的数为99的时候，统计已经随机生成了多少次
	我们为了生成一个随机数，还需要 个rand设置一个种子
		time.Now().Unix(）：返回一个从1970:01:01 的0时0分0秒到现在的秒数
		rand.Seed(time.Now().UnixNano())
	随机的生成1-100整数
		n := rand. Intn (100) + 1  // [0 100）

	*/
	var count int = 0
	for {
		rand.Seed(time.Now().UnixNano()) //纳秒级
		n := rand.Intn(100) + 1
		fmt.Println("n=", n)
		count++
		if n == 99 {
			break //表示跳出for循环
		}
	}
	fmt.Printf("生成99一共随机了%v次\n", count)

	//未使用指定标签的形式
	for i := 0; i < 4; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				break //默认是跳出最近一层的for循环
			}
			fmt.Printf(" i=%v,j=%v\n", i, j)
		}
	}

	//使用指定标签的形式来使用break
label1:
	for i := 0; i < 4; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				break label1 //跳出指定的for循环
			}
			fmt.Printf("i=%v,j=%v\n", i, j)
		}
	}

	// 100以内的数求和，求出当和第一次大于20的当前数
	for {
		a := rand.Intn(100) + 1
		b := rand.Intn(100) + 1
		if a+b > 20 {
			fmt.Println(a, b)
			break
		}
	}

	//实现登录验证，有三次机会，如果用户名为"张无忌"，密码"888"提示登录成功，否则提示还有几次机会.
	var name string = "张无忌"
	var passwd string = "123456"
	for count := 3; count > 0; count-- {
		if name == "张无忌" {
			if passwd == "123456" {
				fmt.Println("配对成功~")
				break
			}
		} else {
			fmt.Printf("还有%v次配对机会\n", count)
		}
	}

	/* 跳转控制语句-continue
	1) continue语句用于结束本次循环，继续执行下一次循环。
	2) continue语句出现在多层嵌套的循环语句体中时，可以通过标签指明要跳过的是哪一层循环，这个和前面的标签的使用的规则一样。
	3) 基本语法：
		{
			...
			continue
			...
		}

	*/
	//continue的使用演示
	for i := 0; i < 4; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				continue //默认是跳出最近一层的for循环
			}
			fmt.Printf(" i=%v,j=%v\n", i, j)
		}
	}

	//实现1-100之内的奇数[for+continue]
	for i := 1; i < 101; i++ {
		if i%2 != 0 {
			fmt.Printf("%v ", i)
			continue
		}
	}

	//从键盘读取一个整数，判断整数的正负，输入为0时结束程序[for、break、continue]
	for {
		var i int
		fmt.Println("请输入一个整数：")
		fmt.Scanln(&i)
		if i > 0 {
			fmt.Printf("%v为positive\n", i)
			continue
		} else if i < 0 {
			fmt.Printf("%v为negative\n", i)
			continue
		} else {
			break
		}
	}

	/* goto
	1）Go语言的goto语句可以无条件地转移到程序中指定的行。
	2）goto语句通常与条件语句配合使用。可用来实现条件转移，跳出循环体等功能。
	3）在Go程序设计中一般不主张使用goto语句，以免造成程序流程的混乱，使理解和调试程序都产生困难
	4）基本语法
		goto label
		...
		label statement
	*/

	// goto 代码演示
	var n int = 10
	fmt.Print("我")
	if n < 20 {
		goto label
	}
	fmt.Print("是")
	fmt.Print("魏更")
	fmt.Print(",")
label:
	fmt.Print("是")
	fmt.Print("中")
	fmt.Print("国")
	fmt.Print("人")

	/* return
	介绍：
		return 使用在方法或者函数中，表示跳出所在的方法或函数
	说明：
		1）如果return是在普通的函数，则表示跳出该函数，即不再执行函数中 return后面代码，也可以理解成终止函数。
		2）如果return是在main函数，表示终止main函数，也就是说终止程序。
	*/

	// return 代码演示
	var m int = 10
	fmt.Print("我")
	if m < 20 {
		return
	}
	fmt.Print("是")
	fmt.Print("魏更")
	fmt.Print(",")
	fmt.Print("是")
	fmt.Print("中")
	fmt.Print("国")
	fmt.Print("人")
}
