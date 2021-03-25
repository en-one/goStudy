package type_test

import "testing"

type MyInt int64

func TestImpLicit(t *testing.T) {
	//var a int=1
	var a int32 = 1
	var b int64

	b = int64(a)
	t.Log(a, b)

	var c MyInt
	c = MyInt(b)
	t.Log(a, b, c)
}

//Go不支持指针运算 其他语言会使用指针去访问数组连续内存空间
func TestPoint(t *testing.T) {
	a := 1
	aPtr := &a
	//aptr = aPtr + 1 //错误
	t.Log(a, aPtr)
	t.Logf("%T %T", a, aPtr)
}
