package main

import (
	"fmt"
)

func main() {
	name := "josiah"
	nameP := &name
	fmt.Println(name)
	fmt.Println(nameP)

	*nameP = "josiahah"
	fmt.Println(*nameP)
	fmt.Println(name)

}
