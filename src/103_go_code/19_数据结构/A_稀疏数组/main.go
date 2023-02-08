package main

import (
	"fmt"
)

/* 稀疏数组
基本介绍：
	当一个数组中大部分元素为0，或者为同一个值的数组时，可以使用稀疏数组来保存该数组。

稀疏数组的处理方法是：
	1) 记录数组一共有几行几列，有多少个不同的值
	2）把具有不同值的元素的行列及值记录在一个小规模的数组中，从而缩小程序的规模

举例说明：
							 ———————————————————
——			 ——					 | 行 | 列 | 值 |
｜ 0 0 8 0 0 ｜				  [0] | 0 | 2  | 8 |
｜ 0 1 0 0 7 ｜				  [1] | 1 | 1  | 1 |
｜ 0 0 0 0 0 ｜          	  [2] | 1 | 4  | 7 |
｜ 2 0 0 3 0 ｜		 --->	  [3] | 2 | 0  | 2 |
｜ 0 0 0 0 0 ｜				  [4] | 2 | 3  | 3 |
｜ 0 0 4 0 0 ｜				  [5] | 5 | 2  | 4 |
｜ 0 0 0 0 5 ｜				  [n] | 6 | 4  | 5 |
——			—— 				 ———————————————————

*/

/* 案例：
1）使用稀疏数组，来保留类似前面的二维数组（棋盘、地图等等）
2）把稀疏数组存盘，并且可以从新恢复原来的二维数组
3）整体思路分析
*/

type ValNode struct {
	row int
	col int
	val int
}

func qipan01() {

	//1. 先创建一个原始数组
	var chessMap [11][11]int
	chessMap[1][2] = 1 //黑子
	chessMap[2][3] = 2 //蓝子

	//2. 输出看看原始的数组
	for _, v := range chessMap {
		for _, v2 := range v {
			fmt.Printf("%d\t", v2)
		}
		fmt.Println()
	}

	//3. 转成稀疏数组。想-> 算法
	//(1). 遍历 chessMap, 如果我们发现有一个元素的值不为0，创建一个node结构体
	//(2). 将其放入到对应的切片即可

	var sparseArr []ValNode
	//标准的一个稀疏数组应该还有一个 记录元素的二维数组的规模(行和列，默认值)
	//创建一个ValNode 值结点
	valNode := ValNode{
		row: 11,
		col: 11,
		val: 0,
	}
	sparseArr = append(sparseArr, valNode)
	for i, v := range chessMap {
		for j, v2 := range v {
			if v2 != 0 {
				//创建一个ValNode 值结点
				valNode := ValNode{
					row: i,
					col: j,
					val: v2,
				}
				sparseArr = append(sparseArr, valNode)
			}
		}
	}
	//输出稀疏数组
	fmt.Println("当前的稀疏数组是:::::")
	for i, valNode := range sparseArr {
		fmt.Printf("%d: %d %d %d\n", i, valNode.row, valNode.col, valNode.val)
	}
	// 先创建一个原始数组
	var chessMap2 [11][11]int
	// 遍历 sparseArr [遍历文件每一行]
	for i, valNode := range sparseArr {
		if i != 0 { //跳过第一行记录值
			chessMap2[valNode.row][valNode.col] = valNode.val
		}
	}
	// 看看chessMap2 是不是恢复.
	fmt.Println("恢复后的原始数据......")
	for _, v := range chessMap2 {
		for _, v2 := range v {
			fmt.Printf("%d\t", v2)
		}
		fmt.Println()
	}
}

func main() {
	qipan01()

}
