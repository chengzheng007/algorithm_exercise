package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("reverse order counting test:")
	list := []int{5,3,6,1,2,4,9}
	count := reverseOrderCounting(list)
	fmt.Printf("reverse order count:%d\n", count)
	fmt.Printf("after sort:%v\n", list)

	list = []int{5,3,6,1,2,4,9}
	count = brouteReverseOrderCountint(list)
	fmt.Printf("validate count:%d\n", count)

	fmt.Printf("\nback tracking test:\n")
	fmt.Printf("n queen:\n")
	n := 4
	list1 := make([]int, n)
	nQueen(list1, n, 0)

	weight := []int{2,2,4,6,3}
	n = 5
	maxW := 9
	var inW int
	packagesProblem(weight, n, 0, 0, maxW, &inW)
	fmt.Printf("max load in package:%d\n", inW)

	fmt.Println("\ndynamic programming test:")
	maxW = 8
	fmt.Printf("all weight:%v\n", weight)
	inw := knapsack(weight, n, maxW)
	fmt.Printf("knapsack max load in package:%d\n", inw)

	inw = knapsack2(weight, n, maxW)
	fmt.Printf("knapsack2 max load in package:%d\n", inw)


	fmt.Println("\n\nnmin dist test:")
	w := [][]int{
		[]int{1,3,5,9},
		[]int{2,1,3,4},
		[]int{5,2,6,7},
		[]int{6,8,4,3},
	}
	n = 4
	minDist := math.MaxInt32
	backTrackMinDist(w, n, 0, 0, 0, &minDist)
	fmt.Printf("backtracking minDist:%d\n", minDist)
	fmt.Printf("dp minDist:%d\n", dpMinDist(w, n))

}
