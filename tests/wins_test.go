package test

import (
	"testing"

	"github.com/Bios-Marcel/tictactoecli/internal"
)

func TestXWins1And2And3(t *testing.T) {
	game := internal.NewTicTacToe()

	//Start
	game.NumEntered(1)

	//Turns
	game.NumEntered(1)
	game.NumEntered(4)
	game.NumEntered(2)
	game.NumEntered(5)
	game.NumEntered(3)

	winner := *game.GetWinner()
	if winner != internal.X {
		t.Errorf("The winner was incorrect, it was %s, but should have been %s", winner, internal.X)
	}
}

func TestXWins4And5And6(t *testing.T) {
	game := internal.NewTicTacToe()

	//Start
	game.NumEntered(1)

	//Turns
	game.NumEntered(4)
	game.NumEntered(1)
	game.NumEntered(5)
	game.NumEntered(2)
	game.NumEntered(6)

	winner := *game.GetWinner()
	if winner != internal.X {
		t.Errorf("The winner was incorrect, it was %s, but should have been %s", winner, internal.X)
	}
}

func TestXWins7And8And9(t *testing.T) {
	game := internal.NewTicTacToe()

	//Start
	game.NumEntered(1)

	//Turns
	game.NumEntered(7)
	game.NumEntered(1)
	game.NumEntered(8)
	game.NumEntered(2)
	game.NumEntered(9)

	winner := *game.GetWinner()
	if winner != internal.X {
		t.Errorf("The winner was incorrect, it was %s, but should have been %s", winner, internal.X)
	}
}

func TestXWins1And4And7(t *testing.T) {
	game := internal.NewTicTacToe()

	//Start
	game.NumEntered(1)

	//Turns
	game.NumEntered(1)
	game.NumEntered(2)
	game.NumEntered(4)
	game.NumEntered(3)
	game.NumEntered(7)

	winner := *game.GetWinner()
	if winner != internal.X {
		t.Errorf("The winner was incorrect, it was %s, but should have been %s", winner, internal.X)
	}
}

func TestXWins2And5And8(t *testing.T) {
	game := internal.NewTicTacToe()

	//Start
	game.NumEntered(1)

	//Turns
	game.NumEntered(2)
	game.NumEntered(1)
	game.NumEntered(5)
	game.NumEntered(4)
	game.NumEntered(8)

	winner := *game.GetWinner()
	if winner != internal.X {
		t.Errorf("The winner was incorrect, it was %s, but should have been %s", winner, internal.X)
	}
}

func TestXWins3And6And9(t *testing.T) {
	game := internal.NewTicTacToe()

	//Start
	game.NumEntered(1)

	//Turns
	game.NumEntered(3)
	game.NumEntered(1)
	game.NumEntered(6)
	game.NumEntered(4)
	game.NumEntered(9)

	winner := *game.GetWinner()
	if winner != internal.X {
		t.Errorf("The winner was incorrect, it was %s, but should have been %s", winner, internal.X)
	}
}

func TestXWins1And5And9(t *testing.T) {
	game := internal.NewTicTacToe()

	//Start
	game.NumEntered(1)

	//Turns
	game.NumEntered(1)
	game.NumEntered(4)
	game.NumEntered(5)
	game.NumEntered(3)
	game.NumEntered(9)

	winner := *game.GetWinner()
	if winner != internal.X {
		t.Errorf("The winner was incorrect, it was %s, but should have been %s", winner, internal.X)
	}
}

func TestXWins3And5And7(t *testing.T) {
	game := internal.NewTicTacToe()

	//Start
	game.NumEntered(1)

	//Turns
	game.NumEntered(3)
	game.NumEntered(1)
	game.NumEntered(5)
	game.NumEntered(4)
	game.NumEntered(7)

	winner := *game.GetWinner()
	if winner != internal.X {
		t.Errorf("The winner was incorrect, it was %s, but should have been %s", winner, internal.X)
	}
}
