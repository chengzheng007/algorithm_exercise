package main

import (
	"fmt"
)

func main() {
	list1 := []int{1,3,9, 100}
	lk1 := initNodeList(list1)
	list2 := []int{2,5,8,11,16}
	lk2 := initNodeList(list2)
	newh := mergeTwoSeqLinkedList(lk1, lk2)
	print(newh)
}


type Node struct {
	Data int
	Next *Node
}

// with head node
func initNodeList(list []int) *Node {
	if len(list) == 0 {
		return nil
	}
	head := new(Node)
	p := head
	for _, val := range list {
		node := new(Node)
		node.Data = val
		p.Next = node
		p = node

	}
	return head
}

func print(head *Node) {
	i := 1
	for head.Next != nil {
		fmt.Printf("node id: %d, value:%d, next:%p, addr:%p\n", i, head.Next.Data, head.Next.Next, head.Next)
		i++
		head = head.Next
	}
}