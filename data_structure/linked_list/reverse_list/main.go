package main

import (
	"fmt"
	"errors"
)

// 单链表倒置
// 时间复杂度：O(n)

type Node struct {
	Data int
	Next *Node
}                       

func main() {
	list := []int{1,3,7,9}
	head, err := initNodeList(list)
	if err != nil {
		fmt.Printf("initNodeList(list) error(%v)\n", err)
		return
	}
	print(head)
	reverseInLoopWay(head)
	fmt.Println("after reverseInLoopWay reverse:")
	print(head)

	fmt.Println("after reverseInRecursiveWay reverse:")
	reverseInRecursiveWay(head, head.Next)
	print(head)
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
	if head == nil {
		fmt.Println("empty linked list")
		return 
	}
	i := 1
	for head.Next != nil {
		fmt.Printf("node id: %d, value:%d, next:%p, addr:%p\n", i, head.Next.Data, head.Next.Next, head.Next)
		i++
		head = head.Next
	}
}

// 直接遍历法倒置链表
func reverseInLoopWay(head *Node) {
	// 没有或只有一个元素
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return
	}
	p := head.Next
	q := p.Next
	for q != nil {
		r := q.Next
		q.Next = p
		p = q
		q = r
	}
	head.Next.Next = nil	// 之前的尾节点Next置空
	head.Next = p			// 头指针指向现在的第一个节点
}

// 递归倒置链表
func reverseInRecursiveWay(head *Node, p *Node) {
	if head == nil || p == nil {
		return
	}
	// 尾节点，需要将头节点的Next指向尾节点
	// 必须将头结点传入，才能最终将带头节点指向最后一个节点
	if p.Next == nil {
		head.Next = p
		return
	}
	reverseInRecursiveWay(head, p.Next)
	// 将后续节点Next指向自己
	p.Next.Next = p
	p.Next = nil
}