package slice

import (
	"testing"
)

func TestSliceInit(t *testing.T) {
	var s0 []int
	t.Log(len(s0), cap(s0))
	s0 = append(s0, 1)
	t.Log(len(s0), cap(s0))

	s1 := []int{1, 2, 3, 4}
	t.Log(len(s1), cap(s1))

	s2 := make([]int, 3, 5)
	t.Log(len(s2), cap(s2))

}

//slice 的存储空间cap是成倍增加的
func TestSliceGrowing(t *testing.T) {
	s1 := []int{}
	for i := 1; i < 10; i++ {
		s1 = append(s1, i)
		t.Log(len(s1), cap(s1))
	}

}

//切片共享内存地址
func TestSliceShareMemory(t *testing.T) {
	year := []string{"Jan", "Feb", "web", "Tues", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "NOv", "Dec"}

	Q2 := year[3:6]
	t.Log(len(Q2), cap(Q2))
	summer := year[5:8]
	t.Log(len(summer), cap(summer))

	summer[0] = "unknow"
	t.Log(year, Q2, summer)

}
