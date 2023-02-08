package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
选择排序基本介绍：
	选择式排序也属于内部排序法，是从欲排序的数据中，按指定的规则选出某一元素，经过和其他元素重整，再依原则交换位置后达到排序的目的

选择排序思想：
	选择排序(select sorting）也是一种简单的排序方法。它的基本思想是：第一次从R[O]~R[n-1]中选取最小值，与R[OJ交换，第二次从
	R[1]～R[n-1]中选取最小值，与R[1]交换，第三次从R[2]~R[n-1]中选取最小值，与R[2]交换，⋯，第i次从R[i-1]~R[n-1]中选取最小
	值，与R[i-1]交换，⋯，第n-1次从R[n-2]~R[n-1]中选取最小值，与R[n-2]交换，总共通过n-1次，得到一个按排序码从小到大排列的有
	序序列。
*/

/* 选择排序应用实例：有一群牛，颜值分别是 10分，34分，19分，100分，80分 请使用选择排序从高到低进行排序 */

//编写函数selectSort 完成排序

func SelectSort(arr *[8]int) {

	//标准的访问方式
	//(*arr)[1] = 600 等价于 arr[1] = 900
	//arr[1] = 900
	//1. 先完成将第一个最大值和 arr[0] => 先易后难

	//1 假设  arr[0] 最大值

	for j := 0; j < len(arr)-1; j++ {

		max := arr[j]
		maxIndex := j
		//2. 遍历后面 1---[len(arr) -1] 比较
		for i := j + 1; i < len(arr); i++ {
			if max < arr[i] { //找到真正的最大值
				max = arr[i]
				maxIndex = i
			}
		}
		//交换
		if maxIndex != j {
			arr[j], arr[maxIndex] = arr[maxIndex], arr[j]
		}

		fmt.Printf("\t第%d次 %v\n", j+1, *arr)
	}

	/*
		max = arr[1]
		maxIndex = 1
		//2. 遍历后面 2---[len(arr) -1] 比较
		for i := 1 + 1; i < len(arr); i++ {
			if max < arr[i] { //找到真正的最大值
				max = arr[i]
				maxIndex = i
			}
		}
		//交换
		if maxIndex != 1 {
			arr[1], arr[maxIndex] = arr[maxIndex], arr[1]
		}

		fmt.Println("第2次 ", *arr)



		max = arr[2]
		maxIndex = 2
		//2. 遍历后面 3---[len(arr) -1] 比较
		for i := 2 + 1; i < len(arr); i++ {
			if max < arr[i] { //找到真正的最大值
				max = arr[i]
				maxIndex = i
			}
		}
		//交换
		if maxIndex != 2 {
			arr[2], arr[maxIndex] = arr[maxIndex], arr[2]
		}

		fmt.Println("第3次 ", *arr)

		max = arr[3]
		maxIndex = 3
		//2. 遍历后面 4---[len(arr) -1] 比较
		for i := 3 + 1; i < len(arr); i++ {
			if max < arr[i] { //找到真正的最大值
				max = arr[i]
				maxIndex = i
			}
		}
		//交换
		if maxIndex != 3 {
			arr[3], arr[maxIndex] = arr[maxIndex], arr[3]
		}

		fmt.Println("第4次 ", *arr)*/
}

func main() {
	//定义一个数组 , 从大到小
	//arr := [6]int{10, 34, 19, 100, 80}

	var arr [8]int
	for i := 0; i < 8; i++ {
		arr[i] = rand.Intn(900)
	}

	fmt.Println("排序前数组：", arr)
	start := time.Now().Unix()
	SelectSort(&arr)
	end := time.Now().Unix()
	fmt.Println("排序后数组：", arr)
	fmt.Printf("main函数~ 选择排序总耗时=%d秒\n", end-start)

}
