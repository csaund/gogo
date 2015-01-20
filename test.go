package main

import (
	"errors"
	"fmt"
	"strings"
)

func main() {
	b := newBoard(10)
	var i,j int
	var err error
	for {
		b.printBoard()
		for{
			i, j = getMove()
			err = b.placeStone(Wpiece, i, j)
			if err == nil{
				break
			}
		}
		b.printBoard()
		for{
			i, j = getMove()
			err = b.placeStone(Bpiece, i, j)
			if err == nil{
				break
			}
		}
	}
}

const (
	Nopiece = iota
	Wpiece
	Bpiece
)

type Board struct {
	cells [][]int
	size  int
}

func getMove() (i, j int) {
	//prompt user
	//get things
	for {
		fmt.Println("please enter i j")
		var i, j int
		_, err := fmt.Scanf("%d %d\n", &i, &j)
		if err == nil {
			return i, j
		} else {
			fmt.Println("Invalid move. Try again")
		}
	}

}

func (b *Board) init() {
	for _, i := range b.cells {
		for ind, _ := range i {
			i[ind] = Nopiece
		}
	}
}

func newBoard(sqsize int) *Board {
	//new different from make. Make only on primitives
	b := new(Board)
	b.cells = make([][]int, sqsize)
	for i := range b.cells {
		b.cells[i] = make([]int, sqsize)
	}
	b.init()
	b.size = sqsize
	return b
}

func (b *Board) printBoard() {
	sep := strings.Repeat("-", b.size*4+1)
	fmt.Println(sep)
	for _, i := range b.cells {
		lnstr := make([]string, b.size)
		for ind, _ := range i {
			if i[ind] == Wpiece {
				lnstr[ind] = " O "
			} else if i[ind] == Bpiece {
				lnstr[ind] = " X "
			} else {
				lnstr[ind] = "   "
			}
		}
		line := strings.Join(lnstr, "|")
		fmt.Printf("|%s|\n", line)
		fmt.Println(sep)
	}
}

func (b *Board) placeStone(Player, i, j int) error {
	if b.cells[i][j] != Nopiece {
		return errors.New("Board cell already occupied")
	} else {
		b.cells[i][j] = Player
	}
	return nil
}

func (b *Board) effectBoard() {

}
