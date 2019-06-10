package main

// size of list equal to (n+1)
// big heap
func buildHeap(list []int, n int) {
	// 只有前一半有叶子节点，只需调整前一半
	// 倒序遍历前一半节点，从上往下堆化
	// 需要调整前一半每一个节点，而不是仅仅从堆顶开始就行，原因是每一次只调整了一个节点
	for i := n/2; i >= 1; i-- {
		heapifyTopToBottom(list, n, i)
	}
}

func heapifyTopToBottom(list []int, n, i int) {
	for {
		maxPos := i
		if 2*i <= n && list[2*i] > list[i] {
			maxPos = 2*i
		}
		if 2*i+1 <= n && list[2*i+1] > list[maxPos] {
			maxPos = 2*i+1
		}
		// satisfy heap condition
		if i == maxPos {
			break
		}
		// swap
		list[i], list[maxPos] = list[maxPos], list[i]

		i = maxPos
	}
}

func heapSort(list []int, n int) {
	heapList := make([]int, n+1)
	copy(heapList[1:], list)

	// 建堆
	buildHeap(heapList, n)

	i := n
	for i > 1 {
		heapList[1], heapList[i] = heapList[i], heapList[1]
		i--
		// 因为与堆顶交换，从堆顶开始需要再做一次堆化
		heapifyTopToBottom(heapList, i, 1)
	}
	// copy back
	copy(list, heapList[1:])
}