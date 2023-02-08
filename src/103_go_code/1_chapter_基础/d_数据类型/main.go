package main

import (
	"fmt"
	"unsafe"
)

/*
Go 的基本类型有
	1) bool
	2) string
	3) int  int8  int16  int32  int64
	4) uint uint8 uint16 uint32 uint64 uintptr
	5) byte // uint8 的别名
	6) rune // int32 的别名,表示一个 Unicode 码点
	7) float32 float64
	8) complex64 complex128
*/

func main() {

	/* 整型的使用细节：
	1）Golang各整数类型分：有符号和无符号，int uint 的大小和系统有关
	2) Golang的整型默认声明为int型
	3) Golang程序中整型变量在使用时，遵守保小不保大的原则，即：在保证程序正确运行下，尽量使用占用空间小的数据类型。
	*/
	//查看在程序查看某个变量的字节大小和数据类型
	var n = 100
	fmt.Printf("变量n的字符类型是%T,并且变量n占用的字节数是%d", n, unsafe.Sizeof(n))

	/* 浮点类型的分类：
	1）单精度float32, 4字节
	2）双精度float64, 8字节
	3）关于浮点数在机器中存放形式的简单说明：
		a. 浮点数 = 符号位 + 指数位 + 尾数位
		b. 尾数部分可能丢失，造成精度损失
	4）浮点型使用细节
		a. Golang 浮点类型有固定的苑围和字段长度，不受具体os（的影响。
		b. Golang 的浮点型默认声明为float64 类型。
		c. 浮点型常量有两神表示形式
			十进制数形式：如：5.12  .512（必須有小数点）
			科学汁数法形式：如：5.1234e2 = 5.12*10的2次方 5.12E-2=5.12/10的2次方
		d. 通常情况下，应该使用 float64，因为它比float32 更精确。
	*/
	var num1 = 3.14159
	fmt.Printf("\n变量num2的字符类型是 %T", num1)
	var num2 float32 = 0.618
	fmt.Println("\nnum1=", num2)
	//科学计数说明
	var num3 = .123e2
	var num4 = 0.123e2
	fmt.Println("num3的大小是", num3, "\tnum4的大小是", num4)

	/* 字符类型基本介绍:
	1) goolang中没有专门的字符类型，如果要存储单个字符(字母)，一般使用byte来保存。
	2) 字符串就是一串固定长度的字符连接起来的字符序列。Go的字符串是由单个字节连接起来的。
	3）对下面代码说明
		a. 如果我们保存的字符在ASCII表的.比如[0-1.2-z.A-乙.]直接可以保存到byte
		b. 如果我们保存的字符对应码值大于255.这时我们可以考虑使用int类型保存
		c. 如果我们需要安装字符的方式输出，这时我们需要格式化输出，即 fmt.Printf("...=%c"）

	字符类型使用细节
	1） 字符常量是用单引号("括起来的单个宇符。例如：var c1 byte ='a' var c2 int = "中' var c3 byte = '9'
	2） Go中允许使用转义字符1来将其后的字符转变为特殊字符型常量。例如：var c3 char = 1n' //^n'表示換行符
	3） Go语言的字符类型使用UTF-8编码【英文宇母-1个字节 汉宇-3个字节】
	4） 在Go中，宇符的本质是一个整数，直接输出时，是该字符对应的UTF-8编码的码值。
	5)  字符类型是可以进行运算的，相当于一个整数，因为它都对应有Unicode码。
	*/
	var c1 byte = 'a'
	var c2 byte = '0'
	//当直接输出byte的值，就是输出了对应字符的ASCII码
	fmt.Println("str1=", c1, "str2=", c2)
	//如果需要输出对应的字符，就需要进行格式化输出
	fmt.Printf("str1=%c str2=%c", c1, c2)

	/* 布尔类型基本介绍
	1） 布尔类型也叫bool类型，bool类型数据只允许取值true和false
	2)  bool类型占1个字节。
	3)  bool类型适于逻辑运算，一般用于程序流程控制
		if;
		for；
	*/
	var b = true
	fmt.Println("\nb =", b)
	fmt.Println("b的大小是", unsafe.Sizeof(b), "字节")

	/* 字符串类型-string
	介绍：字符串就是一串固定长度的字符连接起来的字符序列。Go的字符串是由单个字节连接起来的，Go语言的字符串使用的UTF-8编码
	字符串类型使用细节：
	1）字符串一旦被赋值，就不能修改了
	2）字符串的两种表示形式
		a. 双引号，会识别转义宇符
		b. 反引号，以字符串的原生形式输出，包括换行和特殊字符，可以实现防止攻击、输出源代码等效果
	4) 字符串的拼接 “+”
	*/

	var address string = "长安区 魏寨街道 said hello to me in 1996."
	fmt.Println(address)

	//字符串拼接[分行写时，拼接的“+”需要放在上一行]
	firstname := "wei" + "geng" +
		" =" + "= "
	firstname += "weigeng"
	fmt.Println(firstname)

	/* 基本数据类型的默认值

	数据类型		 默认值
	int				 0
	浮点型			  0
	字符串			  ""
	bool		    false
	*/
}
