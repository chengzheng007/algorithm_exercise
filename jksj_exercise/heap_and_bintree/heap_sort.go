package main

import "fmt"

// 原地排序，时间复杂度为O(n*logn)
func heapSort(list []int) {
	n := len(list)
	buildHeap(list, n)
	fmt.Printf("after buildHeap:%v\n", list)
	// 扫描列表+堆调整，O(n*logn)
	for i := n-1; i >= 0; i-- {
		// list[0]为堆顶，最大元素，和末尾元素交换
		list[0], list[i] = list[i], list[0]
		// 对调整，从0开始由下往上
		heapfy(list, i, 0)
	}
}

func buildHeap(list []int, n int) {
	// 堆列表约定从0开始
	// 只需对堆中的前n个节点做堆化，后n+1个是前n个节点的孩子节点
	// 从下往上堆化（从后往前），需要先堆化后面的元素
	// 注意i的起始下标，这里对列表从0开始，
	for i := n/2-1; i >= 0; i-- {
		heapfy(list, n, i)
	}
}

// 堆调整，也叫堆化，构造大顶堆，从上往下堆化
// n堆中元素总数，i当前堆化位置
func heapfy(list []int, n, i int) {
	var maxPos int
	for {
		maxPos = i
		// 对列表元素从0开始
		if 2*i+1 < n && list[i] < list[2*i+1] {
			maxPos = 2*i+1
		}
		if 2*i+2 < n && list[maxPos] < list[2*i+2] {
			maxPos = 2*i+2
		}
		if maxPos == i {
			break
		}
		// swap
		list[i], list[maxPos] = list[maxPos], list[i]
		i = maxPos
	}
}