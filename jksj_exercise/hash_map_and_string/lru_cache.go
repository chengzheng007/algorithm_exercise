package main

import (
	"errors"
)

var (
	// ErrCacheCap indicate cache size must great than zero
	ErrCacheCap = errors.New("invalid lru cache capacity")
)

type lruCache struct {
	capacity int
	count    int
	hashMap map[string]*dblistNode
	root  *dblistNode
}

type dblistNode struct {
	Val int
	prev *dblistNode
	next *dblistNode
}

func NewLruCache(capacity int) (*lruCache, error) {
	if capacity <= 0 {
		return nil, ErrCacheCap
	}
	lch := &lruCache{
		capacity:    capacity,
		hashMap: make(map[string]*dblistNode, capacity),
		root: &dblistNode{},
	}
	lch.root.next = lch.root
	lch.root.prev = lch.root
	return lch, nil
}

func (lch *lruCache) Put(k string, v int) {
	// 已存在，移动到最前面
	if node, ok := lch.hashMap[k]; ok && node != nil {


		node.prev.next = node.next
		node.next.prev = node.prev
		return
	}
	
}

// 将node节点添加到root节点后
func (lch *lruCache) addNode(node *dblistNode) {
	node.next = lch.root.next
	node.prev = lch.root
	lch.root.next = node
	if lch.root.prev == lch.root {
		lch.root.prev = node
	}
	lch.count++
}

func (lch *lruCache) remove(node *dblistNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
	node.prev = nil
	node.next = nil
	lch.count--
}
