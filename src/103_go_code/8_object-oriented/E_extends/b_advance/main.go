package main

import "fmt"

/* 继承的深入讨论
1） 结构体可以使用嵌套蛋名结构体所有的字段和方法，即：首字母大写或者小写的字段、方法都可以使用
2） 匿名结构体字段访问可以简化
3） 当结构体和匿名结构体有相同的字段或者方法时，编译器采用就近访问原则访问，需访问匿名结构体的字段和方法，可通过匿名结构体名区分
4)  结构体嵌入两个(或多个)匿名结构体，如两个匿名结构体有相同的字段和方法(同时结构体本身没有同名的字段和方法)，在访问时，就必须明确指定匿名结构体名字，否则编译报错
5)  如果一个struct嵌套了一个有名结构体，这种模式就是组合，如果是组合关系，那么在访问组合的结构体的字段或方法时，必须带上结构体的名字
6)  嵌套匿名结构体后，也可以在创建结构体变量时，直接指定各个匿名结构体内字段的值

多重继承说明
	如一个struct嵌套了多个匿名结构体，那么该结构体可以直接访问嵌套的匿名结枃体的字段和方法，从而实现了多重继承。
*/

type A struct {
	Name string
	age  int
}

func (a *A) SayOk() {
	fmt.Println("A SayOk", a.Name)
}

func (a *A) hello() {
	fmt.Println("A hello", a.Name)
}

/*----------------------------------------------------------*/

type B struct {
	A
	Name string
}

func (b *B) SayOk() {
	fmt.Println("B SayOk", b.Name)
}

type C struct {
	A   //匿名结构体
	b B //有名结构体
}

/*----------------------------------------------------------*/
type Goods struct {
	Name  string
	Price float64
}

type Brand struct {
	Name    string
	Address string
}
type TV struct {
	Goods
	Brand
}

type TV2 struct {
	*Goods
	*Brand
}

func main() {

	// var b B
	// b.A.Name = "tom"
	// b.A.age = 19
	// b.A.SayOk()
	// b.A.hello()

	// //上面的写法可以简化

	// b.Name = "smith"
	// b.age = 20
	// b.SayOk()
	// b.hello()

	var b B
	b.Name = "jack"
	b.A.Name = "scott"
	b.age = 100 //ok
	b.SayOk()   //  B SayOk  "jack"
	b.A.SayOk() //  A SayOk  "scott"
	b.hello()   //  A hello  "scott"

	//针对有名结构体，使用时必须带上有名结构体的名称
	var c C
	c.b.Name = "woke"
	fmt.Println(c)

	//嵌套匿名结构体后，也可以在创建结构体变量(实例)时，直接指定各个匿名结构体字段的值
	tv := TV{Goods{"电视机001", 5000.99}, Brand{"海尔", "山东"}}
	fmt.Println(tv)
	tv2 := TV{
		Goods{
			Price: 5000.99,
			Name:  "电视机002",
		},
		Brand{
			Name:    "夏普",
			Address: "北京",
		},
	}
	fmt.Println(tv2)

	tv3 := TV2{&Goods{"电视机003", 5000.99}, &Brand{"海信", "陕西"}}
	fmt.Println(*tv3.Goods, *tv3.Brand)
}
