package operator_test

import "testing"

//按位清0
//右边是1 结果都是0
//右边是0 左边是什么 结果就是什么
func TestCompareArray(t *testing.T) {

	t.Log(1 &^ 1)
	t.Log(0 &^ 1)
	t.Log(1 &^ 0)
	t.Log(0 &^ 0)
}
