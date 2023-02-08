package main

import "fmt"

/* method
基本介绍
	在某些情况下，我们要需要声明(定义)方法。比如Person结构体：除了有一些字段外（年龄，姓名...),Persen结构体还有一些行为比如:
	可以说话、跑步…,通过学习，还可以做算术题，这时就要用方法才能完成
定义：
	Golang中的方法是作用在指定的数据类型上的(即：和指定的数据类型绑定），因此自定义类型，都可以有方法，而不仅仅是struct
		type A struct {
			Num int
		}

		func (a A) test() {
			fmt. Println(a.Num)
		}
	＞对上面的语法的说明
		(1) func (a A) test() {} 表示 A 结构体有一方法，方法名为test
		(2) (a A）体现出test方法和 A 是类型绑定的
*/

type Person struct {
	Name string
}

// 给Person类型绑定一种方法
func (p Person) test() { //  p[可以自定义]，表示某个Person类型的变量调用，这个p就是它的副本，类似于函数传参
	p.Name = "woke"
	fmt.Println("test() name=", p.Name)
}

func main() {
	var p Person
	p.Name = "wiger"
	p.test() //test方法只能通过Person类型的变量来调用，而不能使用其他类型或者直接调用
	fmt.Println("mian() name=", p.Name)
}
