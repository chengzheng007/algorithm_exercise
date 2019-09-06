package main

// 一个二维数组，每一维长度不一致，以轮询的方式将其合并到一个数组中
// 例如 [[1,2,3],[4,5]]，合并的结果是[1,4,2,5,3]
func roundTripMerge(list [][]int) []int {
	totalSize := 0
	for _, lst := range list {
		totalSize += len(lst)
	}
	// 申请一个大数组装所有数据
	newArr := make([]int, totalSize)
	// 申请一个与第一维长度相等的数组，表示这一列数组目前哪个数据应被装入newArr
	// 下标表示list一维的哪一个数组，值表示是哪个数组
	index := make([]int, len(list))
	// cursor游标表示当前放入大数组中的下标位置
	cursor := 0
	// lstIdx表示当前轮询的是哪一个一维数组
	lstIdx := 0
	for cursor < totalSize {
		// 当且仅当第一维数组中的当前索引小于该一维数组元素的个数时
		// 将该一维数组中的元素加入到大数组中
		if index[lstIdx] < len(list[lstIdx]) {
			newArr[cursor] = list[lstIdx][index[lstIdx]]
			// 只有真正加了这个一维数组中的数据时，记录该一维数组的索引值+1，游标+1
			index[lstIdx]++
			cursor++
		}
		// 开始加入下一个第一维数组
		lstIdx++
		// 如果已达到最后一个第一维数组，将其置为第一个
		if lstIdx == len(list) {
			lstIdx = 0
		}
	}
	return newArr
}
