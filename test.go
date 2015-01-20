package main

import (
	"fmt"
	"strings"
)

func main() {
	b := newBoard(10)
	b.printBoard()
}

func helloworld(){
	//no python arrays, just "slices"
	//declare slice using "make" command, like new (slice/map/struct/C things)
	strs := make([]string, 0) //dynamic, make is like new
	strs = append(strs, "Hello,")
	strs = append(strs, "world!")
	//both of these are slices	cause []
	str := strings.Join(strs, " ")
	fmt.Println(str)
}

const (
	Nopiece=iota
	Wpiece
	Bpiece
)

type Board struct{ cells [][]int }

func (b *Board) init() {
	for _,i := range(b.cells){
		for ind,_ := range(i){
			i[ind] = Nopiece
		}
	}
}

func newBoard(sqsize int)(*Board){
	//new different from make. Make only on primitives
	b := new(Board)
	b.cells = make([][]int, sqsize)
	for i := range(b.cells){
		b.cells[i] = make([]int, sqsize)
	}
	b.init()
	return b
}

func (b *Board) printBoard(){
	for _,i := range(b.cells){
		for ind,_ := range(i){
			if i[ind] == Wpiece{
				fmt.Println('O')
			} else if i[ind] == Bpiece{
				fmt.Println('X')
			} else{
				fmt.Println(' ')
			}
		}
	}
}