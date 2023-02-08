package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	/* 2. 打开一个存在的文件中，将原来的内容覆盖成新的内容10句“你好，蜗壳!” */
	filePath := "/Users/weigeng/go/src/Note/woke.txt"
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_TRUNC, 0644) //关键字：os.O_TRUNC
	if err != nil {
		fmt.Printf("open file err=%v\n", err)
		return
	}
	defer file.Close()
	str2 := "你好，蜗壳!\r\n"
	writer2 := bufio.NewWriter(file) //写入时，使用带缓存的 *Writer
	for i := 0; i < 10; i++ {
		writer2.WriteString(str2)
	}
	writer2.Flush()

}
