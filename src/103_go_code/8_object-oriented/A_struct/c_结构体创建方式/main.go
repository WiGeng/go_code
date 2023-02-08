package main

import "fmt"

/* 创建结构体变量和访问结构体字段
1）方式1- 直接声明
	案例演示：var person Person
2）方式2- {}
	案例演示：var person Persan = Personff
3) 方式3- ＆
	案例演示：var person *Person = new(Person)
4) 方式4- {}
	案例演示: var person *Person =&Person{}
＞说明：
	(1) 第3种和第4种方式返回的是 结构体指针。
	(2) 结构体指针访问字段的标准方式应该是：(*结构体指针).字段名，比如(*person).Name = "tom"
	(3) 但go做了一个简化，也支持 结构体指针.字段名，比如person.Name ="tom"。更加符合程序员使用的习惯，go编译器底层对person.Name 做了转化 (*person).Name
*/

type person struct {
	name string
	age  int
}

func main() {

	/* 方式1: */
	var p1 person
	p1.name = "woke"
	p1.age = 26
	fmt.Println(p1)

	/* 方式2 */
	p2 := person{"西安", 70}
	fmt.Println(p2)

	/* 方式3 */
	var p3 *person = new(person) //等价于 p3 := new(person)
	(*p3).name = "smith"         //因为p3是一个指针，因此标准的给字段赋值方式
	p3.name = "WiGeng"           //(*p3).Name = "smith" 也可以这样写 p3.Name = "smith",//原因: go的设计者 为了程序员使用方便，底层会对 p3.Name = "smith" 进行处理, 给p3加上取值运算(*p3).Name = "smith"
	(*p3).age = 30
	p3.age = 100
	fmt.Println(*p3)

	/* 方式4 */
	var person *person = &person{}
	//var person *person = &person{"mary", 60} //也可以直接给字符赋值
	(*person).name = "scott"
	person.name = "wiger" //go的设计者为了程序员使用方便，也可以 person.name = "scott"
	(*person).age = 88
	person.age = 100
	fmt.Println(*person)

}
