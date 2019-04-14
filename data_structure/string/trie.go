package main

import "fmt"

// 字典树/字符树/Trie树

const (
	Char_Size = 26
	Char_Start byte = 'a'
)


type Trie struct {
	Root *TrieNode // 根节点不存储字符
}

// 只允许[a-z]26个字符
type TrieNode struct {
	C byte
	IsEndingChar bool			// 是否是结束节点
	List [Char_Size]*TrieNode 	// 列表下标存储对应的英文字符
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
		if index < 0 || index >= Char_Size {
			return fmt.Errorf("not allowed char exist, must bettewn [a,z]")
		}
		if p.List[index] == nil {
			p.List[index] = &TrieNode{C:word[i]}
		}
		p = p.List[index]
	}
	p.IsEndingChar = true
	return nil
}

// 匹配一个单词
func (t *Trie) Match(pattern string) bool {
	p := t.Root
	if p == nil {
		return false
	}
	for i := 0; i < len(pattern); i++ {
		index := int(pattern[i] - Char_Start)
		if index < 0 || index >= Char_Size {
			return false
		}
		// 不为空一定是对应的字符
		if p.List[index] == nil {
			return false
		}
		p = p.List[index]
	}
	// if p.IsEndingChar is false, that is, pattern is only a prefix
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
		if index < 0 || index >= Char_Size {
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
	for i := 0; i < Char_Size; i++ {
		if p.List[i] == nil {
			continue
		}
		// dfs recursive function to find the longest matched suffix string
		clist := getSuffix(p.List[i], str+string(p.List[i].C))
		// when char is the ended char as a word, put it into a slice, and merge clist which was its child nodes' suffix
		if p.List[i].IsEndingChar {
			suffix := str+string(p.List[i].C)
			list = append(list, suffix)
			list = append(list, clist...)
			//fmt.Printf("suffix:%s, suffix_list:%v clist:%v\n", suffix, list, clist)
		} else if len(clist) > 0 {
			// when passing by non-ended char, and its child nodes' suffix not empty, merge them to slice
			list = append(list, clist...)
		}
	}
	return
}
