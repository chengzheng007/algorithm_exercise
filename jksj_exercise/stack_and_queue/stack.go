package main

import "fmt"

type StackHead struct {
	size int
	Next *StackNode
}

type StackNode struct {
	Val int
	Next *StackNode
}

func NewStack() *StackHead {
	return &StackHead{0, nil}
}

func (stack *StackHead) Push(val int) {
	stackNode := &StackNode{Val:val}
	stackNode.Next = stack.Next
	stack.Next = stackNode
	stack.size++
}


func (stack *StackHead) Pop() *StackNode {
	if stack.Empty() {
		return nil
	}
	node := stack.Next
	stack.Next = node.Next
	stack.size--
	return node
}

func (stack StackHead) Empty() bool {
	return stack.size <= 0
}

func (stack StackHead) Print() {
	list := make([]int, stack.size)
	p := stack.Next
	for i := 0; i < stack.size; i++ {
		list[i] = p.Val
		p = p.Next
	}
	fmt.Printf("elems in stack:%v\n", list)
}

