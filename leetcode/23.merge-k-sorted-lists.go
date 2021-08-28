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

// ##########   分治法（递归实现）   #########
// 本质是转换为两个链表的合并
func mergeKLists(lists []*ListNode) *ListNode {
    n := len(lists)
    if n == 0 {
        return nil
    }
    if n == 1 { // 不用合并
        return lists[0]
    }
    return mergeExec(lists, 0, n-1)
}

func mergeExec(lists []*ListNode, low, high int) *ListNode {
    if low == high {
        return lists[low]
    }
    mid := (low+high)/2
    leftPart := mergeExec(lists, low, mid)
    rightPart := mergeExec(lists, mid+1, high)
    return mergeTwoList(leftPart, rightPart)
}

func mergeTwoList(l1, l2 *ListNode) *ListNode {
    if l1 == nil {
        return l2
    }
    if l2 == nil {
        return l1
    }
    
    head := &ListNode{}
    ph := head
    p := l1
    q := l2
    for p != nil && q != nil {
        if p.Val < q.Val {
            ph.Next = p
            p = p.Next
        } else {
            ph.Next = q
            q = q.Next
        }
        ph = ph.Next
    }
    if p != nil {
        ph.Next = p
    } else {
        ph.Next = q
    }
    return head.Next
}


// #########     每次取最小堆堆顶加到链表后面      #########
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
		// 如何保证永远取的是所有元素中最小的？就在这里，只要取出一个元素，就要把它所在列表
		// 的下一个元素添加进来，如果有的话
		if node.Next != nil {
			heap.Push(&pq, node.Next)
		}
	}
	return head.Next
}

