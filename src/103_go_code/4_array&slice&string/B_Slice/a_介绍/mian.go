package main

import "fmt"

/* 切片
基本介绍：
	1) 切片的英文是slics
	2) 切片是数组的一个引用，因此切片是引用类型，在进行传递时，遵守引用传递的机制。
	3) 切片的使用和数组类似，遍历切片、访问切片的元素和求切片长度len(slice)都一样。
	4) 切片的长度是可以变化的，因此切片是一个可以动态变化数组。
	5) 切片定义的基本语法：
		var 变量名口类型
		比如： var a []int
内存排布：
	1.slice 的确是一个引用类型
	2.slice 从底层来说，其实就是一个数据结构(struct 结构体）
		type slice struct {
			ptr *[2]int
			len int
			cap int
		}
切片使用的三种方式：
	方式1
		定义一个切片，然后让切片去引用一个已经创建好的数组，比如前面的案例就是这样的。
	方式2
		通过make来创建切片.
		基本语法：var 切片名 []type = make([]type, len, [cap])
		參数说明：type: 就是数据类型 len：大小 cap：指定切片容量[可选]
	方式3
		定义一个切片，直接就指定具体数组，使用原理类似make方式

	方式1和方式2的区别【面试】：
		方式1是直接引用数组，这个数组是事先存在的，程序员是可见的。
		方式2是通过make 来创建切片，make也会创建一个数组，是由切片在底层进行维护，程序员是看不见的。
切片的遍历：
	1) for 循环常规方式遍历
	2) for-range 结构遍历切片
*/

func main() {
	fmt.Println("-----------------------切片介绍------------------------------")
	//演示切片的基本使用
	var intArr [5]int = [...]int{1, 22, 33, 66, 99}
	//声明/定义一个切片
	//slice := intArr[1:3]
	//1. slice 就是切片名
	//2. intArr[1:3] 表示 slice 引用到intArr这个数组
	//3. 引用intArr数组的起始下标为 1 , 最后的下标为3(但是不包含3)
	slice := intArr[1:3]
	fmt.Println("intArr=", intArr)
	fmt.Println("slice 的元素是 =", slice)       //  22, 33
	fmt.Println("slice 的元素个数 =", len(slice)) // 2
	fmt.Println("slice 的容量 =", cap(slice))   // 切片的容量是可以动态变化

	fmt.Printf("intArr[1]的地址=%p\n", &intArr[1])
	fmt.Printf("slice[0]的地址=%p slice[0==%v\n", &slice[0], slice[0])
	slice[1] = 34
	fmt.Println()
	fmt.Println("intArr=", intArr)     //  1 22 34 66 99
	fmt.Println("slice 的元素是 =", slice) //  22, 34

	fmt.Println("-----------------------切片创建------------------------------")
	/*切片的基本使用*/
	// 方式1:
	var intArr1 [5]int = [...]int{1, 22, 33, 66, 99}
	slice1 := intArr1[1:3]
	fmt.Println("slice1=", slice1)

	//方式2:
	var slice2 []int = make([]int, 5, 10)
	fmt.Println("初始的slice2=", slice2)
	slice2[0] = 10
	slice2[2] = 20
	fmt.Println("赋值后slice2=", slice2)
	fmt.Println("slice2 size=", len(slice2))
	fmt.Println("slice2 cap=", cap(slice2))

	//方式3:
	var slice3 []string = []string{"tom", "jack", "mary"}
	fmt.Println("slice3=", slice3)
	fmt.Println("slice3 size=", len(slice3)) //3
	fmt.Println("slice3 cap=", cap(slice3))  //3

	//切片遍历
	fmt.Println("-----------------------切片遍历------------------------------")
	//使用常规的for循环遍历切片
	var arr [5]int = [...]int{10, 20, 30, 40, 50}
	//slice := arr[1:4] // 20, 30, 40
	slice4 := arr[1:4]
	for i := 0; i < len(slice4); i++ {
		fmt.Printf("slice4[%v]=%v \n", i, slice4[i])
	}

	fmt.Println()
	//使用for--range 方式遍历切片
	for i, v := range slice4 {
		fmt.Printf("i=%v v=%v \n", i, v)
	}
}
