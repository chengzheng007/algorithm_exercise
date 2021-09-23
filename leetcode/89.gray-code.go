// 此题为编码问题--格雷码（像霍夫曼编码那种一定规律编码方式）
// 需查看解题
func grayCode(n int) []int {
    if n < 0 {
        return nil
    }
    
    res := make([]int, 1)
    for i := 0; i < n; i++ {
        highestBit := 1 << i
        for j := len(res)-1; j >= 0; j-- { //要反着遍历，才能对称
            res = append(res, highestBit|res[j])
        }
    }
    
    return res
}
