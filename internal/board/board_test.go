package board

import "testing"

var emptyBoard = Board{
	BoardState: [9]Cell{
		Cell{Value: ' '},
		Cell{Value: ' '},
		Cell{Value: ' '},
		Cell{Value: ' '},
		Cell{Value: ' '},
		Cell{Value: ' '},
		Cell{Value: ' '},
		Cell{Value: ' '},
		Cell{Value: ' '}}}

var partialBoard = Board{
	BoardState: [9]Cell{
		Cell{Value: ' '},
		Cell{Value: 'O'},
		Cell{Value: ' '},
		Cell{Value: 'X'},
		Cell{Value: 'X'},
		Cell{Value: ' '},
		Cell{Value: 'O'},
		Cell{Value: ' '},
		Cell{Value: ' '}}}

var fullBoardNoWinner = Board{
	BoardState: [9]Cell{
		Cell{Value: 'O'},
		Cell{Value: 'X'},
		Cell{Value: 'O'},
		Cell{Value: 'X'},
		Cell{Value: 'X'},
		Cell{Value: 'O'},
		Cell{Value: 'X'},
		Cell{Value: 'O'},
		Cell{Value: 'X'}}}

var fullBoardWinnerIsO = Board{
	BoardState: [9]Cell{
		Cell{Value: 'X'},
		Cell{Value: 'O'},
		Cell{Value: 'O'},
		Cell{Value: 'X'},
		Cell{Value: 'O'},
		Cell{Value: 'X'},
		Cell{Value: 'O'},
		Cell{Value: 'O'},
		Cell{Value: 'X'}}}

var fullBoardWinnerIsX = Board{
	BoardState: [9]Cell{
		Cell{Value: 'O'},
		Cell{Value: 'X'},
		Cell{Value: 'X'},
		Cell{Value: 'O'},
		Cell{Value: 'X'},
		Cell{Value: 'O'},
		Cell{Value: 'X'},
		Cell{Value: 'X'},
		Cell{Value: 'O'}}}

func TestWhenBoardHasEmptyCellsShouldReturnTrue(t *testing.T) {
	//Given
	board := emptyBoard
	expectedResult := true

	//Test
	result := board.HasEmptyCells()

	//Verify
	if result != expectedResult {
		t.Errorf("Board is full! Expected %t but got %t", expectedResult, result)
	}

	//Given
	board = partialBoard
	expectedResult = true

	//Test
	result = board.HasEmptyCells()

	//Verify
	if result != expectedResult {
		t.Errorf("Board is full! Expected %t but got %t", expectedResult, result)
	}
}

func TestWhenBoardHasNoEmptyCellsShouldReturnFalse(t *testing.T) {
	//Given
	board := fullBoardNoWinner
	expectedResult := false

	//Test
	result := board.HasEmptyCells()

	//Verify
	if result != expectedResult {
		t.Errorf("Board is full! Expected %t but got %t", expectedResult, result)
	}
}

func TestWhenThereIsWinnerShouldReturnTrue(t *testing.T) {
	// //Given
	// players := []rune{'O', 'X'}
	// var boards map[string][2]Board
	// var boardWinner123,	boardWinner147,	boardWinner159,	boardWinner258,	boardWinner357,	boardWinner369,	boardWinner456,	boardWinner789 Board

	// for index, player := range players {
	// 	boards['123'][index] = Board{
	// 		{
	// 			BoardState: [9]Cell{
	// 				{Value: player},
	// 				{Value: player},
	// 				{Value: player},
	// 				{Value: 'X'},
	// 				{Value: player},
	// 				{Value: 'X'},
	// 				{Value: player},
	// 				{Value: 'X'},
	// 				{Value: 'X'},
	// 			}
	// 		},

	// 	}
	// }

}

func TestWhenSettingEmptyCellShouldReturnTrue(t *testing.T) {
	//Given
	expectedPlayerSet := 'O'
	expectedResult := true
	var board = partialBoard

	//Test
	result := board.SetCellValue(0, expectedPlayerSet)
	playerSet := board.BoardState[0].getValue()

	//Verify
	if result != expectedResult || playerSet != expectedPlayerSet {
		t.Errorf("Expecting result = %t and player = %c\nGot result = %t and player = %c",
			expectedResult, expectedPlayerSet,
			result, playerSet)
	}
}

func TestWhenSettingTakenCellShouldReturnFalse(t *testing.T) {
	//Given
	expectedPlayerSet := 'O'
	expectedResult := false
	var board = partialBoard

	//Test
	result := board.SetCellValue(1, expectedPlayerSet)
	playerSet := board.BoardState[1].getValue()

	//Verify
	if result != expectedResult || playerSet != expectedPlayerSet {
		t.Errorf("Expecting result = %t and player = %c\nGot result = %t and player = %c",
			expectedResult, expectedPlayerSet,
			result, playerSet)
	}
}