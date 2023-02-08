package main

import (
	"fmt"
	"math/rand"
)

/*练习1： 循环打印输入月份的天数[使用continue,并且要判断月份是否输入错误]*/
func Monthdays() {
	var year, month int
	for {
		fmt.Println("请输入年份和月份[year>0 && month>0 && month<13]:")
		fmt.Scanln(&year, &month)
		if (year > 0) && (month > 0) && (month < 13) {
			switch month {
			case 1, 3, 5, 7, 8, 10, 12:
				fmt.Printf("公元%v年%v月共有31天\n", year, month)
				continue
			case 2:
				fmt.Printf("公元%v年%v月共有28/29天\n", year, month)
				continue
			default:
				fmt.Printf("公元%v年%v月共有30天\n", year, month)
				continue
			}
		} else {
			fmt.Println("输入年份或月份格式不对，请重新输入：")
			continue

		}
	}
}

/* 练习2：編写一个函数：
随机猜数游戏，随机生成一个 1一100 的整数
有十次机会
如果第一次就猜中，提示 “你真是个天才”
如果第2一3次猜中，提示“你很聪明，赶上我了”
如果第4一9 次猜中，提示 “一般般”
如果最后一次猜中，提示“可算猜对啦”
一次都没猜对，提示 “说你点啥好呢”
*/
func Guessnumber() {
	target := rand.Intn(100) + 1 //生成一个介于1和100之间的整数
	fmt.Println(target)
	var reader int
	for i := 1; i < 11; i++ {
		fmt.Println("输入你要猜的数字[1~100]:")
		fmt.Scan(&reader)
		if reader == target {
			if i == 1 {
				fmt.Printf("谜底是%v,你第%v次就猜中了,真是个天才", target, i)
				break
			} else if i == 3 || i == 2 {
				fmt.Printf("谜底是%v,你第%v次就猜中了,你很聪明，赶上我了", target, i)
				break
			} else if i >= 4 || i <= 9 {
				fmt.Printf("谜底是%v,你第%v次猜中了,一般般", target, i)
				break
			} else {
				fmt.Printf("说你点啥好呢")
				break
			}
		} else {
			fmt.Println("没猜对，请重猜～")
			continue
		}
	}
}

/* 题目3：编写一个函数，输出100以内所有的素数，每行显示5个。并求和*/
func PrimenumberSum() {
	fmt.Println("1-100之间的质数为:")
	var sum int
	var count int
	// i应直接从2开始
	for i := 2; i <= 100; i++ {
		for n := 2; n <= i; n++ {
			// 当走到最后n等于i 了，则说明下面的i%n==0 && n < i 始终没有成立。说是这个数是个质数。
			if n == i {
				fmt.Printf("%d ", i)
				sum += i
				count++
				if count%5 == 0 {
					fmt.Printf("\n")
				}
			}
			// 当满足这个条件的时候就终止里面的循环，不用继续往下走了,因为它已经不是一个质数了。
			if i%n == 0 && n < i {
				break
			}
		}

	}
	fmt.Println("100以内所有的素数和=", sum)
}

/* 题目4：编写一个函数，判断是打鱼还是晒网
中国有句俗语叫“三天打鱼两天晒网”。如果从1990年1月1日起开始执行 “三天打鱼两天晒网”。如何判断在以后的某一天中是“打鱼”还是“晒网”*/

//首先分析：
//1.需要算出以后的某一天和1990.1.1相差多少天数
//2.三天打鱼两天晒网，5天一循环，所以用相差的天数除以5，如果余数是4或者是0，那么这一天就是晒网，否则就是打鱼。

func Days(year, month, day int) int {
	var count int
	for i := 1991; i <= year; i++ {
		if (i%4 == 0 && i%100 != 0) || (i%400 == 0) {
			count += 366
		} else {
			count += 365
		}
	}
	for j := 1; j <= month; j++ {
		switch month {
		case 1, 3, 5, 7, 8, 10, 12:
			count += 31
		case 2:
			if (year%4 == 0 && year%100 != 0) || (year%400 == 0) {
				count += 29
			} else {
				count += 28
			}
		default:
			count += 30
		}
	}
	count += day
	return count
}

func Judge(count int) {
	fmt.Printf("距离1990年1月1日已经%v天,", count)
	if count%5 == 4 || count%5 == 0 {
		fmt.Println("今天应该晒网。")
	} else {
		fmt.Println("今天需要打鱼。")
	}
}

/* 题目5：小小计算器
   --------------小小计算器--------------
   1. 加法
   2. 减法
   3. 乘法
   4. 除法
   5. 取余
请选择：1
 10 + 15 = 25
*/

func Calculator(n int, m int) int {
	for {
		var function int
		fmt.Println("-----小小计算机-----\n1. 加法\n2. 减法\n3. 乘法\n4. 除法\n5. 取余\n请选择：")
		fmt.Scanln(&function)
		switch function {
		case 1:
			fmt.Printf("%v + %v = %v\n", n, m, n+m)
		case 2:
			fmt.Printf("%v - %v = %v\n", n, m, n-m)
		case 3:
			fmt.Printf("%v * %v = %v\n", n, m, n*m)
		case 4:
			fmt.Printf("%v / %v = %v\n", n, m, n/m)
		case 5:
			fmt.Printf("%v 模 %v = %v\n", n, m, n%m)
		}
		continue
	}
}

/* 题目6： 输出小写的a-z及大写的Z-A */
func Alphabet() {
	for i := 97; i < 123; i++ {
		str := string(i)
		fmt.Printf(str)
	}
	fmt.Println()
	for i := 90; i > 64; i-- {
		str := string(i)
		fmt.Printf(str)
	}
}

func main() {
	// Monthdays()
	// Guessnumber()
	// PrimenumberSum()
	// Judge(Days(2022, 12, 11))
	// Calculator(10, 15)
	Alphabet()
}
