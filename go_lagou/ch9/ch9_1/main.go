package main

import (
	"fmt"
	"sync"
)

var sum = 0

func main() {
	run()
}

func run() {
	var wg sync.WaitGroup

	wg.Add(110)

	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			add(10)
		}()
	}

	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			fmt.Println("和为：", readSum())
		}()
	}

	wg.Wait()
}

func add(i int) {
	sum += i
}

func readSum() int {
	b := sum
	return b
}
