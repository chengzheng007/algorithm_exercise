func searchMatrix(matrix [][]int, target int) bool {
    // 每次寻找最右上角的一个的元素，如果相等即返回
    // 如果比查找值大，删除所在列，因为所在的列是递增的
    // 列的一个值比查找值大，同一列都比它大
    // 如果比查找值小，删除所在的行，因为所在行也是递增的
    // 行最后一个元素是最大的，同一行其他元素都比它小
    // 那么查找值不可能出现在该行
    // 以上思路是每次去掉一列或一行，逐步缩小返回
    // 该题目出现在剑指offer
    rows := len(matrix)
    if rows == 0 {
        return false
    }
    cols := len(matrix[0])
    if cols == 0 {
        return false
    }
    
    row := 0
    col := cols-1
    for row < rows && col >= 0 {
        if matrix[row][col] == target {
            return true
        } else if matrix[row][col] > target {
            col--
        } else {
            row++
        }
    }
    
    return false
}

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
