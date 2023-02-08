package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	/* 打开一个存在的文件，将原来的内容读出显示在终端，并且追加5句"hello,西安" */
	filePath := "/Users/weigeng/go/src/Note/woke.txt" // 原文件
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_APPEND, 0644)
	if err != nil {
		fmt.Printf("open file err=%v\n", err)
		return
	}
	defer file.Close() //及时关闭file句柄
	//先读取原文件并输出在终端
	reader := bufio.NewReader(file)
	//循环的读取文件的内容
	for {
		str, err := reader.ReadString('\n') // 读到一个换行就结束
		if err == io.EOF {                  // io.EOF表示文件的末尾
			break
		}
		//输出内容
		fmt.Print(str)
	}
	//向原文件追加内容
	str := "hello,西安\r\n"
	writer := bufio.NewWriter(file) //写入时，使用带缓存的 *Writer
	for i := 0; i < 5; i++ {
		writer.WriteString(str)
	}
	/*  因为writer是带缓存，因此在调用WriterString方法时，其实内容是先写入到缓存的,所以需要调用Flush方法，将缓冲的数据
	真正写入到文件中， 否则文件中会没有数据!!!
	*/
	writer.Flush()
}
