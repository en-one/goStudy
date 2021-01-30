package main

import (
	"fmt"
	"strconv"
)

func main() {
	i, err := strconv.Atoi("a")
	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Println(i)
	}
}
