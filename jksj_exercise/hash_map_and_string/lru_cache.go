package main

import (
	"encoding/json"
	"errors"
	"fmt"
)

var (
	// ErrCacheCap indicate cache size must great than zero
	ErrCacheCap = errors.New("invalid lru cache capacity")
)

type lruCache struct {
	capacity int
	count    int
	hashMap  map[string]*dblistNode
	// 双向链表
	root     *dblistNode
}

type dblistNode struct {
	Key  string
	Val  int
	prev *dblistNode
	next *dblistNode
}

func NewLruCache(capacity int) (*lruCache, error) {
	if capacity <= 0 {
		return nil, ErrCacheCap
	}
	lch := &lruCache{
		capacity: capacity,
		hashMap:  make(map[string]*dblistNode, capacity),
		root:     &dblistNode{},
	}
	lch.root.next = lch.root
	lch.root.prev = lch.root
	return lch, nil
}

func (lch *lruCache) Set(k string, v int) {
	if node, ok := lch.hashMap[k]; ok && node != nil {
		lch.removeNode(node)
		node.Val = v
		lch.addNode(node)
		return
	}

	node := &dblistNode{Key: k, Val: v}
	lch.addNode(node)
	lch.hashMap[k] = node
	lch.count++

	if lch.count > lch.capacity {
		tail := lch.popTail()
		delete(lch.hashMap, tail.Key)
		lch.count--
	}
}

func (lch *lruCache) Get(k string) int {
	node, ok := lch.hashMap[k]
	if !ok {
		return -1
	}
	lch.moveToHead(node)
	return node.Val
}

// addNode put node after root
func (lch *lruCache) addNode(node *dblistNode) {
	node.next = lch.root.next
	node.prev = lch.root
	// 因为这个指向，导致最初第一个节点插入的时候，新增节点永远在head后面，且lch.root.prev指向第一个插入的节点
	// 由于增加节点永远在lch.root.head后面，所以无论增加多少节点，lch.root.prev也不会被修改，永远是指向的尾节点
	// lch.root.pre就是尾节点
	node.next.prev = node
	lch.root.next = node
}

// remove do remove node from double list
func (lch *lruCache) removeNode(node *dblistNode) {
	if node == nil {
		return
	}
	node.next.prev = node.prev
	node.prev.next = node.next
	node.prev = nil
	node.next = nil
}

// moveToHead do move node to head of list
func (lch *lruCache) moveToHead(node *dblistNode) {
	lch.removeNode(node)
	lch.addNode(node)
}

// popTail do pop tail node of list
func (lch *lruCache) popTail() *dblistNode {
	// 双向循环链表，决定了root.prev指向的是链表的最后一个节点
	node := lch.root.prev
	// no node in list
	if node == lch.root {
		return nil
	}
	lch.removeNode(node)
	return node
}

func (lch lruCache) Print() {
	type tempStruct struct {
		K string `json:"key"`
		V int    `json:"val"`
	}
	list := make([]*tempStruct, lch.count)
	i := 0
	p := lch.root.next
	for p != lch.root && p != nil {
		list[i] = &tempStruct{K: p.Key, V: p.Val}
		p = p.next
		i++
	}
	jstr, _ := json.Marshal(list)
	fmt.Printf("elems num:%d, val:%s\n", lch.count, jstr)
}
