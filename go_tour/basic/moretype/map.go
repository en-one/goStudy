package main

import "fmt"

type Vertux struct {
	Lat, Long float64
}

//var m map[string]Vertux

//文法 即初始化
var m = map[string]Vertux{
	"Bell Labs": Vertux{
		40.68433, -74.39967,
	},
	//Vertux 可省略
	"Google": {
		37.42202, -122.08408,
	},
}

func main() {
	//m = make(map[string]Vertux)
	m["josiah"] = Vertux{33.23, 44.32}
	m["Google"] = Vertux{33.33, 55.55}

	fmt.Println(m)
	delete(m, "josiah")
	fmt.Println(m)
	v, ok := m["josiah"]
	fmt.Println(v, ok)

}
