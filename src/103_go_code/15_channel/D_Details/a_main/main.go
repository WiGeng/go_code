package main

import "fmt"

/* 细节
1) channel可以声明为只读或者只只写性质
2）channel只读和只写最佳实践案例
3) 使用select可以解决从管道取数据的阻塞问题
4) goroutine中使用recover，解决协程中出现panic，导致程序崩溃问题
说明：
	如果我们起了一个协程，但是这个协程出现了panic，如果我们没有捕获这个panic,就会造成整个程序崩溃，这时我们可以在goroutine中
	使用recover来捕获panic， 进行处理，这样即使这个协程发生的问题，但是主线程仍然不受影响，可以继续执行。
*/

func main() {
	/* 1.管道可以声明为只读或者只写 */

	/* 1. 在默认情况下下，管道是双向 */
	// var chan1 chan int    //可读可写

	/* 2. 声明为只写 */
	chan2 := make(chan<- int, 3)
	chan2 <- 20
	//num := <-chan2 //error
	fmt.Println("chan2=", chan2)

	/* 3. 声明为只读 */
	var chan3 <-chan int
	num2 := <-chan3
	//chan3<- 30 //err
	fmt.Println("num2", num2)

}
