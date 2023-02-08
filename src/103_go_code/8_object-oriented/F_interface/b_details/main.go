package main

import "fmt"

/* 注意事项和细节
1） 接口本身不能创建实例，但是可以指向一个实现了该接口的自定义类型的变量实例
2） 接口中所有的方法都没有方法体，即都是没有实现的方法
3)  在Golang中，一个自定义类型需要将某个接口的所有方法都实现，我们说这个自定义类型实现了该接口
4） 一个自定义类型只有实现了某个接口，才能将该自定义类型的实例（变量）赋给接口类型
5） 只要是自定义数据类型，就可以实现接口，不仅仅是结构体类型
6)  一个自定义类型可以实现多个接口
7)  Golang接口中不能有任何变量
8） 一个接口(比如A接口)可以继承多个别的接口（比如B，C接口），这时如果要实现A按口，也必须将B,C接口的方法也全部实现
9） interface类型默认是一个指针（引用类型，如果没有对interface初始化就使用，那么会输出nil
10) 空接口interfacef没有任何方法，所以所有类型都实现了空接口
*/

type Stu struct {
	Name string
}

func (stu Stu) Say() {
	fmt.Println("Stu Say()")
}

type integer int

func (i integer) Say() {
	fmt.Println("integer Say i =", i)
}

type AInterface interface {
	Say()
}

type BInterface interface {
	Hello()
}
type Monster struct {
}

func (m Monster) Hello() {
	fmt.Println("Monster Hello()~~")
}

func (m Monster) Say() {
	fmt.Println("Monster Say()~~")
}

func main() {
	var stu Stu //结构体变量，实现了 Say() 实现了 AInterface
	var a AInterface = stu
	a.Say()

	var i integer = 10
	var b AInterface = i
	b.Say() // integer Say i = 10

	//Monster实现了AInterface 和 BInterface
	var monster Monster
	var a2 AInterface = monster
	var b2 BInterface = monster
	a2.Say()
	b2.Hello()
}
