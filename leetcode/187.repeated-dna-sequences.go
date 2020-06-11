func findRepeatedDnaSequences(s string) []string {
	n := len(s)
	if n <= 10 {
		return nil
	}
	// 两个集合，mp1存储第一次出现的子串，如果后边遍历是mp1已经
	// 为true，那么该子串已出现超过1次，加入mp2集合
	// 最后统计mp2结果即可
	mp1 := make(map[string]bool)
	mp2 := make(map[string]bool)

	for i := 0; i+10 <= n; i++ {
		if mp1[s[i:i+10]] {
			mp2[s[i:i+10]] = true
		} else {
			mp1[s[i:i+10]] = true
		}
	}
	res := make([]string, len(mp2))
	i := 0
	for str := range mp2 {
		res[i] = str
		i++
	}
	return res
}