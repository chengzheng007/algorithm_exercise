package main

import (
	"fmt"
)

// 循环队列，队尾进队头出

type arrayQueue struct {
	head, tail int
	capacity int // 最大容量
	list []*Node
}

func NewQueue(cap int) *arrayQueue {
	return &arrayQueue{
		capacity:cap,
		list:make([]*Node, cap),
	}
}

func (q *arrayQueue) Enter(node *Node) bool {
	if q.Full() {
		return false
	}
	q.list[q.tail] = node
	q.tail = (q.tail+1)%q.capacity
	return true
}

func (q *arrayQueue) Pop() (*Node, bool) {
	if q.Empty() {
		return nil, false
	}
	node := q.list[q.head]
	q.list[q.head] = nil
	q.head = (q.head+1)%q.capacity
	return node, true
}

func (q *arrayQueue) Len() int {
	return (q.tail + q.capacity - q.head) % q.capacity
}

func (q *arrayQueue) Empty() bool {
	return q.head == q.tail
}

func (q *arrayQueue) Full() bool {
	return (q.tail + 1) % q.capacity == q.head
}

func (q *arrayQueue) Head() *Node {
	if q.Empty() {
		return nil
	}
	return q.list[q.head]
}

func (q *arrayQueue) String() string {
	var list []string
	for i := q.head; i < q.tail; i = (i+1)%q.capacity {
		list = append(list, q.list[i].Data)
	}
	return fmt.Sprintf("queue size:%d, eles:%v", q.Len(), list)
}
