package main

import "fmt"

func main() {
	s := []int{2, 5, 6, 7, 8, 9, 3}
	printSlice(s)
	s1 := s[:0]
	printSlice(s1)
	s2 := s[:4]
	printSlice(s2)
	s3 := s[2:]
	printSlice(s3)

	//切片的零值是 nil。
	fmt.Println("---------------------")
	var s4 []int
	fmt.Println(s4, len(s4), cap(s4))
	if s4 == nil {
		fmt.Println("nil!")
	}

	//向切片追加元素
	fmt.Println("---------------------")
	s4 = append(s4, 0)
	printSlice(s4)
	s4 = append(s4, 1, 6)
	printSlice(s4)
	s4 = append(s4, 2, 3, 4)
	printSlice(s4)

	//make v创建slice
	fmt.Println("---------------------")
	a := make([]int, 5) //5,5
	printSlice(a)
	b := make([]int, 0, 5) //0,5
	printSlice(b)
	c := b[:2] //2,5
	printSlice(c)
	d := c[2:5] //3,3
	printSlice(d)

	//range循环切片或者映射
	//返回索引以及值
	//可以将下标或值赋予 _ 来忽略它。
	fmt.Println("---------------------")
	for i, v := range s {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
