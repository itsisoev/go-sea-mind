package domain

import "math/rand"

type Player struct {
	Name  string
	Board *Board
}

func RandomAIMove(b *Board) (int, int) {
	for {
		row := rand.Intn(b.Size)
		col := rand.Intn(b.Size)
		if b.Grid[row][col] != "X" && b.Grid[row][col] != "*" {
			return row, col
		}
	}
}
