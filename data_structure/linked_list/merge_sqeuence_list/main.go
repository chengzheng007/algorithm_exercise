package main

import (
	"errors"
	"fmt"
)

// 合并两个顺序相同的有序链表
// 时间复杂度：O(m+n)

type Node struct {
	Data int
	Next *Node
}

func main() {
	list1 := []int{1,2,3,6,8,13,27}
	list2 := []int{4,5,7,8,9,10,15}
	head1, err := initNodeList(list1)
	if err != nil {
		fmt.Printf("initNodeList(list1) error(%v)\n", err)
		return
	}
	head2, err := initNodeList(list2)
	if err != nil {
		fmt.Printf("initNodeList(list2) error(%v)\n", err)
		return
	}

	fmt.Println("list1:")
	print(head1)
	fmt.Println("list2:")
	print(head2)

	fmt.Println("exec merge...")
	mergeSequenceList(head1, head2)
	fmt.Println("list1:")
	print(head1)
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

func mergeSequenceList(head1, head2 *Node) {
	if head1 == nil || head1.Next == nil || head2 == nil || head2.Next == nil {
		return
	}

	p1 := head1.Next
	p2 := head2.Next
	p1Prev := head1
	p2MergeStart := p2

	// 先并到list1链表
	for p1 != nil && p2 != nil {
		if p1.Data <= p2.Data {
			p1Prev = p1
			p1 = p1.Next
			continue
		}

		
		for p2.Next != nil {
			if p2.Next.Data <= p1.Data {
				p2 = p2.Next
			} 
			break  
		}
		tmp := p2.Next

		// 将第二个链表的一部分合并到
		p1Prev.Next = p2MergeStart
		p2.Next = p1
		
		p2 = tmp
		p2MergeStart = p2
	}

	// 如果list2链表还有数据，直接连接到list1后面
	if p2 != nil {
		p1Prev.Next = p2
	}
}
