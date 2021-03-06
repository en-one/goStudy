package main

import (
	"fmt"
	"math"
)

//接口类型是由一组方法签名定义的集合
type Abser interface {
	Abs() float64
}

func main() {
	var a Abser
	f := MyFloatt(-math.Sqrt2)
	//v := &Vertexx{3,4}
	//v := Vertexx{3,4}
	a = f
	//a = v

	fmt.Println(a.Abs())
}

type MyFloatt float64

func (f MyFloatt) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertexx struct {
	X, Y int
}

func (v *Vertexx) Abs() float64 {
	return math.Sqrt(float64(v.X*v.X + v.Y*v.Y))
}
