package main

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

func main() {
	/* 通过go对redis操作list数据类型 */
	//1. 链接到redis
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("redis.Dial err=", err)
		return
	}
	defer conn.Close() //关闭..

	//2. 通过go 向redis写入数据 LPUSH key value [value …]
	_, err = conn.Do("Lpush", "email", "abc@suho.com", "woke@google.com", "wiger@yahoo.com", "weigeng@163.com")
	if err != nil {
		fmt.Println("Lpush  err=", err)
		return
	}
	//3. 获取List的长度
	len, err := conn.Do("LLen", "email")
	if err != nil {
		fmt.Println("Lpush  err=", err)
		return
	}
	fmt.Println("list email len()=", len)

	//4. 通过go 向redis读取数据 LRANGE key start stop
	res, err := redis.Strings(conn.Do("LRANGE", "email", "0", "-1"))
	if err != nil {
		fmt.Println("LRANGE  err=", err)
		return
	}
	for _, i := range res {
		fmt.Println(i)
	}

}
