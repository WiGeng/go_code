package main

import (
	"awesomeProject/src/103_go_code/8_object-oriented/D_Encapsulation/a_introduce/model"
	"fmt"
)

/* 面向对象编程-封装(encapsulation)
封装介绍
	封装就是把抽象出的字段和对字段的操作封装在一起,数据被保护在内部,程序的其它包只有通过被授权的操作(方法),才能对字段进行操作

封装的理解和好处
	1）隐藏实现细节
	2）提可以对数据进行验证，保证安全合理

如何体现封装
	1) 对结构体中的属性进行封装
	2) 通过方法、包实现封装

封装的实现步骤
1）将结构体、字段(属性)的首字母小写（不能导出，其它包不能使用，类似private)
2）给结构体所在包提供一个工厂模式的函数，首字母大写，类似一个构造函数
3）提供一个首字母大写的Set方法(类似其它语言的public)，用于对属性判断并赋值
	func (var 结构体类型名） Setxxx(参数列表)(返回值列表){
		//加入数据验证的业务逻辑
		var.字段 = 参数
	}
4) 提供一个首字母大写的Get方法(类似其它语言的public)，用于获取属性的值
	func (var 结构体类型名）GetXxx(){
		return var.字段;

特别说明：
	在Golang开发中并没有特别强调封装，这点并不像Java,不用总是用java的语法特性来看待Golang，Golang本身对面向对象的特性做了简化的。
*/

func main() {

	p := model.NewPerson("woke")
	p.SetAge(26)
	p.SetSal(15000)
	fmt.Println(*p)
	fmt.Println("name =", p.Name, " age =", p.GetAge(), " sal =", p.GetSal())

}
