package main

import (
	"fmt" // first stdlib imports

	// after empty line 3rd-party imports

	// always use full import paths
	// imternal package is needed only for packages that is explicitly forbidden to use from outside
	"github.com/Kwbmm/TicTacGo/board" // after empty line local module imports
)

func main() {
	fmt.Println("Welcome to TicTacGo")

	currentPlayer := 'O'
	nextPlayer := 'X'

	// it's better not to use variable name the same as package name
	// too complex struct for simple task
	var b board.Board

	for {
		b.Print()

		fmt.Printf("Player '%c' turn\nChoose a cell from 1 to 9\n\n", currentPlayer)

		// add empty lines to separate logical parts and make code more readable

		choice := 0
		fmt.Scanf("%d", &choice)

		if choice < 1 || choice > 9 {
			fmt.Println("Invalid choice!")
			continue
		}

		// always consider flat code over nested

		if !b.MakeMove(choice-1, currentPlayer) {
			fmt.Println("Invalid move!")
			continue
		}

		if b.HasWinner() {
			// print result
			b.Print()

			fmt.Printf("Game won by %c\n", currentPlayer)
			return
		}

		if !b.HasEmptyCells() {
			fmt.Println("Game is draw")
			return
		}

		currentPlayer, nextPlayer = nextPlayer, currentPlayer // there is multiassignments in go
	}
}
