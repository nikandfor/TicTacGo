package board

import (
	"fmt"
)

type cellFunctions interface {
	getValue() rune
	setValue(value rune) bool
	printCell()
}

type Cell struct {
	Value rune
}

func (cell Cell) getValue() rune {
	return cell.Value
}

func (cell *Cell) setValue(value rune) bool{
	if cell.Value != 'X' && cell.Value != 'O' {
		cell.Value = value
		return true
	}
	return false
}

func (cell *Cell) printCell() {
	fmt.Printf("   %c   ", cell.getValue())
}
