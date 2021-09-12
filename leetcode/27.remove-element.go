func removeElement(nums []int, val int) int {
    var index int
    // 假设index以前的元素不等于val，每当遇到不等于val的元素
    // 覆盖到index位置，index自增
    for i := 0; i < len(nums); i++ {
        if nums[i] != val {
            nums[index] = nums[i]
            index++
        }
    }
    return index
}
