// 862. Shortest Subarray with Sum at Least K

func shortestSubarray(nums []int, k int) int {
    n := len(nums)
    // 前缀和数组 sum[i] = nums[0]+...+nums[i-1]
    sum := make([]int, n+1)
    for i := 1; i <= n; i++ {
        sum[i] = sum[i-1]+nums[i-1]
    }
    
    // double ended queue
    // 单调递增队列，存储的sum数组下标，对于queue[i1, i2]
    // 满足sum[i1] < sum[i2]
    minlen := n+1
    dqueue := list.New()
    for i := 0; i <= n; i++ {
        for dqueue.Len() > 0 {
            elem := dqueue.Back()
            index := elem.Value.(int)
            if sum[i] <= sum[index] {
                dqueue.Remove(elem)
            } else {
                break
            }
        }
        for dqueue.Len() > 0 {
            elem := dqueue.Front()
            index := elem.Value.(int)
            if sum[i] - sum[index] >= k {
                if minlen > i-index{
                    minlen = i-index
                }
                dqueue.Remove(elem)
            } else {
                break
            }
        }
        dqueue.PushBack(i)
    }
    
    if minlen == n+1 {
        minlen = -1
    }
    return minlen
}
