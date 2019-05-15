package main

import (
	"fmt"
)

func main() {
	fmt.Println("mergeTwoSeqLinkedList:")
	list1 := []int{1,3,9, 100}
	lk1 := initNodeList(list1)
	list2 := []int{2,5,8,11,16}
	lk2 := initNodeList(list2)
	newh := mergeTwoSeqLinkedList(lk1, lk2)
	print(newh)

	fmt.Println("reverseList:")
	list3 := []int{3,5,8}
	lk3 := initNodeList(list3)
	reverseList(lk3)
	print(lk3)

	fmt.Println("middleOfLinkedList:")
	list4 := []int{3,5,7,9,10,11}
	lk4 := initNodeList(list4)
	print(lk4)
	mid := middleOfLinkedList(lk4)
	fmt.Printf("middle num:%d\n", mid.Data)
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