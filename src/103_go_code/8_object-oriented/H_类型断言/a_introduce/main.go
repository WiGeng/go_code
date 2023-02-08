package main

import (
	"fmt"
)

/* 类型断言
基本介绍：
	类型断言，由于接口是一般类型，不知道具体类型，如果要转成具体类型，就需要使用类型断言
*/

type Point struct {
	x int
	y int
}

func main() {
	var a interface{}
	var point Point = Point{1, 2}
	a = point //oK
	// 如何将 a 赋给一个Point变量?
	var b Point
	// b = a 不可以
	// b = a.(Point) // 可以
	b = a.(Point) //b=a.(point)就是类型断言，表示判断a是否指向point类型的变量，如果是就转成point类型并赋给b变量，否则报错。
	fmt.Println(b)

	//类型断言的其它案例
	// var x interface{}
	// var b2 float32 = 1.1
	// x = b2  //空接口，可以接收任意类型
	// // x=>float32 [使用类型断言]
	// y := x.(float32)
	// fmt.Printf("y 的类型是 %T 值是=%v", y, y)

	//类型断言(带检测的)
	var x interface{}
	var b2 float32 = 2.1
	x = b2 //空接口，可以接收任意类型
	// x=>float32 [使用类型断言]

	//类型断言(带检测的)

	// y, flag := x.(float64)    //也可写成下方的简写版
	// if flag == true {
	// 	fmt.Println("convert success")
	// 	fmt.Printf("y 的类型是 %T 值是=%v\n", y, y)
	// } else {
	// 	fmt.Println("convert fail")
	// }
	// fmt.Println("继续执行...")

	if y, flag := x.(float64); flag {
		fmt.Println("convert success")
		fmt.Printf("y 的类型是 %T 值是=%v\n", y, y)
	} else {
		fmt.Println("convert fail")
	}
	fmt.Println("继续执行...")

}
