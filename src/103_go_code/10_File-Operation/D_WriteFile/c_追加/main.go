package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	/* 打开一个存在的文件，在原来的内容追加内容 "MVP! KOBE!!" */
	filePath := "/Users/weigeng/go/src/Note/woke.txt"
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_APPEND, 0644) //关键字：os.O_APPEND
	if err != nil {
		fmt.Printf("open file err=%v\n", err)
		return
	}
	defer file.Close() //及时关闭file句柄

	str3 := "MVP! KOBE!!\r\n"
	writer := bufio.NewWriter(file)
	for i := 0; i < 2; i++ {
		writer.WriteString(str3)
	}
	writer.Flush()
}
