package main

import (
	"fmt"
	"strconv"
	_ "unsafe"
)

/*  介绍
1）Golang 和ava/c不同，Go 在不同类型的变量之间赋值时需要显式转换。也就是说Golang中数据类型不能自动转换。
2）基本语法
表达式T(v)将值v转换为类型T
	T：数据类型，比如 int32， int64，float32等
	v: 就是需要转换的变量
3)细节说明
	a. Go中，数据类型的转换可以是从表示范围小-›表示范围大，也可以范围大一>范围小
	b. 被转换的是变量存储的数据(即值)，变量本身的数据类型并没有变化
	c. 在转换中，比如将 int64 转成int8[-128-127]，編译时不会报错，只是转换的结果是按溢出处理
*/

func main() {
	var i int32 = 100
	var n1 float32 = float32(i)
	var n2 float64 = float64(i)
	var n3 int8 = int8(i)
	var n4 int64 = int64(i)
	fmt.Printf("i=%v,n1=%v,n2=%v,n3=%v,n4=%v\n", i, n1, n2, n3, n4)

	/* 基本数据类型和String类型的转换
	在程序开发中，我们经常需要将基本数据类型转成string类型或者将string类型转成基本数据类型。
	1. 基本类型转string类型
		方式1：fmt.Sprintf("%参数",表达式）
			参数需要和表达式的数据类型相匹配
			fmt.Sprintf()… 会返回转换后的字符串
			用法：
			func Sprintf：Sprintf根据format参数生成格式化的字符串并返回该字符串。
				func Sprintf(format string, a ...interface{}) string

		方式2：使用strconv包的函数
		func FormatBool(b bool) string
		func FormatInt(i int64, base int) string
		func FormatUint(i uint64, base int) string
		func FormatFloat(f float64, fmt byte, prec, bitSize int) string
		用法：
			func FormatBool
			func FormatBool(b bool) string
			根据b的值返回"true"或"false"。

			func FormatInt
			func FormatInt(i int64, base int) string
			返回i的base进制的字符串表示。base 必须在2到36之间，结果中会使用小写字母'a'到'z'表示大于10的数字。

			func FormatUint
			func FormatUint(i uint64, base int) string
			是FormatInt的无符号整数版本。

			func FormatFloat
			func FormatFloat(f float64, fmt byte, prec, bitSize int) string
	*/

	var num1 int = 99
	var num2 float64 = 23.456
	var b bool = true
	var mychar byte = 'h'
	var str string //空的str
	//使用第一种方式来转换 fmt.Sprintf方法
	str = fmt.Sprintf("%d", num1)
	fmt.Printf("str type %T str=%v\n", str, str)

	str = fmt.Sprintf("%f", num2)
	fmt.Printf("str type %T str=%q\n", str, str)

	str = fmt.Sprintf("%t", b)
	fmt.Printf("str type %T str=%q\n", str, str)

	str = fmt.Sprintf("%c", mychar)
	fmt.Printf("str type %T str=%q\n", str, str)

	//第二种方式 strconv函数
	var num3 int = 99
	var num4 float64 = 23.456
	var b2 bool = false

	str = strconv.FormatInt(int64(num3), 10)
	fmt.Printf("str type %T str=%q\n", str, str)

	str = strconv.FormatFloat(num4, 'f', 10, 64)
	// 说明："f" 格式, 10：表示小数位保留10位, 64：表示这个小数是f1oat64
	fmt.Printf("str type %T str=%q\n", str, str)

	str = strconv.FormatBool(b2)
	fmt.Printf("str type %T str=%q\n", str, str)

	/* 基本数据类型和string的转化
		1）注意事项：在将string 类型转成基本数据类型时，要确保String类型能够转成有效的数据，比如我们可以把“123”，转成一个整数，但是不能把“hello”转成一个整数，如果这样做，Golang直
	接将其转成0
	*/
	//转Int
	var str4 string = "hello"
	var num5 int64
	num5, _ = strconv.ParseInt(str4, 10, 64)
	fmt.Printf("num5 type %T num5= %v\n", num5, num5)
	//转Float
	var str5 string = "123.456"
	var num6 float64
	num6, _ = strconv.ParseFloat(str5, 64)
	fmt.Printf("num6 type %T num6= %v\n", num6, num6)
	//转bool
	var str6 string = "hello"
	var num7 bool
	num7, _ = strconv.ParseBool(str6)
	fmt.Printf("num7 type %T num7= %v\n", num7, num7)

}
