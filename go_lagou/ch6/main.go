package main

import "fmt"

type person struct {
	name string
	age  uint
	add  address
}

type address struct {
	province string
	city     string
}

func main() {
	var p person
	p = person{name: "josiah", age: 26, add: address{
		province: "ss",
		city:     "afds",
	}}
	fmt.Print(p)
}
