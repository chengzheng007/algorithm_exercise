package main

import (
	"errors"
	"fmt"
)

// 回文数-单链表判断
// 时间复杂度：O(n)

type Node struct {
	Data int
	Next *Node
}

func main() {
	//list := []int{1}
	//list := []int{1, 1}
	//list := []int{1, 2}
	//list := []int{1, 2, 1}
	//list := []int{1, 1, 1, 1}
	//list := []int{1,2,3,4}
	list := []int{1, 2, 3, 2, 1}

	head, err := initNodeList(list)
	if err != nil {
		fmt.Printf("initNodeList() error(%v)\n", err)
		return
	}
	print(head)
	fmt.Printf("is isPalindromic Number? %t\n", isPalindromic(head))
	//print(head)
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

func isPalindromic(head *Node) bool {
	// 无元素或只有一个元素
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return false
	}
	pSlow := head.Next
	pFast := head.Next.Next

	for {
		// pSlow到达链表中心
		if pFast == nil || pFast.Next == nil {
			break
		}

		if pSlow.Next != nil {
			pSlow = pSlow.Next
		}
		if pSlow.Next != nil {
			pFast = pSlow.Next.Next
		}
	}

	// 从中间节点往后反转链表
	// 循环结束，p是尾节点
	p := pSlow
	q := pSlow.Next
	for q != nil {
		r := q.Next
		q.Next = p
		p = q
		q = r
	}

	// 从两端向中间遍历比较
	q = head.Next
	tail := p
	// true:是回文数  false:非回文数
	flag := true
	for q != pSlow && tail != pSlow {
		if q.Data != tail.Data {
			flag = false
			break
		}
		q = q.Next
		tail = tail.Next
	}

	// 将中间到尾节点部分调整回来
	tail = p
	q = p.Next
	for {
		if q == pSlow {
			q.Next = p
			break
		}

		node := q.Next
		q.Next = p
		p = q
		q = node
	}

	// 尾节点Next置为空
	tail.Next = nil

	return flag
}

