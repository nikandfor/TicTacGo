package main

import (
	"fmt" // first stdlib imports

	// after empty line 3rd-party imports

	"github.com/Kwbmm/TicTacGo/internal/board" // after empty line local module imports
)

func main() {
	fmt.Println("Welcome to TicTacGo")

	currentPlayer := 'O'
	nextPlayer := 'X'

	board := board.Board{BoardState: [9]board.Cell{ // too complex struct for simple task
		board.Cell{Value: '1'},
		board.Cell{Value: '2'},
		board.Cell{Value: '3'},
		board.Cell{Value: '4'},
		board.Cell{Value: '5'},
		board.Cell{Value: '6'},
		board.Cell{Value: '7'},
		board.Cell{Value: '8'},
		board.Cell{Value: '9'},
	}}

	for {
		board.PrintBoard()

		fmt.Printf("Player '%c' turn\nChoose a cell from 1 to 9\n\n", currentPlayer)

		// add empty lines to separate logical parts and make code more readable

		choice := 0
		fmt.Scanf("%d", &choice)

		if choice < 1 || choice > 9 {
			fmt.Println("Invalid choice!")
			continue
		}

		// always consider flat code over nested

		if board.SetCellValue(choice-1, currentPlayer) {
			fmt.Println("Invalid move!")
			continue
		}

		if board.HasWinner(choice-1, currentPlayer) {
			fmt.Printf("Game won by %c\n", currentPlayer)
			return
		}

		if !board.HasEmptyCells() {
			fmt.Println("Game is draw")
			return
		}

		currentPlayer, nextPlayer = nextPlayer, currentPlayer // there is multiassignments in go
	}
}
