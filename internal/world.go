package internal

import (
	"errors"
	"fmt"
)

//Owner is a constant-type that is used for determining a winner.
type Owner int

const (
	//None means unowned.
	None Owner = 0
	//X stands for the player using the X marks.
	X Owner = 1
	//O stands for the player using the O marks.
	O Owner = 2
)

type world struct {
	board map[int]Owner
}

func (w *world) print() {
	fmt.Println("---")
	fmt.Printf("%s%s%s\n", w.board[1], w.board[2], w.board[3])
	fmt.Printf("%s%s%s\n", w.board[4], w.board[5], w.board[6])
	fmt.Printf("%s%s%s\n", w.board[7], w.board[8], w.board[9])
	fmt.Println("---")
	fmt.Println("")
}

func (w *world) placeStone(field int, player Owner) error {
	if w.board[field] != None {
		return errors.New("This field is already occupied")
	}

	w.board[field] = player

	return nil
}

func (w *world) reset() {
	w.board = make(map[int]Owner)
}

func (o Owner) String() string {
	if o == None {
		return " "
	} else if o == X {
		return "X"
	} else {
		return "O"
	}
}

func (w *world) getWinner() *Owner {
	//TODO Solve this in a less retarded way!
	//It is very stupid, verbose and unperformant, I was kinda lazy.

	//Horizontal 1->2->3
	if w.board[1] != None && w.board[1] == w.board[2] && w.board[1] == w.board[3] {
		winner := w.board[1]
		return &winner
	}

	//Horizontal 4->5->6
	if w.board[4] != None && w.board[4] == w.board[5] && w.board[4] == w.board[6] {
		winner := w.board[4]
		return &winner
	}

	//Horizontal 7->8->9
	if w.board[7] != None && w.board[7] == w.board[8] && w.board[7] == w.board[9] {
		winner := w.board[7]
		return &winner
	}

	//Vertical 1->4->7
	if w.board[1] != None && w.board[1] == w.board[4] && w.board[1] == w.board[7] {
		winner := w.board[1]
		return &winner
	}

	//Vertical 2-->5->8
	if w.board[2] != None && w.board[2] == w.board[5] && w.board[2] == w.board[8] {
		winner := w.board[2]
		return &winner
	}

	//Vertical 3->6->9
	if w.board[3] != None && w.board[3] == w.board[6] && w.board[3] == w.board[9] {
		winner := w.board[3]
		return &winner
	}

	//Diagonal 1->5->9
	if w.board[1] != None && w.board[1] == w.board[5] && w.board[1] == w.board[9] {
		winner := w.board[1]
		return &winner
	}

	//Diagonal 3->5->7
	if w.board[3] != None && w.board[3] == w.board[5] && w.board[3] == w.board[7] {
		winner := w.board[3]
		return &winner
	}

	if w.board[1] != None && w.board[2] != None && w.board[3] != None &&
		w.board[4] != None && w.board[5] != None && w.board[6] != None &&
		w.board[7] != None && w.board[8] != None && w.board[9] != None {
		winner := None
		return &winner
	}

	return nil
}
