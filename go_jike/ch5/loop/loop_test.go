package loop

import "testing"

func TestIfMultiSec(t *testing.T) {
	if v, err := someone(); err == nil {
		t.Log(v)
	} else {
		t.Log(err)
	}
}
