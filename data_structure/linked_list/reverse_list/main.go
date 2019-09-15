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
	
	fmt.Println("reverse list no head:")
	p := &Node{Data:1}
	p.Next = &Node{Data:2}
	p.Next.Next = &Node{Data:3}
	p.Next.Next.Next = &Node{Data:4}
	rep := reverseListNoHead(p)
	printNoHead(rep)
	
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

// 递归倒置链表，带头结点
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
	// 将自己的Next置为空，目的是将第一个节点（翻转后是最后一个）的Next置为空
	// 另一方面，【递归函数其实已经帮我们记住的先后顺序】，所以这里及时把当前节点的
	// Next置为空，回到递归上一层，我们也只需将这一层的这个Next指向上一层的节点
	// 至于这一层的Next指向什么，已经不用关心
	p.Next = nil
}

// 尾递归倒置链表，不带头结点
func reverseListNoHead(h *Node) *Node {
	if h == nil || h.Next == nil {
		return h
	}
	ret := reverseListNoHead(h.Next)
	h.Next.Next = h
	h.Next = nil
	// 重点是需要把倒置前的尾节点返回回来，不然就丢了
	return ret
}

func printNoHead(h *Node) {
	if h == nil {
		fmt.Println("empty linked list")
		return
	}
	for h != nil {
		fmt.Printf("%v ", h.Data)
		h = h.Next
	}
	fmt.Println()
}
