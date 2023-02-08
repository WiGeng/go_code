package main

import (
	"fmt"
)

/* 在程序中，程序运行的流程控制决定了程序事如何执行的，主要有三大流程控制语句
1）顺序控制
	程序从上到下逐行执行，中间没有任何判断和跳转
	⚠️注意事项：变量必须先声明在使用，采用合法的向前引用
2）分枝控制
	分支控制if-else介绍
	让程序有达择的的执行，分支控制有三种：
	<1> 单分支
		if 条件表达式 {
			执行代码块
		}
		说明：当条件表达式为ture 时，就会执行{}的代码。

	<2> 双分支
		if 条件表达式 {
			执行代码块1
		}else{
			执行代码块2
		}
		说明：当条件表达式成立，即执行代码块1，否则执行代码块2。

	<3> 多分支
		if 条件表达式1 {
			执行代码块1
		}else if 条件表达式2 {
			执行代码块2
		}
			...
		else {
			执行代码块n
		}
		说明：当条件表达式1成立，即执行代码块1，当条件表达式2成立，即执行代码块2，以此类推；上述条件表达式都不成立，则执行代码块n。
			 最后的else可以省略

	<4> 分支嵌套
		在一个分支结构中又完整的嵌套了另一个完整的分支结构，里面的分支的结构称为内层分支外面的分支结构称为外层分支。
		if 条件表达式{
			if 条件表达式{

			}else{

			}
		说明：嵌套分支不易过多，最多三层

	<5> switch分支
		switch语句用于基于不同条件执行不同动作，每一个case分支都是唯一的从上到下逐一测试，直到匹配为止;匹配项后不需要加break。
		switch 表达式{
		case 表达式1,表达式2,…:
			语句块1
		case 表达式3,表达式4,…:
			语句块2
		   ...
		default:
			语句块
		}
		switch细节讨论
		1) case后是一个表达式(即：常量值、变量、一个有返回值的函数等都可以）
		2) case后的各个表达式的值的数据类型，必须和switch的表达式数据类型一致
		3) case后面可以带多个表达式，使用逗号间隔。比如case 表达式1，表达式2…
		4) case后面的表达式如果是常量值(字面量)，则要求不能重复
		5) case后面不需要带break，程序匹配到一个case后就会执行对应的代码块，然后退出switch，如果一个都匹配不到，则执行 default
		6) default 语句不是必须的.
		7) switch 后也可以不带表达式，类似if-else分支来使用。
		8) switch 后也可以直接声明/定义一个变量，分号结束，不推荐。
		9) switch 穿透-fallthrough，如果在case语句 块后增加fallthrough，则会继续执行下一个case，也叫switch穿透
		10) switch 语句还可以被用于 type-switch 来判断某个interface变量中实际指向的变量类型
*/

func main() {
	//单分支案例
	var age int
	fmt.Println("请输入你的年龄")
	fmt.Scanln(&age)
	if age > 18 {
		fmt.Printf("你的年龄为%v岁,因此你成年了\n", age)
	}

	//双分支案例
	var grade int
	fmt.Println("请输入你的分数")
	fmt.Scanln(&grade)
	if grade >= 60 {
		fmt.Printf("你的分数为%v。大于等于60分,因此你及格了\n", grade)
	} else {
		fmt.Printf("你的分数为%v。小于60分,因此你未及格了\n", grade)
	}

	//多分支案例
	var score int
	fmt.Println("请输入你的高考成绩")
	fmt.Scanln(&score)
	if score >= 500 {
		fmt.Printf("你的分数为%v。你的分数线大于等于500,恭喜你过一本线\n", score)
	} else if score > 450 {
		fmt.Printf("你的分数为%v。你的分数线大于等于450,恭喜你过二本线\n", score)
	} else if score > 400 {
		fmt.Printf("你的分数为%v。你的分数线大于等于400,恭喜你过三本线\n", score)
	} else {
		fmt.Printf("你的分数为%v。此分数只能上大专\n", score)
	}

	//嵌套分支
	var second float64
	fmt.Println("请输入秒数")
	fmt.Scanln(&second)
	if second <= 11 {
		var gender string
		fmt.Println("请输入性别")
		fmt.Scanln(&gender)
		if gender == "男" {
			fmt.Println("进入决赛的男子组")
		} else {
			fmt.Println("进入决赛的女子组")
		}
	} else {
		fmt.Println("out...")
	}

	// switch分枝
	var key string
	fmt.Println("请输入需要执行的操作[start|stop|status]")
	fmt.Scanln(&key)

	switch key {
	case "start":
		fmt.Println("bash start ...启动中...")
	case "stop":
		fmt.Println("bash stop ...停止中...")
	case "status":
		fmt.Println("bash status ...")
	default:
		fmt.Println("输入有误，请重新输入！！！")
	}

	//switch 的穿透 fal1throught
	var num int = 10
	switch num {
	case 10:
		fmt.Println("ok1")
		fallthrough //默认只能穿透一层
	case 20:
		fmt.Println("ok2")
		fallthrough
	case 30:
		fmt.Println("ok3")
	default:
		fmt.Println("没有匹配到")
	}

	// Type Switch： switch 语句还可以被用于 type-switch 来判断某个interface变量中实际指向的变量类型
	var x interface{}
	var y = 10.0
	x = y
	switch i := x.(type) {
	case nil:
		fmt.Printf("x 的类型~:%T", i)
	case int:
		fmt.Printf("x 是int 型")
	case float64:
		fmt.Printf("x 是float 型")
	case func(int) float64:
		fmt.Printf("x 是func(int)型")
	case bool, string:
		fmt.Printf("x 是bool 或string 型")
	default:
		fmt.Printf("未知型")
	}
}
