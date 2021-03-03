package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

//or 另一种形式
func mul(x, y int) int {
	return x * y
}

func main() {
	fmt.Println(add(45, 3))
	fmt.Println(mul(45, 3))
}
