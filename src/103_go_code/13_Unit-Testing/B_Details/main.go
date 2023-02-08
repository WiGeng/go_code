package main

/*
package testing
import "testing"
testing 提供对 Go 包的自动化测试的支持。通过 `go test` 命令，能够自动执行如下形式的任何函数：

""" func TestXxx(*testing.T) """  // 其中 Xxx 可以是任何字母数字字符串（但第一个字母不能是 [a-z]），用于识别测试例程。

在这些函数中，使用 Error, Fail 或相关方法来发出失败信号。
要编写一个新的测试套件，需要创建一个名称以 _test.go 结尾的文件，该文件包含 `TestXxx` 函数，如上所述。 将该文件放在与被测试
的包相同的包中。该文件将被排除在正常的程序包之外，但在运行 “go test” 命令时将被包含。 有关详细信息，请运行 “go help test”
和 “go help testflag” 了解。

*/

/* 单元测试快速入门总结：
细节说明：
	1）测试用例文件名必须以 _test.go结尾。比如cal_test.go , cal 不是固定的。
	2）测试用例函数必须以Test开头，一般来说就是Test+被测试的函数名，比如TestAddupper。
	3) TestAddupper(t *tesing.T) 的形参类型必须是 *testing.T
	4）一个测试用例文件中，可以有多个测试用例函数，比如 TestAddupper、Testsub
	5) 运行测试用例指令
		<1> cmd > go test [如果运行正确，无日志，错误时，会输出日志]
		<2> cmd > go test -v[ 运行正确或是错误，都输出日志]
	6) 当出现错误时，可以使用t.Fatalf 来格式化输出错误信息，并退出程序
	7） t.Logf 方法可以输出相应的日志
	8）测试用例函数，并没有放在main函数中，也执行了，这就是测试用例的方便之处
	9) PASS 表示测试用例运行成功，FAIL 表示测试用例运行失败
	10) 测试单个文件，一定要带上被测试的原文件
	 	go test -v cal_test.go cal.go
	11) 测试单个方法
		go test -v -test.run TestAddUpper

*/
