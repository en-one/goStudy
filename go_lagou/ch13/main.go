package main

import (
	"fmt"
)

type address struct {
	province string
	city     string
}

func (addr address) String() string {
	return fmt.Sprintf("the addr is %s%s", addr.province, addr.city)
}

func printString(s fmt.Stringer) {
	fmt.Println(s.String())
}

func main() {
	add := address{province: "beijing", city: "beijing"}
	printString(add)
	printString(&add)

}

// go语言中所有的函数传参都是值传递---copy一份，地址是不一样的
// go语言的map建立的时候其实是*hmap 所以修改会成立
// chan引用类型 其实也是一个*hchan
