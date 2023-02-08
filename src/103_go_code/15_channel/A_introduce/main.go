package main

import "fmt"

/*
1. 为什么需要chanrel
前面使用全局变量加锁同步来解决goroutine的通讯，但不完美
	1） 主线程在等待所有goroutine全部完成的时间很难确定，我们这里设置5秒，仅仅是估算。
	2） 如果主线程休眠时间长了，会加长等待时间，如果等待时间短了，可能还有goroutine处于工作状态，这时也会随主线程的退出而销毁
	3） 通过全局变量加锁同步来实现通讯，也并不利用多个协程对全局变量的读写操
	4） 上面种种分析都在呼唤一个新的通讯机制-channel

2. channel的介绍
	1) channle本质就是一个数据结构-队列
	2) 数据是先进先出[FIFO]
	3) 线程安全，多goroutine访问时，不需要加锁，就是说channel本身就是线程安全的
	4) channel是有类型的，一个string的channel只能存放string类型数据。

3. 定义/声明channel
var 变量名chan 数据类型
	举例：
	var intChan chan int (intChan用于存放int数据）
	var mapChan chan maplintJstring (mapChan用于 存放maptint]string类型）
	var perChan chan Person
	var perChan2 chan #Person

	说明
		1) channel是引用类型
		2) channel必须初始化才能写入数据，即make后才能使用
		3) 管道是有类型的，intChan 只能写入 整数 int

4. channel使用的注意事项
	1. channel中只能存放指定的数据类型
	2. channle的数据放满后，就不能再放入了
	3. 如果从channel取出数据后，可以继续放入
	4. 在没有使用协程的情况下，如果channel数据取完了，再取，就会报dead lock
*/

func main() {

	//演示一下管道的使用
	//1. 创建一个可以存放3个int类型的管道
	// var intChan chan int
	intChan := make(chan int, 3)

	//2. 看看intChan是什么
	fmt.Printf("intChan的值=%v intChan本身的地址=%p\n", intChan, &intChan)

	//3. 向管道写入数据
	intChan <- 10
	num := 211
	intChan <- num
	intChan <- 50
	// //如果从channel取出数据后，可以继续放入
	<-intChan
	intChan <- 98 //注意, 当我们给管写入数据时，不能超过其容量

	//4. 看看管道的长度和cap(容量)
	fmt.Printf("channel len= %v cap=%v \n", len(intChan), cap(intChan)) // 3, 3

	//5. 从管道中读取数据

	// var num2 int
	num2 := <-intChan
	fmt.Println("num2=", num2)
	fmt.Printf("channel len= %v cap=%v \n", len(intChan), cap(intChan)) // 2, 3

	//6. 在没有使用协程的情况下，如果我们的管道数据已经全部取出，再取就会报告 deadlock

	num3 := <-intChan
	num4 := <-intChan

	//num5 := <-intChan

	fmt.Println("num3=", num3, "num4=", num4 /*, "num5=", num5*/)

}
