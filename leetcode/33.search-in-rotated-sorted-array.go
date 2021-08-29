/*
 * @lc app=leetcode id=33 lang=golang
 *
 * [33] Search in Rotated Sorted Array
 */
// 即使数组被旋转过，我们仍然可以利用这个数组的递增性，使用二分查找。对于当前的中点，
// 如果它指向的值小于等于右端，那么说明右区间是排好序的；反之，那么说明左区间是排好序的。
// 如果目标值位于排好序的区间内，我们可以对这个区间继续二分查找；反之，我们对于另一半区
// 间继续二分查找。
// 注意，因为数组存在重复数字，如果中点和左端的数字相同，如Input: nums=[2,5,6,0,0,1,2], target=0
// 我们并不能确定是左区间全部相同，还是右区间完全相同。在这种情况下，我们可以简单地将左端点右移一位
// 然后继续进行二分查找
func search(nums []int, target int) int {
    length := len(nums)
    low := 0
    high := length-1
    
    for low <= high {
        mid := (low+high)/2
        if nums[mid] == target {
            return mid
        }
        
        if nums[low] == nums[mid] { // 本题规定元素值均不相同，第一个if可以判断拿掉
            low++
        } else if nums[mid] < nums[high] { // 本题规定所有元素值不同
            if target > nums[mid] && target <= nums[high] {
                low = mid+1
            } else {
                high = mid-1
            }
        } else {
            if target >= nums[low] && target < nums[mid] {
                high = mid-1
            } else {
                low = mid+1
            }
        }
    }
    
    return -1
}

// 确定最小值分界点，重新转换为二分查找
func search(nums []int, target int) int {
    length := len(nums)
    low := 0
    high := length-1
    
    // 找到翻转数组的最小值所在索引
    // 翻转点的后面一个值
    for low < high {
        mid := (low+high)/2
        // 中点数字比高点数字大，说明左边已经有序
        // 那么最小值肯定在右侧
        if nums[mid] > nums[high] {
            low = mid+1
        } else { // 否则往左侧找，注意这里high不能赋为mid-1，可能就是mid
            high = mid 
        }
    }
    // 记录最小值/分界点索引
    pivotIdx := low
    
    // 确定target在分界点的哪一边，转换为普通二分查找！
    if nums[pivotIdx] <= target && nums[length-1] >= target {
        low = pivotIdx
        high = length-1
    } else {
        low = 0
        high = pivotIdx
    }
    for low <= high {
        mid := (low+high)/2
        if nums[mid] == target {
            return mid
        } else if nums[mid] < target {
            low = mid+1
        } else {
            high = mid-1
        }
    }
    
    return -1
}

func search(nums []int, target int) int {
	// 先找出翻转后的分界点索引
	// 或者说找到数组中最小数的索引
	size := len(nums)
	lo := 0
	hi := size - 1
	for lo < hi {
		mid := (lo + hi) / 2
		// 当中间数比最后一个值大时，想象
		// 应该是左边的翻转的数比较的多，且都是偏大的数
		// 比分界点右边的大(3,4,5,6,1,2)，也就是说分界点在
		// 偏向右侧的地方，因此lo右移
		// 否则hi左移，左移时hi要等于mid
		if nums[mid] > nums[hi] {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	// 分界点为lo的地方
	pivotIdx := lo
	lo = 0
	hi = size - 1
	// 下面开始用二分法定位元素位置
	for lo <= hi {
		mid := (lo + hi) / 2
		// 注意，因为实际的分界点不在中间（当没有翻转时的mid）
		// 因此这里的mid只是索引大小上的中点（非值排列的中点）
		// 那么偏移了中间多少呢，就是这个pivotIdx个位置
		// 所以要用计算的中间点+偏移量，因为可能大于整个数组长度
		// 所以要对长度取模，也就是realMid
		// realMid从数值上来说，才是真正的中点
		realMid := (mid + pivotIdx) % size
		if nums[realMid] == target {
			return realMid
		} else if nums[realMid] < target {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	return -1
}

