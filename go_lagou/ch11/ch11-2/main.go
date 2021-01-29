package main

import (
	"fmt"
)

func main() {
	coms := buy(10)       //采购10套配件
	phones := build(coms) //组装10部手机
	packs := pack(phones) //打包它们以便售卖
	//输出测试，看看效果
	for p := range packs {
		fmt.Println(p)
	}
}

func buy(n int) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for i := 1; i < 10; i++ {
			out <- fmt.Sprint("配件", i)
		}
	}()
	return out
}

//工序2组装
func build(in <-chan string) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for c := range in {
			out <- "组装(" + c + ")"
		}
	}()
	return out
}

//工序3打包
func pack(in <-chan string) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for c := range in {
			out <- "打包(" + c + ")"
		}
	}()
	return out
}
