package main

import "fmt"

/* 切片注意事项
细节说明:
	1)切片初始化时 var slice = arr[startindex:endlndex]
		说明：从arr数组下标为startindex，取到下标为endindex的元素(不含arr(endIndex]).
	2)切片初始化时，仍然不能越界。范围在 [O-len(arr)]之间，但是可以动态增长．
		a) var slice = arr[0:end]可以简写 var slice = arr[:end]
		b) var slice = arr[start:len(arr]可以简写：var slice =arr[start:]
		c) var slice = arr[0:len(arr)]可以简写：var slice = arr[:]
	3)cap是一个内置函数，用于统计切片的容量，即最大可以存放多少个元素。
	4)切片定义完后，还不能使用，因为本身是一个空的，需要让其引用到一个数组，或者make一个空间供切片来使用
	5)切片可以继续切片
	6)使用append内置函数，可以对切片进行动态追加
		切片append操作的底层原理分析：
		a) 切片append操作的本质就是对数组扩容
		b) go底层会创建一下新的数组newArr(安装扩容后大小)
		c) 将slice原来包含的元素拷贝到新的数组newArr
		d) slice重新引用到newArr
		e) 注意newArr是在底层来维护的，程序员不可见
	7)切片的拷贝：切片使用copy内置函数完成拷贝
		a) copy(para1,para2)参数的数据类型时切片
		b) 两个切片的数据空间是独立的，相互不影响
		c）切片的拷贝不存在溢出



*/

func main() {
	var arr [5]int = [...]int{10, 20, 30, 40, 50}
	slice := arr[1:4]
	fmt.Println("slice=", slice)

	//切片可以继续切片
	slice1 := slice[0:2]
	fmt.Println("slice1=", slice1)

	//用append内置函数，可以对切片进行动态追加
	var slice3 []int = []int{100, 200, 300}
	//通过append直接给slice3追加具体的元素
	slice3 = append(slice3, 400, 500, 600)
	fmt.Println("slice3", slice3) //100, 200, 300,400, 500, 600

	//通过append将切片slice3追加给slice3
	slice3 = append(slice3, slice3...) // 100, 200, 300,400, 500, 600 100, 200, 300,400, 500, 600
	fmt.Println("slice3", slice3)

	//切片的拷贝操作
	//切片使用copy内置函数完成拷贝
	fmt.Println()
	var slice4 []int = []int{1, 2, 3, 4, 5}
	var slice5 = make([]int, 10)
	copy(slice5, slice4)
	fmt.Println("slice4=", slice4) // 1, 2, 3, 4, 5
	fmt.Println("slice5=", slice5) // 1, 2, 3, 4, 5, 0, 0, 0, 0, 0
}
