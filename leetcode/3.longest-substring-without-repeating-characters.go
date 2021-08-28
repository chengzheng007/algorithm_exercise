/*
 * @lc app=leetcode id=3 lang=golang
 *
 * [3] Longest Substring Without Repeating Characters
 */
// 将不重复的字符记录到集合里
// 两个下标i,j前进，j-i+1即为当前不重复长度
func lengthOfLongestSubstring(s string) int {
    n := len(s)
    if n <= 1 {
        return n
    }
    var (
        set = make(map[byte]struct{})
        maxlen = 0 // 记录最大长度
        i, j = 0, 0
    )
    for i < n && j < n {
        if _, ok := set[s[j]]; !ok {
            set[s[j]] = struct{}{}
            // j右移，不重复子串长度增加
            // 只有长度增加时和maxlen比较
            j++ 
            if j-i > maxlen { // 注意是j自增之后算长度，无需再加1
                maxlen = j-i
            }
        } else { // 出现重复字符，删掉i对应字符
            delete(set, s[i])
            i++
        }
    }
    
    return maxlen
}
