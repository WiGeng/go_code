package main

import "fmt"

/* 多态介绍
基本介绍
	变量(实例)具有多种形态。面向对象的第三大特征，在Go语言，多态特征是通过接口实现的。可以按照统一的接口来调用不同的实现。这时接口变量就呈现不同的形态。
快速入门
	在前面的Usb接口察例，Usb usb，既可以接收手机变量，又可以接收相机变量，就体现了usb 接口多态特性
接口体现多态特征
1) 多态参数
	在前面的Usb接口案例，Usb usb，即可以接收手机变量，又可以接收相机变量，就体现了Usb 接口多态
2）多态数组
	演示一个案例：给Usb数组中，存放Phone 结构体 和 Camera结构体变量，Phone还有一个特有的方法cal(），请遍历Usb数组，如果是Phone变量，除了
调用Usb 接口声明的方法外，还需要调用Phone 特有方法 call.
*/

//声明/定义一个接口
type Usb interface {
	//声明了两个没有实现的方法
	Start()
	Stop()
}

type Phone struct {
	name string
}

//让Phone 实现 Usb接口的方法
func (p Phone) Start() {
	fmt.Println("手机开始工作。。。")
}
func (p Phone) Stop() {
	fmt.Println("手机停止工作。。。")
}

type Camera struct {
	name string
}

//让Camera 实现   Usb接口的方法
func (c Camera) Start() {
	fmt.Println("相机开始工作。。。")
}
func (c Camera) Stop() {
	fmt.Println("相机停止工作。。。")
}

func main() {
	//定义一个Usb接口数组，可以存放Phone和Camera的结构体变量
	//这里就体现出多态数组
	var usbArr [3]Usb
	usbArr[0] = Phone{"vivo"}
	usbArr[1] = Phone{"小米"}
	usbArr[2] = Camera{"索尼"}

	fmt.Println(usbArr)

}
