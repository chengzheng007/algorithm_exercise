// 方法三：不排序给定单词，遍历每个单词直接与s
func findLongestWord(s string, dictionary []string) string {
    res := ""
    for _, word := range dictionary {
        if !isSubStr(s, word) {
            continue
        }
        
        if len(res) < len(word) || len(res) == len(word) && word < res {
            res = word
        }
    }
    return res
}

// 判断word是否为s的子串，【双指针】的变种
func isSubStr(s, word string) bool {
    i := 0
    j := 0
    for i < len(s) && j < len(word) {
        if s[i] == word[j] {
            i++
            j++
        } else {
            i++
        }
    }
    return j >= len(word)
}

// ################################## //
// 方法二：将给定字典排序，那每一个单词与字符串比较，看是否是一个子串
// note: 不会超时，但速度不快
func findLongestWord(s string, dictionary []string) string {
    // 靠前的是符合最要求的单词
    sort.Sort(words(dictionary))
    for _, word := range dictionary {
        if isSubStr(s, word) {
            return word
        }
    }
              
    return ""
}

// 将给定字典中的单词进行排序，长度长的排在前面，长度相等前面字母小的排在前面
type words []string

func (ws words) Len() int { return len(ws) }
func (ws words) Less(i, j int) bool { 
    len1 := len(ws[i])
    len2 := len(ws[j])
    if len1 == len2 {
        return ws[i] < ws[j]
    } else {
        return len1 > len2
    }
}
func (ws words) Swap(i, j int) { ws[i], ws[j] = ws[j], ws[i] }


// ################################################### //
// 回溯法，生成s的所有子串，跟给定字典中单词比对
// 遍历寻找判断长度最长、字典序最小的那个
// note: 回溯法会超时
func findLongestWord(s string, dictionary []string) string {
    words := make(map[string]struct{})
    for _, word := range dictionary {
        words[word] = struct{}{}
    }
    
    allSubs := make([]string, 0)
    generate(s, 0, nil, &allSubs)
    res := ""
    for _, subs := range allSubs {
        if _, ok := words[subs]; !ok {
            continue
        }
        if len(subs) > len(res) || len(subs) == len(res) && subs < res {
            res = subs
        }
    }
    
    return res
}

// 递归（回溯）获取所有s能组成的子字符串
func generate(s string, i int, curr []byte, subs *[]string) {
    if i == len(s) {
        *subs = append(*subs, string(curr))
        return
    }
    // 选择第i个字符
    curr = append(curr, s[i])
    generate(s, i+1, curr, subs)
    
    // 不选第i个字符
    curr = curr[:len(curr)-1]
    generate(s, i+1, curr, subs)
}
