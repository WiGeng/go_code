package main

import "fmt"

/* map
介绍：
	map是key-value的数据结构，又称为字段或者关联数组，类似于其他编程语言的集合
基本语法
	1）map的声明：
		var map变量名 map[keytype]valuetype
	2）key 可以是什么类型?
		golang中map的key 可以是很多种类型，比如bool、数字、string、指针、channel，还可以是只包含前面几个类型的
			接口、结构体、数组。通常为int、string
		注意：slice、map 、function不可以，因为这几个没法用 == 来判断
	3）valuetype 可以是什么类型
		valuetype的类型和key基本一样，这里我就不再赘述了
		通常为：数字(整数,浮点数),string.map,struct
	4）map的声明举例：
		var a map[string]string
		var a map[string]int
		var a map[int]string
		var a map[string]map[string]string
		注意：声明是不分配内存的，初始化需要make,分配内存后才能赋值和使用
*/

func main() {
	//map的声明和注意事项
	var a map[string]string
	//在使用map前，需要先make , make的作用就是给map分配数据空间
	a = make(map[string]string, 10)
	a["no1"] = "宋江"
	a["no2"] = "吴用"
	a["no1"] = "武松"
	a["no3"] = "吴用"
	fmt.Println(a)
	/* 总结：对上面代码的说明
	1）map 在使用前一定要 make
	2）map的key 是不能重复，如果重复了，则以最后这个key-value为准
	3）map 的value 是可以相同
	4）map 的 key-value 是无序
	*/
}
