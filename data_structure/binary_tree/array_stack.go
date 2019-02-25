package main

import (
	"fmt"
)

type arrayStack struct {
	list []*Node
	size int
}

func InitArrayStack() *arrayStack {
	stack := &arrayStack{
		list:nil,
		size:0,
	}
	return stack
}

func (s *arrayStack) Push(node *Node) {
	if len(s.list) > s.size {
		s.list[s.size] = node
	} else {
		s.list = append(s.list, node)
	}
	if len(s.list) - s.size > 10 {
		list := make([]*Node, s.size)
		for i := 0; i < s.size; i++ {
			list[i] = s.list[i]
		}
		s.list = list
	}
	s.size++
}

func (s *arrayStack) Pop() (*Node, bool) {
	if s.size <= 0 {
		return nil, false
	}
	node := s.list[s.size-1]
	s.list[s.size-1] = nil
	s.size--
	return node, true
}

func (s *arrayStack) Empty() bool {
	return s.size <= 0
}

func (s *arrayStack) Top() *Node {
	if s.size <= 0 {
		return nil
	}
	return s.list[s.size-1]
}

func (s *arrayStack) String() string {
	list := make([]string, 0)
	for i := 0; i < s.size; i++ {
		if s.list[i] != nil {
			list = append(list, s.list[i].Data)
		}
	}
	return fmt.Sprintf("elems size:%d, real size:%d, elems list:%v", s.size, len(s.list), list)
}
