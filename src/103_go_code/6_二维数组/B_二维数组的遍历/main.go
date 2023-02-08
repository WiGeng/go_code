package main

import (
	"fmt"
)

/* 二维数组的遍历
1）方式1: 双层for循环
2）方式2: for-range遍历
*/

func main() {
	// 方式1: 双层for循环
	var twoArr = [...][3]int{{1, 2, 3}, {2, 3, 4}}
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			fmt.Print(twoArr[i][j], "       ")
		}
		fmt.Println()
	}

	// 方式2: for-range遍历
	for i, v := range twoArr {
		for j, v2 := range v {
			fmt.Printf("twoArr[%v][%v]=%v \t", i, j, v2)
		}
		fmt.Println()
	}
}
