package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

/*读文件操作应用实例
读取文件的内容并品示在终端(带缓冲区的方式），使用 os.Open, file.Close, bufio.NewReader(),reader.ReadString 函数和方法
*/

func main() {
	file1, err := os.Open("/Users/weigeng/go/src/Note/hello.txt")
	if err != nil {
		fmt.Println("open file erro=", err)
	}
	fmt.Println("文件读取方式1..........")
	//当函数退出时，要及时的关闭file
	defer file1.Close() //要及时关闭file句柄，否则会有内存泄漏.
	/* 方式1: 使用bufio.NewReader()创建一个*Reader来读取文件, 是带缓冲的*/
	/*
	   const (
	   defaultBufSize = 4096 //默认的缓冲区为4096
	   )
	*/
	reader := bufio.NewReader(file1)
	//循环的读取文件的内容
	for {
		str, err := reader.ReadString('\n') // 读到一个换行就结束
		if err == io.EOF {                  // io.EOF表示文件的末尾
			break
		}
		//输出内容
		fmt.Print(str)
	}
 
	fmt.Println()
	fmt.Println("文件读取方式2..........")
	/* 方式2: 使用ioutil.ReadFile一次性将文件读取到位[适用于小文件]]*/
	file2 := "/Users/weigeng/go/src/Note/hello.txt"
	content, err := ioutil.ReadFile(file2)
	if err != nil {
		fmt.Printf("read file err=%v", err)
	}
	//把读取到的内容显示到终端
	//fmt.Printf("%v", content) // []byte
	fmt.Printf("%v", string(content)) // []byte
	//没有显式的Open文件，因此也不需要显式的Close文件
	//因为，文件的Open和Close被封装到 ReadFile 函数内部
}
