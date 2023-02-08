package main

import (
	"fmt"
)

/* Map
1. 使用方式介绍：
	1）方式1:
		var a map[string]string
		a = make(map[string]string, 10)
		a["name"] = "woke"

	2）方式2：
		a1 := make(map[string]string)
		a1["name"] = "weigeng"

	3）方式3:
		a2 := map[string]string{
		"name": "wiger",
		}

2. map的增删改查操作：
	map增加和更新：
		maprkeyy=valuve 1/如果key 还没有，就是增加，如果key存在就是修改。
	map删除：
		delete(map,"key")， delete是一个内置函数，如果key存在，就删除该key-value,如果key不存在，不操作，但是也不会报错
		＞细节说明
		1) 如果我们要制除map的所有ey，没有一个专门的方法一次删除，可以遍历一下key， 逐个删除
		2) 或者 map =  make(...)，make一个新的，让原来的成为垃圾，被gc回收
	map查找：
		说明：如果一个map中存在某个key,那么就会返回一个true，否则返回一个false

3. map的遍历
	map的遍历是使用for-range的结构进行遍历

4. 统计map的长度
	len
*/

func main() {
	/*
		1. mapde使用方式：
	*/

	// 方式1
	var a map[string]string
	// 在使用map前，需要先make , make的作用就是给map分配数据空间
	a = make(map[string]string, 10)
	a["name"] = "woke"
	fmt.Println(a)

	//方式2
	a1 := make(map[string]string)
	a1["name"] = "weigeng"
	fmt.Println(a1)

	//方式3
	a2 := map[string]string{
		"name": "wiger",
	}
	fmt.Println(a2)

	/*
		练习：演示一个key-value 的value是map的案例
		比如：我们要存放3个学生信息, 每个学生有 name和sex 信息
		思路:   map[string]map[string]string
	*/
	StudentMap := make(map[string]map[string]string)
	// 方式1
	StudentMap["stu01"] = make(map[string]string, 3)
	StudentMap["stu01"]["name"] = "woke"
	StudentMap["stu01"]["sex"] = "f"
	StudentMap["stu01"]["age"] = "26"
	fmt.Println(StudentMap["stu01"])
	// 方式2
	StudentMap["stu02"] = make(map[string]string)
	StudentMap["stu02"]["name"] = "mengqing"
	StudentMap["stu02"]["sex"] = "m"
	StudentMap["stu02"]["age"] = "24"
	fmt.Println(StudentMap["stu02"])

	/*
		2. map的增删改查
	*/
	//1）map的增加
	cities := make(map[string]string)
	cities["no1"] = "北京"
	cities["no2"] = "天津"
	cities["no3"] = "上海"
	fmt.Println(cities)

	//2）map的修改，因为 no3这个key已经存在
	cities["no3"] = "上海~"
	fmt.Println(cities)

	//3）map的删除
	delete(cities, "no1")
	fmt.Println(cities)
	//当delete指定的key不存在时，删除不会操作，也不会报错
	delete(cities, "no4")
	fmt.Println(cities)

	//如果希望一次性删除所有的key
	//a.遍历所有的key,如何逐一删除 [遍历]
	//b.直接make一个新的空间
	cities = make(map[string]string)
	fmt.Println(cities)

	//4）map的查找
	val, ok := cities["no2"]
	if ok {
		fmt.Printf("有no1 key 值为%v\n", val)
	} else {
		fmt.Printf("没有no1 key\n")
	}

	/*
		map的遍历
	*/
	cities1 := make(map[string]string)
	cities1["no1"] = "西安"
	cities1["no2"] = "杭州"
	cities1["no3"] = "成都"
	for k, v := range cities1 {
		fmt.Printf("Key = %v , Value = %v\n", k, v)
	}

	StudentMap1 := make(map[string]map[string]string)
	StudentMap1["stu01"] = make(map[string]string, 3)
	StudentMap1["stu01"]["name"] = "woke"
	StudentMap1["stu01"]["sex"] = "m"
	StudentMap1["stu01"]["age"] = "26"

	StudentMap1["stu02"] = make(map[string]string)
	StudentMap1["stu02"]["name"] = "mengqing"
	StudentMap1["stu02"]["sex"] = "f"
	StudentMap1["stu02"]["age"] = "24"

	for k1, v1 := range StudentMap {
		fmt.Printf("k1=%v\n", k1)
		for k2, v2 := range v1 {
			fmt.Printf("  k2=%v,v2=%v", k2, v2)
		}
		fmt.Println()
	}

	/*
		map的长度统计
	*/
	b := len(cities1)
	fmt.Printf("map_citices有 %v 对儿 key-value", b)
}
