package main

import (
	"fmt"
	"time"
)

//func fibonacci(c, quit chan int)  {
//	x ,y :=  0,1
//	for  {
//		select {
//		case c <- x:
//			x,y = y, x+y
//		case <-quit:
//			fmt.Println("quit")
//			return
//		}
//	}
//}

func main() {
	//	select 可以使一个go等待多个通信操作
	//	c := make(chan int)
	//	quit := make(chan int)
	//	go func() {
	//		for i :=0; i<10;i++ {
	//			fmt.Println(<-c)
	//		}
	//		quit <- 0
	//	}()
	//	fibonacci(c, quit)

	//	当select其他分支没有准备好的时候 就会执行默认分支
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)
	for {
		select {
		case <-tick:
			fmt.Println("tick")
		case <-boom:
			fmt.Println("boom")
			return
		default:
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)

		}
	}
}
