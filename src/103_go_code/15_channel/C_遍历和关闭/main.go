package main

import "fmt"

/* channel的遍历和关闭
channel的关闭
	使用内置函数close可以关闭channel， 当channel关闭后，就不能再向channel写数据了，但是仍然可以从该channel读取数据。
channel的遍历
	channel支持for-range的方式进行遍历，请注意两个細节
	1) 在遍历时，如果channel设有关闭，则回出现deadlock的错误
	2) 在遍历时，如果channel己经关闭，则会正常遍历数据，遍历完后，就会退出遍历。
*/

func main() {
	/* 1. 关闭管道 */
	intChan1 := make(chan int, 3)
	intChan1 <- 100
	intChan1 <- 200
	close(intChan1) // close
	//这是不能够再写入数到channel
	//intChan<- 300
	//当管道关闭后，读取数据是可以的
	n1 := <-intChan1
	fmt.Println("n1=", n1)

	/* 2.遍历管道 */
	intChan2 := make(chan int, 100)
	for i := 0; i < 100; i++ {
		intChan2 <- i * 2 //放入100个数据到管道
	}

	/* 遍历管道不能使用普通的 for 循环
	for i := 0; i < len(intChan2); i++ {
			...
	} */

	close(intChan2) //在遍历时，如果channel没有关闭，则会出现deadlock的错误;在遍历时，如果channel已经关闭，则会正常遍历数据，遍历完后，就会退出遍历

	/* for .. range 遍历管道*/
	for v := range intChan2 {
		fmt.Println("v=", v)
	}

}
