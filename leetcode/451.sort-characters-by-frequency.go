func frequencySort(s string) string {
    // 统计各个字符频率
    count := make(map[byte]int)
    // 记录最大频次
    maxFreq := 0
    for _, char := range s {
        bchar := byte(char)
        count[bchar]++
        if count[bchar] > maxFreq {
            maxFreq = count[bchar]
        }
    }
    
    // 将频次相同的放到一个桶中
    bucket := make([][]byte, maxFreq+1)
    for char, frequency := range count {
        bucket[frequency] = append(bucket[frequency], char)
    }
    
    // 遍历桶，拼接字符串
    res := make([]byte, len(s))
    pos := 0
    for i := maxFreq; i > 0; i-- {
        // 将对应频次桶的字符依次拼接
        for _, char := range bucket[i] {
            for j := 0; j < i; j++ {
                res[pos] = char
                pos++
            }
        }
    }
    
    // go技巧：不拷贝一个新字符串数组
    return *(*string)(unsafe.Pointer(&res))
}
