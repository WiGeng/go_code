package main

import "fmt"

/*
实现接口，可以在不破坏继承的结构上，对结构体的功能进行扩展，实际上接口就是对继承关系的延展扩充

接口和继承解决的解决的问题不同
1) 继承的价值主要在于：解决代码的复用性和可维护性。
2) 接口的价值生要在于：设计，设计好各种规范(方法)，让其它自定义类型去实现这些方法。
3) 接口比继承更加灵活：继承是满足 is - a的关系，而接口只需满足 like - a的关系。
4) 接口在一定程度上实现代码解辖
*/

//Monkey结构体
type Monkey struct {
	Name string
}

func (this *Monkey) climbing() {
	fmt.Println(this.Name, " 生来会爬树...")
}

//LittleMonkey结构体
type LittleMonkey struct {
	Monkey //继承
}

//声明接口
type BirdAble interface {
	Flying()
}

//让LittleMonkey实现BirdAble
func (this *LittleMonkey) Flying() {
	fmt.Println(this.Name, " 通过学习，会飞翔...")
}

type FishAble interface {
	Swimming()
}

//让LittleMonkey实现FishAble
func (this *LittleMonkey) Swimming() {
	fmt.Println(this.Name, " 通过学习，会游泳...")
}

func main() {
	//创建一个LittleMonkey 实例
	monkey := LittleMonkey{
		Monkey{
			Name: "悟空",
		},
	}
	monkey.climbing()
	monkey.Flying()
	monkey.Swimming()

}
