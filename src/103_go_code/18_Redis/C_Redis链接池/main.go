package main

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

/* redis链接池
说明：
	通过Golang对Redis操作，还可以通过Redis链接池，流程如下：
	1） 事先初始化一定数量的链接，放入到链接池
	2） 当Go需要操作Redis时，直接从Redis链接池取出链接即可
	3）这样可以节省临时获取Redis链接的时间，从而提高效率

	核心代码：
	var pool *redis.Pool
	pool = &redis.Pool{
		Maxidle: 8,							//最大空闲链接数
		MaxActive: 0,						//表示和数据库的最大链接数,0表示没有限制
		IdleTimeout: 100,					//最大空闲时间
		Dial: func()(redis.Conn, error){	//初始化链接代码
			return redis.Dial("tcp", "localhost:6379")
		},
	}
	c := pool.Get()							//从链接池中取出一个链接
	pool.Close()							//关闭链接池，一旦关闭链接池，就不能在从链接池中再取出链接
*/

//定义一个全局的pool
var pool *redis.Pool

//当启动程序时，就初始化连接池
func init() {

	pool = &redis.Pool{
		MaxIdle:     8,   //最大空闲链接数
		MaxActive:   0,   // 表示和数据库的最大链接数， 0 表示没有限制
		IdleTimeout: 100, // 最大空闲时间
		Dial: func() (redis.Conn, error) { // 初始化链接的代码， 链接哪个ip的redis
			return redis.Dial("tcp", "localhost:6379")
		},
	}

}

func main() {
	//先从pool 取出一个链接
	conn := pool.Get()
	defer conn.Close()

	_, err := conn.Do("Set", "name", "woke")
	if err != nil {
		fmt.Println("conn.Do err=", err)
		return
	}

	//取出
	res, err := redis.String(conn.Do("Get", "name"))
	if err != nil {
		fmt.Println("conn.Do err=", err)
		return
	}

	fmt.Println("res=", res)

	//如果我们要从pool 取出链接，一定保证链接池是没有关闭
	//pool.Close()
	conn2 := pool.Get()
	_, err1 := conn2.Do("Set", "name2", "WiGENG")
	if err != nil {
		fmt.Println("conn.Do err1=", err1)
		return
	}
	//取出
	res2, err := redis.String(conn2.Do("Get", "name2"))
	if err != nil {
		fmt.Println("conn.Do err=", err)
		return
	}

	fmt.Println("res2=", res2)
}
