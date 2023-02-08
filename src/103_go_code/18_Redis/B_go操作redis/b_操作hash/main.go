package main

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

func main() {
	/* 通过go对redis操作hash数据类型 */
	//1. 链接到redis
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("redis.Dial err=", err)
		return
	}
	defer conn.Close() //关闭..

	//2. 通过go 向redis写入数据 HMSET key field value [field value …]
	_, err = conn.Do("HMSET", "website", "google", "www.google.com", "yahoo", "www.yahoo.com")
	if err != nil {
		fmt.Println("hmset  err=", err)
		return
	}

	//3. 通过go 向redis读取数据 HGET hash field
	res1, err := redis.String(conn.Do("HGET", "website", "google")) //单个返回res是interface{} ;通过redis.String()方法转换
	if err != nil {
		fmt.Println("hmget  err=", err)
		return
	}
	fmt.Println("读取website的google=", res1)

	//4. 通过go 向redis读取数据 HGET hash field
	res2, err := redis.String(conn.Do("HGET", "website", "yahoo")) //单个返回res是interface{} ;通过redis.String()方法转换
	if err != nil {
		fmt.Println("hmget  err=", err)
		return
	}
	fmt.Println("读取website的yahoo=", res2)

	//5. 通过go 向redis读取数据 HGETALL website
	res3, err := redis.Strings(conn.Do("HGetAll", "website")) //多个返回res是interface{} ;通过redis.Strings()方法转换
	if err != nil {
		fmt.Println("hmget  err=", err)
		return
	}
	fmt.Println("读取website的value=", res3)

	//6. 通过go 向redis读取数据 HMGet website
	res4, err := redis.Strings(conn.Do("HMGet", "website", "google", "yahoo")) //多个返回res是interface{} ;通过redis.Strings()方法转换
	if err != nil {
		fmt.Println("hmget  err=", err)
		return
	}
	for i, v := range res4 { //返回的res4类似是切片，因此需要对其遍历读取
		fmt.Printf("r[%d]=%s\n", i, v)
	}

}
