package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

var (
	v1 = Vertex{1, 2}  // 创建一个 Vertex 类型的结构体
	v2 = Vertex{X: 1}  // Y:0 被隐式地赋予
	v3 = Vertex{}      // X:0 Y:0
	p1 = &Vertex{1, 2} // 创建一个 *Vertex 类型的结构体（指针）
)

func main() {
	fmt.Println(Vertex{1, 2})

	//结构体类似于php对象 使用.来访问
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X)

	//指针p指向结构体v 可以直接使用指针
	v0 := Vertex{1, 2}
	p := &v0
	p.X = 1e9
	fmt.Println(v)

	fmt.Println(v1, p1, v2, v3)
}
