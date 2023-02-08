package main

import (
	"encoding/json"
	"fmt"
)

/* 介绍
json反序列化是指将json字符串反序列化成对应的数据类型(比如结构体、map、切片)的操作

对下面代码的小结说明
	1）在反序列化一个json字符串时，要确保反序列化后的数据类型和原来序列化前的数据类型一致。
	2）如果json字符串是通过程序获取到的，则不需要再对""进行转义处理。
*/

//定义一个结构体
type Person struct {
	Name     string
	Age      int
	Birthday string
	Sal      float64
	Skill    string
}

/* 1. 将json字符串，反序列化成struct */
func unmarshalStruct() {
	str := "{\"Name\":\"woke\",\"Age\":26,\"Birthday\":\"2011-08-08\",\"Sal\":18000,\"Skill\":\"DevOps\"}"
	var person Person //定义一个Person实例
	err := json.Unmarshal([]byte(str), &person)
	if err != nil {
		fmt.Printf("unmarshal err=%v\n", err)
	}
	fmt.Printf("反序列化后 person=%v person.Name=%v \n", person, person.Name)

}

/*  将map进行序列化 */
func testMap() string {
	var person2 map[string]interface{}     //定义一个map
	person2 = make(map[string]interface{}) //使用map,需要make
	person2["name"] = "WiGENG"
	person2["age"] = 26
	person2["address"] = "长安区"

	data, err := json.Marshal(person2)
	if err != nil {
		fmt.Printf("序列化错误 err=%v\n", err)
	}
	//fmt.Printf("a map 序列化后=%v\n", string(data)) //输出序列化后的结果
	return string(data)

}

/* 2. 演示将json字符串，反序列化成map */
func unmarshalMap() {
	//str := "{\"address\":\"长安区\",\"age\":26,\"name\":\"WiGENG\"}"
	str := testMap()
	var person2 map[string]interface{} //定义一个map
	//反序列化   注意：反序列化map,不需要make,因为make操作被封装到 Unmarshal函数
	err := json.Unmarshal([]byte(str), &person2)
	if err != nil {
		fmt.Printf("unmarshal err=%v\n", err)
	}
	fmt.Printf("反序列化后 person2=%v\n", person2)

}

/* 3. 演示将json字符串，反序列化成切片 */
func unmarshalSlice() {
	str := "[{\"address\":\"北京\",\"age\":\"7\",\"name\":\"jack\"}," +
		"{\"address\":[\"墨西哥\",\"夏威夷\"],\"age\":\"20\",\"name\":\"tom\"}]"

	var slice []map[string]interface{}         //定义一个slice
	err := json.Unmarshal([]byte(str), &slice) //反序列化，不需要make,因为make操作被封装到 Unmarshal函数
	if err != nil {
		fmt.Printf("unmarshal err=%v\n", err)
	}
	fmt.Printf("反序列化后 slice=%v\n", slice)
}

func main() {
	unmarshalStruct()
	unmarshalMap()
	unmarshalSlice()
}
