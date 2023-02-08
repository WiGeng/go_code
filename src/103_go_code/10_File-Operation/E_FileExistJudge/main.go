package main

import (
	"fmt"
	"os"
)

/* golang判断文件或文件夹是否存在的方法为使用os.stat()函数返回的错误值进行判断：
1）如果返回的错误为nil,说明文件或文件夹存在
2）如果返回的错误类型使用os.IsNotExist()判断为true,说明文件或文件夹不存在
3）如果返回的错误为其它类型,则不确定是否在存在

func PathExists(path string) (bool, error) {
	_,err := os. Stat(path)
	if err == nil {			//文件/目录存在
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

*/
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		fmt.Println("File exist")
		return true, nil
	}
	if os.IsNotExist(err) {
		fmt.Println("File not exist")
		return false, nil

	}
	return false, err
}

func main() {
	PathExists("/Users/weigeng/go/src/Note/")
}
