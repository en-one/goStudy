package array

import "testing"

func TestArrayInit(t *testing.T) {
	var arr [3]int
	arr1 := [4]int{1, 2, 3, 4}
	arr3 := [...]int{12, 14}

	arr1[2] = 5
	t.Log(arr[1])
	t.Log(arr1, arr3)
}

func TestArrayTravel(t *testing.T) {
	arr3 := [...]int{12, 14}

	for idx, e := range arr3 {
		t.Log(idx, e)
	}
}

func TestArratSection(t *testing.T) {
	arr3 := [...]int{12, 14}
	arr4 := arr3[1:]
	t.Log(arr4)
}
