package _func

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

//import "testing"

func returnMultiValues() (int, int) {
	return rand.Intn(10), rand.Intn(20)
}

func TestFn(t *testing.T) {
	a, b := returnMultiValues()
	t.Log(a, b)
}

func timeSpend(inner func(op int) int) func(op int) int {
	return func(n int) int {
		start := time.Now()
		ret := inner(n)
		fmt.Println("time spend", time.Since(start).Seconds())
		return ret
	}
}

func Sum(ops ...int) int {
	ret := 0
	for _, op := range ops {
		ret += op
	}
	return ret
}

func TestVarParam(t *testing.T) {
	t.Log(Sum(1, 2, 3))
	t.Log(Sum(2, 3, 5))

}

func TestDefer(t *testing.T) {
	defer func() {
		t.Log("clear resources")
	}()
	t.Log("started")
	panic("Fatal error")
}
