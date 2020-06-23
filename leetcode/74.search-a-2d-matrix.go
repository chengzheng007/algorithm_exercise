func searchMatrix(matrix [][]int, target int) bool {
	// 二分查找定位可能存在的行
	m := len(matrix)
	if m == 0 {
		return false
	}
	n := len(matrix[0])
	if n == 0 {
		return false
	}

	l := 0
	r := m - 1
	rowBound := 0

	for l <= r {
		mid := (l + r) / 2
		if matrix[mid][0] <= target {
			if matrix[mid][0] == target {
				return true
			}
			if mid == m-1 || matrix[mid+1][0] > target {
				rowBound = mid
				break
			} else {
				l = mid + 1
			}
		} else {
			r = mid - 1
		}
	}

	// 从rowBound开始从下往上，每一行二分查找
	for row := rowBound; row >= 0; row-- {
		// 如果末尾元素比target小，跳出
		if matrix[row][n-1] < target {
			break
		}

		l = 0
		r = n - 1
		for l <= r {
			mid := (l + r) / 2
			if matrix[row][mid] == target {
				return true
			} else if target > matrix[row][mid] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
	}

	return false
}