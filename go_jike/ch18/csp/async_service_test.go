package csp

import (
	"fmt"
	"testing"
	"time"
)

func otherTask() {
	fmt.Println("working on something else")
	time.Sleep(time.Millisecond * 100)
	fmt.Println("task is done")
}

func service() string {
	time.Sleep(time.Millisecond * 50)
	return "Done"
}

//func TestService(t *testing.T) {
//	fmt.Print(service())
//	otherTask()
//}

func ASynService() chan string {
	retCh := make(chan string, 1)
	go func() {
		ret := service()
		fmt.Println("returned result.")
		retCh <- ret
		fmt.Println("service exited")
	}()
	return retCh
}

func TestASynService(t *testing.T) {
	retCh := ASynService()
	otherTask()
	fmt.Println(<-retCh)
}

/*
csp模式 是通过channel进行通讯的
go中的channel是有容量限制并且独立于处理器groutine

有缓存（buffer channel）   可以在容量足够情况先 先push 而不需要pop
无缓存（channel）          有push 有pop 要不然会阻塞

多路选择
select {
case
case ： time.After   //超时机制
}


close(chan) //chan关闭
if data, ok := <-ch; ok {}
*/
