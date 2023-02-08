package main

import "fmt"

/* interface
基本介绍
	interface 类型可以定义一组方法，但是这些不需要实现。并且interface不能包含任何变量。到某个自定义类型（比如结构体Phone）要使用的时候，在根据具体情况把这些方法写出来。
基本浯法
	type 接口名 interface{
		method1(参数列表） 返回値列表
		method2(参数列表)  返回值列表
	}
小结说明：
	1） 接口里的所有方法都没有方法体，即接口的方法都是没有实现的方法。接口体现了程序设计的多态和高内聚低偶合的思想。
	2） Golang 中的接口，不需要显式的实现。只要一个变量，含有接口类型中的所有方法，那么这个变量就实现这个接口。因此，Golang中没有implement 这样的关键字
*/

//声明一个接口
type Usb interface {
	//声明了两个没有实现的方法
	Start()
	Stop()
}

type Phone struct {
}

//让Phone 实现Usb接口的方法
func (p *Phone) Start() {
	fmt.Println("手机开始工作...请稍等～")
}
func (p Phone) Stop() {
	fmt.Println("手机停止工作...")
}

type Camera struct {
}

//让Camera 实现Usb接口的方法
func (c Camera) Start() {
	fmt.Println("相机开始工作...请稍等～")
}
func (c Camera) Stop() {
	fmt.Println("相机停止工作...")
}

//计算机
type Computer struct {
}

//编写一个方法Working方法，接收一个Usb接口类型变量
//只要是实现了 Usb接口 （所谓实现Usb接口，就是指实现了Usb接口声明所有方法
func (c Computer) Working(usb Usb) {
	//通过usb接口变量来调用Start和Stop方法
	usb.Start()
	usb.Stop()
}

func main() {
	//先创建结构体变量
	computer := Computer{}
	phone := Phone{}
	camera := Camera{}

	//关键点
	computer.Working(&phone)
	computer.Working(camera)
}
