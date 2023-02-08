package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

/* 文件拷贝：
将文件[图片/mp3/视频]拷贝搭配另外一个目录

````
func Copy(dst Writer, src Reader) (written int64, err error)
````

将src的数据拷贝到dst，直到在src上到达EOF或发生错误。返回拷贝的字节数和遇到的第一个错误。对成功的调用，返回值err为nil而非EOF，
因为Copy定义为从src读取直到EOF，它不会将读取到EOF视为应报告的错误。如果src实现了WriterTo接口，本函数会调用src.WriteTo(dst)
进行拷贝；否则如果dst实现了ReaderFrom接口，本函数会调用dst.ReadFrom(src)进行拷贝。
*/

//自己编写一个函数，接收两个文件路径 srcFileName destFileName
func CopyFile(dstFileName string, srcFileName string) (written int64, err error) {

	srcFile, err := os.Open(srcFileName)
	if err != nil {
		fmt.Printf("open file err=%v\n", err)
	}
	defer srcFile.Close()
	//通过srcfile ,获取到 Reader
	reader := bufio.NewReader(srcFile)

	//打开dstFileName
	dstFile, err := os.OpenFile(dstFileName, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("open file err=%v\n", err)
		return
	}

	//通过dstFile, 获取到 Writer
	writer := bufio.NewWriter(dstFile)
	defer dstFile.Close()

	return io.Copy(writer, reader)

}

func main() {

	//将d:/flower.jpg 文件拷贝到 e:/abc.jpg

	//调用CopyFile 完成文件拷贝
	srcFile := "/Users/weigeng/Pictures/R-C.jpeg"
	destFile := "/Users/weigeng/Desktop/aim.jpeg"
	_, err := CopyFile(destFile, srcFile)
	if err == nil {
		fmt.Printf("拷贝完成\n")
	} else {
		fmt.Printf("拷贝错误 err=%v\n", err)
	}

}
