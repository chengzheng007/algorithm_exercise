// 时间复杂度，O(n)
func longestConsecutive(nums []int) int {
    // hash map 存储出现过的数字
	numsMap := make(map[int]struct{})
    for _, num := range nums {
        numsMap[num] = struct{}{}
    }
    
    maxlong := 0
    for _, num := range nums {
        // 以当前数字为起点查找，即num-1不存在时才找
        if _, ok := numsMap[num-1]; ok {
            continue
        }
        currNum := num
        currlong := 1
        for {
			_, ok := numsMap[currNum+1] // 单次查找O(1)，整个循环可看成O(1)
            if !ok {
                break
            }
            currNum++
            currlong++
        }
        if currlong > maxlong {
            maxlong = currlong
        }
    }
    return maxlong
}
