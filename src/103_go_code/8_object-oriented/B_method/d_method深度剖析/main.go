package main

import "fmt"

/*
方法的声明（定义）：

func (recevier type) methodName(参数列表) (返回值列表）{
		方法体
	    return 返回值
}
	1) 参数列表：表示方法输入
	2) recevier type :表示这个方法和type这个类型进行绑定，或者说该方法作用于type类型
	3) receiver type : type 可以是结构体，也可以其它的自定义类型
	4) receiver：就是type类型的一个变量(实例，比如：Person结构体的一个变量(实例）
	5) 参数列表：表示方法输入
	6) 返回值列表：表示返回的值，可以多个
	7) 方法主体：表示为了实现某一功能代码块
	8) return 语句不是必须的。

方法注意事项和细节讨论
1) 结构体类型是值类型，在方法调用中，遵守值类型的传递机制，是值拷贝传递方式
2) 如程序员希望在方法中，修改结构体变量的值，可以通过结构体指针的方式来处理
3) Golang中的方法作用在指定的数据类型上的(即：和指定的数据类型鄉定)，因此自定义类型，都可以有方法，而不仅仅是struct，比如int,float32等都可以有方法
4) 方法的访问范围控制的规则，和函数一样。方法名首字母小写，只能在本包访问，方法首字母大写，可以在本包和其它包访问
5) 如果一个变量实现了String()这个方法，那么fmt.Printin默认会调用这个变量的String()进行输出
*/

type integer int

func (i integer) print() {
	fmt.Println("i=", i)
}

//编写一个方法，可以改变i的值
func (i *integer) change() {
	*i = *i + 1
}

type Student struct {
	Name string
	Age  int
}

//给*Student实现方法String()
func (stu *Student) String() string {
	str := fmt.Sprintf("Name=[%v] Age=[%v]", stu.Name, stu.Age)
	return str
}

func main() {
	var i integer = 10
	i.print()
	i.change()
	fmt.Println("i=", i)

	//定义一个Student变量
	stu := Student{
		Name: "tom",
		Age:  20,
	}
	//如果你实现了 *Student 类型的 String方法，就会自动调用
	fmt.Println(&stu)
}
