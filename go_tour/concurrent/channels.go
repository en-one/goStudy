package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}

func fib(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)

	x, y := <-c, <-c
	fmt.Println(x, y, x+y)

	//带缓冲的channel
	ch := make(chan int, 2)
	ch <- 8
	ch <- 9
	fmt.Println(<-ch)
	fmt.Println(<-ch)

	//range 和close
	//循环 for i := range c 会不断从信道接收值，直到它被关闭
	//close关闭channel（只有发送者才能关闭），向关闭的channel发送数据会引起panic
	//一般情况下channel不需要关闭，除非为了中止tange
	z := make(chan int, 10)
	go fib(cap(z), z)
	//v, ok := <-ch
	for i := range z {
		fmt.Println(i)
	}

}
