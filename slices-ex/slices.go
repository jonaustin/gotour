package main

import (
	//"fmt"
	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
	board := make([][]uint8, dy)

	for y := range board {
		board[y] = make([]uint8, dx)
		for x := range board[y] {
			board[y][x] = uint8(x * y)
		}
	}

	//fmt.Println(board)

	return board
}

func main() {
	Pic(2, 4)
	pic.Show(Pic)
}
