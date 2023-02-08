package main

import (
	"fmt"
	"reflect"
)

//专门演示反射
func reflectTest(b interface{}) {

	//通过反射获取的传入的变量的 type , kind, 值
	//1. 先获取到 reflect.Type
	rTyp := reflect.TypeOf(b)
	fmt.Println("rType=", rTyp)

	//2. 获取到 reflect.Value
	rVal := reflect.ValueOf(b)
	fmt.Printf("rVal=%v rVal type=%T\n", rVal, rVal)
	//3. 将rVal的reflect.Value数据类型转换为int
	n2 := 2 + rVal.Int()
	fmt.Println("n2=", n2)
	// n3 := 2 + rVal.Float()   //要求数据类型匹配,否则报错
	// fmt.Println("n3=", n3)

	//下面将 rVal 转成 interface{}
	iV := rVal.Interface()
	//将 interface{} 通过断言转成需要的类型
	num2 := iV.(int)
	fmt.Printf("num2=%v num2 type=%T\n", num2, num2)
}

func main() {

	//请编写一个案例，
	//演示对(基本数据类型、interface{}、reflect.Value)进行反射的基本操作
	//1. 先定义一个int
	var num int = 100
	reflectTest(num)

}
