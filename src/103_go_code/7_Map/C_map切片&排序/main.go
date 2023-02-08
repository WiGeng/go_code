package main

import (
	"fmt"
	"sort"
)

/* map切片&排序
切片介绍：
	切片的数据类型如果是map，就称之为slice of map [map切片],这样使用则map的个数就可以动态变化了
排序介绍：
	1) golang中没有一个专门的方法针对map的key进行排序
	2) golang中的map默认是无序的，注意也不是按照添加的顺序存放的，每次遍历，得到的输出可能不一样
	3) golang中map的排序，是先将key进行排序，然后根据key值遍历输出即可
*/

func main() {
	/*
		演示map切片的使用

		要求：使用一个map来记录monster的信息 name 和 age, 也就是说一个
		monster对应一个map,并且妖怪的个数可以动态的增加=>map切片
	*/
	// 声明一个map切片
	var monsters []map[string]string
	monsters = make([]map[string]string, 2) //准备放入两个妖怪

	if monsters[0] == nil {
		monsters[0] = make(map[string]string, 2)
		monsters[0]["name"] = "牛魔王"
		monsters[0]["age"] = "500"
	}

	if monsters[1] == nil {
		monsters[1] = make(map[string]string, 2)
		monsters[1]["name"] = "玉兔精"
		monsters[1]["age"] = "400"
	}

	// if monsters[2] == nil {       //这个写法越界报错
	// 	monsters[2] = make(map[string]string, 2)
	// 	monsters[2]["name"] = "狐狸精"
	// 	monsters[2]["age"] = "300"
	// }

	/* 这里我们需要使用到切片的append函数，可以动态的增加monster */

	// 定义个newMonster信息
	newMonster := map[string]string{
		"name": "新的妖怪~火云邪神",
		"age":  "200",
	}
	monsters = append(monsters, newMonster)
	fmt.Println(monsters)

	/*
		演示map排序的使用

		golang中map的排序，是先将key进行排序，然后根据key值遍历输出即可
	*/

	//map的排序
	map1 := make(map[int]int, 10)
	map1[10] = 100
	map1[1] = 13
	map1[4] = 56
	map1[8] = 90
	fmt.Println(map1)
	//如果按照map的key的顺序进行排序输出
	//1. 先将map的key 放入到 切片中
	//2. 对切片排序
	//3. 遍历切片，然后按照key来输出map的值
	var keys []int
	for k := range map1 {
		keys = append(keys, k)
	}
	//排序
	sort.Ints(keys)
	fmt.Println(keys)

	for _, k := range keys {
		fmt.Printf("map1[%v]=%v \n", k, map1[k])
	}
}
