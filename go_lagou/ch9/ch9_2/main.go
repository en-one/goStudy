package main

import (
	"fmt"
	"sync"
)

func main() {
	doOnce()
}

func doOnce() {
	var once sync.Once

	onceBody := func() {
		fmt.Println("once only")
	}

	done := make(chan bool)

	for i := 0; i < 10; i++ {
		go func() {
			once.Do(onceBody)
			done <- true
		}()
	}

	for i := 0; i < 10; i++ {
		<-done
	}
}
