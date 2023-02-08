package main

import (
	"fmt"
	"reflect"
)

/* 给你一个变量 var score float64 = 145.5, 请使用反射来得到它的reflect.Value，然后获取对应的Type, Kind 和值，
并将reflect.Value转换成interfacef{}，再将interface{}转换成float64，并修改为150.0
*/

func reflectTest(b interface{}) {
	//获取到 reflect.Type
	tVal := reflect.TypeOf(b)
	// 获取到 reflect.Value
	rVal := reflect.ValueOf(b)
	fmt.Printf("rVal.Type= %v\nrVal.Value= %v\n", tVal, rVal)

	//将rVal的reflect.Value数据类型转换为Interface{}
	iV := rVal.Interface()
	res := iV.(float64)
	fmt.Println("interface{}转换成float64=", res)

}

func main() {
	var score float64 = 145.5
	reflectTest(score)
}
