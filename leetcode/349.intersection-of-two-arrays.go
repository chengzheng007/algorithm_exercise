func intersection(nums1 []int, nums2 []int) []int {
	mp := make(map[int]int)
	for _, num := range nums1 {
		// val记录次数
		mp[num] = 1
	}
	res := make([]int, 0)
	for _, num := range nums2 {
		if times, ok := mp[num]; ok {
			// 防止nums2中重复添加到res
			if times == 1 {
				res = append(res, num)
			}
			mp[num] = times + 1
		}
	}

	return res
}