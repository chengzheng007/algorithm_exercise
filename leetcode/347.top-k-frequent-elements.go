// 思想：将数组中的数按照num->频次统计到一个map中
// 统计最大频次值（因为有负数不能用数组）
// 新建一个桶，将对应频次值的数据放在桶后面，桶可以用
// 二维数组，一维是频次（频次永远>=0），二维是频次对应的
// 所有原始数值。然后从最大频次值开始递减循环，统计前k个数值，
// 有可能某个桶加到一半多出k个，需要break出循环
func topKFrequent(nums []int, k int) []int {
    n := len(nums)
    if n <= 1 {
        return nums
    }
    
    maxNum := nums[0]
    for i := 1; i < n; i++ {
        if maxNum < nums[i] {
            maxNum = nums[i]
        }
    }
    // count统计nums每个数字出现的次数，即频次
    count := make(map[int]int)
    maxCount := 0 // 统计最大频次
    for i := 0; i < n; i++ {
        count[nums[i]]++
        if maxCount < count[nums[i]] {
            maxCount = count[nums[i]]
        }
    }
    
    // bucket: 一维下标为频次，二维是一个数组，对应该频次对应的所有原始数值
    bucket := make([][]int, maxCount+1)
    for i, cnt := range count {
        bucket[cnt] = append(bucket[cnt], i)
    }
    
    ans := make([]int, k)
    pos := 0
    loop:
    for i := maxCount; i >= 0; i-- {
        for j := 0; j < len(bucket[i]); j++ {
            ans[pos] = bucket[i][j]
            pos++
            if pos >= k {
                break loop
            }
        }
    }
    
    return ans
}
