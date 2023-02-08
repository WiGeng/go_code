package main

import (
	"fmt"
)

/*	运算符的基本介绍
	运算符是一种特殊的符号，用以表示数据的运算、赋值和比较等运算符是一种特殊的符号，用以表示数据的运算、赋值和比较等
		1）算术运算符
		2）赋值运算符
		3）关系运算符
		4）逻辑运算符
		5）位运算符
		6) 其它运算符
*/

func main() {

	/* 算数运算符：针对数值类型的变量进行运算，比如：加减乘除
	细节说明
	1） 对手除号""，它的整数除和小数除是有区别的：整数之间做除法时，只保留整数部分而舍弃小数部分。例如：x:=19/5,结果是3
	2） 当对一个数取模时，可以等价 a%a=a-a/b*b ，这样我们可以看到 取模的一个本质运算
	3)  Golang的自增自诚只能当做一个独立语言使用时，不能这样使用 b := a++ 或者 b := a--
	4)  Golang 的++和--只能写在变量的后面，不能写在变量的前面，即。只有a++,a--没有++a,--a
	5)  Golang的设计者去掉C/java中的 自增自减 容易混淆的写法，让Golang更加简洁
	*/
	var num1 float32 = 10 / 4 //结果为2，保留整数部分，去除小数部分
	fmt.Println(" 10 / 4 的结果是:", num1)
	var num2 float32 = 10.0 / 4 //结果为2.5，如果需要保留小数部分，则需要有浮点数参与运算
	fmt.Println(" 10.0 / 4 的结果是:", num2)

	//取模运算(运算公式：a % b = a - a / b * b)
	fmt.Println("10%3的结果是:", 10%3)
	fmt.Println("-10%3的结果是:", -10%3)
	fmt.Println("10%-3的结果是:", 10%-3)

	// ++ | --
	var i int = 10
	i++ //等价于 i=i+1
	fmt.Println("i =", i)
	i-- //等价于 i=i-1
	fmt.Println("i =", i)

	var a int = 100
	var b int
	a++
	b = a
	fmt.Println("b =", b)

	/* 关系运算符
	关系运算符的结果都是 bool型，也就是要么是true，要么是false
	关系表达式经常用在if结构的条件中或循环结构的条件中

	   细节说明
		1）关系运算符的结果都是bool型，也就是要么是true，要么是false。
		2）关系运算符組成的表达式，我们称为关系表达式：a>b
		3）比较运算符"=="不能误写成"="
	*/
	var n1 int = 9
	var n2 int = 8
	fmt.Println(n1 == n2) //false
	fmt.Println(n1 != n2) //true
	fmt.Println(n1 > n2)  //true
	fmt.Println(n1 >= n2) //true
	fmt.Println(n1 < n2)  //flase
	fmt.Println(n1 <= n2) //flase
	flag := n1 > n2
	fmt.Println("flag =", flag)

	/* 逻辑运算符（&& 、|| 、！）
	用于连接多个条件（一来讲是关系表达式），最终的结果也是一个bool值
	*/
	// 逻辑&&演示
	var age int = 40
	if age > 30 && age < 50 {
		fmt.Println("ok1") //为真，输出
	}

	if age > 30 && age < 40 {
		fmt.Println("ok2") //为假，不输出
	}
	//逻辑!演示
	var age1 int = 50
	if age1 > 40 {
		fmt.Println("ok3—1")
	}
	if !(age1 > 40) {
		fmt.Println("ok3-2")
	}

	/* 赋值运算符（ =、+= 、 -= 、*=  、/= 、%= ）
	赋值运算就是将某个运算后的值，赋值指定变量
	*/
	//有两个变量，a和b，要求将其进行交换，最终打印结果
	// x=9,y=2 ==> x=2,y=9
	x := 9
	y := 2
	fmt.Printf("交换前的情况是x = %v, y = %v\n", x, y)
	//定义一个临时变量
	w := x
	x = y
	y = w
	fmt.Printf("交换后的情况是x = %v, y = %v\n", x, y)

	//复合赋值运算
	x += 99
	fmt.Println("x的最新值 =", x)

	/*移位运算
	>> <<右格和左移，运算规则：
	右移运算符 >>：低位溢出,符号位不变 并用符号位补溢出的高位
	左移运算符 <<：符号位不变,低位补0
	*/
	h1 := 1 >> 2 // 0000 0001 => 0000 0000 = 0
	h2 := 1 << 2 // 0000 0001 => 0000 0100 = 2
	fmt.Printf("h1 = %v,h2 = %v\n", h1, h2)

	/* 面试题：
	有两个变量n和m,要求对变量的值进行交换，但不能使用中间变量，并答应最终结果
	*/
	n := 10
	m := 1
	fmt.Printf("结果置换前n = %v,n = %v\n", n, m)
	n = n + m
	m = n - m
	n = n - m
	fmt.Printf("结果置换后n = %v,n = %v\n", n, m)

	/*运算符优先级
	1)运算符有不同的优先級，所谓优先級就是表达式运算中的运算顺序。如右表，上一行运算符总优先于下一行。
	2)只有单目运算符、赋值运算符是从右向中左运算的。
	*/

	var m1 int = 66
	var m2 int = 88
	var max int
	if m1 > m2 {
		max = m1
	} else {
		max = m2
	}
	fmt.Println("max的值是：", max)
}
