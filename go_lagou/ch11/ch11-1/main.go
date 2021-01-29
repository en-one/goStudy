package main

import (
	"fmt"
	"time"
)

func main() {
	result := make(chan string)
	go func() {
		time.Sleep(time.Second * 8)
		result <- "服务端结果"
	}()

	select {
	case c := <-result:
		fmt.Println(c)
	case <-time.After(time.Second * 5):
		fmt.Println("超时了")
	}
}
