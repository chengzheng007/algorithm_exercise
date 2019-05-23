package main

import "fmt"

// 基于链表法冲突的散列表

type hashMap struct {
	slots int
	list  []*linkedNode
}

type linkedNode struct {
	Val  int
	Next *linkedNode
}

func NewHashMap(size int) *hashMap {
	return &hashMap{
		slots: size,
		list:  make([]*linkedNode, size),
	}
}

func (hm *hashMap) Add(val int) bool {
	slot := val % hm.slots
	p := hm.list[slot]
	var pre *linkedNode
	for p != nil {
		if p.Val == val {
			return false
		}
		pre = p
		p = p.Next
	}
	// slot is empty
	if pre == nil {
		hm.list[slot] = &linkedNode{Val: val, Next: nil}
	} else {
		pre.Next = &linkedNode{Val: val, Next: nil}
	}
	return true
}

func (hm *hashMap) Del(val int) bool {
	slot := val % hm.slots
	p := hm.list[slot]
	var pre *linkedNode
	for p != nil {
		if p.Val == val {
			// find it
			break
		}
		pre = p
		p = p.Next
	}
	// cannot find val
	if p == nil {
		return false
	}
	// val to delete is the first elem in the slot
	if pre == nil {
		hm.list[slot] = p.Next
	} else { // middle node to delete
		pre.Next = p.Next
	}
	p = nil
	return true
}

func (hm hashMap) Print() {
	valList := make([]int, 0)
	for i := 0; i < hm.slots; i++ {
		var p *linkedNode
		p = hm.list[i]
		for p != nil {
			valList = append(valList, p.Val)
			p = p.Next
		}
	}
	fmt.Printf("val in hash map:%v\n", valList)
}
