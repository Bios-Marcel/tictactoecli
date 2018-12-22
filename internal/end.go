package internal

type end struct {
	context *TicTacToe
}

func (state *end) numEntered(num int) {
	state.context.currentState = &splash{context: state.context}
}
