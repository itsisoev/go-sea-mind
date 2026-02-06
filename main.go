package main

import (
	"fmt"
	"math/rand"
	"time"

	"SeaMind/domain"
	"SeaMind/service"
	"SeaMind/storage"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	exp, _ := storage.LoadExperience("storage/ai_experience.json")

	player := domain.Player{
		Name:  "Игрок",
		Board: domain.NewBoard(),
	}
	enemy := domain.Player{
		Name:  "ИИ",
		Board: domain.NewBoard(),
	}

	shipSizes := []int{4, 3, 3, 2, 2, 1}
	for _, size := range shipSizes {
		player.Board.PlaceShipAuto(size)
		enemy.Board.PlaceShipAuto(size)
	}

	fmt.Println("Добро пожаловать в SeaMind!")
	fmt.Println("Доска игрока:")
	player.Board.PrintBoard()

	for {
		var input string
		var row, col int
		for {
			fmt.Print("Введите координаты (например B5): ")
			fmt.Scan(&input)
			if len(input) < 2 {
				fmt.Println("Некорректный ввод!")
				continue
			}
			col = int(input[0] - 'A')
			row = int(input[1] - '1')
			if row < 0 || row >= 10 || col < 0 || col >= 10 {
				fmt.Println("Координаты вне диапазона!")
				continue
			}
			break
		}

		playerResult := enemy.Board.Shoot(row, col)

		aiRow, aiCol := service.SmartAIMove(player.Board, exp)
		aiResult := player.Board.Shoot(aiRow, aiCol)

		if aiResult == "Попадание!" {
			exp.Hits = append(exp.Hits, [2]int{aiRow, aiCol})
		} else if aiResult == "Промах!" {
			exp.Misses = append(exp.Misses, [2]int{aiRow, aiCol})
		}
		storage.SaveExperience("storage/ai_experience.json", exp)

		fmt.Printf("Ваш ход: %c%d → %s | ИИ ход: %c%d → %s\n",
			'A'+col, row+1, playerResult,
			'A'+aiCol, aiRow+1, aiResult)

		if enemy.Board.AllShipsSunk() {
			fmt.Println("Поздравляем! Игрок победил!")
			break
		}
		if player.Board.AllShipsSunk() {
			fmt.Println("ИИ победил! Попробуй в следующий раз.")
			break
		}

		fmt.Println("\nДоска игрока:")
		player.Board.PrintBoard()
	}
}
