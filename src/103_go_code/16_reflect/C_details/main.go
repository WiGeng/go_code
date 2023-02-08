package main

import (
	"fmt"
	"reflect"
)

/* 反射的注意事项和细节
1) reflect Value.Kind，获取变量的类别，返回的是一个常量
2) Type是类型，Kind是类别，Type 和 Kind 可能是相同的，也可能是不同的.
	比如：var num int = 10 num的Type是int, Kind也是int
	比如：var stu Student stu的Type是 包名.Student, Kind是struct
3） 通过反射可以在让 变量 在interface{} 和Reflect.Value之间相互转换
		变量  <----> interface{}  <----> Reflect.Value
4） 使用反射的方式来获取变量的值(并返回对应的类型），要求数据类型匹配，比如x是int,那么就应该使用 reflect.Value(x).Int()，
    而不能使用其它的，否则报panic
5） 通过反射来修改变量，注意当使用SetXxx方法来设置需要通过对应的指针类型来完成，这样才能改变传入的变量的值，同时需要使用到
	reflect.Value.Elem()方法
*/

//通过反射，修改,
// num int 的值
// 修改 student的值

func reflect01(b interface{}) {
	//2. 获取到 reflect.Value
	rVal := reflect.ValueOf(b)
	// 看看 rVal的Kind是
	fmt.Printf("rVal kind=%v\n", rVal.Kind())
	//3. rVal
	//Elem返回v持有的接口保管的值的Value封装，或者v持有的指针指向的值的Value封装
	rVal.Elem().SetInt(20)
}

func main() {

	var num int = 10
	reflect01(&num)
	fmt.Println("num=", num) // 20

	/* 可以这样理解rVal.Elem() */
	// num := 9
	// ptr *int = &num
	// num2 := *ptr  //=== 类似 rVal.Elem()
}
