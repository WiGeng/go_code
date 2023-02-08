package cal

//一个被测试函数
func addUpper(n int) int {
	res := 0
	for i := 1; i <= n; i++ {
		res += i
	}
	return res
}

//求两个数的查
func getSub(n1 int, n2 int) int {
	return n1 - n2
}

/* 执行测试用例步骤：
(1) 进入单元测试用例所在目录[/Users/weigeng/go/src/VsCode/awesomeProject/src/103_go_code/13_Unit-Testing/A_introduce/b_testing]
(2) 执行命令：
[weigeng@WEIGENGdeMacBook-Pro b_testing % go test -v
	=== RUN   TestAddUpper
		cal_test.go:19: AddUpper(10) 执行正确...
	--- PASS: TestAddUpper (0.00s)
	=== RUN   TestHello
	TestHello被调用..
	--- PASS: TestHello (0.00s)
	=== RUN   TestGetSub
		sub_test.go:18: getSub(10, 3) 执行正确!!!!...
	--- PASS: TestGetSub (0.00s)
	PASS
	ok  	awesomeProject/src/103_go_code/13_Unit-Testing/A_introduce/b_testing	0.423s

*/
