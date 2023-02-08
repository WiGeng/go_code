package main

import "fmt"

/*
数组概念：数组可以存放多个同一类型的数据。数组也是一种数据类型，在go中，数组是值类型
数组定义：
	var 数组名 [数组大小]数据类型
	var a [5]int
	赋初始值：a[0] = 11  a[1] = 20 ...

	内存排布：
	1. 数组的地址可以通过数组名来获取&intArr
	2. 数组中第一个元素的地址就是数组的首地址
	3. 数组中各个元素的地址间隔是根据定义数组时数组的值类型决定的

	数组的使用：
	访问数组元素：
		数组名[下标]

数组定义的四种方式：
	方式1:
	var numArr01 [3]int = [3]int{1, 2, 3}
	方式2:
	var numArr02 = [3]int{5, 6, 7}
	方式3:
	var numArr03 = [...]int{8, 9, 10}
	方式4:
	var numArr04 = [...]int{1: 800, 0: 900, 2: 999}

数组的遍历：
	方式1-常规遍历：
		前面for循环的方式
	方式2-for-range结构遍历
		这是Go语言一种独有的结构，可以用来遍历访问数组的元素。
	基本语法
			for index, value := range array01 {
				...
			}
		说明：
		1)  第一个返回值 index 是数组的下标
		2)  第二个value是在该下标位置的值
		3)  他们都是仅在for循环内部可见的局部变量
		4)  遍历数组元素的时候，如果不想使用下标index，可以直接把下标index标为下划线 _
		5)  index和value的名称不是固定的，即程应员可以自行指定，一般命名为 index 和 value

悦明
*/

func main() {
	//使用数组的方式来解决问题

	//1.定义一个数组
	var hens [6]float64
	//2.给数组的每个元素赋值， 元素的下标是从0开始的  0-5
	hens[0] = 3.0 //hens数组的第1个元素 hens[0]
	hens[1] = 5.0 //hens数组的第2个元素 hens[1]
	hens[2] = 1.0
	hens[3] = 3.4
	hens[4] = 2.0
	hens[5] = 50.0
	//3.遍历数组求出总体重
	totalweight := 0.0
	for i := 0; i < len(hens); i++ {
		totalweight += hens[i]
	}
	//4.求出平均体重
	avfwight := fmt.Sprintf("%.2f", totalweight/float64(len(hens)))
	fmt.Printf("所有鸡一共的重量是%vKG\n", totalweight)
	fmt.Printf("每一只鸡的平均重量为%v", avfwight)

	/*数组在内存中的排布*/
	var intArr [3]int64 //int占8个字节
	//当我们定义完数组后，其实数组的各个元素有默认值 0
	fmt.Println(intArr)
	intArr[0] = 10
	intArr[1] = 20
	intArr[2] = 30
	fmt.Println(intArr)
	fmt.Printf("intArr的地址=%p intArr[0] 地址%p intArr[1] 地址%p intArr[2] 地址%p\n",
		&intArr, &intArr[0], &intArr[1], &intArr[2])

	//数组的使用
	/*练习1：循环输入5个值，保存到float64的地址，并输出*/
	// var score [5]float64

	// for i := 0; i < len(score); i++ {
	// 	fmt.Printf("请输入第%v位同学的成绩：", i+1)
	// 	fmt.Scanln(&score[i])
	// }
	// fmt.Println("五位同学的成绩依次为:")

	// for j := 0; j < len(score); j++ {
	// 	fmt.Println(score[j])
	// }

	//四种初始化数组的方式
	//方式1:
	var numArr01 [3]int = [3]int{1, 2, 3}
	fmt.Println("numArr01=", numArr01)

	//方式2:
	var numArr02 = [3]int{5, 6, 7}
	fmt.Println("numArr02=", numArr02)

	//方式3:
	//这里的 [...] 是规定的写法
	var numArr03 = [...]int{8, 9, 10}
	fmt.Println("numArr03=", numArr03)

	//方式4:
	var numArr04 = [...]int{1: 800, 0: 900, 2: 999}
	fmt.Println("numArr04=", numArr04)

	//方式4 + 类型推导
	strArr05 := [...]string{1: "tom", 0: "jack", 2: "mary"}
	fmt.Println("strArr05=", strArr05)

	//for-range遍历数组
	strArr06 := [...]string{1: "woke", 0: "weigeng", 3: "wiger"}
	for index, value := range strArr06 {
		fmt.Printf("%v  %v\n", index, value)
	}
}
