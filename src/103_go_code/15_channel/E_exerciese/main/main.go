package main

import (
	"fmt"
)

/* 应用实例1
请完成goroutine 和channel协同工作的案例，具体要求：
	1） 开启一个writeData协程，向管道intChan中写入50个整数.
	2） 开启一个readData协程，从管道intChan中读取writeData写入的数据。
	3） 注意：writeData和readDate操作的是同一个管道
	4） 主线程需要等待writeData 和readDate协程都完 成工作才能退出管道
*/

//write Data
func writeData(intChan chan int) {
	for i := 1; i <= 50; i++ {
		//放入数据
		intChan <- i
		fmt.Println("writeData 写入数据=", i)
		//time.Sleep(time.Second)
	}
	close(intChan) //关闭
}

//read data
func readData(intChan chan int, exitChan chan bool) {
	for {
		v, ok := <-intChan
		if !ok {
			break
		}
		// time.Sleep(time.Second)
		fmt.Printf("readData 读到数据=%v\n", v)
	}
	//readData 读取完数据后，即任务完成
	exitChan <- true
	close(exitChan)

}

func main() {

	//创建两个管道
	intChan := make(chan int, 10)
	exitChan := make(chan bool, 1)

	go writeData(intChan)
	go readData(intChan, exitChan)

	//time.Sleep(time.Second * 10)
	for {
		_, ok := <-exitChan
		if !ok {
			break
		}
	}

}
