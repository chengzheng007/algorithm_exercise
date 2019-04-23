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
	
	// kmp test
	mains = "ababaeabac"
	pattern = "abac"
	//pattern = "ab"
	idx = kmp(mains, pattern)
	fmt.Printf("kmp idx:%d\n", idx)

	// trie test
	fmt.Println()
	fmt.Println("trie test")
	trie := NewTrie()
	words := []string{"how", "hi", "her", "hello", "so", "see", "hea", "heas", "heb", "head", "hero", "hetls", "slice"}
	for _, word := range words {
		_ = trie.Insert(word)
	}
	fmt.Printf("words:%v\n", words)
	pattern = "hello"
	fmt.Printf("match %s? %v\n", pattern, trie.Match(pattern))
	pattern = "so"
	fmt.Printf("match %s? %v\n", pattern, trie.Match(pattern))
	pattern = "he"
	fmt.Printf("match %s? %v\n", pattern, trie.Match(pattern))
	pattern = "her"
	fmt.Printf("match %s? %v\n", pattern, trie.Match(pattern))
	pattern = "herr"
	fmt.Printf("match %s? %v\n", pattern, trie.Match(pattern))
	fmt.Println("-----")


	prefix := "he"
	list := trie.MatchByPrefix(prefix)
	fmt.Printf("prefix:%s, MatchByPrefix:%v\n", prefix, list)
	prefix = "s"
	list = trie.MatchByPrefix(prefix)
	fmt.Printf("prefix:%s, MatchByPrefix:%v\n", prefix, list)

	fmt.Println("ac automation test:")
	acAuto := NewAcAutomation()
	words = []string{"abcd", "bcd", "c"}
	acAuto.BuildTrieTree(words)
	acAuto.BuildFailPointer()
	acAuto.MatchFind("abcd")
}
