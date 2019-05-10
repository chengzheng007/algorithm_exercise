package main

import "fmt"

func main() {
	//fmt.Println("n queen test:")
	//n := 4
	//nqueen := NewNQueen(n)
	//nqueen.CalPos()

	ksnapsack(0, 0)
	fmt.Println()
	ksnapsack2(weight,n,w)

	fmt.Println()
	ksnapsack3(weight,n,w)

	fmt.Println()
	maxVal := ksnapsack4(weight, value, n, w)
	fmt.Printf("ksnapsack4 max val:%d\n", maxVal)

	fmt.Println()
	fmt.Println("mindst backtracking:")
	minDistBT(0, 0, 0, matrix, matrixSize)
	fmt.Printf("min dist value sum:%d\n", minDist)

	fmt.Println("mindst dynamic program:")
	minDistDP(matrix, matrixSize)

	fmt.Println()
	fmt.Println("dynamic programming exercise:")
	fmt.Printf("fibonacciNumDP:%d fibonacciNumBT:%d\n", fibonacciNumDP(10), fibonacciNumBT(10))

	fmt.Println()
	findChangeBT(0, 0, 0, rmbs, len(rmbs), 15)
	fmt.Printf("findChangeBT:%d\n", currNeedNum)
	changeDPNum := findChangeDP(rmbs, len(rmbs), 15)
	fmt.Printf("findChangeDP:%d\n", changeDPNum)

	fmt.Println()
	fmt.Printf("lisDP:%d\n", lisDP(series))
}

