package main

import (
	"fmt"
	"time"
)

func main() {
	firstChan := make(chan string)
	secondChan := make(chan string)
	thirdChan := make(chan string)

	go func() {
		firstChan <- downloadFile("firstCh")
	}()

	go func() {
		secondChan <- downloadFile("secondCh")
	}()

	go func() {
		thirdChan <- downloadFile("thirdCh")
	}()

	select {
	case filepath := <-firstChan:
		fmt.Println(filepath)
	case filepath := <-secondChan:
		fmt.Println(filepath)
	case filepath := <-thirdChan:
		fmt.Println(filepath)

	}
}

func downloadFile(chanName string) string {
	time.Sleep(time.Second)
	return chanName + ":filepath"
}
