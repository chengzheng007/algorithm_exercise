package main

import "fmt"

// Trie树/字典树/字符树

const (
	List_Size = 26
	Char_Start byte = 'a'
)


type Trie struct {
	Root *TrieNode // root node doesn't store any significant char
}

// only allow[a-z] 26 chars
type TrieNode struct {
	C byte
	// true: the char is ending for a word, false indicate this is a middle char for a word
	IsEndingChar bool
	// store its child nodes, for example, char a's ascii is 97, then it store at
	// List[97-97], index is 0, and b will store in List[98-97], index is 1
	List [List_Size]*TrieNode
}

func NewTrie() *Trie {
	return &Trie{
		Root:&TrieNode{C:'/'},
	}
}

func (t *Trie) Insert(word string) error {
	p := t.Root
	if p == nil {
		return fmt.Errorf("invalid object value, root is nil")
	}
	for i := 0; i < len(word); i++ {
		index := int(word[i] - Char_Start)
		if index < 0 || index >= List_Size {
			return fmt.Errorf("not allowed char exist, must between [a,z]")
		}
		if p.List[index] == nil {
			p.List[index] = &TrieNode{C:word[i]}
		}
		p = p.List[index]
	}
	p.IsEndingChar = true
	return nil
}

// match a word
func (t *Trie) Match(pattern string) bool {
	p := t.Root
	if p == nil {
		return false
	}
	for i := 0; i < len(pattern); i++ {
		index := int(pattern[i] - Char_Start)
		if index < 0 || index >= List_Size {
			return false
		}
		// when not null, it must be a corresponding char
		if p.List[index] == nil {
			return false
		}
		p = p.List[index]
	}
	// if p.IsEndingChar is false, that is, pattern is only a prefix that
	// can't match completely
	return p.IsEndingChar
}

func (t *Trie) MatchByPrefix(prefix string) (words []string) {
	p := t.Root
	if p == nil {
		return
	}
	index := -1
	for i := 0; i < len(prefix); i++ {
		index = int(prefix[i] - Char_Start)
		if index < 0 || index >= List_Size {
			words = []string{}
			return
		}
		if p.List[index] == nil {
			words = []string{}
			return
		}
		p = p.List[index]
	}
	// get all words in trie that prefix equal to param "prefix"
	words = getSuffix(p, prefix)
	return
}

func getSuffix(p *TrieNode, str string) (list []string) {
	if p == nil {
		return
	}
	for i := 0; i < List_Size; i++ {
		if p.List[i] == nil {
			continue
		}
		// dfs recursive function to find the longest matched suffix string
		clist := getSuffix(p.List[i], str+string(p.List[i].C))
		// when char is the ended char as a word, put it into a slice, and merge clist which was its child nodes' suffix
		if p.List[i].IsEndingChar {
			suffix := str + string(p.List[i].C)
			list = append(list, suffix)
		}
		// when passing by non-ended char or an ending char, and its child nodes' suffix not empty
		// merge them to the slice
		if len(clist) > 0 {
			list = append(list, clist...)
		}
	}
	return
}
