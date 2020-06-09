func intersect(nums1 []int, nums2 []int) []int {
	mp := make(map[int]int)
	// appear as many times as it shows in both arrays
	// 尽可能出现多次，但不能一个数组中包含的次数，nums1出现了2次2
	// nums2出现了3次2，那么最多出现2次2
	for _, num := range nums1 {
		times, ok := mp[num]
		if ok {
			mp[num] = times + 1
		} else {
			mp[num] = 1
		}
	}
	var res []int
	for _, num := range nums2 {
		times, ok := mp[num]
		if ok && times > 0 {
			res = append(res, num)
			mp[num] = times - 1
		}
	}
	return res
}