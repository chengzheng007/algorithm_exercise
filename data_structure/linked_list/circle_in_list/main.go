package main

import (
	"fmt"
	"errors"
)

// 检测链表中有没有环

type Node struct {
	Data int
	Next *Node
}

func main() {

	list := []int{1,2,3,4,5}
	nonCircleListHead, err := initNodeList(list)
	if err != nil {
		fmt.Printf("initNodeList(list) error(%v)\n", err)
		return 
	}

	fmt.Printf("Is nonCircleListHead a circle linked list? %t\n", isCircleLinkedList(nonCircleListHead))

	circleList := initCircleLinkedList()
	fmt.Printf("Is circleList a circle linked list? %t\n", isCircleLinkedList(circleList))


	// list3为单循环链表，Circular List
	list3 := new(Node)
	node := &Node{Data:1, Next:list3}
	list3.Next = node
	fmt.Printf("Is list3 a circle linked list? %t\n", isCircleLinkedList(list3))
}

func initCircleLinkedList() *Node {
	head := new(Node)
	p := head
	var addr *Node
	
	for i := 1; i < 9; i++ {
		node := &Node{Data:i}

		if i == 5 {
			addr = node
		}

		p.Next = node
		p = node
	}

	p.Next = addr

	return head
}


// 带头结点
func initNodeList(list []int) (*Node, error) {
	if len(list) == 0 {
		return nil, nil
	}
	head := new(Node)
	p := head
	for _, val := range list {
		node := new(Node)
		if node == nil {
			return nil, errors.New("no spare space")
		}
		node.Data = val
		p.Next = node
		p = node

	}
	return head, nil
}

func print(head *Node) {
	i := 1
	for head.Next != nil {
		fmt.Printf("node id: %d, value:%d, next:%p, addr:%p\n", i, head.Next.Data, head.Next.Next, head.Next)
		i++
		head = head.Next
	}
}

func isCircleLinkedList(head *Node) bool {
	// 快慢指针遍历链表，慢指针走一步，快指针走两步
	pSlow := head
	pFast := head

	for {
		// 到链表末尾，说明没有环
		if pFast == nil {
			return false
		}

		if pSlow != nil {
			pSlow = pSlow.Next
		}

		if pFast.Next != nil {
			pFast = pFast.Next.Next
		}

		if pSlow == pFast {
			return true
		}
	}
	
	return false
}


