package main

import (
	"fmt"
	"math/rand"
	"time"
)

/* 插入排序
插入排序法介绍：
	插入式排序属于内部排序法，是对于欲排序的元素以插入的方式找寻该元素的适当位置，以达到排序的目的

插入排序法思想：
	插入排序 (Insertion Sorting）的基本思想是：把n个待排序的元素看成为一个有序表和一个无序表，开始时有序表中只包含一个元素，
	无序表中包含有n-1个元素，排序过程中每次从无序表中取出第一个元素，把它的排序码依次与有序表元素的排序码进行比较，将它插入到
	有序表中的适当位置，使之成为新的有序表
*/

func InsertSort(arr *[8]int) {

	//完成第一次，给第二个元素找到合适的位置并插入

	for i := 1; i < len(arr); i++ {

		insertVal := arr[i]
		insertIndex := i - 1 // 下标

		//从大到小
		for insertIndex >= 0 && arr[insertIndex] < insertVal {
			arr[insertIndex+1] = arr[insertIndex] // 数据后移
			insertIndex--
		}
		//插入
		if insertIndex+1 != i {
			arr[insertIndex+1] = insertVal
		}
		fmt.Printf("\t第%d次插入后 %v\n", i, *arr)
	}

	/*

		//完成第2次，给第3个元素找到合适的位置并插入
		insertVal = arr[2]
		insertIndex = 2 - 1 // 下标

		//从大到小
		for insertIndex >= 0 && arr[insertIndex] < insertVal {
			arr[insertIndex + 1] = arr[insertIndex] // 数据后移
			insertIndex--
		}
		//插入
		if insertIndex + 1 != 2 {
			arr[insertIndex + 1] = insertVal
		}
		fmt.Println("第2次插入后", *arr)

		//完成第3次，给第4个元素找到合适的位置并插入
		insertVal = arr[3]
		insertIndex = 3 - 1 // 下标

		//从大到小
		for insertIndex >= 0 && arr[insertIndex] < insertVal {
			arr[insertIndex + 1] = arr[insertIndex] // 数据后移
			insertIndex--
		}
		//插入
		if insertIndex + 1 != 3 {
			arr[insertIndex + 1] = insertVal
		}
		fmt.Println("第3次插入后", *arr)

		//完成第4次，给第5个元素找到合适的位置并插入
		insertVal = arr[4]
		insertIndex = 4 - 1 // 下标

		//从大到小
		for insertIndex >= 0 && arr[insertIndex] < insertVal {
			arr[insertIndex + 1] = arr[insertIndex] // 数据后移
			insertIndex--
		}
		//插入
		if insertIndex + 1 != 4 {
			arr[insertIndex + 1] = insertVal
		}
		fmt.Println("第4次插入后", *arr)*/
}

func main() {

	//arr := [7]int{23, 0, 12, 56,  34, -1, 55}

	var arr [8]int
	for i := 0; i < 8; i++ {
		arr[i] = rand.Intn(90)
	}

	fmt.Println("排序前数组：", arr)
	start := time.Now().Unix()
	InsertSort(&arr)
	end := time.Now().Unix()
	fmt.Println("排序后数组：", arr)
	fmt.Printf("main函数~ 选择排序总耗时=%d秒\n", end-start)
}
