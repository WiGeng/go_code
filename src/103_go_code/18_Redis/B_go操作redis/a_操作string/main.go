package main

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

func setstring() {
	/* 通过go 向redis 写入数据和读取数据 */
	//1. 链接到redis
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("redis.Dial err=", err)
		return
	}
	defer conn.Close() //关闭..

	//2. 通过go 向redis写入数据 string [key-val]
	_, err = conn.Do("Set", "name", "woke")
	if err != nil {
		fmt.Println("set  err=", err)
		return
	}

	//3. 通过go 向redis读取数据 string [key-val]
	res, err := redis.String(conn.Do("Get", "name")) //因为返回res是interface{};而name对应的值是string ,因此我们需要转换,通过redis.String()方法
	if err != nil {
		fmt.Println("Get  err=", err)
		return
	}
	fmt.Println("读取name的value=", res)
}

func expiretime() {
	/* 设置有效时间 */
	// 1. 连接redis
	conn, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println("redis Dail err=", err)
		return
	}
	defer conn.Close()
	//2. 通过go 向redis写入数据 string [key-val]
	_, err = conn.Do("Setex", "exkey", 10, "10s后过期")
	if err != nil {
		fmt.Println("Setex err=", err)
		return
	}
	//3. 通过go 向redis读取数据
	res, err := redis.String(conn.Do("Get", "exkey"))
	if err != nil {
		fmt.Println("Get  err=", err)
		return
	}
	fmt.Println("读取exkey的value=", res)

}

func main() {
	setstring()
	expiretime()
}
