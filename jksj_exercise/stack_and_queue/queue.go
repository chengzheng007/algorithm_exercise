package main

import "fmt"

type queueHead struct {
	head *queueNode
	tail *queueNode
	size int
}

type queueNode struct {
	Val int
	Next *queueNode
}

func NewQueue() *queueHead {
	return &queueHead{nil, nil, 0}
}

// push from queue tail
func (q *queueHead) Push(val int) {
	node := &queueNode{Val:val}
	// empty queue
	if q.tail == nil {
		q.tail = node
		q.head = node
		q.size = 1
		return
	}

	q.tail.Next = node
	q.tail = node
	q.size++
}

func (q *queueHead) Pop() *queueNode {
	if q.Empty() {
		return nil
	}
	node := q.head
	q.head = node.Next
	q.size--
	return node
}

func (q queueHead) Empty() bool {
	return q.size <= 0
}

func (q queueHead) Print() {
	list := make([]int, q.size)
	i := 0
	for p := q.head; p != nil; p = p.Next {
		list[i] = p.Val
		i++
	}
	fmt.Printf("elems in queue:%v\n", list)
}
