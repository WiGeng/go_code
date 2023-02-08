package main

import (
	"fmt"
	"os"
)

/* 常用的文件操作函数和方法
func Open (打开一个文件进行读操作)
	func Open(name string) (file *File, err error)
	Open打开一个文件用于读取。如果操作成功，返回的文件对象的方法可用于读取数据；对应的文件描述符具有O_RDONLY模式。如果出错，错误底层类型是*PathError。

	func (*File) Close (关闭一个文件)
	func (f *File) Close() error
	Close关闭文件f，使文件不能用于读写。它返回可能出现的错误。
*/

func main() {
	//打开文件
	//概念说明: file 的叫法
	//1. file 叫 file对象
	//2. file 也叫 file指针
	//3. file 也叫 file 文件句柄
	file, err := os.Open("/Users/weigeng/go/src/VsCode/awesomeProject/")
	if err != nil {
		fmt.Println("open file err=", err)
	}
	//输出下文件，看看文件是什么, 看出file 就是一个指针 *File
	fmt.Printf("file=%v", file)
	//关闭文件
	err = file.Close()
	if err != nil {
		fmt.Println("close file err=", err)
	}
}
