package encapsolution

import "testing"

type Employee struct {
	Id   string
	Name string
	Age  int
}

func TestCreateEmployeeObj(t *testing.T) {
	e := Employee{"0", "BOB", 20}
	e1 := Employee{Name: "josiah", Age: 34}
	t.Log(e, e1)
}
