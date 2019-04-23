package main

import (
	"fmt"
	"container/list"
)

// ac-automation/ac自动机

type AcNode struct {
	C byte
	IsEndingChar bool
	List [List_Size]*AcNode
	Length int     	// when char is ending char, length of the word
	Fail *AcNode	// fail pointer point to the longest prefix of the word which this node matched in other words in tree
}

func (acn AcNode) String() string {
	var failPointChar string
	if acn.Fail != nil {
		failPointChar = string(acn.Fail.C)
	} else {
		failPointChar = "nil"
	}
	return fmt.Sprintf("{C:%s IsEndingChar:%t Length:%d FailPointChar:%s}", string(acn.C), acn.IsEndingChar, acn.Length, failPointChar)
}

type AcAutomation struct {
	Root *AcNode
}

func NewAcAutomation() *AcAutomation {
	return &AcAutomation{
		Root:&AcNode{C:'/'},
	}
}

// build trie tree
func (ac *AcAutomation) BuildTrieTree(words []string) error {
	if ac.Root == nil {
		return fmt.Errorf("root pointer is nil")
	}
	for _, word := range words {
		p := ac.Root
		length := len(word)
		for i := 0; i < length; i++ {
			index := int(word[i] - Char_Start)
			if index < 0 || index >= List_Size {
				return fmt.Errorf("not allowed char exist, must between [a,z]")
			}
			if p.List[index] == nil {
				p.List[index] = &AcNode{C:word[i]}
			}
			p = p.List[index]
		}
		if length > 0 {
			p.IsEndingChar = true
		}
		p.Length = length
	}
	return nil
}

// buid fail pointer 
func (ac *AcAutomation) BuildFailPointer() bool {
	if ac.Root == nil {
		return false
	}
	dblist := list.New()
	dblist.PushBack(ac.Root)

	for dblist.Len() > 0 {
		elem := dblist.Front()
		p := elem.Value.(*AcNode)
		for i := 0; i < List_Size; i++ {
			pc := p.List[i]
			if pc == nil {
				continue
			}
			if p == ac.Root {
				pc.Fail = ac.Root
			} else {
				q := p.Fail
				for q != nil {
					qc := q.List[int(pc.C - Char_Start)]
					// find fail pointer
					if qc != nil {
						pc.Fail = qc
						break
					}
					q = q.Fail
				}
				// cannot find, point to root
				if q == nil {
					pc.Fail = ac.Root
				}
			}	
			dblist.PushBack(pc)
		}
		dblist.Remove(elem)
	}
	return true
}

func (ac *AcAutomation) MatchFind(pattern string) {
	p := ac.Root
	if p == nil {
		return 
	}
	length := len(pattern)
	for i := 0; i < length; i++ {
		idx := int(pattern[i] - Char_Start)
		for p.List[idx] == nil && p != ac.Root {
			p = p.Fail
		}
		// find one node match pattern char
		p = p.List[idx]
		if p == nil {
			p = ac.Root
		}
		tmp := p
		for tmp != ac.Root {
			// matched one word
			if tmp.IsEndingChar {
				fmt.Printf("matched str:%s\n", pattern[i-tmp.Length+1:i+1])
			}
			tmp = tmp.Fail
		}
	}
	return
}
