package main

import "fmt"

//func main() {
//	sum := sum(7,8,8)
//	fmt.Print(sum)
//}
//
//func sum(params ...int) int {
//	sum :=0;
//	for _,i :=range params {
//		sum +=i
//	}
//	return sum
//}
func main() {
	cl := colsure()
	fmt.Println(cl())
	fmt.Println(cl())
	fmt.Println(cl())
	fmt.Println(colsure())
	fmt.Println(colsure())
	fmt.Println(colsure())
}

func colsure() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
