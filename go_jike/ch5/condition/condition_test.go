package condition

import "testing"

func TestSwitchMultiCase(t *testing.T) {
	for i := 0; i < 5; i++ {
		switch i {
		case 0, 2:
			t.Log("Even")
		case 1, 3:
			t.Log("odd")
		default:
			t.Log(" it is not 0-3")
		}
	}
}
