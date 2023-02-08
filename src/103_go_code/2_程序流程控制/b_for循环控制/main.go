package main

import (
	"fmt"
)

/*
	for循环基本语法
		for 循环变量初始化;循环条件;循环变量迭代{
			循环体
		}
		for循环四要素：循环变量初始化、循环条件、循环变量迭代、循环体

	for循环的注意事项和细节说明
		1）循环条件是返回一个布尔值的表达式
		2）for循环的第二种使用方式
				for 循环判断条件{
					//循环执行语句
				}
			将变量初始化和变量迭代写到其它位置
		3）for循环的第三种使用方式
				for {
					//循环执行语句
				}
			上面的写法等价 for;;{} 是一个无限循环，通常需要配合break语句使用。
		4) Golang提供for-range的方式可以方便遍历字符串和数组
*/

func main() {
	// for循环的第一种写法
	for i := 1; i < 11; i++ {
		fmt.Println("你好，世界", i)
	}

	// for循环的第二种写法
	j := 1       //变量初始化
	for j < 11 { //循环条件控制
		fmt.Println("你好，世界！", j)
		j++ //循环变量迭代
	}

	// for循环的第三种写法(死循环)，通常是要配合break来跳出死循环
	k := 1
	for { //这里等价于 for ; ;{
		if k < 11 {
			fmt.Println("你好，世界～", k)
		} else {
			break
		}
		k++
	}

	//for循环遍历字符串的两种方式
	//方式1-传统方式
	var str string = "hello,world! hello 蜗壳"
	str2 := []rune(str)              //string类型转成切片
	for i := 0; i < len(str2); i++ { //按照字节来遍历，string中出现中文的话会报错，因为汉字在utf-8编码中占3个字节，需要转成切片
		fmt.Printf("%c", str2[i])
	}
	//方式2-for-range
	str = "weigeng,你好"
	for index, val := range str { //将str中的字符逐一遍历出来存放到val中
		fmt.Printf("index=%d,val=%c\n", index, val)
	}

	//练习～
	//打印1~100之间所有的7的倍数的整数个数及总和
	fmt.Println("------------------------------")
	var count int = 0
	var sum int = 0
	for i := 1; i < 101; i++ {
		if i%7 == 0 {
			count++
			sum += i
		}
	}
	fmt.Printf("count=%v,sum=%v\n", count, sum)

	/*完成下列表达式的输出
	0 + 6 = 6
	1 + 5 = 6
	2 + 4 = 6
	3 + 3 = 6
	4 + 2 = 6
	5 + 1 = 6
	6 + 0 = 6
	*/
	fmt.Println("------------------------------")
	var m int = 6
	for i := 0; i < 7; i++ {
		fmt.Printf("%v + %v = %v\n", i, m-i, m)
	}

}
