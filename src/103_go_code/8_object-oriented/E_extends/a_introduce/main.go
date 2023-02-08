package main

import "fmt"

/* 继承
1. 继承基本介绍
	1）继承可以解决代码复用，让我们的编程更加靠近人类思维。
	2）当多个结构体存在相同的属性(字段)和方法时,可以从这些结构体中抽象出结构体(比如下边的Examinee),在该结构体中定义
	这些相同的属性和方法。其它的结构体不需要重新定义这些属性和方法，只需嵌套一个Examinee匿名结构体即可。
	3）也就是说：在Golang中，如果一个struct嵌套了另一个匿名结构体，那么这个结构体可以直接访问蛋名结构体的字段和方法，从
	而实现了继承特性。

2. 嵌套匿名结构体的基本语法
	type Goods struct {
		Name string
		Price int
	}

	type Book struct {
		Goods //这里就是嵌套匿名结构体Goods
		Writer string
	}

*/

// 考生信息结构体
type Examinee struct {
	Name  string
	Age   int
	Score float64
}

// 方法 显示学生信息[姓名 年龄 成绩]
func (stu *Examinee) ShowInfo() {
	fmt.Printf("学生姓名=%v 年龄=%v 成绩=%v\n", stu.Name, stu.Age, stu.Score)
}

// 方法 成绩计算
func (stu *Examinee) SetScore(score float64) {
	stu.Score = score
}

/* ------------------------小学生---------------------------- */
type Pupil struct {
	Examinee //嵌入[Examinee]匿名结构体
}

// 方法 模拟考试
func (stu1 *Pupil) examinating() {
	fmt.Println("小学生考试中......")
}

/* ------------------------大学生---------------------------- */
type Undergraduate struct {
	Examinee //嵌入[Examinee]匿名结构体
}

// 方法 模拟考试
func (stu2 *Undergraduate) examinating() {
	fmt.Println("大学生考试中......")
}
func main() {
	pupil := &Pupil{}
	pupil.Examinee.Name = "woke"
	pupil.Examinee.Age = 8
	pupil.examinating()
	pupil.Examinee.SetScore(90)
	pupil.Examinee.ShowInfo()

	undergraduate := &Undergraduate{}
	undergraduate.Examinee.Name = "wigeng"
	undergraduate.Examinee.Age = 27
	undergraduate.examinating()
	undergraduate.Examinee.SetScore(145)
	undergraduate.Examinee.ShowInfo()
}
