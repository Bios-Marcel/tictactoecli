package internal

import (
	"errors"
	"fmt"
)

type owner int

const (
	none owner = 0
	x    owner = 1
	o    owner = 2
)

type world struct {
	board map[int]owner
}

func (w *world) print() {
	fmt.Println("---")
	fmt.Printf("%s%s%s\n", w.board[1], w.board[2], w.board[3])
	fmt.Printf("%s%s%s\n", w.board[4], w.board[5], w.board[6])
	fmt.Printf("%s%s%s\n", w.board[7], w.board[8], w.board[9])
	fmt.Println("---")
	fmt.Println("")
}

func (w *world) placeStone(field int, player owner) error {
	if w.board[field] != none {
		return errors.New("This field is already occupied")
	}

	w.board[field] = player

	return nil
}

func (w *world) reset() {
	w.board = make(map[int]owner)
}

func (o owner) String() string {
	if o == none {
		return " "
	} else if o == x {
		return "X"
	} else {
		return "O"
	}
}

func (w *world) getWinner() *owner {
	//TODO Solve this in a less retarded way!
	//It is very stupid, verbose and unperformant, I was kinda lazy.

	//Horizontal 1->2->3
	if w.board[1] == w.board[2] && w.board[1] == w.board[3] {
		winner := w.board[1]
		return &winner
	}

	//Horizontal 4->5->6
	if w.board[4] == w.board[5] && w.board[4] == w.board[6] {
		winner := w.board[1]
		return &winner
	}

	//Horizontal 7->8->9
	if w.board[7] == w.board[8] && w.board[7] == w.board[9] {
		winner := w.board[1]
		return &winner
	}

	//Vertical 1->4->7
	if w.board[1] == w.board[4] && w.board[1] == w.board[7] {
		winner := w.board[1]
		return &winner
	}

	//Vertical 2-->5->8
	if w.board[2] == w.board[5] && w.board[2] == w.board[8] {
		winner := w.board[1]
		return &winner
	}

	//Vertical 3->6->9
	if w.board[3] == w.board[6] && w.board[3] == w.board[9] {
		winner := w.board[1]
		return &winner
	}

	//Diagonal 1->5->9
	if w.board[1] == w.board[5] && w.board[1] == w.board[9] {
		winner := w.board[1]
		return &winner
	}

	//Diagonal 3->5->7
	if w.board[3] == w.board[5] && w.board[3] == w.board[7] {
		winner := w.board[1]
		return &winner
	}

	return nil
}
