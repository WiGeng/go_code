package main

import (
	"fmt"
)

/*🚩🚩方法和函数的区别
1)调用方式不一样
    函数的调用方式：函数名(实参列表)
    方法的调用方式：变量.方法名(实参列表）
2)对于普通函数，接收者为值类型时，不能将指针类型的数据直接传递，反之亦然
3)对于方法（如struct的方法），接收者为值类型时，可以直接用指针类型的变量调用方法，反过来同样也可以
*/

type Person struct {
	Name string
}

//函数
//对于普通函数，接收者为值类型时，不能将指针类型的数据直接传递，反之亦然

func test01(p Person) {
	fmt.Println(p.Name)
}

func test02(p *Person) {
	fmt.Println(p.Name)
}

//对于方法（如struct的方法），
//接收者为值类型时，可以直接用指针类型的变量调用方法，反过来同样也可以

func (p Person) test03() {
	p.Name = "jack"
	fmt.Println("test03() =", p.Name) // jack
}

func (p *Person) test04() {
	p.Name = "mary"
	fmt.Println("test03() =", p.Name) // mary
}

/* 总结
1）不管调用形式如何，真正决定是值拷贝还是地址拷贝，看这个方法是和哪个类型鄉定
	A. 如果是和值类型，比如 (p Person) ,则是值拷贝，
	B. 如果是指针类型，比如 (p *Person) ,则是地址拷贝。
*/

func main() {

	p := Person{"tom"}
	test01(p)
	test02(&p)

	p.test03()
	fmt.Println("main() p.name=", p.Name) // tom

	(&p).test03() // 从形式上是传入地址，但是本质仍然是值拷贝

	fmt.Println("main() p.name=", p.Name) // tom

	(&p).test04()
	fmt.Println("main() p.name=", p.Name) // mary
	p.test04()                            // 等价 (&p).test04 , 从形式上是传入值类型，但是本质仍然是地址拷贝

}
