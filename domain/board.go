package domain

import (
	"fmt"
	"math/rand"
)

type Board struct {
	Size int
	Grid [10][10]string
}

func NewBoard() *Board {
	b := &Board{Size: 10}

	for i := 0; i < b.Size; i++ {
		for j := 0; j < b.Size; j++ {
			b.Grid[i][j] = "~"
		}
	}
	return b
}

func (b *Board) PrintBoard() {
	fmt.Println("  A B C D E F G H I J")
	for r := 0; r < b.Size; r++ {
		fmt.Printf("%d ", r+1)
		for c := 0; c < b.Size; c++ {
			cell := b.Grid[r][c]
			switch cell {
			case "X":
				fmt.Print("\033[31mX \033[0m")
			case "*":
				fmt.Print("\033[34m* \033[0m")
			case "■":
				fmt.Print("\033[32m■ \033[0m")
			default:
				fmt.Print(". ")
			}
		}
		fmt.Println()
	}
}

func (b *Board) PlaceShipAuto(shipSize int) {
	placed := false
	for !placed {
		direction := rand.Intn(2)
		row := rand.Intn(b.Size)
		col := rand.Intn(b.Size)

		if direction == 0 && col+shipSize <= b.Size {
			canPlace := true
			for i := 0; i < shipSize; i++ {
				if b.Grid[row][col+i] != "~" {
					canPlace = false
					break
				}
			}
			if canPlace {
				for i := 0; i < shipSize; i++ {
					b.Grid[row][col+i] = "O"
				}
				placed = true
			}
		} else if direction == 1 && row+shipSize <= b.Size {
			canPlace := true
			for i := 0; i < shipSize; i++ {
				if b.Grid[row+i][col] != "~" {
					canPlace = false
					break
				}
			}
			if canPlace {
				for i := 0; i < shipSize; i++ {
					b.Grid[row+i][col] = "O"
				}
				placed = true
			}
		}
	}
}

func (b *Board) Shoot(row, col int) string {
	if b.Grid[row][col] == "O" {
		b.Grid[row][col] = "X"
		return "Попадание!"
	} else if b.Grid[row][col] == "~" {
		b.Grid[row][col] = "*"
		return "Промах!"
	} else {
		return "Уже стрелял сюда!"
	}
}

func (b *Board) AllShipsSunk() bool {
	for i := 0; i < b.Size; i++ {
		for j := 0; j < b.Size; j++ {
			if b.Grid[i][j] == "O" {
				return false
			}
		}
	}
	return true
}
