package main

import (
	"fmt"
	"net" //做网络socket开发时,net包含有我们需要所有的方法和函数
)

/*
服务端的处理流程：
	1) 监听端口
	2) 接收客户端的tcp链接，建立客户端和服务器端的链接
	3) 创建goroutine，处理该链接的请求(通常客户端会通过链接发送请求包）

客户端的处理流程：
	1) 建立与服务端的链接
	2) 发送请求数据，接收服务器端返回的结果数据
	3) 关闭链接
*/

/* 应用案例：
程序框架图示意图（略）
＞服务器端功能：
	1）编写一个服务器端程序，在8888端口监听
	2）可以和多个客户端创建链接
	3）链接成功后，客户端可以发送数据，服务器端接受数据，并显示在终端上
	4）先使用telnet来测试，然后编写客户端程序来测试
*/

func process(conn net.Conn) {

	//这里我们循环的接收客户端发送的数据
	defer conn.Close() //关闭conn

	for {
		//创建一个新的切片
		buf := make([]byte, 1024)
		//conn.Read(buf)
		//1. 等待客户端通过conn发送信息
		//2. 如果客户端没有wrtie[发送]，那么协程就阻塞在这里
		//fmt.Printf("服务器在等待客户端%s 发送信息\n", conn.RemoteAddr().String())
		n, err := conn.Read(buf) //从conn读取
		if err != nil {

			fmt.Printf("客户端退出 err=%v", err)
			return //!!!
		}
		//3. 显示客户端发送的内容到服务器的终端
		fmt.Print(string(buf[:n]))
	}

}

func main() {

	fmt.Println("服务器开始监听....")
	//net.Listen("tcp", "0.0.0.0:8888")
	//1. tcp 表示使用网络协议是tcp
	//2. 0.0.0.0:8888 表示在本地监听 8888端口
	listen, err := net.Listen("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("listen err=", err)
		return
	}
	defer listen.Close() //延时关闭listen

	//循环等待客户端来链接我
	for {
		//等待客户端链接
		fmt.Println("等待客户端来链接....")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("Accept() err=", err)

		} else {
			fmt.Printf("Accept() suc con=%v 客户端ip=%v\n", conn, conn.RemoteAddr().String())
		}
		//这里准备其一个协程，为客户端服务
		go process(conn)
	}

	//fmt.Printf("listen suc=%v\n", listen)
}
