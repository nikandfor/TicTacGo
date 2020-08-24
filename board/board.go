package board

import (
	"fmt"
)

// Do not use interfaces everywhere
type Board struct {
	State [9]Cell
}

func (b *Board) MakeMove(i int, p rune) bool {
	if i < 0 || i >= len(b.State) || b.State[i] != 0 {
		return false
	}

	b.State[i] = Cell(p)

	return true
}

// It's idiomatic to use short name for receiver arg
func (b *Board) HasWinner() bool {
	for i := 0; i < 3; i++ {
		// split big task on smaller and simpler tasks
		if b.checkCol(i) {
			return true
		}
		if b.checkRow(i) {
			return true
		}
	}

	// 012
	// 345
	// 678

	s := b.State

	if s[0] != 0 && s[0] == s[4] && s[0] == s[8] {
		return true
	}

	if s[2] != 0 && s[2] == s[4] && s[2] == s[6] {
		return true
	}

	return false
}

func (b *Board) checkRow(i int) bool {
	i *= 3
	return b.State[i] != 0 &&
		b.State[i] == b.State[i+1] && b.State[i] == b.State[i+2]
}

func (b *Board) checkCol(i int) bool {
	return b.State[i] != 0 &&
		b.State[i] == b.State[i+3] && b.State[i] == b.State[i+6]
}

// It's idiomatic to use short name for receiver arg
func (b *Board) HasEmptyCells() bool {
	for _, c := range b.State {
		if c != 'O' && c != 'X' {
			return true
		}
	}
	return false
}

func (b *Board) printHeaderFooter() {
	fmt.Println("+-------|-------|-------+")
}

func (b *Board) printFiller() {
	fmt.Println("|       |       |       |")
}

// board.PrintBoard // board is two times, it's excessively
func (b *Board) Print() {
	// consider flat code over nested
	// it's now more readable
	b.printHeaderFooter()

	b.printFiller()

	for i, c := range b.State {
		if c == 0 {
			c = Cell('0' + i + 1)
		}

		fmt.Printf("|   %c   ", c)

		if i%3 == 2 {
			fmt.Printf("|\n")
			b.printFiller()
		}
	}

	b.printHeaderFooter()
}
