func findKthLargest(nums []int, k int) int {
    size := len(nums)
    if size == 0 || k <= 0 || k > size {
        return 0
    }
    // Note: 我们是把nums以升序排列，然后寻找k大数，所以实际上k大
    // 的数会出现在排序后size-k的位置，这个位置的下标为排序后数组的
    // 中的n-k，可以举例验证
    targetK := size-k
    l := 0
    r := size-1
    for l < r {
        // 使用快排的分区函数排序
        idx := partition(nums, l, r)
        if idx == targetK {
            break
        } else if idx > targetK { // 分区点下标较大，k大数落在左边区间
            r = idx-1
        } else {
            l = idx+1
        }
    }
    
    return nums[targetK]
}

func partition(nums []int, l, r int) int {
    i := l
    j := r
    pivot := nums[l]
    for i < j {
        // 右边小的移到左边
        for i < j && nums[j] > pivot {
            j--
        }
        if i < j {
            nums[i] = nums[j]
            i++
        }
        
        // 左边大的移到右边
        for i < j && nums[i] < pivot {
            i++
        }
        if i < j {
            nums[j] = nums[i]
            j--
        }
    }
    nums[i] = pivot
    return i
}
