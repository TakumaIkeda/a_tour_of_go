package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	s := [][]uint8{}
	for i := 0; i < dx; i++ {
		s = append(s, []uint8{})
		for j := 0; j < dy; j++ {
			s[i] = append(s[i], uint8((i+j)/2))
		}
	}
	return s
}

func main() {
	pic.Show(Pic)
}