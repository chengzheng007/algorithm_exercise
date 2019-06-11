package main

import (
	"container/heap"
)

// 合并k和有序列表

func mergeKList(list [][]int) []int {
	if len(list) == 0 {
		return []int{}
	}

	itemList := make(heapItemList, 0)
	// indexList记录每个子列表放入堆中的元素后，下一个从子列表取出的元素的下标
	indexList := make([]int, len(list))
	for i, numList := range list {
		if len(numList) > 0 {
			itemList = append(itemList, &heapItem{Idx:i, Val:numList[0]})
			indexList[i]++
		}
	}

	heap.Init(&itemList)
	
	ret := make([]int, 0)
	for itemList.Len() > 0 {
		item := heap.Pop(&itemList).(*heapItem)
		ret = append(ret, item.Val)
		// 将子列表剩余的数据取一个放进去，如果还有
		if indexList[item.Idx] < len(list[item.Idx]) {
			heap.Push(&itemList, &heapItem{Idx:item.Idx, Val:list[item.Idx][indexList[item.Idx]]})
			indexList[item.Idx]++
		}			
	}
	return ret
}


type heapItem struct {
	Idx int    // 子列表的在待排序列表list中的索引
	Val int    // 值
}


type heapItemList []*heapItem

func (h heapItemList) Len() int           { return len(h) }
func (h heapItemList) Less(i, j int) bool { return h[i].Val < h[j].Val }
func (h heapItemList) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *heapItemList) Push(x interface{}) {
	*h = append(*h, x.(*heapItem))
}

func (h *heapItemList) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

