package main

import (
	"errors"
	"fmt"
)

/* 错误处理
基本说明
	1）Go语言追求简洁优雅，所以，Go语言不支持传统的try...catch...finally这种处理
	2）go中引入的处理方式为：defer, panic, recover
	3）这几个异常的使用场景可以这么简单描述：Go中可以抛出一个panic的异常，然后在defer中通过recover捕获这个异常，最后正常处理

错误处理的好处：
	进行错误处理，程序不会立即挂掉，如果加入预警代码，则可以使程序更加健壮

自定义错误：
Go程序中，也支持自定义错误，使用errors.New 和 panic 内置函数。
	1) errors.New("错误说明")，会返回一个error类型的值，表示一个错误
	2) panic内置函数，接收一个interface{}类型的值（也就是任何值了）作为参数。可以接收error类型的变量，输出错误信息，并退出程序.
*/

func test() {
	//使用defer + recover 来捕获和处理异常
	defer func() {
		err := recover() // recover()内置函数，可以捕获到异常
		if err != nil {  // 说明捕获到错误
			fmt.Println("err=", err)
			//这里就可以将错误信息发送给管理员....
			fmt.Println("发送邮件给admin@sohu.com~")
		}
	}()
	num1 := 10
	num2 := 0
	res := num1 / num2
	fmt.Println("res=", res)
}

//函数去读取以配置文件init.conf的信息
//如果文件名传入不正确，我们就返回一个自定义的错误
func readConf(name string) (err error) {
	if name == "config.ini" {
		//读取...
		return nil
	} else {
		//返回一个自定义错误
		return errors.New("读取文件错误..")
	}
}

func test02() {
	err := readConf("config1.ini")
	if err != nil {
		//如果读取文件发送错误，就输出这个错误，并终止程序
		panic(err)
	}
	fmt.Println("test02()继续执行....")
}

func main() {

	//测试
	// test()
	// for {
	// 	fmt.Println("main()下面的代码...")
	// 	time.Sleep(time.Second)
	// }

	//测试自定义错误的使用

	test02()
	fmt.Println("main()下面的代码...")
}
