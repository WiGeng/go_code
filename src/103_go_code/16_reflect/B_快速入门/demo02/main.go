package main

import (
	"fmt"
	"reflect"
)

//专门演示反射[对结构体的反射]
func reflectTest(b interface{}) {
	//通过反射获取的传入的变量的 type , kind, 值
	//1. 先获取到 reflect.Type
	rTyp := reflect.TypeOf(b)
	fmt.Println("rType=", rTyp)

	//2. 获取到 reflect.Value
	rVal := reflect.ValueOf(b)
	fmt.Println(rVal)

	//3. 获取 变量对应的Kind
	//(1) rVal.Kind() ==>
	kind1 := rVal.Kind()
	//(2) rTyp.Kind() ==>
	kind2 := rTyp.Kind()
	fmt.Printf("kind = %v  kind = %v\n", kind1, kind2)

	//下面我们将 rVal 转成 interface{}
	iV := rVal.Interface()
	fmt.Printf("iV = %v  iV type = %T \n", iV, iV)
	//将 interface{} 通过断言转成需要的类型
	//这里就简单使用了一带检测的类型断言
	//可以使用 swtich 的断言形式来做的更加的灵活 (参考: 103_go_code/8_object-oriented/H_类型断言/c_最佳实践)
	stu, ok := iV.(Student)
	if ok {
		fmt.Printf("stu.Name=%v\n", stu.Name)
	}
}

type Student struct {
	Name string
	Age  int
}

type Monster struct {
	Name string
	Age  int
}

func main() {
	//1. 定义一个Student的实例
	stu := Student{
		Name: "tom",
		Age:  20,
	}
	reflectTest(stu)
}
