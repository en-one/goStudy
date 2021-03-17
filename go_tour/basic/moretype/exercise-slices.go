package main

//
//import "golang.org/x/tour/pic"
//
////一个 Pic 函数来生成一个 [][]uint8 的 2D 图片（即可说是 Array of Array）。它的大小由参数 (dx, dy int) 决定，这个有 dy 个数组，每个数组里又有一个长度为 dx 的数组。而相关的位置上 pic[y][x] 是这个图片的 bluescale（只有蓝色）数值，格式为 uint8。
//
//func Pic(dx, dy int) [][]uint8 {
//	pic := make([][]uint8, dy)
//	for i:= range pic {
//		pic[i] = make([]uint8, dx)
//		for j := range pic[i] {
//			pic[i][j] = uint8(i*i + j*j)
//		}
//	}
//	return pic
//}
//
//func main() {
//	pic.Show(Pic)
//}
