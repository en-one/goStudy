package main

import (
	"errors"
	"fmt"
)

func main() {
	e := errors.New("原始错误e")
	w := fmt.Errorf("Wrap了一个错误:%w", e)
	fmt.Println(w)
}
