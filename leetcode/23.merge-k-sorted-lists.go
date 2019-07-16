import "container/heap"

/*
 * @lc app=leetcode id=23 lang=golang
 *
 * [23] Merge k Sorted Lists
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type SmallTopHeap []*ListNode

func (bth SmallTopHeap) Len() int           { return len(bth) }
func (bth SmallTopHeap) Less(i, j int) bool { return bth[i].Val < bth[j].Val } // 小顶堆
func (bth SmallTopHeap) Swap(i, j int)      { bth[i], bth[j] = bth[j], bth[i] }

// Push、Pop必须为指针，因为要更改内部值
func (bth *SmallTopHeap) Push(x interface{}) {
	item := x.(*ListNode)
	*bth = append(*bth, item)
}
func (bth *SmallTopHeap) Pop() interface{} {
	old := *bth
	n := len(old)
	item := old[n-1]
	*bth = old[0 : n-1]
	return item
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	pq := make(SmallTopHeap, 0)
	// 初始化将数组中第一维度列表的第一个元素加入堆中
	for _, l := range lists {
		if l != nil {
			pq = append(pq, l)
		}
	}

	heap.Init(&pq)
	head := &ListNode{Val: -1}
	p := head
	for pq.Len() > 0 {
		node := heap.Pop(&pq).(*ListNode)
		p.Next = node
		p = p.Next
		// 每次选出堆中最小值节点后，需将该节点所在数组的下一个节点加入堆中，进行下一轮比较
		if node.Next != nil {
			heap.Push(&pq, node.Next)
		}
	}
	return head.Next
}

