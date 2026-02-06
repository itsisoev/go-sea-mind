package service

import (
	"SeaMind/domain"
	"SeaMind/storage"
	"math/rand"
)

func SmartAIMove(board *domain.Board, exp *storage.Experience) (int, int) {
	size := board.Size
	used := map[[2]int]bool{}

	for _, h := range exp.Hits {
		used[h] = true
	}
	for _, m := range exp.Misses {
		used[m] = true
	}

	for _, hit := range exp.Hits {
		dirs := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
		for _, d := range dirs {
			nr, nc := hit[0]+d[0], hit[1]+d[1]
			if nr >= 0 && nr < size && nc >= 0 && nc < size {
				if !used[[2]int{nr, nc}] && board.Grid[nr][nc] != "X" && board.Grid[nr][nc] != "*" {
					return nr, nc
				}
			}
		}
	}

	for {
		row := rand.Intn(size)
		col := rand.Intn(size)
		pos := [2]int{row, col}
		if !used[pos] && board.Grid[row][col] != "X" && board.Grid[row][col] != "*" {
			return row, col
		}
	}
}
