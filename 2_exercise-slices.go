package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	res := make([][]uint8, dx)
	for x := 0; x < dx; x++ {
		res[x] = make([]uint8, dy)
		for y := 0; y < dy; y++ {
			res[x][y] = uint8((x ^ y) * (x + y) / 13)
		}
	}
	return res
}

func main() {
	pic.Show(Pic)
}
