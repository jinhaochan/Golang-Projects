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
func (b board) addMove(p string, m move) {
	b[m.x][m.y] = p
}

// function to check if the slice contains 3 of the same elements other than '_'
// returns 1 if its does
// returns 0 if its does not
func checkThrees(collection []string) int {
	isThrees := 0

	// creating a map to store the moves in the collection
	elements := make(map[string]int)

	for _, pat := range collection {
		if pat != "_" {
			// checking if the element exists
			_, ok := elements[pat]
			if !ok {
				elements[pat] = 1
			} else {
				elements[pat] += 1
				// if the total element count adds up to 3, set the flag = 1
				if elements[pat] == 3 {
					isThrees = 1
				}
			}
		}
	}
	return isThrees
}

// checks if there is a winner on the board
// happens after making a move
// checks the rows, columns and diagonals
// retuns 1 if there is a winner
// else returns 0 if there is none
func (b board) checkWin() int {
	winConditions := []int{}

	// checking rows
	for _, val := range b {
		winConditions = append(winConditions, checkThrees(val))
	}

	// Transpose to get columns
	tb := transpose(b)

	for _, val := range tb {
		winConditions = append(winConditions, checkThrees(val))
	}

	// Checking diagonals
	d1 := []string{b[0][0], b[1][1], b[2][2]}
	d2 := []string{b[0][2], b[1][1], b[2][0]}
	winConditions = append(winConditions, checkThrees(d1))
	winConditions = append(winConditions, checkThrees(d2))

	maxVal := 0

	// if there is a winner, the value would 1, which is larger maxVal = 0
	// we return that value to indicate if there is a winner or not
	for _, val := range winConditions {
		if val > maxVal {
			maxVal = val
		}
	}

	return maxVal
}

func transpose(s board) board {
	t := newBoard()

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			t[i][j] = s[j][i]
		}
	}

	return t
}

// checks if the move is valid
// checks if the coordinates given are within range 0 to 2
// checks if the position is empty
func (b board) validCheck(m move) int {
	isValid := 1

	if m.x > 2 || m.y > 2 {
		fmt.Println("Invalid coordinates!")
		isValid = 0
	} else if b[m.x][m.y] != "_" {
		fmt.Println("Position already taken")
		isValid = 0
	}

	return isValid
}

func inputCheck(s []string) int {
	isValid := 1

	if len(s) != 2 {
		fmt.Println("Wrong number of coordinates given. Expected 2")
		isValid = 0
	}
	return isValid
}

func (b board) printBoard() {
	for i := range b {
		fmt.Println(b[i])
	}
}
