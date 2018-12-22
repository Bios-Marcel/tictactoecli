package internal

type splash struct {
	context *TicTacToe
}

func (state *splash) numEntered(num int) {
	state.context.currentState = &xTurn{context: state.context}
}
