package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(time.Microsecond * 100)
		fmt.Println(s)
	}
}

func main() {
	go say("go")
	say("hello")
}
