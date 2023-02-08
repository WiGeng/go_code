package main

import (
	"encoding/json"
	"fmt"
)

/* 结构体的细节说明和注意事项
1）结构体的所有字段在内存中是连续的
2）结构体是用户单独定义的类型，和其它类型进行转换时 需要有完全相同的字段(名字、个数和类型）
3）结构体进行type重定义（相当于取别名），GoLang会认为是新的数据类型，但是相互直接可以强制转换
4）struct的每个字段上，可以写一个tag,该tag可以通过反射机制获取，常见的使用场景就是序列化和反序列化
*/

func main() {
	/* 结构体的所有字段在内存中是连续的 */
	//结构体
	type Point struct {
		x int
		y int
	}

	//结构体
	type Rect struct {
		leftUp, rightDown Point
	}

	//结构体
	type Rect2 struct {
		leftUp, rightDown *Point
	}

	r1 := Rect{Point{1, 2}, Point{3, 4}}
	//r1有四个int, 在内存中是连续分布
	//打印地址
	fmt.Printf("r1.leftUp.x 地址=%p r1.leftUp.y 地址=%p r1.rightDown.x 地址=%p r1.rightDown.y 地址=%p \n",
		&r1.leftUp.x, &r1.leftUp.y, &r1.rightDown.x, &r1.rightDown.y)

	//r2有两个 *Point类型，这个两个*Point类型的本身地址也是连续的，
	//但是他们指向的地址不一定是连续

	r2 := Rect2{&Point{10, 20}, &Point{30, 40}}

	//打印地址
	fmt.Printf("r2.leftUp 本身地址=%p r2.rightDown 本身地址=%p \n",
		&r2.leftUp, &r2.rightDown)

	//他们指向的地址不一定是连续...， 这个要看系统在运行时是如何分配
	fmt.Printf("r2.leftUp 指向地址=%p r2.rightDown 指向地址=%p \n",
		r2.leftUp, r2.rightDown)

	/* 结构体类型进行转换 */
	type A struct {
		Num int
	}
	type B struct {
		Num int
	}
	var a A
	var b B
	a = A(b) // ? 可以转换，但是有要求，就是结构体的的字段要完全一样(包括:名字、个数和类型！)
	fmt.Println(a, b)

	/* struct tag */
	type Monster struct {
		Name string `json:"name"` // `json:"name"` 就是 struct tag
		Age  int    `json:"age"`
	}
	//1. 创建一个Monster变量
	monster := Monster{"牛魔王", 500}

	//2. 将monster变量序列化为 json格式字串
	//   json.Marshal 函数中使用反射
	jsonStr, err := json.Marshal(monster)
	if err != nil {
		fmt.Println("json 处理错误 ", err)
	}
	fmt.Println("jsonStr", string(jsonStr))

}
