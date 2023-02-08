package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	//将d:/abc.txt 文件内容导入到  e:/kkk.txt
	//1. 首先将  d:/abc.txt 内容读取到内存
	//2. 将读取到的内容写入 e:/kkk.txt
	srcfilePath := "/Users/weigeng/go/src/Note/woke.txt"
	destfilePath := "/Users/weigeng/go/src/Note/copy.txt"

	content, err := ioutil.ReadFile(srcfilePath)
	if err != nil {
		//说明读取文件有错误
		fmt.Printf("read file err=%v\n", err)
		return
	}
	err = ioutil.WriteFile(destfilePath, content, 0666)
	if err != nil {
		fmt.Printf("write file error=%v\n", err)
	}
}
