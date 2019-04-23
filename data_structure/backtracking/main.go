package main

import "fmt"

func main() {
	fmt.Println("n queen test:")
	n := 4
	nqueen := NewNQueen(n)
	nqueen.CalPos()
}
