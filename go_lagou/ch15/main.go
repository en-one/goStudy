package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	p := person{Name: "josiah", Age: 24}
	//ppv := reflect.ValueOf(&p)
	//ppv.Elem().Field(0).SetString("zhangsan")
	//ppv.Elem().Field(1).SetInt(34)
	//fmt.Print(p)
	jsonB, err := json.Marshal(p)
	if err == nil {
		fmt.Println(string(jsonB))
	}

	respJSON := "{\"Name\":\"李四\",\"Age\":40}"
	json.Unmarshal([]byte(respJSON), &p)
	fmt.Println(p)

}

type person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}
