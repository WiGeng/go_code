package main

import (
	"encoding/json"
	"fmt"
)

/* 1. 结构体序列化 */
type Person struct {
	Name     string `json:"person_name"` //将结构体的Name映射为person_name在结果中显示
	Age      int    `json:"Person_age"`
	Birthday string
	Sal      float64
	Skill    string
}

func testStruct() {
	person1 := Person{
		Name:     "woke",
		Age:      26,
		Birthday: "2011-08-08",
		Sal:      18000.0,
		Skill:    "DevOps",
	}

	//将Person 序列化
	data, err := json.Marshal(&person1)
	if err != nil {
		fmt.Printf("序列化错误 err=%v\n", err)
	}
	//输出序列化后的结果
	fmt.Printf("person1序列化后=%v\n", string(data))

}

/* 2. 将map进行序列化 */
func testMap() {
	//定义一个map
	var person2 map[string]interface{}
	person2 = make(map[string]interface{}) //使用map,需要make
	person2["name"] = "WiGENG"
	person2["age"] = 26
	person2["address"] = "长安区"

	//将person这个map进行序列化
	data, err := json.Marshal(person2)
	if err != nil {
		fmt.Printf("序列化错误 err=%v\n", err)
	}
	//输出序列化后的结果
	fmt.Printf("person2 map 序列化后=%v\n", string(data))

}

/* 3. 演示对切片进行序列化, 这个切片为 []map[string]interface{} */
func testSlice() {
	var slice []map[string]interface{}

	var m1 map[string]interface{}
	//使用map前，需要先make
	m1 = make(map[string]interface{})
	m1["name"] = "jack"
	m1["age"] = "7"
	m1["address"] = "北京"
	slice = append(slice, m1)

	var m2 map[string]interface{}
	//使用map前，需要先make
	m2 = make(map[string]interface{})
	m2["name"] = "tom"
	m2["age"] = "20"
	m2["address"] = [2]string{"墨西哥", "夏威夷"}
	slice = append(slice, m2)

	//将切片进行序列化操作
	data, err := json.Marshal(slice)
	if err != nil {
		fmt.Printf("序列化错误 err=%v\n", err)
	}
	//输出序列化后的结果
	fmt.Printf("slice 序列化后=%v\n", string(data))

}

/* 4. 对基本数据类型序列化，对基本数据类型进行序列化意义不大 */
func testFloat64() {
	var num1 float64 = 2345.67

	//对num1进行序列化
	data, err := json.Marshal(num1)
	if err != nil {
		fmt.Printf("序列化错误 err=%v\n", err)
	}
	//输出序列化后的结果
	fmt.Printf("num1 序列化后=%v\n", string(data))
}

func main() {
	testStruct()
	testMap()
	testSlice()
	testFloat64()
}
