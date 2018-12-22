package internal

import "fmt"

//TicTacToe contains the games state.
type TicTacToe struct {
	currentState turns
	currentWorld world
}

//NewTicTacToe creates a new instance of the tictactoe game.
func NewTicTacToe() *TicTacToe {
	newTicTacToe := TicTacToe{}
	newTicTacToe.currentState = &splash{context: &newTicTacToe}
	newTicTacToe.currentWorld = world{board: make(map[int]owner)}

	return &newTicTacToe
}

//NumEntered is used for making a turn.
func (t *TicTacToe) NumEntered(field int) {
	t.currentState.numEntered(field)
	t.currentWorld.print()
}

func printEndResult(player owner) {
	if player == none {
		fmt.Println("Noone has won, it is a tie!")
	} else if player == x {
		fmt.Println("Player X has won!")
	} else if player == o {
		fmt.Println("Player O has won!")
	} else {
		fmt.Println("The game is still going, keep fighting fiercly!")
	}
}
