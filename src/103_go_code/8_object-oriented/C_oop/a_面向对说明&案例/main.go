package main

import "fmt"

/*
Golang语言面向对象编程说明
	1) Golang也支持面向对象编程(OOP)，但是和传统的面向对象编程有区别，并不是纯粹的面向对象语言。所以我们说Golang支持面向对象编程特性是比较准确的。
	2) Golang没有类(class)，Go语言的结构体(struct) 和其它编程语言的类(class)有同等的地位，你可以理解Golang是基于struct来实现OOP特性的。
	3) Golang面向对象编程非常简洁，去掉了传统OOP语言的继承、方法重載、构造函数和析构函数、隐藏的this指针等等
	4) Golang仍然有面向对象编程的继承，封装和多态的特性，只是实现的方式和其它OOP语言不一样，比如继承：Golang没有extends 关键宇，继承是通过匿名字段来实现。
	5) Golang 面向对象(OOP)很优雅，OOP本身就是语言类型系统(type system)的一部分，通过接口(interface)关联，糯合性低，也非常灵活。后面同学们会充分体会到这个特点。也就是说在Golang中面向接口编程是非常重要的特性。
*/

/*
1. 学生案例：
编写一个Student结构体，包含name、gender、age、id、score字段，分别为string、string、int、int、float64类型。
结构体中声明一个say方法，返回string类型，方法返回信息中包含所有字段值。
在main方法中，创建Student结构体实例(变量)，并访问say方法，并将调用结果打印输出。
*/

type Student struct {
	name   string
	gender string
	age    int
	id     int
	score  float64
}

func (student *Student) info() string {
	stu := fmt.Sprintf("学生信息  姓名:[%v] 性别:[%v] 年龄:[%v] id:[%v] 成绩:[%v]",
		student.name, student.gender, student.age, student.id, student.score)
	return stu
}

/*
2. 盒子案例
1） 编程创建一个Box结构体，在其中声明三个字段表示一个立方体的长、宽和高，长宽高要从终端获取
2） 声明一个方法获取立方体的体积。
3） 创建一个Box结构体变量，打印给定尺寸的立方体的体积。
*/

type Box struct {
	long float64
	wide float64
	high float64
}

func (box Box) volume() float64 {
	volume := box.long * box.wide * box.high
	return volume

}

/*
3. 景区门票案例
1）一个景区根据游人的年龄收取不同价格的门票，比如年龄大于18，收费20元，其它情况门票免费.
2）请编写Visitor结构体，根据年龄段决定能够购买的门票价格并输出
*/
type Visitor struct {
	Name string
	Age  int
}

func (visitor *Visitor) showPrice() {
	if visitor.Age >= 90 || visitor.Age <= 8 {
		fmt.Println("考虑到安全，就不要玩了")
		return
	}
	if visitor.Age > 18 {
		fmt.Printf("游客的名字为 %v 年龄为 %v 收费20元 \n", visitor.Name, visitor.Age)
	} else {
		fmt.Printf("游客的名字为 %v 年龄为 %v 免费 \n", visitor.Name, visitor.Age)
	}
}

func main() {
	var student = Student{
		name:   "woke",
		gender: "male",
		age:    26,
		id:     17091115,
		score:  99.99,
	}
	fmt.Println(student.info())

	var box = Box{
		long: 11.11,
		wide: 8.88,
		high: 3.33,
	}
	fmt.Printf("长方体的体积为%v立方米\n", box.volume())

	var v Visitor
	for {
		fmt.Println("请输入你的名字")
		fmt.Scanln(&v.Name)
		if v.Name == "n" {
			fmt.Println("退出程序....")
			break
		}
		fmt.Println("请输入你的年龄")
		fmt.Scanln(&v.Age)
		v.showPrice()

	}
}
