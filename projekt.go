package main

import (
	"fmt"
	"math/rand"
	"time"
)

var board [3][3]string

func main() {
	var endGame string
	var mode int
	var checkMode bool
	var rounds int
	fmt.Printf("Welcome to the Tic-Tac-Toe game\n")
	fmt.Printf("This is how the board looks lke\n")
	fmt.Printf(" (0,0) | (0,1) | (0,2)\n")
	fmt.Printf("-------+-------+------\n")
	fmt.Printf(" (1,0) | (1,1) | (1,2)\n")
	fmt.Printf("-------+-------+------\n")
	fmt.Printf(" (2,0) | (2,1) | (2,2)\n")
	fmt.Printf("\nWe have 3 mode types to choose from\n")
	fmt.Printf("1. Player 1 vs Player 2\n2. Player 1 vs AI\n3. AI vs AI \n\nChoose the mode: ")
	fmt.Scanf("%d", &mode)
	fmt.Printf("Let's begin!\n")

	checkMode = false
	for checkMode == false {
		if mode != 1 && mode != 2 && mode != 3 {
			fmt.Printf("Invalid mode. Try once again:\n")
			fmt.Scanf("%d", &mode)
		} else {
			checkMode = true
		}
	}

	rounds = 0
	endGame = " "
	inicialiseBoard()
	switch mode {
	case 1:
		for endGame == " " {
			printBoard()
			fmt.Printf("\nPlayer 1 move:\n")
			playerMove()
			printBoard()
			rounds++
			endGame = checkForWin() // checking if player 1 already won
			if endGame != " " {
				break // if he won -> end game
			}
			if rounds == 9 {
				break // if we've got to the end and none has won, we are ending it as a draw
			}
			fmt.Printf("\nPlayer 2 move:\n")
			playerMove2()
			fmt.Printf("\n")
			endGame = checkForWin() // checking if player 2 already won
			rounds++
		}
		if endGame == "X" {
			fmt.Printf("*Player 1 wins!*\n")
		} else if endGame == "O" {
			fmt.Printf("*Player 2 wins*\n")
			printBoard()
		} else {
			fmt.Printf("\n\n*DRAW!*")
		}

	case 2:
		for endGame == " " {
			printBoard()
			playerMove()
			printBoard()
			rounds++
			endGame = checkForWin() 
			if endGame != " " {
				break 
			}
			if rounds == 9 {
				break
			}
			AIMove()
			fmt.Printf("\n")
			endGame = checkForWin() // checking if AI won
			rounds++
		}
		if endGame == "X" {
			fmt.Printf("*Player 1 wins!*\n")
		} else if endGame == "O" {
			fmt.Printf("*AI wins!*\n")
			printBoard()
		} else {
			fmt.Printf("\n\n*DRAW!*")
		}

	case 3:
		for endGame == " " {
			printBoard()
			time.Sleep(1200 * time.Millisecond)
			AIMove2()
			printBoard()
			rounds++
			endGame = checkForWin() 
			if endGame != " " {
				break 
			}
			if rounds == 9 {
				break
			}
			time.Sleep(1200 * time.Millisecond)
			AIMove()
			fmt.Printf("\n")
			endGame = checkForWin() 
			rounds++
		}
		if endGame == "X" {
			fmt.Printf("*X wins!*\n")
		} else if endGame == "O" {
			fmt.Printf("*O wins*\n")
			printBoard()
		} else {
			fmt.Printf("\n\nDRAW!")
		}
	}
}


func inicialiseBoard() {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			board[i][j] = " "
		}
	}
}


func printBoard() {
	for i := 0; i < 3; i++ {
		fmt.Printf(" %s | %s | %s ", board[i][0], board[i][1], board[i][2])
		if i != 2 {
			fmt.Printf("\n---|---|---\n")
		}
	}
	fmt.Printf("\n")
}

// Get player input 
func playerMove() {
	var x int
	var y int
	var correctMove bool
	fmt.Printf("Enter (x,y): ")
	fmt.Scanf("%d%d", &x, &y)
	fmt.Printf("\n")

	correctMove = false

	for correctMove == false {
		if x > 2 || x < 0 || y < 0 || y > 2 {
			fmt.Printf("Invalid move. Try once again: ")
			fmt.Scanf("%d%d", &x, &y)
			fmt.Printf("\n")
		}
		if board[x][y] != " " {
			fmt.Printf("This position is taken. Try once again: ")
			fmt.Scanf("%d%d", &x, &y)
			fmt.Printf("\n")
		} else {
			board[x][y] = "X"
			correctMove = true
		}
	}
}

func playerMove2() {
	var x int
	var y int
	var correctMove bool
	fmt.Printf("Enter (x,y): ")
	fmt.Scanf("%d%d", &x, &y)
	fmt.Printf("\n")

	correctMove = false

	for correctMove == false {
		if x > 2 || x < 0 || y < 0 || y > 2 {
			fmt.Printf("Invalid move. Try once again: ")
			fmt.Scanf("%d%d", &x, &y)
			fmt.Printf("\n")
		}
		if board[x][y] != " " {
			fmt.Printf("This position is taken. Try once again: ")
			fmt.Scanf("%d%d", &x, &y)
			fmt.Printf("\n")
		} else {
			board[x][y] = "O"
			correctMove = true
		}
	}
}

// Get AI move
func AIMove() {
	fmt.Printf("AI move\n")
	for {
		rand.Seed(time.Now().UTC().UnixNano())
		choice := rand.Int()%9 + 1
		x := (choice - 1) / 3
		y := (choice - 1) % 3 
		var boardPosition string = board[x][y]

		if boardPosition == "X" || boardPosition == "O" {
			continue
		} else {
			fmt.Printf("AI chose: %d\n", choice)
			board[x][y] = "O"
			break
		}
	}
}

func AIMove2() {
	fmt.Printf("AI move:\n")
	for {
		rand.Seed(time.Now().UTC().UnixNano())
		choice := rand.Int()%9 + 1
		x := (choice - 1) / 3
		y := (choice - 1) % 3 
		var boardPosition string = board[x][y]

		if boardPosition == "X" || boardPosition == "O" {
			continue
		} else {
			fmt.Printf("AI chose: %d\n", choice)
			board[x][y] = "X"
			break
		}
	}
}


func checkForWin() string {
	// check rows 
	for i := 0; i < 3; i++ {
		if board[i][0] == board[i][1] && board[i][1] == board[i][2] {
			return board[i][0]
		}
	} // check columns
	for i := 0; i < 3; i++ {
		if board[0][i] == board[1][i] && board[1][i] == board[2][i] {
			return board[0][i]
		}
	} //test diagonal
	if board[0][0] == board[1][1] && board[1][1] == board[2][2] {
		return board[0][0]
	}
	if board[0][2] == board[1][1] && board[1][1] == board[2][0] {
		return board[0][2]
	}
	return " "
}


