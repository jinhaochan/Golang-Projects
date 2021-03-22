package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	movePattern := [2]string{"X", "O"}

	var pat string

	b := newBoard()

	b.printBoard()

	winner := 0
	moveCounter := 0

	scanner := bufio.NewScanner(os.Stdin)

	for winner == 0 && moveCounter < 9 {
		pat = movePattern[moveCounter%2]

		fmt.Println(pat, "Turn")

		scanner.Scan()
		input := scanner.Text()

		s := strings.Split(input, " ")

		if inputCheck(s) == 1 {
			xc, _ := strconv.Atoi(s[0])
			yc, _ := strconv.Atoi(s[1])

			m := move{x: xc, y: yc}

			// only adds the move and advances to the next player if the current move is valid
			if b.validCheck(m) == 1 {

				b.addMove(pat, m)
				winner = b.checkWin()
				moveCounter += 1
			}
		}

		b.printBoard()

	}

	if winner == 1 {
		fmt.Println(pat, "is the winner!")
	} else if moveCounter == 9 {
		fmt.Println("Ended in a draw!")
	}

}
