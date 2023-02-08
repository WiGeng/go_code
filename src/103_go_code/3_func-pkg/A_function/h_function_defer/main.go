package main

import (
	"fmt"
)

/* defer
为什么需要defer:
	在函数中，程序员经常需要创建资源（如数据库连接、创建文件句柄、锁等）为了在函数执行完毕后及时释放资源，go引入了defer(延迟机制)
defer的细节说明
	1) 当go执行到一个defer时，不会立即执行defer后的语句，而是将defer 后的语句压入到一个栈中，暂时称该栈为defer栈，然后继续执行函数下一个语句
	2）当函数执行完毕后，在从defer栈中，依次从栈顶取出语句执行(注，遵守栈 先入后出的机制
	3）在defer将语句放入到栈时，也会将相关的值 defer fmt.Printin('ok1 n1=',n1)拷贝同时入栈。
总结
	1在golang编程中的通常做法是，创建资源后，比如(打开了文件，获取了数据库的链接，或者是锁资源)，可以执行 defer file.CloseO defer connect.Closeo
	2）在defer后，可以继续使用创建资源.
	3）当函数完毕后，系统会依次从defer栈中，取出语句，关闭资源
	4) 这种机制，非带商洁，程序员水用再为在什么时机关闭资源而烦心。
*/

//当执行到defer时，暂时不执行，会将defer后面的语句压入到独立的栈(defer栈)
//当[sum]函数执行完毕后，再从defer栈，按照先入后出的方式出栈，执行
func sum(n1 int, n2 int) int {
	defer fmt.Println("ok1 n1=", n1)
	defer fmt.Println("ok2 n2=", n2)
	res := n1 + n2
	fmt.Println("ok3 resp[sum]=", res)
	return res
}

func sum1(n1 int, n2 int) int {
	defer fmt.Println("ok1 n1=", n1) // n1 = 10
	defer fmt.Println("ok2 n2=", n2) // n2 = 20
	n1++
	n2++
	res1 := n1 + n2 //n1 =11,n2 = 21
	fmt.Println("ok3 res1[sum]=", res1)
	return res1
}

func main() {
	res := sum(10, 20)
	fmt.Println("res[main]=", res)

	res1 := sum1(10, 20)
	fmt.Println("res1[main]=", res1)

}
