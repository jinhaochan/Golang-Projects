package main

import "fmt"

type board [][]string

type move struct {
	x int
	y int
}

// creates a new board
func newBoard() board {
	return board{
		{"_", "_", "_"},
		{"_", "_", "_"},
		{"_", "_", "_"}}
}

// adds a move to the board
// returns 1 if the move is valid
// returns -1 if the move is invalid
func (b board) addMove(p string, m move) int {
	// board placement already taken
	if b[m.x][m.y] == "_" {
		fmt.Println("Position already taken")
		return -1
	} else {
		b[m.x][m.y] = p
		return 0
	}
}

// function to check if the slice contains 3 of the same elements
// returns 1 if its does
// returns 0 if its does not
func checkThrees(moves []string) int {
	isThrees := 0

	// creating a map to store the moves
	elements := make(map[string]int)

	for _, move := range moves {
		// checking if the element exists
		_, ok := elements[move]
		if ok != true {
			elements[move] = 1
		} else {
			elements[move] += 1
			// if the total element count adds up to 3, set the flag = 1
			if elements[move] == 3 {
				isThrees = 1
			}
		}
	}
	return isThrees
}

// checks if there is a winner on the board
// happens after making a move
// checks the rows, columns and diagonals
// retuns 1 if there is a winner
// else returns -1 if there is none
func (b board) checkWin() int {
	winCondition := 0

	// checking rows
	for _, val := range b {
		winCondition = checkThrees(val)
		if winCondition == 1 {
			break
		}
	}

	// Transpose to get columns

	// Checking diagonals

	return winCondition

}
