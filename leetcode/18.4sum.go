func fourSum(nums []int, target int) [][]int {
    n := len(nums)
    if n < 4 {
        return nil
    }
    
    sort.Sort(sort.IntSlice(nums))
    var res [][]int

    for i := 0; i < n-3; i++ {
        // 跳过和第一个相等的
        if i > 0 && nums[i] == nums[i-1] {
            continue
        }
        for j := i+1; j < n-2; j++ {
            // 跳过和第一个j相等的数
            if j > i+1 && nums[j] == nums[j-1] {
                continue
            }
            
            sum := target - nums[i] - nums[j]
            low := j+1
            high := n-1
            for low < high {
                remain := nums[low] + nums[high]
                if remain == sum {
                    res = append(res, []int{nums[i], nums[j], nums[low], nums[high]})
                    for low < high && nums[low] == nums[low+1] {
                        low++
                    }
                    for low < high && nums[high] == nums[high-1] {
                        high--
                    }
                    low++
                    high--
                } else if remain > sum {
                    high--
                } else {
                    low++
                }
            }
        }       
    }
    
    return res
}
