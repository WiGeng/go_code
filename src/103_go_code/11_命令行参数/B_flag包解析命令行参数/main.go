package main

import (
	"flag"
	"fmt"
)

/* 通过flag包来完成指定参数来获取参数的值
func String：
	func String(name string, value string, usage string) *string：String用指定的名称、默认值、使用信息注册一个string类型flag。返回一个保存了该flag的值的指针。
func IntVar：
	func IntVar(p *int, name string, value int, usage string)：IntVar用指定的名称、默认值、使用信息注册一个int类型flag，并将flag的值保存到p指向的变量。
func Parse：
	func Parse()：从os.Args[1:]中解析注册的flag。必须在所有flag都注册好而未访问其值时执行。未注册却使用flag -help时，会返回ErrHelp。
*/

func main() {

	//定义几个变量，用于接收命令行的参数值
	var user string
	var pwd string
	var host string
	var port int

	//&user 就是接收用户命令行中输入的 -u 后面的参数值
	//"u" ,就是 -u 指定参数
	//"" , 默认值
	//"用户名,默认为空" 说明
	flag.StringVar(&user, "u", "", "用户名:默认为空")
	flag.StringVar(&pwd, "pwd", "", "密码:默认为空")
	flag.StringVar(&host, "h", "localhost", "主机名:默认为localhost")
	flag.IntVar(&port, "port", 3306, "端口号:默认为3306")

	flag.Parse() //重要的操作 - 转换， 必须调用该方法

	//输出结果
	fmt.Printf("user=%v pwd=%v host=%v port=%v",
		user, pwd, host, port)

}
