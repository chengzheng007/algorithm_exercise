package main

import "fmt"

func main() {
	lruCache := NewLRU(2)
	lruCache.Put(1, 1)
	lruCache.Put(2, 2)
	fmt.Printf("1 cached value:%v\n", lruCache.GetAllValues())

	val := lruCache.Get(1)
	fmt.Printf("key 1, val:%v\n", val)
	fmt.Printf("2 cached value:%v\n", lruCache.GetAllValues())

	lruCache.Put(3, 3)
	fmt.Printf("3 cached value:%v\n", lruCache.GetAllValues())

	val = lruCache.Get(2)
	fmt.Printf("key 2, val:%v\n", val)

	// 覆盖key为1的值
	lruCache.Put(1, 99)
	fmt.Printf("4 cached value:%v\n", lruCache.GetAllValues())
}

// lru-最近最少使用，使用链表（双向）和hashmap组成的结构，主要对外暴露方法为get、put
// 最近访问的将提升到链表尾部，最不常访问的将退到链表表头（尾部加入节点，头部删除节点）

type LRU struct {
	hashmap map[int]*doubleListNode
	dblist  *doubleList // 双向链表
	cap     int // lru缓存的容量
}

func NewLRU(cap int) *LRU {
	if cap <= 0 {
		panic("invalid cap")
	}
	return &LRU{
		hashmap: make(map[int]*doubleListNode),
		dblist:  newDoubleList(),
		cap:     cap,
	}
}

// 找不到key时返回-1
func (lru *LRU) Get(key int) int {
	node, ok := lru.hashmap[key]
	if !ok {
		return -1
	}
	// 将结点调至最近访问的位置
	lru.dblist.remove(node)
	lru.dblist.addLast(node)
	return node.val
}

func (lru *LRU) Put(key, val int) {
	node, ok := lru.hashmap[key]
	// 如果lru缓存中存在结点，将其调整到最近访问的位置
	if ok {
		lru.dblist.remove(node)
		// 复用该结点，更新value值
		node.val = val
		lru.dblist.addLast(node)
		return
	}

	// 若不存在，尝试加入结点
	// 超过lru缓存容量，则删除最老结点，包括hashmap中的key
	if lru.cap == lru.dblist.nodeSize() {
		delNode := lru.dblist.removeFirst()
		delete(lru.hashmap, delNode.key)
	}

	newNode := &doubleListNode{key: key, val: val}
	lru.dblist.addLast(newNode)
	lru.hashmap[key] = newNode
}

func (lru LRU) GetAllValues() [][]int {
	return lru.dblist.getNodeValue()
}

type doubleList struct {
	head *doubleListNode // 头结点指向双向链表中第一个结点，尾节点指向最后一个
	tail *doubleListNode // 最近访问/设置的结点在链表尾部
	size int             // 节点个数
}

func newDoubleList() *doubleList {
	dblist := &doubleList{
		head: &doubleListNode{},
		tail: &doubleListNode{},
	}
	// 为空时头结点的next指向尾结点
	// 尾节点的prev指向头结点
	dblist.head.next = dblist.tail
	dblist.tail.prev = dblist.head
	return dblist
}

// 双向链表结构如下，新添加在尾部，作为最近访问的数据，待删除的在头部，作为老数据
// |head| -> |k1| -> |k2| -> |k3| -> |tail|
// |    | <- |v1| <- |v2| <- |v3| <- |    |
// 添加结点至尾部
func (dblist *doubleList) addLast(node *doubleListNode) {
	if node == nil {
		panic("add nil node")
	}
	// 待加入结点链接到尾部
	node.prev = dblist.tail.prev
	node.next = dblist.tail
	// 修改原来尾节点处相关的前后指针
	dblist.tail.prev.next = node
	dblist.tail.prev = node

	dblist.size++
}

func (dblist *doubleList) remove(node *doubleListNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
	node.prev = nil
	node.next = nil
	dblist.size--
}

// 删除第一个结点
func (dblist *doubleList) removeFirst() *doubleListNode {
	// 链表为空
	if dblist.head.next == dblist.tail {
		return nil
	}
	firstNode := dblist.head.next
	dblist.remove(firstNode)
	return firstNode
}

func (dblist doubleList) nodeSize() int {
	return dblist.size
}

func (dblist doubleList) getNodeValue() [][]int {
	if dblist.size == 0 {
		return nil
	}
	values := make([][]int, dblist.size)
	i := 0
	p := dblist.head.next
	for p != nil && p != dblist.tail {
		values[i] = []int{p.key, p.val}
		p = p.next
		i++
	}
	return values
}

// 双向链表结点
type doubleListNode struct {
	key, val   int
	prev, next *doubleListNode
}
