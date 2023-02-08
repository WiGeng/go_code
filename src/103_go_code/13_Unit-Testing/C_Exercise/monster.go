package store

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

/* 单元测试综合案例要求：
1）编写一个Monster结构体，字段 Name, Age, skill
2) 给Monster忘方法Store，可以将一个Monster变量（对象),序列化后保存到文件中
3) 给Monster绑定方法Restore，可以将一个序列化的Monster,从文件中读取，并反序列化为Monster对象
4) 编程测试用例文件 store_test.go，编写测试用例函数 Teststore和TestRestore进行测试。
*/

type Monster struct {
	Name  string
	Age   int
	Skill string
}

//给Monster绑定方法Store, 可以将一个Monster变量(对象),序列化后保存到文件中
func (this *Monster) Store() bool {

	// 序列化
	monster_data, err1 := json.Marshal(this)
	if err1 != nil {
		fmt.Printf("序列号错误 err1=%v\n", err1)
		return false
	}

	// 写入文件
	filePath := "/Users/weigeng/go/src/Note/test.txt"
	err2 := ioutil.WriteFile(filePath, monster_data, 0666)
	if err2 != nil {
		fmt.Printf("write file err=%v\n", err2)
		return false
	}
	return true
}

func (this *Monster) ReStore() bool {
	// 读取文件
	file := "/Users/weigeng/go/src/Note/test.txt"
	content, err1 := ioutil.ReadFile(file)
	if err1 != nil {
		fmt.Printf("read file err=%v", err1)
		return false
	}

	// 使用读取到data []byte ,并反序列化
	err2 := json.Unmarshal(content, this)
	if err2 != nil {
		fmt.Printf("unmarshal err=%v\n", err2)
		return false
	}
	return true
}
