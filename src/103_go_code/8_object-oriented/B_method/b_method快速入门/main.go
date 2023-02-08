package main

import "fmt"

type Person struct {
	Name string
}

// 1）给Person结构体添加speak方法，输出 xxx是一个好人
func (p Person) speak() {
	fmt.Printf("%v is a good man !\n", p.Name)
}

// 2）给person结构体添加jisuan 方法,可以计算从1+.+1000的结果
func (p Person) jisuan() { //方法体中可以和函数一样进行多种运算
	sum := 0
	for i := 0; i <= 1000; i++ {
		sum += i
	}
	fmt.Printf("1+2+3+...+999+1000=%v\n", sum)
}

// 3) 给Person结构体jisuan2 方法,该方法可以接收一个数n，计算从1+.+n的结果
func (p Person) jisuan2(n int) { //方法体中可以和函数一样进行多种运算
	sum := 0
	for i := 0; i <= n; i++ {
		sum += i
	}
	fmt.Printf("1+2+3+...+%v=%v\n", n, sum)
}

// 4) 给Person结构体添加getsum方法,可以计算两个数的和，并返回结果。
func (p Person) getsum(n int, m int) int {
	return n + m

}

func main() {
	var p Person
	p.Name = "woke"
	p.speak()
	p.jisuan()
	p.jisuan2(24)
	fmt.Printf("10+20=%v", p.getsum(10, 20))
}
