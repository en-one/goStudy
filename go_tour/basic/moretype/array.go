package main

import "fmt"

func main() {

	var a [2]string
	a[0] = "hello"
	a[1] = "go"

	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 14}
	fmt.Println(primes)

	//切片, 不定义长度
	//切片的默认行为忽略上下界
	var s []int = primes[1:4]
	s1 := primes[:2]
	fmt.Println(s, s1)

	//切片并不存储任何数据 他只是描述了底层数组的一段
	//更改切片元素会更改底层数组中对应的元素
	names := [4]string{"josiah", "jenny", "greogre", "ringo"}
	fmt.Println(names)
	c := names[0:2]
	d := names[1:3]
	fmt.Println(c, d)
	d[0] = "xxx"
	fmt.Println(c, d)
	fmt.Println(names)

	//切片文法：初始化的时候复制
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	e := []struct {
		i int
		f bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(e)
}
