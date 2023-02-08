package main

import (
	"bufio"
	"fmt"
	"os"
)

/* 基本介绍
func OpenFile(name string, flag int, perm FileMode) (file *File, err error)
说明：
	os.OpenFile是一个更一般性的文件打开函数，它会使用指定的选项（如O_RDONLY等）、指定的模式（如0666等）打开指定名称的文件。
	 如果操作成功，返回的文件对象可用于I/O。如果出错，错误底层类型是*PathError。
文件的几种模式：
	const (
    O_RDONLY int = syscall.O_RDONLY // 只读模式打开文件
    O_WRONLY int = syscall.O_WRONLY // 只写模式打开文件
    O_RDWR   int = syscall.O_RDWR   // 读写模式打开文件
    O_APPEND int = syscall.O_APPEND // 写操作时将数据附加到文件尾部
    O_CREATE int = syscall.O_CREAT  // 如果不存在将创建一个新文件
    O_EXCL   int = syscall.O_EXCL   // 和O_CREATE配合使用，文件必须不存在
    O_SYNC   int = syscall.O_SYNC   // 打开文件用于同步I/O
    O_TRUNC  int = syscall.O_TRUNC  // 如果可能，打开时清空文件
)
*/

/* 文件写操作应用实例
基本应用实例-方式一
	1）新建一个新文件，写入内容5句"hello, woke"
	2）打开一个存在的文件中，将原来的内容覆盖成新的内容10句“你好，蜗壳!
	3）打开一个存在的文件，在原来的内容追加内容‘'ABC! ENGLISH!!
	4）打开一个存在的文件，将原来的内容读出显示在终端，并且追加5句"hello,西安"
	使用os.0penFile(),bufio.Newwriter(,*Writer的方法Write string完成上面的任务．
*/

func main() {
	/* 1. 新建一个文件，写入内容5句"hello, woke" */
	filePath := "/Users/weigeng/go/src/Note/woke.txt" // 创建新文件
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Printf("open file err=%v\n", err)
		return
	}
	defer file.Close() //及时关闭file句柄
	str := "hello,woke\r\n"
	writer := bufio.NewWriter(file) //写入时，使用带缓存的 *Writer
	for i := 0; i < 5; i++ {
		writer.WriteString(str)
	}
	/*  因为writer是带缓存，因此在调用WriterString方法时，其实内容是先写入到缓存的,所以需要调用Flush方法，将缓冲的数据
	真正写入到文件中， 否则文件中会没有数据!!!
	*/
	writer.Flush()
}
