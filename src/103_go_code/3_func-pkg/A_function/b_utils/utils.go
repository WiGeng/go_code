package utils

import (
	"fmt"
)

//输入两个数，在输入一个运算符，得到运算结果
func Cal(n1 float64, n2 float64, operator byte) float64 {
	var res float64
	switch operator {
	case '+':
		res = n1 + n2
	case '-':
		res = n1 - n2
	case '*':
		res = n1 * n2
	case '/':
		res = n1 / n2
	default:
		fmt.Println("操作符号输入错误～")
	}
	return res
}
