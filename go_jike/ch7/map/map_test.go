package map_test

import "testing"

func TestInitMap(t *testing.T) {
	m1 := map[int]int{1: 1, 2: 2, 3: 4}
	t.Log(m1[2])
	t.Log(len(m1))

	m2 := map[int]int{}
	m2[4] = 6
	t.Log(m2)
	t.Log(len(m2))
}

//通过两个变量判断该值为0 是不存在go默认为0 还是本身就是0

func TestAccessNotExistingKey(t *testing.T) {
	m1 := map[int]int{}
	t.Log(m1[1])

	//m1[2] = 0
	if v, ok := m1[2]; ok {
		t.Log("存在是0", v)
	} else {
		t.Log("不存在 默认为0")
	}

}
