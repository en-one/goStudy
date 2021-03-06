package main

import "fmt"

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

func main() {
	//类型断言 提供了访问接口之底层具体值的方法	···
	//格式 i.(string)
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)
	s, ok := i.(string) // hello  true
	fmt.Println(s, ok)
	f, ok := i.(float64) //有ok  0 false
	fmt.Println(f, ok)
	f = i.(float64) //没有ok 报错(panic)
	fmt.Println(f)

	//格式v := i.(type)
	do(21)
	do("hello")
	do(true)

}
