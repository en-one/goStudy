package main

import (
	"fmt"
	"math"
)

func main() {

	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))

	var z uint = uint(f)
	fmt.Println(x, y, z)

	//未指定变量类型时，由右键推导
	v := 42
	fmt.Printf("v is of type %T\n", v)

}
