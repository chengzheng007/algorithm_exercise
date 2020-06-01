package main

// top k problem

func topK(list []int, k int) []int {
	size := len(list)
	if size <= k {
		return list
	}

	heapList := make([]int, k+1) // 从1开始存数据，非0

	// 将堆中数据填充至k个并建小顶堆
	i := 0
	for i < k {
		heapList[i+1] = list[i]
		i++
	}
	buildSmallTopHeap(heapList, k)

	// 处理剩余数据
	for i < size {
		// 如果比堆顶元素大，替换为堆顶元素，并从堆顶开始堆化
		if heapList[1] < list[i] {
			heapList[1] = list[i]
			smallTopHeapify(heapList, k, 1)
		}
		i++
	}
	return heapList[1:]
}

func buildSmallTopHeap(list []int, k int) {
	for i := k / 2; i >= 1; i-- {
		smallTopHeapify(list, k, i)
	}
}

func smallTopHeapify(list []int, n, i int) {
	for {
		pos := i
		if 2*i <= n && list[i] > list[2*i] {
			pos = 2 * i
		}
		if 2*i+1 <= n && list[pos] > list[2*i+1] {
			pos = 2*i + 1
		}
		if i == pos {
			break
		}
		list[i], list[pos] = list[pos], list[i]
		i = pos
	}
}
