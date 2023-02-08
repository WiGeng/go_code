package main

import "fmt"

/*
1） 将一个循环放在另一个循环体内，就形成了嵌套循环。在外边的for称为外层循环在里面的for循环称为内层循环。【建议一般使用两层，最多不要超过3层】
2） 实质上，嵌套循环就是把内层循环当成外层储环的循环体。当只有内层循环的循环条件为false时，才会完全跳出内层循环，才可结束外层的当次循环，开始下一次的循环。
3） 设外层循环次数为m次，内层为n次，则内层循环体实际上需要执行m*n=mn
次。
*/

func main() {

	// 1）统计3个班成绩情况，每个班有5名同学，求出各个班的平均分和所有班级的平均分[学生的成绩从键盘输入]
	var totalsum float64
	for j := 1; j <= 3; j++ {
		var scroe float64
		var sum float64
		for i := 1; i <= 5; i++ {
			fmt.Printf("请输入%v班第%v位同学的成绩\n", j, i)
			fmt.Scanln(&scroe)
			sum += scroe
		}
		fmt.Println("班级平均分为", sum/5)
		totalsum += sum
	}
	fmt.Printf("3个班的总成绩为%v  年级平均分为%v\n", totalsum, totalsum/(5*3))

	/*---------------------------------------------------------------*/
	// 2）统计三个班及格人数，每个班有5名同学
	var totalcount int
	for j := 1; j <= 3; j++ {
		var count int
		var score float32
		for i := 1; i <= 5; i++ {
			fmt.Printf("请输入第%v班第%v位同学的成绩\n", j, i)
			fmt.Scanln(&score)
			if score >= 60 {
				count++
			}
		}
		totalcount += count
	}
	fmt.Println("全年级及格的人数为：", totalcount)

	/*---------------------------------------------------------------*/
	// 3) 打印一个矩形
	for i := 1; i <= 15; i++ {
		for j := 1; j <= 10; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}

	/*---------------------------------------------------------------*/
	// 4）打印金宇塔经典案例
	for i := 1; i <= 10; i++ {
		//在打印*前先打印空格，空格的规律为 总层数-当前层数
		for j := 1; j <= 10-i; j++ {
			fmt.Print(" ")
		}
		//k表示每层打印多少*，规律为 2 * i - 1
		for k := 1; k <= 2*i-1; k++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	/*---------------------------------------------------------------*/
	// 5）打印实心菱形
	var n int = 10
	//打印上半部分
	for i := 1; i <= 10; i++ {
		for k := 1; k <= n-i; k++ {
			fmt.Print(" ")
		}

		for j := 1; j <= 2*i-1; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
	//打印下半部分
	for i := n - 1; i >= 1; i-- {
		for k := 1; k <= n-i; k++ {
			fmt.Print(" ")
		}

		for j := 1; j <= 2*i-1; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
	/*---------------------------------------------------------------*/
	// 6) 打印空心棱形
	//打印上半部分
	for i := 1; i <= 10; i++ {
		for k := 1; k <= n-i; k++ {
			fmt.Print(" ")
		}

		for j := 1; j <= 2*i-1; j++ {
			if j == 1 || j == 2*i-1 {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
	//打印下半部分
	for i := n - 1; i >= 1; i-- {
		for k := 1; k <= n-i; k++ {
			fmt.Print(" ")
		}

		for j := 1; j <= 2*i-1; j++ {
			if j == 1 || j == 2*i-1 {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}

		}
		fmt.Println()
	}

	/*---------------------------------------------------------------*/
	// 7）打印九九乘法表
	// 遍历, 决定处理第几行
	for y := 1; y <= 9; y++ {
		// 遍历, 决定这一行有多少列
		for x := 1; x <= y; x++ {
			fmt.Printf("%d*%d=%d ", x, y, x*y)
		}
		// 手动生成回车
		fmt.Println()
	}
}
