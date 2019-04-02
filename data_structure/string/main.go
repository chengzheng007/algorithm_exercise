package main

import "fmt"

func main() {
	//bf test

	mains := "abdbcef"
	pattern := "a"

	idx := bf(mains, pattern)
	fmt.Printf("bf find pattern(%s) index:%d\n", pattern, idx)

	pattern = "bdcefabc"
	idx = bf(mains, pattern)
	fmt.Printf("bf find pattern(%s) index:%d\n", pattern, idx)
}
