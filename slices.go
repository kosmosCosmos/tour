package main

import "code.google.com/p/go-tour/pic"

// Implement Pic.
//
// It should return a slice of length dy, each element of which
// is a slice of dx 8-bit unsigned integers.
// When you run the program, it will display your picture,
// interpreting the integers as grayscale (well, bluescale) values.
// The choice of image is up to you. Interesting functions include x^y, (x+y)/2, and x*y.
// (You need to use a loop to allocate each []uint8 inside the [][]uint8.)
// (Use uint8(intValue) to convert between types.)
func Pic(dx, dy int) [][]uint8 {

	picture := make([][]uint8, dy)
	for y_index := range picture {
		picture[y_index] = make([]uint8, dx)
	}

	for y_index, picture_row := range picture {
		for x_index := range picture_row {
			picture_row[x_index] = uint8(x_index ^ y_index)
		}
	}
	return picture
}

func main() {
	pic.Show(Pic)
}
