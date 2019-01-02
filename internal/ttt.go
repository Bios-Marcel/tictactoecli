package internal

import "fmt"

//TicTacToe contains the games state.
type TicTacToe struct {
	currentState   turns
	currentWorld   world
	inputAndOutput bool
}

//NewTicTacToe creates a new instance of the tictactoe game.
func NewTicTacToe() *TicTacToe {
	newTicTacToe := TicTacToe{}
	newTicTacToe.currentState = &splash{context: &newTicTacToe}
	newTicTacToe.currentWorld = world{board: make(map[int]Owner)}

	return &newTicTacToe
}

//LaunchTicTacToeGame creates a new instance of the tictactoe game.
//This instance allows human input and shows output. The game starts
//immediately after calling this method.
func LaunchTicTacToeGame() {
	newTicTacToe := TicTacToe{}
	newTicTacToe.inputAndOutput = true
	newTicTacToe.currentState = &splash{context: &newTicTacToe}
	newTicTacToe.currentWorld = world{board: make(map[int]Owner)}

	fmt.Println("Welcome to TicTacToe CLI, please input anything to start the game.")
	var trash string
	fmt.Scan(&trash)
	newTicTacToe.NumEntered(1)
}

func getInput() int {
	var input int
	fmt.Scan(&input)
	return input
}

//NumEntered is used for making a turn.
func (t *TicTacToe) NumEntered(field int) {
	t.currentState.numEntered(field)
	if t.inputAndOutput {
		t.currentWorld.print()
		t.NumEntered(getInput())
	}
}

//GetWinner primarily exists for calling it from a unit test
func (t *TicTacToe) GetWinner() *Owner {
	return t.currentWorld.getWinner()
}

func printEndResult(player Owner) {
	if player == None {
		fmt.Println("Noone has won, it is a tie!")
	} else if player == X {
		fmt.Println("Player X has won!")
	} else if player == O {
		fmt.Println("Player O has won!")
	} else {
		fmt.Println("The game is still going, keep fighting fiercly!")
	}
}
