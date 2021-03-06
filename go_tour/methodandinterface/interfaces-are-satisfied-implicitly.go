package main

import (
	"fmt"
	"math"
)

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

type IF float64

func (interF IF) M() {
	fmt.Println(interF)
}

func main() {
	var i I
	//接口内的具体值为 nil，方法仍然会被 nil 接收者调用
	var t *T
	i = t
	//nil 接口调用方法会产生运行时错误,若没有上面两个赋值
	describe(i)
	i.M()

	i = &T{"hello"}
	describe(i)
	i.M()

	i = IF(math.Pi)
	describe(i)
	i.M()

	//空接口
	var emptyI interface{}
	describeEmpty(emptyI)

	emptyI = 42
	describeEmpty(emptyI)
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func describeEmpty(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
