package constant_test

import "testing"

const (
	monday = 1 + iota
	Tuesday
	Wednesday
)

const (
	Readable = 1 << iota
	Writable
	Executable
)

func TestConstantTry(t *testing.T) {
	t.Log(monday, Wednesday)
	t.Log(Writable, Executable)
}

func TestConstantTry1(t *testing.T) {
	a := 7
	t.Log(a & Readable)
}
