package main

/* 工厂模式
说明: Golang的结构体没有构造函数，通常可以使用工厂模式来解决这个问题。
*/

import (
	"awesomeProject/src/103_go_code/8_object-oriented/C_oop/b_Factory-mode/model"
	"fmt"
)

func main() {

	//定student结构体是首字母小写，我们可以通过工厂模式来解决
	var stu = model.NewStudent("tom~", 98.8)

	fmt.Println(*stu) //&{....}
	fmt.Println("name=", stu.Name, " score=", stu.GetInfo())
}
