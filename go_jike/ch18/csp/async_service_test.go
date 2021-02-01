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

func TestService(t *testing.T) {
	fmt.Print(service())
	otherTask()
}

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
