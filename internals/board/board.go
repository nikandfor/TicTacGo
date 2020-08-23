package board

import (
	"fmt"
)

type boardFunctions interface {
	PrintBoard()
	printHeaderFooter()
	printFiller()
	printCell(index int) string
	SetCellValue(index int, value rune) bool
	HasWinner(index int, player rune) bool
	HasEmptyCells() bool
}

type Board struct {
	BoardState [9]Cell
}

func (board *Board) SetCellValue(index int, value rune) bool {
	return board.BoardState[index].setValue(value)
}

func (board *Board) HasWinner(index int, player rune) bool {
	state := board.BoardState
	switch index {
	case 0: //Top left
		//123
		if state[1].getValue() == player && state[2].getValue() == player {
			return true
		}
		//147
		if state[3].getValue() == player && state[6].getValue() == player {
			return true
		}
		//159
		if state[4].getValue() == player && state[8].getValue() == player {
			return true
		}
		return false
	case 1: //Top middle
		//123
		if state[0].getValue() == player && state[2].getValue() == player {
			return true
		}
		//258
		if state[4].getValue() == player && state[7].getValue() == player {
			return true
		}
		return false
	case 2: // Top right
		//123
		if state[0].getValue() == player && state[1].getValue() == player {
			return true
		}
		//369
		if state[5].getValue() == player && state[8].getValue() == player {
			return true
		}
		//357
		if state[4].getValue() == player && state[6].getValue() == player {
			return true
		}
		return false
	case 3: // Middle left
		//456
		if state[4].getValue() == player && state[5].getValue() == player {
			return true
		}
		//147
		if state[0].getValue() == player && state[6].getValue() == player {
			return true
		}
		return false
	case 4: //Middle middle
		//456
		if state[3].getValue() == player && state[5].getValue() == player {
			return true
		}
		//258
		if state[1].getValue() == player && state[7].getValue() == player {
			return true
		}
		//159
		if state[0].getValue() == player && state[8].getValue() == player {
			return true
		}
		//357
		if state[2].getValue() == player && state[6].getValue() == player {
			return true
		}
		return false
	case 5: //Middle right
		//456
		if state[3].getValue() == player && state[4].getValue() == player {
			return true
		}
		//369
		if state[2].getValue() == player && state[8].getValue() == player {
			return true
		}
		return false
	case 6: //Bottom left
		//147
		if state[0].getValue() == player && state[3].getValue() == player {
			return true
		}
		//789
		if state[7].getValue() == player && state[8].getValue() == player {
			return true
		}
		//753
		if state[4].getValue() == player && state[2].getValue() == player {
			return true
		}
		return false
	case 7: //Bottom middle
		//789
		if state[6].getValue() == player && state[8].getValue() == player {
			return true
		}
		//852
		if state[4].getValue() == player && state[1].getValue() == player {
			return true
		}
		return false
	case 8: //Bottom right
		//369
		if state[2].getValue() == player && state[5].getValue() == player {
			return true
		}
		//789
		if state[6].getValue() == player && state[7].getValue() == player {
			return true
		}
		//159
		if state[0].getValue() == player && state[4].getValue() == player {
			return true
		}
		return false
	default:
		return false
	}
}

func (board *Board) HasEmptyCells() bool {
	for _, cell := range board.BoardState {
		if cell.getValue() != 'O' && cell.getValue() != 'X' {
			return true
		}
	}
	return false
}

func (board *Board) printHeaderFooter() {
	fmt.Println("+-------|-------|-------+")
}

func (board *Board) printFiller() {
	fmt.Println("|       |       |       |")
}

func (board *Board) PrintBoard() {
	for i, c := range board.BoardState {
		indexInRow := i % 3
		switch indexInRow {
		case 0:
			board.printHeaderFooter()
			board.printFiller()
			fmt.Print("|")
			c.printCell()
		case 1:
			fmt.Print("|")
			c.printCell()
			fmt.Print("|")
		case 2:
			c.printCell()
			fmt.Print("|\n")
			board.printFiller()

		}
	}
	board.printHeaderFooter()
}
