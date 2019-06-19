package main

// 逆序度统计
// 在一个数组序列中，如果前面的数大于后面的数，称为一个逆序度
// 逆序度即为有多少个这样的逆序对
// 时间复杂度：nO(n)
// 空间复杂度：O(n)
func reverseOrderCounting(list []int) int {
	size := len(list)
	if size <= 1 {
		return 0
	}
	count := 0
	mergeSort(list, 0, size-1, &count)
	return count
}

func mergeSort(list []int, low, high int, count *int) {
	if low >= high {
		return
	}
	mid := (low+high)/2
	mergeSort(list, low, mid, count)
	mergeSort(list, mid+1, high, count)
	merge(list, low, mid, high, count)
}

func merge(list []int, low, mid, high int, count *int) {
	tmp := make([]int, high-low+1)
	i := low
	j := mid+1
	k := 0           
	for i <= mid && j <= high {
		if list[i] <= list[j] {
			tmp[k] = list[i]
			i++
		} else {
			// 因为合并的两个序列都是有序序列（这里是从小到大），所以前一个序列中从i往后的数肯定均比后一个序列中的数大
			// 故逆序度是
			*count += mid - i + 1
			tmp[k] = list[j]
			j++
		}
		k++
	}
	for i <= mid {
		tmp[k] = list[i]
		i++
		k++
	}
	for j <= high {
		tmp[k] = list[j]
		j++
		k++
	}
	// 复制会原数组
	copy(list[low:high+1], tmp)
}

// 暴力统计，O(n^2)
func brouteReverseOrderCountint(list []int) int {
	size := len(list)
	count := 0
	if size <= 1 {
		return count
	}
	for i := 0; i < size-1; i++ {
		for j := i+1; j < size; j++ {
			if list[i] > list[j] {
				count++
			}
		}
	}
	return count
}
