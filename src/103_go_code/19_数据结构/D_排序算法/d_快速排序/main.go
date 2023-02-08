package main

import (
	"fmt"
	"math/rand"
	"time"
)

/* 快速排序
快速排序法介绍：
	快速排序（Quicksort）是对冒泡排序的一种改进。基本思想是：通过一趟排序将要排序的数据分割成独立的两部分，其中一部分的所有数据都
	比另外一部分的所有数据都要小，然后再按此方法对这两部分数据分别进行快速排序，整个排序过程可以递归进行，以此达到整个数据变成有序
	序列
*/

/* 快速排序法应用实例：
	要求：对 [-9,78,0,23,-567,70]进行从小到小的排序，要求使用快速排序法
＞说明[验证分析〕：
	1) 如果职消左右递归，结果是 [-9 -567 0 23 78 70]
	2) 如果取消右递归,结果是[-567 -9 0 23 78 70]
	3) 如果职消左递归,結果是[-9 -567 0 23 70 78]
*/

//快速排序
//说明
//1. left 表示 数组左边的下标
//2. right 表示数组右边的下标
//3  array 表示要排序的数组
func QuickSort(left int, right int, array *[20]int) {
	l := left
	r := right
	// pivot 是中轴， 支点
	pivot := array[(left+right)/2]
	temp := 0

	//for 循环的目标是将比 pivot 小的数放到 左边
	//  比 pivot 大的数放到 右边
	for l < r {
		//从  pivot 的左边找到大于等于pivot的值
		for array[l] < pivot {
			l++
		}
		//从  pivot 的右边边找到小于等于pivot的值
		for array[r] > pivot {
			r--
		}
		// 1 >= r 表明本次分解任务完成, break
		if l >= r {
			break
		}
		//交换
		temp = array[l]
		array[l] = array[r]
		array[r] = temp
		//优化
		if array[l] == pivot {
			r--
		}
		if array[r] == pivot {
			l++
		}
	}
	// 如果  1== r, 再移动下
	if l == r {
		l++
		r--
	}
	// 向左递归
	if left < r {
		QuickSort(left, r, array)
	}
	// 向右递归
	if right > l {
		QuickSort(l, right, array)
	}
}

func main() {

	// arr := [9]int {-9,78,0,23,-567,70, 123, 90, -23}
	// fmt.Println("初始数组", arr)

	var arr [20]int
	for i := 0; i < 20; i++ {
		arr[i] = rand.Intn(90)
	}

	fmt.Println("排序前数组：", arr)
	start := time.Now().Unix()
	QuickSort(0, len(arr)-1, &arr)
	end := time.Now().Unix()
	fmt.Println("排序后数组：", arr)
	fmt.Printf("main函数~ 选择排序总耗时=%d秒\n", end-start)

}
