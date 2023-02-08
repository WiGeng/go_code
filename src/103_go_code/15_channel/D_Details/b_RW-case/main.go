package main

import "fmt"

/* 2. channel只读和只写最佳实践案例 */

// ch chan<- int，这样ch就只能写樂作了
func send(ch chan<- int, exitChan chan struct{}) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
	var a struct{}
	exitChan <- a
}

// ch s-chan int,这样ch就只能读操作了
func recv(ch <-chan int, exitChan chan struct{}) {
	for {
		v, ok := <-ch
		if !ok {
			break
		}
		fmt.Println(v)
	}
	var a struct{}
	exitChan <- a

}

func main() {

	ch := make(chan int, 10)
	exitChan := make(chan struct{}, 2)
	go send(ch, exitChan)
	go recv(ch, exitChan)

	var total = 0
	for _ = range exitChan {
		total++
		if total == 2 {
			break
		}
		fmt.Println("结束...")
	}
}
