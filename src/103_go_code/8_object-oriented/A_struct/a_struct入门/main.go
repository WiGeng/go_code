package main

import (
	"fmt"
)

/* 结构体
结构体和结构体变量(实例)的区别和联系
	通过下面的案例可以看出：
	1) 结构体是自定义的数据类型，代表一类事物。
	2) 结构体变量(实例)是县体的，实际的，代表一个具体变量
如何声明结构体：
	type 标识符 struct{
		field1 type
		field2 type
		field3 type
	}
字段/属性
	＞基本介绍
	1） 从概念或叫法上看：结构体字段 = 属性 = field  （后面统一叫字段）
	2） 字段是结构体的一个组成部分，一般是基本数据类型、数组，也可是引用类型。比如我们下边定义Cat结构体的 Name string 就是属性
	>注意事项和细节说明
	1） 字段声明语法同变量，示例：字段名 字段类型
	2） 字段的类型可以为：基本类型、数组或引用类型
	3） 在创建一个结构体变量后，如果没有给字段赋值，都对应一个零值(默认值)，规则同前面讲的一样：
		布尔类型是false，数值是0，字符串是"。
		数组类型的默认值和它的元素类型相关，比如 score I3jint 则为10,0,0]
		指针，slice， 和map的零值都是nil，即还没有分配空间。
	4)  不同结构体变量的字段是独立，互不影响，一个结构体变量字段的更改，不影响另外一个。因为结构体是值类型
*/

//定义一个Cat结构体，将Cat的各个字段/属性信息，放入到Cat结构体进行管理
type Cat struct {
	Name   string
	Age    int
	Color  string
	Hobby  string
	Scores [3]int // 字段是数组...
}

func main() {

	// 张老太养了20只猫猫:一只名字叫小白,今年3岁,白色。还有一只叫小花,
	// 今年100岁,花色。请编写一个程序，当用户输入小猫的名字时，就显示该猫的名字，
	// 年龄，颜色。如果用户输入的小猫名错误，则显示 张老太没有这只猫猫。

	// 1. 使用变量的处理
	// var cat1Name string = "小白"
	// var cat1Age int = 3
	// var cat1Color string = "白色"

	// var cat2Name string = "小花"
	// var cat2Age int = 100
	// var cat2Color string = "花色"

	// 2. 使用数组解决
	// var catNames [2]string = [...]string{"小白", "小花"}
	// var catAges [2]int = [...]int{3, 100}
	// var catColors [2]string = [...]string{"白色", "花色"}

	// 使用struct来完成案例
	// 创建一个Cat的变量
	var cat1 Cat // var a int

	fmt.Printf("cat1的地址=%p\n", &cat1)
	cat1.Name = "小白"
	cat1.Age = 3
	cat1.Color = "白色"
	cat1.Hobby = "吃<・)))><<"

	fmt.Println("cat1=", cat1)

	fmt.Println("猫猫的信息如下：")
	fmt.Println("name=", cat1.Name)
	fmt.Println("Age=", cat1.Age)
	fmt.Println("color=", cat1.Color)
	fmt.Println("hobby=", cat1.Hobby)

}
