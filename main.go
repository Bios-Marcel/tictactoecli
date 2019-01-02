package main

import "github.com/Bios-Marcel/tictactoecli/internal"

func main() {
	game := internal.NewTicTacToeWithInputAndOutput()

	//Start
	game.NumEntered(0)

	//Turns
	game.NumEntered(5)
	game.NumEntered(4)
	game.NumEntered(3)
	game.NumEntered(7)
	game.NumEntered(1)
	game.NumEntered(2)
	game.NumEntered(9)
}
