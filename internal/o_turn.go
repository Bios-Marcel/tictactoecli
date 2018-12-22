package internal

import "fmt"

type oTurn struct {
	context *TicTacToe
}

func (state *oTurn) numEntered(num int) {
	turnError := state.context.currentWorld.placeStone(num, o)

	if turnError != nil {
		fmt.Println(turnError.Error())
		return
	}

	winner := state.context.currentWorld.getWinner()
	if winner != nil {
		printEndResult(*winner)
		state.context.currentState = &end{context: state.context}
	} else {
		state.context.currentState = &xTurn{context: state.context}
	}
}
