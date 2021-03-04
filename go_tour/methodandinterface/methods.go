package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

type MyFloat float64

//方法就是一类带特殊的（接收者参数）的函数
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

//没有接受者就是个普通函数
func NormalAbs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

//非结构体类型声明方法
func (f MyFloat) NotStructAbs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

//指针接收者
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {

	//方法就是一类带特殊的（接收者参数）的函数
	v := Vertex{3, 4}
	fmt.Println(v.Abs())

	//没有接受者就是个普通函数
	fmt.Println("------------------------")
	v1 := Vertex{3, 4}
	fmt.Println(NormalAbs(v1))

	fmt.Println("------------------------")
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.NotStructAbs())

	fmt.Println("------------------------")
	v2 := Vertex{3, 4}
	v2.Scale(10)
	fmt.Println(v2.Abs())

	//go语言内部做了处理
	//指针为接收者的方法调用时， 接收者既能为值又能为指针
	//普通函数的一个参数，则要保持指针对指针，值对值

}
