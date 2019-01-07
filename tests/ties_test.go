package test

import (
	"testing"

	"github.com/Bios-Marcel/tictactoecli/internal"
)

func TestTie(t *testing.T) {
	game := internal.NewTicTacToe()

	//Start
	game.NumEntered(1)

	//Turns
	game.NumEntered(1)
	game.NumEntered(2)
	game.NumEntered(3)
	game.NumEntered(5)
	game.NumEntered(4)
	game.NumEntered(7)
	game.NumEntered(6)
	game.NumEntered(9)
	game.NumEntered(8)

	winner := *game.GetWinner()
	if winner != internal.None {
		t.Errorf("The winner was incorrect, it was %s, but should have been %s", winner, internal.None)
	}
}
