package internal

import "fmt"

type xTurn struct {
	context *TicTacToe
}

func (state *xTurn) numEntered(num int) {
	turnError := state.context.currentWorld.placeStone(num, x)

	if turnError != nil {
		fmt.Println(turnError.Error())
		return
	}

	winner := state.context.currentWorld.getWinner()
	if winner != nil {
		printEndResult(*winner)
		state.context.currentState = &end{context: state.context}
	} else {
		state.context.currentState = &oTurn{context: state.context}
	}
}
