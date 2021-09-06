// 解题思路: https://mp.weixin.qq.com/s/-afUNplTiJ7LeeAvoH7w0g
// 每个区间由两部分组成 [start,end]，如果两个区间 [a1,b1]和[a2,b2] 不重叠需要满足：a1 != a2 且
// 	若 a1 < a2，则 b1 < a2；
// 	若 a1 > a2，则 b2 < a1； 

// 题目描述中说明，区间集合已经按每个区间的开始升序排列，所以给定区间集合 [a1,b1],[a2,b2],[a3,b3],...,[an,bn]，一定满足以下条件： 
// 	a1 < b1 < a2 < b2 < a3 < b3 < ... < an < bn。

// 要插入的新的区间是 [start,end] 会将区间集合分成三个部分left set，overlapped set和right set： 
// left set：[a1,b1],...,[ai,bi]，0 <= i <= n；
// overlapped set：[ai+1,bi+1],...,[aj,bj]，i <= j <= n；
// right set：[aj+1,bj+1],...,[an,bn]，i <= j <= n；

// 注意：三个集合都有可能为空！  并有：
// 	bi < start < end < aj+1;
// 	start <= bi+1 and aj <= end。

// overlapped set 中的集合要和新插入的区间融合，融合之后的结果是 [ min(ai+1, start), max(bj, end) ]。  
// 算法流程  所以解题的思路就是把集合中的区间按条件分成上述的三个部分：
// 	1.假设集合中的某个区间是 [am,bm]，新插入的区间是 [start,end]。 
// 	如果 bm < start，[am,bm] 属于 left set；
// 	如果 end < am，[am,bm] 属于 right set；
// 	否则 [am,bm] 属于 overlapped set。 
// 	2.overlapped set 中的集合和新插入的区间 [start,end] 融合，融合成一个区间 [min_start, max_end]。  
// 	3.将三个部分组合成新的区间集合：left set + [min_start, max_end] + right set。 

// 分成左边、重合区间、右边三部分，最终结果将三部分拼接在一起
func insert(intervals [][]int, newInterval []int) [][]int {
    var leftPart, rightPart [][]int
    start := newInterval[0]
    end := newInterval[1]

    for _, pair := range intervals {
        // 位于新插入区间的左边
        if pair[1] < start {
            leftPart = append(leftPart, []int{pair[0],pair[1]})
        } else if pair[0] > end {
            // 位于新插入区间的右边
            rightPart = append(rightPart, []int{pair[0],pair[1]})
        } else {
            // 有重合发生，取左边左边界最小，右边界最大值
            if pair[0] < start {
                start = pair[0]
            }
            if pair[1] > end {
                end = pair[1]
            }
        }
    }
    
    var res [][]int
    res = append(res, leftPart...)
    res = append(res, []int{start, end})
    res = append(res, rightPart...)
    
    return res
}
