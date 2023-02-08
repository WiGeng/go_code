package main

import (
	"fmt"
	"time"
)

/* 日期和时间函数
说明：
	在编程中，程序员会经常使用到日期相关的函数，比如：统计某段代码执行花费的时间等等。
操作使用细节：
	1） 时间和日期相关函数，需要导入time包
	2） 时间常量：在程序中用于获取指定时间单位的时间，比如100毫秒 100 * time.Millisecond
		const (
			Nanosecond  Duration = 1	               //纳秒
			Microsecond          = 1000 * Nanosecond   //微秒
			Millisecond          = 1000 * Microsecond  //毫秒
			Second               = 1000 * Millisecond  //秒
			Minute               = 60 * Second	       //分钟
			Hour                 = 60 * Minute         //小时
		)
	3） 休眠： func Sleep(d Duration) // Sleep阻塞当前go程至少d代表的时间段。d<=0时，Sleep会立刻返回
	4）	时间戳： 获取unix和unixNano时间戳

*/

func main() {
	//看看日期和时间相关函数和方法使用
	//1. 获取当前时间
	now := time.Now()
	fmt.Printf("now=%v now type=%T\n", now, now)

	//2.通过now可以获取到年月日，时分秒
	fmt.Printf("年=%v\n", now.Year())
	fmt.Printf("月=%v\n", now.Month())
	fmt.Printf("月=%v\n", int(now.Month()))
	fmt.Printf("日=%v\n", now.Day())
	fmt.Printf("时=%v\n", now.Hour())
	fmt.Printf("分=%v\n", now.Minute())
	fmt.Printf("秒=%v\n", now.Second())

	//格式化日期时间

	fmt.Printf("当前年月日 %d-%d-%d %d:%d:%d \n",
		now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())

	dateStr := fmt.Sprintf("当前年月日 %d-%d-%d %d:%d:%d \n",
		now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
	fmt.Printf("dateStr=%v\n", dateStr)

	//格式化日期时间的第二种方式 [固定格式，数字请勿修改]
	fmt.Printf(now.Format("2006-01-02 15:04:05"))
	fmt.Println()
	fmt.Printf(now.Format("2006-01-02"))
	fmt.Println()
	fmt.Printf(now.Format("15:04:05"))
	fmt.Println()
	fmt.Printf(now.Format("2006"))
	fmt.Println()

	//需求，每隔1秒中打印一个数字，打印到100时就退出
	//需求2: 每隔0.1秒中打印一个数字，打印到100时就退出
	i := 0
	for {
		i++
		fmt.Println(i)
		//休眠0.1s
		//time.Sleep(time.Second)
		time.Sleep(time.Millisecond * 100)
		if i == 100 {
			break
		}
	}

	//Unix和UnixNano的使用
	fmt.Printf("unix时间戳=%v unixnano时间戳=%v\n", now.Unix(), now.UnixNano())
	fmt.Printf("unix时间戳=%v unixnano时间戳=%v\n", time.Now().Unix(), now.UnixNano())

}
