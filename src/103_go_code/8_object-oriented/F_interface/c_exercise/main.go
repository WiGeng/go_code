package main

import (
	"fmt"
	"math/rand"
	"sort"
)

/*
练习： 实现对Hero结构体的切片的排序：sort:Sort(data Interface)
*/

//1.声明Hero结构体
type Hero struct {
	Name string
	Age  int
}

//2.声明一个Hero结构体切片类型
type HeroSlice []Hero

//3.实现Interface 接口
func (hs HeroSlice) Len() int {
	return len(hs)
}

//Less方法就是决定你使用什么标准进行排序
//1. 按Hero的年龄从小到大排序!!
func (hs HeroSlice) Less(i, j int) bool {
	return hs[i].Age < hs[j].Age
	//修改成对Name排序
	//return hs[i].Name < hs[j].Name
}

func (hs HeroSlice) Swap(i, j int) {
	//交换
	// temp := hs[i]
	// hs[i] = hs[j]
	// hs[j] = temp
	//下面的一句话等价于三句话
	hs[i], hs[j] = hs[j], hs[i]
}

func main() {

	//测试看看我们是否可以对结构体切片进行排序
	var heroes HeroSlice
	for i := 0; i < 10; i++ {
		hero := Hero{
			Name: fmt.Sprintf("英雄-%d", rand.Intn(100)),
			Age:  rand.Intn(100),
		}
		//将 hero append到 heroes切片
		heroes = append(heroes, hero)
	}

	//看看排序前的顺序
	fmt.Println("-----------排序前------------")
	for _, v := range heroes {
		fmt.Println(v)
	}

	//调用sort.Sort
	sort.Sort(heroes)
	fmt.Println("-----------排序后------------")
	//看看排序后的顺序
	for _, v := range heroes {
		fmt.Println(v)
	}

}
