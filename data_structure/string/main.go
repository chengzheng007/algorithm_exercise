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
	fmt.Println()

	fmt.Printf("kmp next array test:\n")
	pattern = "ABCABCBA"
	fmt.Printf("pattern:%s\n", pattern)
	next := getNext(pattern, len(pattern))
	for k, v := range next {
		fmt.Printf("k:%d v:%d\n", k, v)
	}
	
	pattern = "ABABABC"
	fmt.Printf("pattern:%s\n", pattern)
	next = getNext(pattern, len(pattern))
	for k, v := range next {
		fmt.Printf("k:%d v:%d\n", k, v)
	}
	fmt.Println()


	// kmp test
	mains = "ababaeabac"
	pattern = "abac"
	//pattern = "ab"
	idx = kmp(mains, pattern)
	fmt.Printf("kmp idx:%d\n", idx)

}
