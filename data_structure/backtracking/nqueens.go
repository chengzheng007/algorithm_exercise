package main

import "fmt"

// n queen problem

type nQueen struct {
	pos []int  // index indicate row number, array value show which column to place
	queenNum int
}

func NewNQueen(n int) *nQueen {
	if n <= 0 {
		return nil
	}
	nq := &nQueen{
		queenNum:n,
		pos:make([]int, n),
	}
	//for i := 0; i < n; i++ {
	//	nq.pos[i] = -1
	//}
	return nq
}

// calculate all queens' position on chessboard
func (nq *nQueen) CalPos() {
	nq.calPos(0)
}

func (nq *nQueen) calPos(row int) {
	if row == nq.queenNum {
		nq.PrintQueen()
		return
	}
	for column := 0; column < nq.queenNum; column++ {
		if nq.isPlaceablePos(row, column) {
			nq.pos[row] = column
			//fmt.Printf("true row:%d column:%d true pos:%v\n", row, column, nq.pos)
			nq.calPos(row+1)
		}
	}
}	

func (nq *nQueen) isPlaceablePos(row, column int) bool {
	leftup := column-1
	rightup := column+1
	for i := row-1; i >= 0; i-- {
		// does the i(th) row column(th) column has chess ?
		if nq.pos[i] == column {
			return false
		}
		// does left up pos has chess ?
		if leftup >= 0 && nq.pos[i] == leftup {
			return false
		}
		// does right up pos has chess ?
		if rightup < nq.queenNum && nq.pos[i] == rightup {
			return false
		}
		leftup--
		rightup++
	}
	return true
}

func (nq *nQueen) PrintQueen() {
	fmt.Printf("pos:%v\n", nq.pos)
	for row := 0; row < nq.queenNum; row++ {
		for column := 0; column < nq.queenNum; column++ {
			if nq.pos[row] == column {
				fmt.Print("Q ")
			} else {
				fmt.Print("* ")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}