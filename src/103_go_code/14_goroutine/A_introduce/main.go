package main

import (
	"fmt"
	"strconv"
	"time"
)

/* goroutine 基本介绍
进程和线程：
	1） 进程就是程序程序在操作系统中的一次执行过程，是系统进行资源分配和调度的基本单位
	2） 线程是进程的一个执行实例，是程序执行的最小单元，它是比进程更小的能独立运行的基本单位
	3) 一个进程可以创建核销毁多个线程，同一个进程中的多个线程可以并发执行。
4) 一个程序至少有一个进程，一个进程至少有一个线程

并发和并行
	1） 多线程程序在单核上运行，就是并发
		并发：因为是在一个cpu上，比如有10个线程，每个线程执行10毫秒(进行轮询操作)，从人的角度看，好像这10个线程都在运行，但是从微观上看，
	在某一个时间点看，其实只有一个线程在执行，这就是并发。
	2） 多线程程序在多核上运行，就是并行
		并行：因为是在多个cpu上(比如有10个cpu)，比如有10个线程，每个线程执行10毫1秒(各自在不同cpu上执行)，从人的角度看，这10个线程都
	在运行，但是从微观上看，在某一个时间点看，也同时有10个线程在执行，这就是并行

Go协程和Go主线程
	1） Go主线程(有程序员直接称为线程/也可以理解成进程)：一个Go线程上，可以起多个协程，你可以这样理解，协程是轻量级的线程。
	2） Go协程的特点***[面试题]:
		• 有独立的栈空间
		• 共享程序堆空间
		• 调度由用户控制
		• 协程是轻量級的线程
*/

/* 案例入门
在主线程(可以理解成进程)中，开启一个goroutine, 该协程每隔1秒输出 "hello,world"
在主线程中也每隔一秒输出"hello,golang", 输出5次后，退出程序，要求主线程和goroutine同时执行
	main() hello,golang-0
	test() hello,world-0
	test() hello,world-1
	main() hello,golang-1
	main() hello,golang-2
	test() hello,world-2
	main() hello,golang-3
	test() hello,world-3
	test() hello,world-4
	main() hello,golang-4
	main() hello,golang-5
	test() hello,world-5
*/
func test() {
	for i := 0; i <= 5; i++ {
		fmt.Printf("test() hello,world-%v\n", strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}

func main() {
	go test() //开启一个协程
	for i := 0; i <= 5; i++ {
		fmt.Printf("main() hello,golang-%v\n", strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}
